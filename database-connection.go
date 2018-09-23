package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

//DatabaseConnection represents the settings and inputs for the LoadTest
type DatabaseConnection struct {
	Type             string
	ParsedConnection string
}

func getDefaultPort(connection string) int {
	if strings.HasPrefix(connection, "postgres") {
		return 5432
	} else if strings.HasPrefix(connection, "mysql") {
		return 3306
	}
	return -1
}

// GetConnectionInfo determine db type and formats connectionstring
func GetConnectionInfo(connection string) DatabaseConnection {
	u, err := url.Parse(connection)
	if err != nil {
		fmt.Println("error", err)
	}
	p, _ := u.User.Password()
	hostParts := strings.Split(u.Host, ":")
	var port int
	if len(hostParts) == 2 {
		port, _ = strconv.Atoi(hostParts[1])
	} else {
		port = getDefaultPort(connection)
	}
	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostParts[0], port, u.User.Username(), p, u.Path[1:])
	dbConn := DatabaseConnection{
		Type:             u.Scheme,
		ParsedConnection: connString,
	}
	return dbConn
}
