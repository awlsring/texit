package db

import "fmt"

func CreatePostgresConnectionURI(host string, port int, user, pass, dbName string, ssl bool) string {
	sslMode := "disable"
	if ssl {
		sslMode = "enable"
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", user, pass, host, port, dbName, sslMode)
}
