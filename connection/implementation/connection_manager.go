package implementation

import (
	"database/sql"
	"fmt"
	dns2 "github.com/nsdash/mysql-lib/connection/dns"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type ConnectionManager struct {
	openedConnection *sql.DB
}

var connectionManager *ConnectionManager

func GetConnectionManagerSingleton() ConnectionManager {
	if connectionManager == nil {
		connectionManager = &ConnectionManager{openedConnection: nil}
	}

	return *connectionManager
}

func (cm *ConnectionManager) GetConnection() *sql.DB {

	if cm.openedConnection != nil {
		return cm.openedConnection
	}

	cm.openedConnection = openConnection()

	return cm.openedConnection
}

func (cm *ConnectionManager) CloseConnection() {
	if cm.openedConnection == nil {
		return
	}

	err := cm.openedConnection.Close()

	if err != nil {
		log.Fatal(err)
	}
}

func openConnection() *sql.DB {
	dnsData := dns2.NewManager().GetData()

	pattern := "%s:%s@tcp(%s:%s)/%s"

	dns := fmt.Sprintf(
		pattern,
		dnsData.User,
		dnsData.Password,
		dnsData.Host,
		dnsData.Port,
		dnsData.Database,
	)

	connection, err := sql.Open(dnsData.DriverName, dns)

	if err != nil {
		log.Fatal(err)
	}

	return connection
}
