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

type Command struct {
	Name  string
	Args  []string
	Dir   string
	Stdin string
}

type DeployCmdRunner struct {
	ProcessState *os.ProcessState
	Process      *os.Process
	Stdout       io.Writer
	Stderr       io.Writer
}

func NewDeployCmdRunner() DeployCmdRunner {
	return DeployCmdRunner{}
}

func (this *DeployCmdRunner) RunCommand(cmdName, input string, out *bytes.Buffer, args ...string) {
	cmd := exec.Command(cmdName, args...)
	if input != "" {
		cmd.Stdin = strings.NewReader(input)
	}
	cmdString := strings.Join(cmd.Args, " ")
	cmd.Stdout = out
	cmd.Stderr = out
	cmd.Dir = "/home/ubuntu/bosh-workspace/deploy/"
	err := cmd.Start()
	this.Process = cmd.Process
	err = cmd.Wait()
	this.ProcessState = cmd.ProcessState
	if err != nil {
		errMessage := fmt.Sprintf("Starting command %s : err: %s \n", cmdString, err)
		out.WriteString(errMessage)
	}
}

// 执行cmd，将输出结果写入out
func (this *DeployCmdRunner) RunCommandCmd(command Command, out *bytes.Buffer) {
	cmd := exec.Command(command.Name, command.Args...)
	if command.Stdin != "" {
		cmd.Stdin = strings.NewReader(command.Stdin)
	}
	cmdString := strings.Join(cmd.Args, " ")
	cmd.Stdout = out
	cmd.Stderr = out
	cmd.Dir = command.Dir
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

func (this *DeployCmdRunner) Exited() bool {
	if this.ProcessState != nil {
		return this.ProcessState.Exited()
	}
	return false
}

func (this *DeployCmdRunner) String() string {
	if this.ProcessState != nil {
		return this.ProcessState.String()
	}
	return ""
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

func (this *DeployCmdRunner) RunCommandAsyncCmd(command Command, out *bytes.Buffer) {
	cmd := exec.Command(command.Name, command.Args...)
	if command.Stdin != "" {
		cmd.Stdin = strings.NewReader(command.Stdin)
	}
	cmdString := strings.Join(cmd.Args, " ")
	cmd.Stdout = out
	cmd.Stderr = out
	cmd.Dir = command.Dir
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
