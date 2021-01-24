package reqres

import (
	"fmt"
	"time"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

// UserData ...
type Users struct {
	Data []Data `json:"data"`
}

// User ...
type User struct {
	Data Data `json:"data"`
}

// Data ...
type Data struct {
	ID     uint   `json:"id"`
	Email  string `json:"email"`
	First  string `json:"first_name"`
	Last   string `json:"last_name"`
	Avatar string `json:"avatar"`
}

type support struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

func usrToString(usr User) string {
	s := ""

	s += fmt.Sprintln("Id....: ", usr.Data.ID)
	s += fmt.Sprintln("Email.: ", usr.Data.Email)
	s += fmt.Sprintln("First.: ", usr.Data.First)
	s += fmt.Sprintln("Last..: ", usr.Data.Last)
	s += fmt.Sprintln("Avatar: ", usr.Data.Avatar)

	return s
}

var baseURL = "https://reqres.in/"

// Header ...
func Header() (string, string) {
	return "Accept", "application/json"
}

func showRes(resp *resty.Response, err error, ti time.Duration) string {
	s := ""

	if(err != nil) {
		s += fmt.Sprintln("Error....: ", err)
	}

	s += fmt.Sprintln("Status...: ", resp.Status())
	s += fmt.Sprintln("Resp Time: ", ti)

	return s
}

// Status ...
func Status(assert *assert.Assertions, status int, resp int, name string) {
	msg := fmt.Sprintf("%s - Status isnt %v == %v", name, status, resp)

	assert.Equal(status, resp, msg)
}

// RespTime ...
func RespTime(assert *assert.Assertions, tempo int, resp time.Duration, name string) {
	t := time.Duration(tempo) * time.Microsecond
	msg := fmt.Sprintf("%s - Response Time isnt %v < %v", name, resp, t)

	assert.LessOrEqual(resp, t, msg)
}
