package dockerApi

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os"
	"os/exec"
)

// StartContainer 启动容器
func StartContainer(containerID string) error {
	ctx := context.Background()
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		return err
	} else {
		return nil
	}
}

// StopContainer 停止容器
func StopContainer(containerID string) {
	ctx := context.Background()
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	err := cli.ContainerStop(ctx, containerID, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("stop成功：", containerID)
	}
}

// ListImages 获取docker镜像列表
func ListImages() []string {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	var imagesLib []string
	for _, image := range images {
		imagesLib = append(imagesLib, image.ID)
	}
	return imagesLib
}

func Test() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
	}
}

// ListContainer 容器列表
func ListContainer() []string {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	var Container []string
	for _, container := range containers {
		Container = append(Container, container.ID)
	}
	return Container
}

// CreatContainer 创建容器
func CreatContainer(alias string, images string, port1 string, port2 string, SqlPort string) string {
	cmd := exec.Command("docker", "run", "-d", "--name", alias, "-p", port1, "-p", port2, "-p", SqlPort, images)
	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))
	return string(output)
}

// ImportContainer 导入容器
func ImportContainer(ImagesUrlName string, ContainerNameVersion string) bool {
	cmd := exec.Command("docker", "import", ImagesUrlName, ContainerNameVersion)
	output, _ := cmd.CombinedOutput()
	if output == nil {
		return true
	}
	return false
}
