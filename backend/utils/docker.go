package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/LuizFelipeVPCrema/minecraft-dashboard-go/models"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func RunMinecraftContainer(containerName, port, project string, req models.ServerRequest) error {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("erro ao criar cliente Docker: %v", err)
	}

	mcPort, _ := nat.NewPort("tcp", "25565")

	containerConfig := &container.Config{
		Image: "itzg/minecraft-server",
		Env: []string{
			"EULA=TRUE",
			"VERSION=1.19",
			fmt.Sprintf("MODE=%s", strings.ToUpper(req.GameMode)),
			fmt.Sprintf("LEVEL_TYPE=%s", strings.ToUpper(req.LevelType)),
			"ONLINE_MODE=FALSE",
		},
		ExposedPorts: nat.PortSet{
			mcPort: struct{}{},
		},
		Labels: map[string]string{
			"app":     "minecraft-dashboard",
			"project": project,
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			mcPort: []nat.PortBinding{
				{
					HostPort: port,
				},
			},
		},
	}

	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, containerName)
	if err != nil {
		return fmt.Errorf("erro ao criar o contêiner: %v", err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("erro ao iniciar o contêiner: %v", err)
	}

	log.Printf("Contêiner %s iniciado com sucesso na porta %s", containerName, port)
	return nil
}

func ListMinecraftContainers() ([]types.Container, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}

	var mcContainers []types.Container
	for _, container := range containers {
		log.Printf("Container encontrado: %s | Labels: %v", container.Image, container.Labels)

		if container.Labels["app"] == "minecraft-dashboard" {
			mcContainers = append(mcContainers, container)
		}
	}

	return mcContainers, nil
}

func StopAndRemoveContainer(containerName string) error {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("erro ao criar cliente Docker: %v", err)
	}

	_, err = cli.ContainerInspect(ctx, containerName)
	if err != nil {
		if client.IsErrNotFound(err) {
			return fmt.Errorf("contêiner %s não encontrado", containerName)
		}
		return err
	}

	timeout := 10
	err = cli.ContainerStop(ctx, containerName, container.StopOptions{
		Timeout: &timeout,
	})
	if err != nil {
		log.Printf("Erro ao parar container (talvez já parado): %v", err)
	}

	err = cli.ContainerRemove(ctx, containerName, types.ContainerRemoveOptions{
		Force: true,
	})
	if err != nil {
		return fmt.Errorf("erro ao remover container: %v", err)
	}

	log.Printf("Contêiner %s removido com sucesso", containerName)
	return nil
}

func UpdateVelocityConfig(servers []models.ServerInfo) error {
	var config strings.Builder
	config.WriteString("[servers]\n")
	for _, server := range servers {
		config.WriteString(fmt.Sprintf("  [servers.%s]\n", server.Name))
		config.WriteString(fmt.Sprintf("  address = \"%s:25565\"\n", server.Name))
		config.WriteString("  retricted = false \n\n")
	}

	return os.WriteFile("../../velocity/velocity.toml", []byte(config.String()), 0644)
}

func StartContainer(name string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	return cli.ContainerStart(ctx, name, types.ContainerStartOptions{})
}

func StopContainer(name string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	timeout := 10
	return cli.ContainerStop(ctx, name, container.StopOptions{Timeout: &timeout})
}
