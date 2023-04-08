package config

import "fmt"

type Web struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (w *Web) Address() string {
	return fmt.Sprintf("%s:%d", w.Host, w.Port)
}
