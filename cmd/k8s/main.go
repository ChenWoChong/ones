package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "strings"

    "github.com/BurntSushi/toml"
    "golang.org/x/crypto/ssh"
)

type Config struct {
    Machines []Machine `toml:"machine"`
}

type Machine struct {
    IP       string `toml:"ip"`
    Username string `toml:"username"`
    Password string `toml:"password"`
}

func main() {
    // 读取配置文件
    var config Config
    _, err := toml.DecodeFile("config.toml", &config)
    if err != nil {
        log.Fatalf("unable to decode config file: %s", err)
    }

    // 部署k8s集群
    for _, machine := range config.Machines {
        // SSH连接配置
        config := &ssh.ClientConfig{
            User: machine.Username,
            Auth: []ssh.AuthMethod{
                ssh.Password(machine.Password),
            },
            HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        }

        // 连接SSH服务器
        conn, err := ssh.Dial("tcp", machine.IP+":22", config)
        if err != nil {
            log.Fatalf("unable to connect: %s", err)
        }
        defer conn.Close()

        // 执行命令
        session, err := conn.NewSession()
        if err != nil {
            log.Fatalf("unable to create session: %s", err)
        }
        defer session.Close()

        // 执行sudo命令
        cmd := "sudo kubeadm init"
        output, err := session.CombinedOutput(cmd)
        if err != nil {
            log.Fatalf("unable to execute command: %s", err)
        }

        fmt.Println(string(output))
    }

    // 卸载k8s集群
    for _, machine := range config.Machines {
        // SSH连接配置
        config := &ssh.ClientConfig{
