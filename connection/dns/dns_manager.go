package dns

import (
	"github.com/nsdash/config-lib"
)

type Manager struct {
	configManager config.Manager
}

type Data struct {
	User       string
	Password   string
	Host       string
	Port       string
	Database   string
	DriverName string
}

func NewManager() Manager {
	return Manager{
		configManager: config.NewManager(),
	}
}

func (m Manager) GetData() Data {
	return Data{
		User:       m.configManager.Get("MYSQL_DB_USER"),
		Password:   m.configManager.Get("MYSQL_DB_PASSWORD"),
		Host:       m.configManager.Get("MYSQL_DB_HOST"),
		Port:       m.configManager.Get("MYSQL_DB_PORT"),
		Database:   m.configManager.Get("MYSQL_DB_NAME"),
		DriverName: m.configManager.Get("MYSQL_DB_DRIVER"),
	}
}
