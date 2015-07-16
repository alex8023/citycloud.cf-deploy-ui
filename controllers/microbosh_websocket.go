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
var out bytes.Buffer

var running bool = false

func (this *MicroBOSHWebSocketController) Get() {

	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		return
	}
	fmt.Println("no err")
	fmt.Println("he")
	for {
		mesgType, message, _ := ws.ReadMessage()
		if message != nil && mesgType == websocket.TextMessage {
			fmt.Println(string(message))
			fmt.Println("come here")
		}
	}
	fmt.Println("has go")
	if this.GetString("action") == "MicroBOSH" {
		var success bool = false
		success = this.setMicroBOSHDeployment(ws)
		if !success {
			writeStringMessage(ws, "设置deployment出现了错误，需要检查MicroBOSH Deployment的配置")
			ws.WriteMessage(websocket.CloseMessage, []byte("Closed"))
			return
		}
		success = this.deployMicroBOSH(ws)
		if !success {
			writeStringMessage(ws, "部署 MicroBOSH 实例出现了错误！需要检查MicroBOSH Deployment的配置")
			ws.WriteMessage(websocket.CloseMessage, []byte("Closed"))
			return
		}
		success = this.targetMicroBOSH(ws)
		if !success {
			writeStringMessage(ws, "设置bosh target失败！")
			ws.WriteMessage(websocket.CloseMessage, []byte("Closed"))
			return
		}
		success = this.loginMicroBOSH(ws)
		if !success {
			writeStringMessage(ws, "登录失败！")
			ws.WriteMessage(websocket.CloseMessage, []byte("Closed"))
			return
		}
		ws.WriteMessage(websocket.CloseMessage, []byte("Closed"))
		return

	}
	if this.GetString("action") == "SetDeployment" {
		var success bool = false
		success = this.setMicroBOSHDeployment(ws)
		if !success {
			writeStringMessage(ws, "设置deployment出现了错误，需要检查MicroBOSH Deployment的配置")
			ws.WriteMessage(websocket.CloseMessage, []byte("Closed"))
			return
		}
		ws.WriteMessage(websocket.CloseMessage, []byte("Closed"))
		return
	}
	if this.GetString("action") == "Login" {
		var success bool = false
		success = this.targetMicroBOSH(ws)
		if !success {
			writeStringMessage(ws, "设置bosh target失败！")
			ws.WriteMessage(websocket.CloseMessage, []byte("Closed"))
			return
		}
		success = this.loginMicroBOSH(ws)
		if !success {
			writeStringMessage(ws, "登录失败！")
			ws.WriteMessage(websocket.CloseMessage, []byte("Closed"))
			return
		}
		ws.WriteMessage(websocket.CloseMessage, []byte("Closed"))
		return
	}

}

//set deployment
func (this *MicroBOSHWebSocketController) setMicroBOSHDeployment(ws *websocket.Conn) bool {
	writeStringMessage(ws, "============================================")
	writeStringMessage(ws, "Set MicroBosh Deployment")

	var out bytes.Buffer
	deployment := "microbosh/micro_bosh.yml"
	dir := "/home/ubuntu/bosh-workspace/deploy"
	cmdCommand := utils.Command{Name: "bosh", Args: []string{"micro", "deployment", deployment}, Dir: dir, Stdin: ""}

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
	dir := "/home/ubuntu/bosh-workspace/deploy"
	stemcells := "/home/ubuntu/bosh-workspace/stemcells/bosh-stemcell-2719-openstack-kvm-ubuntu-lucid-go_agent.tgz"
	cmdCommand := utils.Command{Name: "bosh", Args: []string{"micro", "deploy", stemcells}, Dir: dir, Stdin: "yes"}
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
	dir := "/home/ubuntu/bosh-workspace/deploy"
	var ip string = "192.168.133.108"
	target := fmt.Sprintf("https://%s:25555", ip)

	loginCommand := utils.Command{Name: "bosh", Args: []string{"target", target}, Dir: dir, Stdin: "admin\nadmin\n"}

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
	dir := "/home/ubuntu/bosh-workspace/deploy"

	loginCommand := utils.Command{Name: "bosh", Args: []string{"login"}, Dir: dir, Stdin: "admin\nadmin\n"}

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
