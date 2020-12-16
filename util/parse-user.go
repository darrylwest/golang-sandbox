package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type User struct {
	id          string
	dateCreated time.Time
	lastUpdated time.Time
	version     int64
	username    string
	session     string
	status      string
}

func parseDate(json map[string]interface{}, attr string, dflt time.Time) time.Time {
	if sdt, ok := json[attr].(string); ok {
		if dt, err := time.Parse(time.RFC3339Nano, sdt); err != nil {
			return dt
		}
	}

	return dflt
}

func parseInt64(json map[string]interface{}, attr string, dflt int64) int64 {
	if value, ok := json[attr].(float64); ok {
		return int64(value)
	} else {
		return dflt
	}
}

func parseUser(json map[string]interface{}) *User {
	user := new(User)

	user.id = json["id"].(string)
	user.dateCreated = parseDate(json, "dateCreated", time.Now())
	user.lastUpdated = parseDate(json, "lastUpdated", time.Now())

	user.version = parseInt64(json, "version", 0)

	user.username = json["username"].(string)
	user.session = json["session"].(string)
	user.status = json["status"].(string)

	return user
}

func main() {

	file, err := os.Open("./user.json")
	if err != nil {
		panic(err)
	}
	stat, _ := file.Stat()

	bytes := make([]byte, stat.Size())
	count, err := file.Read(bytes)
	if stat.Size() != int64(count) {
		fmt.Printf("size: %d != count: %d\n", stat.Size(), count)
	}

	var data map[string]interface{}

	if err := json.Unmarshal(bytes, &data); err != nil {
		panic(err)
	}

	user := parseUser(data["user"].(map[string]interface{}))

	fmt.Printf("%v\n", user)

}
