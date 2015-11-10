package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
)

type ServiceDeployWebScoketController struct {
	beego.Controller
}

func (service *ServiceDeployWebScoketController) Get() {
	ws, err := upgrader.Upgrade(service.Ctx.ResponseWriter, service.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(service.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		return
	}
	ws.ReadMessage()
}
