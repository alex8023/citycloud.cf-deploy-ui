package controllers

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/entity"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ServiceDeployWebSocketController struct {
	beego.Controller
}

func (service *ServiceDeployWebSocketController) Get() {
	ws, err := upgrader.Upgrade(service.Ctx.ResponseWriter, service.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(service.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		return
	}
	for {
		time.Sleep(2 * time.Second)
		mesgType, message, _ := ws.ReadMessage()
		if message != nil && mesgType == websocket.TextMessage {
			var parameters = string(message)
			param := strings.Split(parameters, "-")
			if len(param) != 2 {
				writeStringMessage(ws, "Error parameter")
				writeStringMessage(ws, EndEOF)
				continue
			}
			switch param[0] {
			case "deploy":
				var serviceId = param[1]
				service.Deploying(ws, serviceId)
				writeStringMessage(ws, EndEOF)
			case "start":
				var serviceId = param[1]
				service.OperateService(ws, serviceId, utils.Service_Start)
				writeStringMessage(ws, EndEOF)
			case "stop":
				var serviceId = param[1]
				service.OperateService(ws, serviceId, utils.Service_Stop)
				writeStringMessage(ws, EndEOF)
			case "restart":
				var serviceId = param[1]
				service.OperateService(ws, serviceId, utils.Service_Restart)
				writeStringMessage(ws, EndEOF)
			default:
				writeStringMessage(ws, "Error parameter")
				writeStringMessage(ws, EndEOF)
			}
		}
	}
}

func (serviceDeploy *ServiceDeployWebSocketController) Deploying(ws *websocket.Conn, serviceIds string) {
	serviceId, err := strconv.ParseInt(serviceIds, 10, 64)
	if err != nil {
		writeStringMessage(ws, fmt.Sprintf("Error ServiceId,%s", err))
		return
	}
	//加载service
	writeStringMessage(ws, "Load Service")
	service := entity.Service{}
	service.Id = serviceId
	err = service.Load()

	if err != nil {
		writeStringMessage(ws, fmt.Sprintf("Error,%s", err))
		return
	}

	switch service.Where {
	case utils.Deploy_On_PaaS:
		serviceDeploy.Deploy2PaaS(ws, service)
	case utils.Deploy_On_Vms:
		serviceDeploy.Deploy2Vms(ws, service)
	default:
		writeStringMessage(ws, fmt.Sprintf("Unknow Platform %s", service.Where))
	}
}

// deploy 2 vms
func (serviceDeploy *ServiceDeployWebSocketController) Deploy2Vms(ws *websocket.Conn, service entity.Service) {
	//加载serviceDto
	//	writeStringMessage(ws, "Load ServiceDto")
	serviceDto := entity.ServiceDto{}
	serviceDto.Service = service
	err := serviceDto.Load()
	if err != nil {
		writeStringMessage(ws, fmt.Sprintf("Error,%s", err))
		return
	}

	//加载TemplateList
	writeStringMessage(ws, "Load TemplateList")
	template, templateLoadErr := entity.LoadTemplateList(service.Id)

	if templateLoadErr != nil {
		writeStringMessage(ws, fmt.Sprintf("Error,%s", templateLoadErr))
		return
	}

	//模板输出文件
	result := serviceDeploy.ParseTemplate(ws, service)
	if !result {
		return
	}
	//模板输出文件
	//传输文件
	writeStringMessage(ws, "传输文件...")

	vms := serviceDto.OnCustom

	sshRunner := utils.NewDeploySSHRunner(vms.Ip, vms.User, vms.Passwd)

	for index, temp1 := range template {
		if temp1.FileType != utils.FileTypes_War {
			continue
		}
		serviceDeploy.pushFiles(ws, index, temp1, sshRunner)

	}
	writeStringMessage(ws, "Finish deploying")

}

func (serviceDeploy *ServiceDeployWebSocketController) pushFiles(ws *websocket.Conn, index int, template entity.Template, sshRunner utils.DeploySSHRunner) {
	var out bytes.Buffer
	//	writeStringMessage(ws, fmt.Sprintf("%d,检测文件%s", index, template.Name))
	dir, filename, isDir, err := utils.Detechfile(template.TemplateFile)
	if err != nil {
		writeStringMessage(ws, fmt.Sprintf("%d,检测文件%s 发生错误：Error %s", index, template.Name, err))
		return
	}
	//	writeStringMessage(ws, fmt.Sprintf("%s,%s,%s,%s", dir, filename, isDir, err))
	//开始处理文件，压缩打包文件，scp传送
	//压缩文件，主要解决scp不拷贝symbolic link
	if isDir {
		//		writeStringMessage(ws, fmt.Sprintf("%d,检测文件%s是文件夹，准备打包", index, template.Name))
		cmdCommand := utils.Command{Name: "tar", Args: []string{"-czf", filename + ".tgz", filename}, Dir: dir, Stdin: ""}
		cmdRunner := utils.NewDeployCmdRunner()
		cmdRunner.RunCommandCmd(cmdCommand, &out)
		//		writeStringMessage(ws, fmt.Sprintf("%d,检测文件%s是文件夹，打包完毕", index, template.Name))
	}

	writeStringMessage(ws, fmt.Sprintf("deploy index:%d,fileName:%s,filePath:%s", index, template.Name, template.TemplateFile))

	filetgz := template.TemplateFile
	if isDir {
		filetgz = filetgz + ".tgz"
	}
	ssherror := sshRunner.Push(filetgz, "~/", &out)
	if ssherror != nil {
		writeStringMessage(ws, fmt.Sprintf("传输出错了。%s", ssherror))
		return
	}
	writeBufferMessage(ws, &out)

	if isDir {
		cmd := fmt.Sprintf("cd ~/ && tar -xzf %s", filename+".tgz")
		ssherror = sshRunner.RunCommand(cmd, &out)
		if ssherror != nil {
			writeStringMessage(ws, fmt.Sprintf("解压出错了。%s", ssherror))
		}
	}
}

//deploy 2 vms

func (serviceDeploy *ServiceDeployWebSocketController) Deploy2PaaS(ws *websocket.Conn, service entity.Service) {
	//加载serviceDto
	//	writeStringMessage(ws, "Load ServiceDto")
	serviceDto := entity.ServiceDto{}
	serviceDto.Service = service
	err := serviceDto.Load()

	if err != nil {
		writeStringMessage(ws, fmt.Sprintf("Error,%s", err))
		return
	}

	//模板输出文件
	result := serviceDeploy.ParseTemplate(ws, service)
	if !result {
		return
	}
	//模板输出文件
	onPaaS := serviceDto.OnPaaS

	var deployDir = customServiceDir + "/" + service.Name

	success := serviceDeploy.setApi(ws, onPaaS, deployDir)

	if !success {
		writeStringMessage(ws, "Set endpoint failed")
	}

	if success {
		success = serviceDeploy.loginPaaS(ws, onPaaS, deployDir)
		if !success {
			writeStringMessage(ws, "Login PaaS failed")
		}
	}

	if success {
		success = serviceDeploy.targetPaaS(ws, onPaaS, deployDir)
		if !success {
			writeStringMessage(ws, "Target PaaS failed")
		}
	}

	if success {
		success = serviceDeploy.pushApp(ws, onPaaS, deployDir)
		if !success {
			writeStringMessage(ws, "Deploy app failed")
		}
	}

	if success {
		writeStringMessage(ws, "Push app successful")
	}

}

// operate service
func (serviceDeploy *ServiceDeployWebSocketController) OperateService(ws *websocket.Conn, serviceIds string, operate string) {
	serviceId, err := strconv.ParseInt(serviceIds, 10, 64)
	if err != nil {
		writeStringMessage(ws, fmt.Sprintf("Error ServiceId,%s", err))
		return
	}
	//加载service
	writeStringMessage(ws, "Load Service")
	service := entity.Service{}
	service.Id = serviceId
	err = service.Load()

	if err != nil {
		writeStringMessage(ws, fmt.Sprintf("Error,%s", err))
		return
	}

	switch service.Where {
	case utils.Deploy_On_PaaS:
		serviceDeploy.OperateOnPaaS(ws, service, operate)
	case utils.Deploy_On_Vms:
		serviceDeploy.OperateOnVms(ws, service, operate)
	default:
		writeStringMessage(ws, fmt.Sprintf("Unknow Platform %s", service.Where))
	}
}

// operate service

// operate service on paas

func (this *ServiceDeployWebSocketController) setApi(ws *websocket.Conn, onPaaS entity.OnPaaS, deployDir string) bool {
	writeStringMessage(ws, "============================================")
	writeStringMessage(ws, "Set endpoint ")

	var out bytes.Buffer

	cfEndpointCommand := utils.Command{Name: "cf", Args: []string{"api", onPaaS.Api}, Dir: deployDir, Stdin: "", Env: []string{"CF_COLOR=false"}}

	cmdRunner := utils.NewDeployCmdRunner()

	cmdRunner.RunCommandAsyncCmd(cfEndpointCommand, &out)

	writeBytesBufferMessage(&out, &cmdRunner, ws)

	writeStringMessage(ws, "Set endpoint successful")
	writeStringMessage(ws, "============================================")

	return cmdRunner.Success()
}

func (this *ServiceDeployWebSocketController) loginPaaS(ws *websocket.Conn, onPaaS entity.OnPaaS, deployDir string) bool {
	writeStringMessage(ws, "============================================")
	writeStringMessage(ws, "Login PaaS ")

	var out bytes.Buffer

	cfLoginCommand := utils.Command{Name: "cf", Args: []string{"login", onPaaS.Api}, Dir: deployDir, Stdin: onPaaS.User + "\n" + onPaaS.Passwd + "\n", Env: []string{"CF_COLOR=false"}}

	cmdRunner := utils.NewDeployCmdRunner()

	cmdRunner.RunCommandAsyncCmd(cfLoginCommand, &out)

	writeBytesBufferMessage(&out, &cmdRunner, ws)

	writeStringMessage(ws, "Login PaaS successful")
	writeStringMessage(ws, "============================================")

	return cmdRunner.Success()
}

func (this *ServiceDeployWebSocketController) targetPaaS(ws *websocket.Conn, onPaaS entity.OnPaaS, deployDir string) bool {
	writeStringMessage(ws, "============================================")
	writeStringMessage(ws, "Target PaaS ")

	var out bytes.Buffer

	cfTargetCommand := utils.Command{Name: "cf", Args: []string{"target", "-o", onPaaS.Org, "-s", onPaaS.Space}, Dir: deployDir, Stdin: "", Env: []string{"CF_COLOR=false"}}

	cmdRunner := utils.NewDeployCmdRunner()

	cmdRunner.RunCommandAsyncCmd(cfTargetCommand, &out)

	writeBytesBufferMessage(&out, &cmdRunner, ws)

	writeStringMessage(ws, "Target PaaS successful")
	writeStringMessage(ws, "============================================")

	return cmdRunner.Success()
}

func (this *ServiceDeployWebSocketController) pushApp(ws *websocket.Conn, onPaaS entity.OnPaaS, deployDir string) bool {
	writeStringMessage(ws, "============================================")
	writeStringMessage(ws, "Deploying app ")

	var out bytes.Buffer

	cfDeployCommand := utils.Command{Name: "cf", Args: []string{"push", "-f", "manifest.yml"}, Dir: deployDir, Stdin: "", Env: []string{"CF_COLOR=false"}}

	cmdRunner := utils.NewDeployCmdRunner()

	cmdRunner.RunCommandAsyncCmd(cfDeployCommand, &out)

	writeBytesBufferMessage(&out, &cmdRunner, ws)

	writeStringMessage(ws, "Deploy app successful")
	writeStringMessage(ws, "============================================")

	return cmdRunner.Success()
}

func (this *ServiceDeployWebSocketController) operateAppOnPaaS(ws *websocket.Conn, onPaaS entity.OnPaaS, deployDir string, operate string) bool {
	writeStringMessage(ws, "============================================")
	writeStringMessage(ws, "Operate app ")

	var out bytes.Buffer
	mapps, err := utils.ReadYamlFile(deployDir + "/manifest.yml")
	var apps []string
	if err != nil {
		writeStringMessage(ws, fmt.Sprintf("Operate app failed ,errors :%v", err))
		return false
	}

	if err == nil {
		apps, err = utils.GetValuesByKey(mapps, "name")
		if err != nil {
			writeStringMessage(ws, fmt.Sprintf("Operate app failed ,errors :%v", err))
			return false
		}
		if err == nil {
			for _, app := range apps {
				if app != "" {
					cfOperateCommand := utils.Command{Name: "cf", Args: []string{operate, app}, Dir: deployDir, Stdin: "", Env: []string{"CF_COLOR=false"}}
					cmdRunner := utils.NewDeployCmdRunner()
					cmdRunner.RunCommandAsyncCmd(cfOperateCommand, &out)
					writeBytesBufferMessage(&out, &cmdRunner, ws)
				}
			}
			writeStringMessage(ws, "Operate app successful")
		}
	}

	writeStringMessage(ws, "============================================")

	return true
}

func (serviceDeploy *ServiceDeployWebSocketController) OperateOnPaaS(ws *websocket.Conn, service entity.Service, operate string) {
	//加载serviceDto
	//	writeStringMessage(ws, "Load ServiceDto")
	serviceDto := entity.ServiceDto{}
	serviceDto.Service = service
	err := serviceDto.Load()

	if err != nil {
		writeStringMessage(ws, fmt.Sprintf("Error,%s", err))
		return
	}

	onPaaS := serviceDto.OnPaaS

	var deployDir = customServiceDir + "/" + service.Name

	success := serviceDeploy.setApi(ws, onPaaS, deployDir)

	if !success {
		writeStringMessage(ws, "Set endpoint failed")
	}

	if success {
		success = serviceDeploy.loginPaaS(ws, onPaaS, deployDir)
		if !success {
			writeStringMessage(ws, "Login PaaS failed")
		}
	}

	if success {
		success = serviceDeploy.targetPaaS(ws, onPaaS, deployDir)
		if !success {
			writeStringMessage(ws, "Target PaaS failed")
		}
	}

	if success {
		switch operate {
		case utils.Service_Restart:
			success = serviceDeploy.operateAppOnPaaS(ws, onPaaS, deployDir, "rs")
		case utils.Service_Start:
			success = serviceDeploy.operateAppOnPaaS(ws, onPaaS, deployDir, "start")
		case utils.Service_Stop:
			success = serviceDeploy.operateAppOnPaaS(ws, onPaaS, deployDir, "stop")
		}
		if !success {
			writeStringMessage(ws, fmt.Sprintf("%s app failed", operate))
		}
	}

	if success {
		switch operate {
		case utils.Service_Restart:
			writeStringMessage(ws, fmt.Sprintf("%s app successful", "Restart"))
		case utils.Service_Start:
			writeStringMessage(ws, fmt.Sprintf("%s app successful", "Start"))
		case utils.Service_Stop:
			writeStringMessage(ws, fmt.Sprintf("%s app successful", "Stop"))
		}
	}

}

// operate service on paas

// operate service on vms
func (serviceDeploy *ServiceDeployWebSocketController) OperateOnVms(ws *websocket.Conn, service entity.Service, operate string) {
	//加载serviceDto
	//	writeStringMessage(ws, "Load ServiceDto")
	serviceDto := entity.ServiceDto{}
	serviceDto.Service = service
	err := serviceDto.Load()

	if err != nil {
		writeStringMessage(ws, fmt.Sprintf("Error,%s", err))
		return
	}
	writeStringMessage(ws, "Load custom vms info")
	onCustom := serviceDto.OnCustom
	sshRunner := utils.NewDeploySSHRunner(onCustom.Ip, onCustom.User, onCustom.Passwd)

	operation := entity.Operation{}
	operation.Sid = service.Id
	err = operation.LoadBySid()

	var out bytes.Buffer

	if err == nil {
		switch operate {
		case utils.Service_Restart:
			writeStringMessage(ws, fmt.Sprintf("Start run command： %s app, please wait for a monent", "Restart"))
			err = sshRunner.RunCommand(operation.Restart, &out)
			if err == nil {
				writeStringMessage(ws, fmt.Sprintf("%s app successful", "Restart"))
			}
		case utils.Service_Start:
			writeStringMessage(ws, fmt.Sprintf("Start run command： %s app, please wait for a monent", "Start"))
			err = sshRunner.RunCommand(operation.Start, &out)
			if err == nil {
				writeStringMessage(ws, fmt.Sprintf("%s app successful", "Start"))
			}
		case utils.Service_Stop:
			writeStringMessage(ws, fmt.Sprintf("Start run command： %s app, please wait for a monent", "Stop"))
			err = sshRunner.RunCommand(operation.Stop, &out)
			if err == nil {
				writeStringMessage(ws, fmt.Sprintf("%s app successful", "Stop"))
			}
		}
		if err != nil {
			writeStringMessage(ws, fmt.Sprintf("%s app failed ,errors: %v", operate, err))
		}
	}

}

// operate service on vms

// parse template 2 file
func (servicedDeploy *ServiceDeployWebSocketController) ParseTemplate(ws *websocket.Conn, service entity.Service) bool {

	writeStringMessage(ws, "Load templateList")

	template, templateLoadErr := entity.LoadTemplateList(service.Id)

	if templateLoadErr != nil {
		writeStringMessage(ws, fmt.Sprintf("Error,%s", templateLoadErr))
		return false
	}
	writeStringMessage(ws, "Load componentList")
	component, componentErr := entity.LoadComponentList(service.Id)

	if componentErr != nil {
		writeStringMessage(ws, fmt.Sprintf("Error,%s", componentErr))
		return false
	}
	writeStringMessage(ws, "Parse template")

	data := make(map[string]string)

	if len(component) > 0 {
		for _, item := range component {
			data[item.Name] = item.Value
		}
	}

	for _, templatefile := range template {
		if templatefile.FileType == utils.FileTypes_Template {
			if templatefile.TargetFile == "" {
				writeStringMessage(ws, fmt.Sprintf("%s TargetFile is null,please set the value", templatefile.Name))
			}
			if templatefile.TemplateFile == "" {
				writeStringMessage(ws, fmt.Sprintf("%s TemplateFile is null,please set the value", templatefile.Name))
			}
			result, err := utils.ParseTemplateFile2File(templatefile.TargetFile, data, templatefile.TemplateFile)
			if !result {
				writeStringMessage(ws, fmt.Sprintf("Error,%s", err))
				return result
			}
		}
	}
	writeStringMessage(ws, "Parse template successful")
	return true
}

// parse template 2 file
