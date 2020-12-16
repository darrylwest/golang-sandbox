package main

import (
	"encoding/json"
	"fmt"
	"github.com/pborman/uuid"
	"strings"
	"time"
)

type DocumentIdentifier struct {
	id          string
	dateCreated time.Time
	lastUpdated time.Time
	version     int64
}

type User struct {
	doi      DocumentIdentifier
	username string
	session  string
	status   string
}

func NewModelId() string {
	parts := strings.Split(uuid.New(), "-")

	return strings.Join(parts, "")
}

func (doi DocumentIdentifier) ToMap() map[string]interface{} {
	var m = map[string]interface{}{
		"id":          doi.id,
		"dateCreated": doi.dateCreated,
		"lastUpdated": doi.lastUpdated,
		"version":     doi.version,
	}

	return m
}

func (u User) ToMap() map[string]interface{} {
	var m = u.doi.ToMap()

	m["username"] = u.username
	m["session"] = u.session
	m["status"] = u.status

	return m
}

func (ptr *DocumentIdentifier) FromMap(v map[string]interface{}) {
	doi := *ptr

	doi.id = v["id"].(string)

	dt, _ := time.Parse(time.RFC3339Nano, v["dateCreated"].(string))

	doi.dateCreated = dt

	dt, _ = time.Parse(time.RFC3339Nano, v["lastUpdated"].(string))
	doi.lastUpdated = dt

	doi.version = int64(v["version"].(float64))

	*ptr = doi
}

func (ptr *User) FromMap(v map[string]interface{}) {
	u := *ptr

	u.doi.FromMap(v)

	u.username = v["username"].(string)
	u.session = v["session"].(string)
	u.status = v["status"].(string)

	*ptr = u
}

func ToJson(v map[string]interface{}) ([]byte, error) {
	// fmt.Printf("%v\n", v)
	return json.MarshalIndent(v, "", "  ")
}

func main() {
	user := new(User)

	user.doi.id = NewModelId()
	user.doi.dateCreated = time.Now().UTC()
	user.doi.lastUpdated = time.Now().UTC()
	user.doi.version = 1234

	user.username = "dpw@rcs.com"
	user.session = uuid.New()
	user.status = "active"

	fmt.Printf("type: %T %v\n", *user, user)

	fmt.Println("user model:")
	fmt.Printf("%s %v %v %d %s\n", user.doi.id, user.doi.dateCreated, user.doi.lastUpdated, user.doi.version, user.username)

	b, err := ToJson(user.ToMap())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf("%s\n", b)

	// now unmarshal

	v := make(map[string]interface{})
	err = json.Unmarshal(b, &v)

	// fmt.Printf("%v\n", v)

	u := new(User)
	u.FromMap(v)

	fmt.Printf("%s %v %v %d %s\n", u.doi.id, u.doi.dateCreated, u.doi.lastUpdated, u.doi.version, u.username)
}
