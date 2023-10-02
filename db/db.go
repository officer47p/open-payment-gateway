package db

import "fmt"

func CreateDBUrl(url string, port string, name string, user string, password string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, url, port, name)
}
