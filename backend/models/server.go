package models

type ServerRequest struct {
	Name      string `json:"name"`
	Port      string `json:"port"`
	Project   string `json:"project"`
	GameMode  string `json:"gameMode"`
	LevelType string `json:"levelType"`
}

type ServerInfo struct {
	Name   string `json:"name"`
	Port   string `json:"port"`
	Status string `json:"status"`
}
