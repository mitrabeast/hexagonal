package config

import "fmt"

type Database struct {
	Name         string `json:"name"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	MaxOpenConns int    `json:"maxOpenConns"`
}

func (d *Database) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.User, d.Password, d.Name,
	)
}
