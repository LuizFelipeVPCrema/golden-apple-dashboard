package services

import (
	"fmt"
	"log"
	"strings"

	"github.com/LuizFelipeVPCrema/minecraft-dashboard-go/models"
	"github.com/LuizFelipeVPCrema/minecraft-dashboard-go/utils"
)

func CreateServer(req models.ServerRequest) error {
	containerName := req.Name
	port := req.Port

	return utils.RunMinecraftContainer(containerName, port, req.Project, req)
}

func ListServersWithFilters(statusFilter, projectFilter string) ([]models.ServerInfo, error) {
	containers, err := utils.ListMinecraftContainers()
	if err != nil {
		return nil, err
	}

	var servers []models.ServerInfo
	for _, c := range containers {
		name := strings.TrimPrefix(c.Names[0], "/")

		port := ""
		if len(c.Ports) > 0 {
			port = fmt.Sprintf("%d", c.Ports[0].PublicPort)
		}

		status := "stopped"
		if c.State == "running" {
			status = "running"
		}

		project := c.Labels["project"]

		// Aplica filtros se foram definidos
		if statusFilter != "" && status != statusFilter {
			continue
		}
		if projectFilter != "" && project != projectFilter {
			continue
		}

		servers = append(servers, models.ServerInfo{
			Name:   name,
			Port:   port,
			Status: status,
		})

		log.Printf("Container: %s | status=%s | project=%s", name, status, project)
		log.Printf("Filtro: status=%s | project=%s", statusFilter, projectFilter)

	}

	return servers, nil
}

func DeleteServer(name string) error {
	containerName := name

	return utils.StopAndRemoveContainer(containerName)
}

func StartServer(containerName string) error {
	return utils.StartContainer(containerName)
}

func StopServer(containerName string) error {
	return utils.StopContainer(containerName)
}
