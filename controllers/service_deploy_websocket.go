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
			case "restart":
				writeStringMessage(ws, "Unrelease method")
				writeStringMessage(ws, EndEOF)
			case "stop":
				writeStringMessage(ws, "Unrelease method")
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
		writeStringMessage(ws, fmt.Sprintf("Unknow target %s", service.Where))
	}
}

func (serviceDeploy *ServiceDeployWebSocketController) Deploy2Vms(ws *websocket.Conn, service entity.Service) {
	//加载serviceDto
	writeStringMessage(ws, "Load ServiceDto")
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
	writeStringMessage(ws, "Parse Template")
	//TODO

	//TODO
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

	writeStringMessage(ws, fmt.Sprintf("%d,传输%s,路径%s", index, template.Name, template.TemplateFile))

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

func (serviceDeploy *ServiceDeployWebSocketController) Deploy2PaaS(ws *websocket.Conn, service entity.Service) {
	//加载serviceDto
	writeStringMessage(ws, "Load ServiceDto")
	serviceDto := entity.ServiceDto{}
	serviceDto.Service = service
	err := serviceDto.Load()

	if err != nil {
		writeStringMessage(ws, fmt.Sprintf("Error,%s", err))
		return
	}
	//加载TemplateList
	//	writeStringMessage(ws, "Load TemplateList")
	//	template, templateLoadErr := entity.LoadTemplateList(service.Id)

	//	if templateLoadErr != nil {
	//		writeStringMessage(ws, fmt.Sprintf("Error,%s", templateLoadErr))
	//		return
	//	}

	//模板输出文件
	writeStringMessage(ws, "Parse Template")
	//TODO

	//TODO

	onPaaS := serviceDto.OnPaaS

	var deployDir = "/home/ubuntu/deploy/example"

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

func (this *ServiceDeployWebSocketController) setApi(ws *websocket.Conn, onPaaS entity.OnPaaS, deployDir string) bool {
	writeStringMessage(ws, "============================================")
	writeStringMessage(ws, "Set endpoint ")

	var out bytes.Buffer

	cfEndpointCommand := utils.Command{Name: "cf", Args: []string{"api", onPaaS.Api}, Dir: deployDir, Stdin: ""}

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

	cfLoginCommand := utils.Command{Name: "cf", Args: []string{"login", onPaaS.Api}, Dir: deployDir, Stdin: onPaaS.User + "\n" + onPaaS.Passwd + "\n"}

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

	cfTargetCommand := utils.Command{Name: "cf", Args: []string{"target", "-o", onPaaS.Org, "-s", onPaaS.Space}, Dir: deployDir, Stdin: ""}

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

	cfDeployCommand := utils.Command{Name: "cf", Args: []string{"push", "-f", "manifest.yml"}, Dir: deployDir, Stdin: ""}

	cmdRunner := utils.NewDeployCmdRunner()

	cmdRunner.RunCommandAsyncCmd(cfDeployCommand, &out)

	writeBytesBufferMessage(&out, &cmdRunner, ws)

	writeStringMessage(ws, "Deploy app successful")
	writeStringMessage(ws, "============================================")

	return cmdRunner.Success()
}
