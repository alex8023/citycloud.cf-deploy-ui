package utils

import (
	"errors"
	"fmt"
	"github.com/nats-io/nats"
	gouuid "github.com/nu7hatch/gouuid"
	"time"
)

func GenerateUUID() (uuidStr string) {
	uuid, err := gouuid.NewV4()
	if err != nil {
		return
	}
	return uuid.String()
}

func GetMonitor(natsServerIp, agentId string, timeout time.Duration) (string, error) {

	natsUrl := fmt.Sprintf("nats://nats:nats@%s:4222", natsServerIp)

	// if err != nil nc return nil
	nc, err := nats.Connect(natsUrl)

	if err != nil {
		return "", err
	}

	defer nc.Close()

	msg := make(chan *nats.Msg, 1)

	subject := fmt.Sprintf("director.%s.%s", GenerateUUID(), GenerateUUID())
	sub, suberr := nc.Subscribe(subject, func(m *nats.Msg) {
		msg <- m
	})

	if suberr != nil {
		return "", fmt.Errorf("Subccribe %s error", subject)
	}
	defer sub.Unsubscribe()

	message := fmt.Sprintf(`{"method":"get_state","arguments":["full"],"reply_to":"%s"}`, subject)

	agentSubject := fmt.Sprintf("agent.%s", agentId)
	nc.Publish(agentSubject, []byte(message))

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case mmsg, ok := <-msg:
		if !ok {
			return "", errors.New("Received message error !")
		}
		return string(mmsg.Data), nil
	case <-timer.C:
		return "", errors.New("Timeout!")
	}
}
