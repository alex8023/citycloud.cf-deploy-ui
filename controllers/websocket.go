package controllers

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"os/exec"
	"time"
)

type WebSocketController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (this *WebSocketController) Get() {

	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		return
	}
	cmd := exec.Command("ping", "127.0.0.1", "-c 20")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Start()

	for {
		//		time.Sleep(5 * time.Second)
		abc, err := out.ReadBytes('\n')
		if err == nil {
			fmt.Print(string(abc)) //读取输出的日志

			ws.WriteMessage(websocket.TextMessage, abc)
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	cmd.Wait()
	time.Sleep(time.Second)
	fmt.Println("end")
	ws.WriteMessage(websocket.TextMessage, []byte("Close"))
}
