package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

type DeployCmdRunner struct {
	ProcessState *os.ProcessState
	Process      *os.Process
	Stdout       io.Writer
	Stderr       io.Writer
}

func NewDeployCmdRunner() (deployCmdRunner DeployCmdRunner) {
	return
}

func (this *DeployCmdRunner) RunCommand(cmdName, input string, out *bytes.Buffer, args ...string) {
	cmd := exec.Command(cmdName, args...)
	if input != "" {
		cmd.Stdin = strings.NewReader(input)
	}
	cmdString := strings.Join(cmd.Args, " ")
	cmd.Stdout = out
	cmd.Stderr = out
	err := cmd.Start()
	this.Process = cmd.Process
	err = cmd.Wait()
	this.ProcessState = cmd.ProcessState
	if err != nil {
		errMessage := fmt.Sprintf("Starting command %s : err: %s \n", cmdString, err)
		out.WriteString(errMessage)
	}
}

func (this *DeployCmdRunner) Success() bool {
	if this.ProcessState != nil {
		return this.ProcessState.Success()
	}
	return false
}

func (this *DeployCmdRunner) RunCommandAsync(cmdName, input string, out *bytes.Buffer, args ...string) {
	cmd := exec.Command(cmdName, args...)

	if input != "" {
		cmd.Stdin = strings.NewReader(input)
	}
	cmdString := strings.Join(cmd.Args, " ")
	cmd.Stdout = out
	cmd.Stderr = out
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	err := cmd.Start()
	this.Process = cmd.Process
	if err != nil {
		errMessage := fmt.Sprintf("Starting command %s : err: %s \n", cmdString, err)
		out.WriteString(errMessage)
	}
	go func() {
		waiterr := cmd.Wait()
		if waiterr != nil {
			out.WriteString(fmt.Sprintf("Running command err：%s \n", waiterr))
		}
		this.ProcessState = cmd.ProcessState
	}()
}
