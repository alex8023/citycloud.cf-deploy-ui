package controllers

import (
	"bytes"
	"github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type WebSocketController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var out bytes.Buffer

var running bool = false

func (this *WebSocketController) Get() {

	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		return
	}

	if this.GetString("action") == "MicroBOSH" {
		this.deployMicroBOSH(ws)
		this.deployDeployment(ws)
	}
	ws.WriteMessage(websocket.TextMessage, []byte("Wait for closed"))
	ws.WriteMessage(websocket.CloseMessage, []byte("Closed"))
}

func (this *WebSocketController) deployMicroBOSH(ws *websocket.Conn) {
	if running {
		ws.WriteMessage(websocket.TextMessage, []byte(""))
	}

	ws.WriteMessage(websocket.TextMessage, []byte("============================================"))
	ws.WriteMessage(websocket.TextMessage, []byte("Set MicroBosh Deployment "))

	filePath := "/home/ubuntu/bosh-workspace/deploy/microbosh/micro_bosh.yml"
	cmdRunner := utils.NewDeployCmdRunner()
	var out bytes.Buffer

	go cmdRunner.RunCommand("bosh", "", &out, "micro", "deployment", filePath)

	for {
		time.Sleep(time.Second)
	reread:
		message, err := out.ReadBytes('\n')
		if err != nil && cmdRunner.Success() {
			break
		} else {
			if err != nil {
				goto reread
			}
			ws.WriteMessage(websocket.TextMessage, message)
		}
	}

	ws.WriteMessage(websocket.TextMessage, []byte("Finished Set MicroBosh Deployment "))
	ws.WriteMessage(websocket.TextMessage, []byte("============================================"))
}

func (this *WebSocketController) deployDeployment(ws *websocket.Conn) {

	var out bytes.Buffer

	cmdRunner := utils.NewDeployCmdRunner()

	stemcell := "/home/ubuntu/bosh-workspace/stemcells/bosh-stemcell-2719-openstack-kvm-ubuntu-lucid-go_agent.tgz"
	go cmdRunner.RunCommand("bosh", "yes", &out, "micro", "deploy", stemcell)
	for {
		time.Sleep(time.Second)
	reread:
		message, err := out.ReadBytes('\n')
		if err != nil && cmdRunner.Success() {
			break
		} else {
			if err != nil {
				goto reread
			}
			ws.WriteMessage(websocket.TextMessage, message)
		}
	}
}
