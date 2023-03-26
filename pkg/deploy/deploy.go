package deploy

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os"
	"strings"
)

const (
	sshUser = "corey"     // 远程SSH用户
	sshPass = "corey1996" // 远程SSH用户密码
	sshHost = "127.0.0.1" // 远程SSH主机
	sshPort = 2230        // 远程SSH端口
	sudoPsw = "corey1996" // 远程SSH用户sudo的密码
	script  = `#!/bin/bash           # 要执行的shell脚本
cd ~
mkdir $1
echo "Hello, $2!" > $1/hello.txt
ls -la`
)

func connect() (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPass),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", sshHost, sshPort), config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func runScript(client *ssh.Client, script string, args []string) error {
	// 创建一个新的session
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	sessionStdin, err := session.StdinPipe()
	if err != nil {
		return err
	}
	defer sessionStdin.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	// 把脚本写入session的标准输入中
	// 如果存在参数，则把参数添加到脚本中
	cmd := fmt.Sprintf("echo '%s' | sudo -S /bin/bash\n", sudoPsw)
	// echo 'corey1996' | sudo -S /bin/bash\n
	if len(args) > 0 {
		script += " " + strings.Join(args, " ")
	}
	err = session.Start(cmd)
	if err != nil {
		fmt.Println("err")
		return err
	}

	// 把脚本写入到session的标准输入中
	_, err = sessionStdin.Write([]byte("ls"))
	if err != nil {
		return err
	}

	err = session.Wait()
	if err != nil {
		return err
	}

	return nil
}

func deployMain() {
	// 连接到远程SSH主机
	client, err := connect()
	if err != nil {
		log.Fatalf("Failed to connect to remote host: %v", err)
	}
	defer client.Close()

	// 准备要执行的shell脚本
	scriptBytes := []byte(script)

	// 准备参数
	args := []string{"test", "World"}

	// 开始执行脚本
	err = runScript(client, string(scriptBytes), args)
	if err != nil {
		log.Fatalf("Failed to execute script: %v", err)
	}

	// 执行结束，退出
	fmt.Println("Script executed successfully.")
}
