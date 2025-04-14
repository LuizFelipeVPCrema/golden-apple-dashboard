package models

type ServerRequest struct {
	Name    string `json:"name"`
	Port    string `json:"port"`
	Project string `json:"project"`
}

type ServerInfo struct {
	Name   string `json:"name"`
	Port   string `json:"port"`
	Status string `json:"status"`
}
