package utils

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	SCP_SEND_FILE_BEGIN   = "C"    //传输文件前缀
	SCP_SEND_FOLDER_BEGIN = "D"    //传输目录前缀
	SCP_SEND_FOLDER_SIZE  = 0      //传输目录大小
	SCP_SEND_FOLDER_END   = "E"    //传输目录结束
	SCP_SEND_FILE_END     = "\x00" //传输文件结束
)

type DeploySSHRunner struct {
	Ip       string
	User     string
	Password string
}

func NewDeploySSHRunner(ip string, userName string, password string) DeploySSHRunner {
	return DeploySSHRunner{
		Ip:       ip,
		User:     userName,
		Password: password,
	}
}

func (this *DeploySSHRunner) clientRemote() (*ssh.Client, error) {
	conf := ssh.ClientConfig{User: this.User, Auth: []ssh.AuthMethod{ssh.Password(this.Password)}}
	return ssh.Dial("tcp", fmt.Sprintf("%s:22", this.Ip), &conf)
}

func (this *DeploySSHRunner) Push(localPath, remotePath string, out *bytes.Buffer) error {
	client, err := this.clientRemote()
	if err != nil {
		return err
	}
	defer client.Close()

	session, sessionerr := client.NewSession()

	if sessionerr != nil {
		return sessionerr
	}
	defer session.Close()

	session.Stdout = out
	session.Stderr = out

	info, infoerr := os.Lstat(localPath)
	if infoerr != nil {
		return infoerr
	}
	go func() {
		writeCloser, _ := session.StdinPipe()
		defer writeCloser.Close()
		if info.IsDir() {
			prepareDir(writeCloser, localPath)
		} else {
			prepareFile(writeCloser, localPath)
		}
	}()

	runerr := session.Run("/usr/bin/scp -qrt " + remotePath)

	if runerr != nil && fmt.Sprintf("%s", runerr) != "Process exited with: 1. Reason was:  ()" {
		return runerr
	}

	return nil
}

func (this *DeploySSHRunner) RunCommand(cmd string, out *bytes.Buffer) error {
	client, err := this.clientRemote()
	if err != nil {
		return err
	}
	defer client.Close()

	session, sessionerr := client.NewSession()

	if sessionerr != nil {
		return sessionerr
	}
	defer session.Close()

	session.Stdout = out
	session.Stderr = out

	runerr := session.Run(cmd)
	if runerr != nil && fmt.Sprintf("%s", runerr) != "Process exited with: 1. Reason was:  ()" {
		return runerr
	}
	return nil
}

//获取4位文件权限
//e.g 0644
func GetPerm(f *os.File) (perm string) {
	fileStat, _ := f.Stat()

	mod := fileStat.Mode().Perm()
	return fmt.Sprintf("%04o", uint32(mod))
}

func readDirAndWrite(w io.WriteCloser, dir string) error {
	fi, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	//parcours des dossiers
	for _, f := range fi {
		filename := filepath.Join(dir, f.Name())
		if f.IsDir() {
			folderSrc, err := os.Open(filename)
			if err != nil {
				return err
			}
			defer folderSrc.Close()
			mode := SCP_SEND_FOLDER_BEGIN + GetPerm(folderSrc)
			fmt.Fprintln(w, mode, SCP_SEND_FOLDER_SIZE, f.Name())
			readDirAndWrite(w, filename)
			fmt.Fprintln(w, SCP_SEND_FOLDER_END)
		} else {
			if err := prepareFile(w, filename); err != nil {
				return err
			}
		}
	}
	return nil
}

func prepareFile(w io.WriteCloser, src string) error {
	fileSrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fileSrc.Close()
	srcStat, err := fileSrc.Stat()
	if err != nil {
		return err
	}

	mode := SCP_SEND_FILE_BEGIN + GetPerm(fileSrc)
	fmt.Fprintln(w, mode, srcStat.Size(), filepath.Base(src))
	io.Copy(w, fileSrc)
	fmt.Fprint(w, SCP_SEND_FILE_END)
	return nil
}

func prepareDir(w io.WriteCloser, src string) error {
	fileSrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fileSrc.Close()
	srcStat, err := fileSrc.Stat()
	if err != nil {
		return err
	}
	mod := srcStat.Mode().Perm()
	mode := SCP_SEND_FOLDER_BEGIN + fmt.Sprintf("%04o", uint32(mod))
	fmt.Fprintln(w, mode, SCP_SEND_FOLDER_SIZE, filepath.Base(src))
	readDirAndWrite(w, src)
	fmt.Fprint(w, SCP_SEND_FOLDER_END)
	return nil
}
