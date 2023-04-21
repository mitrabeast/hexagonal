package server

import "fmt"

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (w *Config) Address() string {
	return fmt.Sprintf("%s:%d", w.Host, w.Port)
}
