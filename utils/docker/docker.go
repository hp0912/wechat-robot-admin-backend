package docker

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// networkExists 检查 Docker 网络是否存在
func NetworkExists(networkName string) bool {
	cmd := exec.Command("docker", "network", "ls", "--filter", fmt.Sprintf("name=%s", networkName), "--format", "{{.Name}}")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error checking network:", err)
		return false
	}

	// 检查输出中是否包含网络名称
	return strings.Contains(out.String(), networkName)
}

// createNetwork 创建 Docker 网络
func CreateNetwork(networkName string) error {
	cmd := exec.Command("docker", "network", "create", networkName)
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Printf("Network '%s' created successfully.\n", networkName)
	return nil
}
