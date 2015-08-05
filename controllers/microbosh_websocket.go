package controllers

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type MicroBOSHWebSocketController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// becareful
//此处未判断任务是否完成，需前端判断控制任务流程
func (this *MicroBOSHWebSocketController) Get() {

	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		return
	}

	for {
		mesgType, message, _ := ws.ReadMessage()
		if message != nil && mesgType == websocket.TextMessage {

			var action = string(message)
			switch {
			case action == "AllStep":
				var success bool = false
				success = this.setMicroBOSHDeployment(ws)
				if !success {
					writeStringMessage(ws, "设置deployment出现了错误，需要检查MicroBOSH Deployment的配置")
				}
				if success {
					success = this.deployMicroBOSH(ws)
					if !success {
						writeStringMessage(ws, "部署 MicroBOSH 实例出现了错误！需要检查MicroBOSH Deployment的配置")
					}
				}
				if success {
					success = this.targetMicroBOSH(ws)
					if !success {
						writeStringMessage(ws, "设置bosh target失败！")
					}
				}

				if success {
					success = this.loginMicroBOSH(ws)
					if !success {
						writeStringMessage(ws, "登录失败！")
					}
				}
			case action == "SetDeploy":
				var success bool = false
				success = this.setMicroBOSHDeployment(ws)
				if !success {
					writeStringMessage(ws, "设置deployment出现了错误，需要检查MicroBOSH Deployment的配置")
				}
			case action == "Deploy":
				var success bool = false
				success = this.deployMicroBOSH(ws)
				if !success {
					writeStringMessage(ws, "部署 MicroBOSH 实例出现了错误！需要检查MicroBOSH Deployment的配置")
				}
			case action == "Login":
				var success bool = false
				success = this.targetMicroBOSH(ws)
				if !success {
					writeStringMessage(ws, "设置bosh target失败！")
				}
				if success {
					success = this.loginMicroBOSH(ws)
					if !success {
						writeStringMessage(ws, "登录失败！")
					}
				}
			default:
				writeStringMessage(ws, fmt.Sprintf("未知的执行命令！%s", action))
			}
		}
		//write a message to tell me running over
		writeStringMessage(ws, "da39a3ee5e6b4b0d3255bfef95601890afd80709")
	}
}

//set deployment
func (this *MicroBOSHWebSocketController) setMicroBOSHDeployment(ws *websocket.Conn) bool {
	writeStringMessage(ws, "============================================")
	writeStringMessage(ws, "Set MicroBosh Deployment")

	var out bytes.Buffer
	cmdCommand := utils.Command{Name: "bosh", Args: []string{"micro", "deployment", microManiest}, Dir: workDir, Stdin: ""}

	cmdRunner := utils.NewDeployCmdRunner()
	cmdRunner.RunCommandAsyncCmd(cmdCommand, &out)
	writeBytesBufferMessage(&out, &cmdRunner, ws)

	writeStringMessage(ws, "Finished Set MicroBosh Deployment")
	writeStringMessage(ws, "============================================")

	return cmdRunner.Success()
}

//deploy microbosh
func (this *MicroBOSHWebSocketController) deployMicroBOSH(ws *websocket.Conn) bool {
	writeStringMessage(ws, "============================================")
	writeStringMessage(ws, "Deploying MicroBosh instances")
	var out bytes.Buffer
	cmdCommand := utils.Command{Name: "bosh", Args: []string{"micro", "deploy", stemcellRelease}, Dir: workDir, Stdin: "yes"}
	cmdRunner := utils.NewDeployCmdRunner()
	cmdRunner.RunCommandAsyncCmd(cmdCommand, &out)
	writeBytesBufferMessage(&out, &cmdRunner, ws)

	writeStringMessage(ws, "Finished deploying MicroBosh instances")
	writeStringMessage(ws, "============================================")
	return cmdRunner.Success()
}

//target to microbosh
func (this *MicroBOSHWebSocketController) targetMicroBOSH(ws *websocket.Conn) bool {
	writeStringMessage(ws, "============================================")
	writeStringMessage(ws, "Target to MicroBosh instances")
	var out bytes.Buffer
	var ip string = mi.Network.Vip
	target := fmt.Sprintf("https://%s:25555", ip)

	loginCommand := utils.Command{Name: "bosh", Args: []string{"target", target}, Dir: workDir, Stdin: "admin\nadmin\n"}

	cmdRunner := utils.NewDeployCmdRunner()

	cmdRunner.RunCommandAsyncCmd(loginCommand, &out)

	writeBytesBufferMessage(&out, &cmdRunner, ws)

	writeStringMessage(ws, "Finished target to MicroBosh instances")
	writeStringMessage(ws, "============================================")
	return cmdRunner.Success()
}

//login into
func (this *MicroBOSHWebSocketController) loginMicroBOSH(ws *websocket.Conn) bool {
	writeStringMessage(ws, "============================================")
	writeStringMessage(ws, "Login MicroBosh instances")
	var out bytes.Buffer

	loginCommand := utils.Command{Name: "bosh", Args: []string{"login"}, Dir: workDir, Stdin: "admin\nadmin\n"}

	cmdRunner := utils.NewDeployCmdRunner()

	cmdRunner.RunCommandAsyncCmd(loginCommand, &out)

	writeBytesBufferMessage(&out, &cmdRunner, ws)

	writeStringMessage(ws, "Finished login MicroBosh instances")
	writeStringMessage(ws, "============================================")
	return cmdRunner.Success()
}

func writeStringMessage(ws *websocket.Conn, message string) {
	ws.WriteMessage(websocket.TextMessage, []byte(message))
}

func writeBytesMessage(ws *websocket.Conn, message []byte) {
	ws.WriteMessage(websocket.TextMessage, message)
}

func writeBytesBufferMessage(out *bytes.Buffer, cmdRunner *utils.DeployCmdRunner, ws *websocket.Conn) {
	for {
	loop:
		if out.Len() == 0 {
			if cmdRunner.Exited() {
				break
			} else {
				time.Sleep(time.Second)
				goto loop
			}
		}
	again:
		line, _ := out.ReadBytes('\n') //按行读取
		if line != nil {
			writeBytesMessage(ws, line)
			goto again
		} else {
			goto loop
		}
	}
}

//	if action == "MicroBOSH" {
//		var success bool = false
//		success = this.setMicroBOSHDeployment(ws)
//		if !success {
//			writeStringMessage(ws, "设置deployment出现了错误，需要检查MicroBOSH Deployment的配置")
//		}
//		if success {
//			success = this.deployMicroBOSH(ws)
//			if !success {
//				writeStringMessage(ws, "部署 MicroBOSH 实例出现了错误！需要检查MicroBOSH Deployment的配置")
//			}
//		}

//		if success {
//			success = this.targetMicroBOSH(ws)
//			if !success {
//				writeStringMessage(ws, "设置bosh target失败！")
//			}

//		}

//		if success {
//			success = this.loginMicroBOSH(ws)
//			if !success {
//				writeStringMessage(ws, "登录失败！")
//			}
//		}
//	}
//	if action == "SetDeployment" {
//		var success bool = false
//		success = this.setMicroBOSHDeployment(ws)
//		if !success {
//			writeStringMessage(ws, "设置deployment出现了错误，需要检查MicroBOSH Deployment的配置")
//		}
//	}

//	if action == "Deploy" {
//		var success bool = false
//		success = this.deployMicroBOSH(ws)
//		if !success {
//			writeStringMessage(ws, "部署 MicroBOSH 实例出现了错误！需要检查MicroBOSH Deployment的配置")
//			ws.WriteMessage(websocket.CloseMessage, []byte("Closed"))
//			return
//		}
//	}

//	if action == "Login" {
//		var success bool = false
//		success = this.targetMicroBOSH(ws)
//		if !success {
//			writeStringMessage(ws, "设置bosh target失败！")
//		}
//		if success {
//			success = this.loginMicroBOSH(ws)
//			if !success {
//				writeStringMessage(ws, "登录失败！")
//			}
//		}
//	}
