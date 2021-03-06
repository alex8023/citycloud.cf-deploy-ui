package controllers

import (
	"bytes"
	"fmt"
	_ "github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type CloudFoundryWebSocketController struct {
	BaseController
}

// becareful
//此处未判断任务是否完成，需前端判断控制任务流程
func (this *CloudFoundryWebSocketController) Get() {

	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		return
	}

	for {
		time.Sleep(2 * time.Second)
		mesgType, message, _ := ws.ReadMessage()
		if message != nil && mesgType == websocket.TextMessage {

			var action = string(message)
			switch {
			case action == "AllStep":
				var success bool = false

				success = this.uploadRelease(ws)
				if !success {
					writeStringMessage(ws, "上传Release失败！检查是否登录BOSH")
				}

				if success {
					success = this.uploadStemcell(ws)
					if !success {
						writeStringMessage(ws, "上传Stemcell失败！检查是否登录BOSH")
					}
				}

				if success {
					success = this.setCloudFoundryDeployment(ws)
					if !success {
						writeStringMessage(ws, "设置deployment出现了错误!")
					}
				}

				if success {
					success = this.deployCloudFoundry(ws)
					if !success {
						writeStringMessage(ws, "部署PaaS实例出现了错误！")
					}
				}
			case action == "UpRelease":
				var success bool = false
				success = this.uploadRelease(ws)
				if !success {
					writeStringMessage(ws, "上传Release失败！检查是否登录BOSH")
				}
			case action == "UpStemcell":
				var success bool = false
				success = this.uploadStemcell(ws)
				if !success {
					writeStringMessage(ws, "上传Stemcell失败！检查是否登录BOSH")
				}
			case action == "SetDeploy":
				var success bool = false
				success = this.setCloudFoundryDeployment(ws)
				if !success {
					writeStringMessage(ws, "设置deployment出现了错误!")
				}
			case action == "Deploy":
				var success bool = false
				success = this.deployCloudFoundry(ws)
				if !success {
					writeStringMessage(ws, "部署PaaS实例出现了错误！")
				}
			case action == "Stats":
				var success bool = false
				success = this.statsCloudFoundry(ws)
				if !success {
					writeStringMessage(ws, "获取PaaS实例信息出现了错误！")
				}
			default:
				writeStringMessage(ws, fmt.Sprintf("未知的执行命令！%s", action))
			}
			//write a message to tell me running over
			writeStringMessage(ws, EndEOF)
		}
	}
}

//set deployment
func (this *CloudFoundryWebSocketController) setCloudFoundryDeployment(ws *websocket.Conn) bool {
	writeStringMessage(ws, "============================================")
	writeStringMessage(ws, "Set PaaS Deployment")

	var out bytes.Buffer
	cmdCommand := utils.Command{Name: "bosh", Args: []string{"deployment", cloudFoundryManiest}, Dir: workDir, Stdin: ""}

	cmdRunner := utils.NewDeployCmdRunner()
	cmdRunner.RunCommandAsyncCmd(cmdCommand, &out)
	writeBytesBufferMessage(&out, &cmdRunner, ws)

	writeStringMessage(ws, "Finished Set PaaS Deployment")
	writeStringMessage(ws, "============================================")

	return cmdRunner.Success()
}

//deploy CloudFoundry
func (this *CloudFoundryWebSocketController) deployCloudFoundry(ws *websocket.Conn) bool {
	writeStringMessage(ws, "============================================")
	writeStringMessage(ws, "Deploying PaaS Instances")
	var out bytes.Buffer
	cmdCommand := utils.Command{Name: "bosh", Args: []string{"deploy"}, Dir: workDir, Stdin: "yes"}
	cmdRunner := utils.NewDeployCmdRunner()
	cmdRunner.RunCommandAsyncCmd(cmdCommand, &out)
	writeBytesBufferMessage(&out, &cmdRunner, ws)

	writeStringMessage(ws, "Finished PaaS CloudFoundry Instances")
	writeStringMessage(ws, "============================================")
	return cmdRunner.Success()
}

//bosh vms
func (this *CloudFoundryWebSocketController) statsCloudFoundry(ws *websocket.Conn) bool {
	writeStringMessage(ws, "============================================")
	writeStringMessage(ws, "Status PaaS Instances")
	var out bytes.Buffer
	cmdCommand := utils.Command{Name: "bosh", Args: []string{"vms", cf.CloudFoundryProperties.Name, "--details"}, Dir: workDir, Stdin: ""}
	cmdRunner := utils.NewDeployCmdRunner()
	cmdRunner.RunCommandAsyncCmd(cmdCommand, &out)
	writeBytesBufferMessage(&out, &cmdRunner, ws)

	writeStringMessage(ws, "Finished Status PaaS CloudFoundry Instances")
	writeStringMessage(ws, "============================================")
	return cmdRunner.Success()
}

//upload CloudFoundryv2 release
func (this *CloudFoundryWebSocketController) uploadRelease(ws *websocket.Conn) bool {
	writeStringMessage(ws, "============================================")
	writeStringMessage(ws, "Uploading PaaS v2 release")
	var out bytes.Buffer
	cmdCommand := utils.Command{Name: "bosh", Args: []string{"upload", "release", cloudFoundryRelease, "--skip-if-exists"}, Dir: workDir, Stdin: ""}
	cmdRunner := utils.NewDeployCmdRunner()
	cmdRunner.RunCommandAsyncCmd(cmdCommand, &out)
	writeBytesBufferMessage(&out, &cmdRunner, ws)

	writeStringMessage(ws, fmt.Sprintf("Finished uploading PaaS v2 release %s ", cloudFoundryRelease))
	writeStringMessage(ws, "============================================")
	return cmdRunner.Success()
}

//upload stemcell
func (this *CloudFoundryWebSocketController) uploadStemcell(ws *websocket.Conn) bool {
	writeStringMessage(ws, "============================================")
	writeStringMessage(ws, "Uploading stemcells  ")
	var out bytes.Buffer

	cmdCommand := utils.Command{Name: "bosh", Args: []string{"upload", "stemcell", stemcellRelease, "--skip-if-exists"}, Dir: workDir, Stdin: ""}
	cmdRunner := utils.NewDeployCmdRunner()
	cmdRunner.RunCommandAsyncCmd(cmdCommand, &out)

	writeBytesBufferMessage(&out, &cmdRunner, ws)

	writeStringMessage(ws, fmt.Sprintf("Finished uploading stemcells %s ", stemcellRelease))
	writeStringMessage(ws, "============================================")
	return cmdRunner.Success()
}
