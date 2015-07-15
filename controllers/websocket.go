package controllers

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/citycloud/citycloud.cf-deploy-ui/utils"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
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
	stdout, stderr := this.initMicroBOSH()
	ws.WriteMessage(websocket.TextMessage, []byte(stdout))
	ws.WriteMessage(websocket.TextMessage, []byte(stderr))
	ws.WriteMessage(websocket.TextMessage, []byte("Finished Set MicroBosh Deployment "))
	ws.WriteMessage(websocket.TextMessage, []byte("============================================"))
}

func (this *WebSocketController) initMicroBOSH() (string, string) {
	filePath := "/home/ubuntu/bosh-workspace/deploy/microbosh/micro_bosh.yml"
	cmdRunner := utils.NewExecCmdRunner()
	stdout, stderr, status, err := cmdRunner.RunCommand("bosh", "micro", "deployment", filePath)
	if status == -1 {
		return "bash err !", fmt.Sprintf("ErrorMessage : %s", err)
	}
	return stdout, stderr
}

func (this *WebSocketController) deployDeployment(ws *websocket.Conn) {
	cmdRunner := utils.NewExecCmdRunner()

	cmd := utils.Command{
		Name:   "bosh",
		Args:   []string{"micro", "deployment", "/home/ubuntu/bosh-workspace/stemcells/bosh-stemcell-2719-openstack-kvm-ubuntu-lucid-go_agent.tgz"},
		Stdin:  strings.NewReader("yes"),
		Stdout: &out,
		Stderr: &out,
	}
	process, errors := cmdRunner.RunComplexCommandAsync(cmd)

	go process.Wait()

	ws.WriteMessage(websocket.TextMessage, []byte("Command is Running: pid"+process.Pid()))
	var count int = 0
	for {
		time.Sleep(5 * time.Second)
		result, err := out.ReadBytes('\n')
		if err != nil {
			time.Sleep(5 * time.Second)
			count = count + 1
		}
		ws.WriteMessage(websocket.TextMessage, []byte(result))
	}
}
