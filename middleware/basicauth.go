package middleware

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type usernameType string
type passwordType string

var basicAuthUsers = map[usernameType]passwordType{}

const csvUsers = "users.csv"

func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(basicAuthUsers) == 0 {
			var err error
			basicAuthUsers, err = readUsersFromCSV(csvUsers)
			if err != nil {
				log.Fatalln(fmt.Sprintf("BasicAuth.readUsersFromCSV returned error: %s", err))
			}
		}

		if username, password, ok := r.BasicAuth(); ok {
			if expectedPass, ok := basicAuthUsers[usernameType(username)]; ok && expectedPass == passwordType(password) {
				next.ServeHTTP(w, r)
				return
			}
		}

		// If the Authentication header is not present, is invalid, or the
		// username or password is wrong, then set a WWW-Authenticate
		// header to inform the client that we expect them to use basic
		// authentication and send a 401 Unauthorized response.
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func readUsersFromCSV(csvFilePath string) (map[usernameType]passwordType, error) {
	users := map[usernameType]passwordType{}
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		return users, err
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	rawCSVData, err := reader.ReadAll()
	if err != nil {
		return users, err
	}

	header := []string{} // holds first row (header)
	for lineNum, record := range rawCSVData {
		// for first row, build the header slice
		if lineNum == 0 {
			for i := 0; i < len(record); i++ {
				header = append(header, strings.TrimSpace(record[i]))
			}
			if len(header) != 2 || header[0] != "username" || header[1] != "password" {
				return users, errors.New(fmt.Sprintf("%s must start with headers: username,password", csvUsers))
			}
		} else {
			users[usernameType(record[0])] = passwordType(record[1])
		}
	}

	return users, nil
}
