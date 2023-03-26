package deploy

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"

// 	"github.com/BurntSushi/toml"
// 	"golang.org/x/crypto/ssh"
// )

// // Server 是配置文件中的服务器信息
// type Server struct {
// 	IP       string `toml:"ip"`
// 	Username string `toml:"username"`
// 	Password string `toml:"password"`
// }

// // Config 是配置文件
// type Config struct {
// 	Servers []Server `toml:"servers"`
// }

// func conect() {

// 	// 从命令行参数获取配置文件路径
// 	if len(os.Args) < 2 {
// 		fmt.Printf("Usage: %s <config file>\n", os.Args[0])
// 		os.Exit(1)
// 	}
// 	configPath := os.Args[1]

// 	// 读取配置文件
// 	configBytes, err := ioutil.ReadFile(configPath)
// 	if err != nil {
// 		log.Fatalf("Failed to read config file: %v", err)
// 	}

// 	// 解析配置文件
// 	var config Config
// 	if err := toml.Unmarshal(configBytes, &config); err != nil {
// 		log.Fatalf("Failed to parse config file: %v", err)
// 	}

// 	// 连接每个服务器并执行命令
// 	for _, server := range config.Servers {
// 		log.Printf("Connecting to %s...\n", server.IP)

// 		// 建立 SSH 连接
// 		client, err := ssh.Dial("tcp", server.IP+":22", &ssh.ClientConfig{
// 			User: server.Username,
// 			Auth: []ssh.AuthMethod{ssh.Password(server.Password)},
// 		})
// 		if err != nil {
// 			log.Printf("Failed to connect to %s: %v\n", server.IP, err)
// 			continue
// 		}
// 		defer client.Close()

// 		// 部署 Kubernetes 集群
// 		if err := deployKubernetes(client); err != nil {
// 			log.Printf("Failed to deploy Kubernetes on %s: %v\n", server.IP, err)
// 			continue
// 		}

// 		log.Printf("Kubernetes deployed successfully on %s\n", server.IP)
// 	}
// }

// // 使用 exec.Command 执行命令，并输出执行的命令
// func runCommand(client *ssh.Client, command string) error {
// 	fmt.Printf("Executing command: %s\n", command)

// 	session, err := client.NewSession()
// 	if err != nil {
// 		return err
// 	}
// 	defer session.Close()

// 	session.Stdout = os.Stdout
// 	session.Stderr = os.Stderr

// 	return session.Run(command)
// }

// // 部署 Kubernetes
// func deployKubernetes(client *ssh.Client) error {
// 	// 安装 Docker
// 	if err := runCommand(client, "sudo apt-get update"); err != nil {
// 		return err
// 	}
// 	if err := runCommand(client, "sudo apt-get install -y docker.io"); err != nil {
// 		return err
// 	}

// 	// 安装 Kubernetes
// 	if err := runCommand(client, "sudo apt-get update && sudo apt-get install -y apt-transport-https"); err != nil {
// 		return err
// 	}
// }
