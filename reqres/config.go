package reqres

import (
	"fmt"
	"time"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

// Users ...
type Users struct {
	Data []User `json:"data"`
}

// User ...
type User struct {
	ID     uint   `json:"id"`
	Email  string `json:"email"`
	First  string `json:"first_name"`
	Last   string `json:"last_name"`
	Avatar string `json:"avatar"`
}

func (usr User) toString() string {
	s := ""

	s += fmt.Sprintln("Id....: ", usr.ID)
	s += fmt.Sprintln("Email.: ", usr.Email)
	s += fmt.Sprintln("First.: ", usr.First)
	s += fmt.Sprintln("Last..: ", usr.Last)
	s += fmt.Sprintln("Avatar: ", usr.Avatar)

	return s
}

// UserBody ...
type UserBody struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

// PostUserResp ...
type PostUserResp struct {
	Name string `json:"name"`
	Job  string `json:"job"`
	ID   string `json:"id"`
	Date string `json:"createdAt"`
}

func (usrResp PostUserResp) toString() string {
	s := ""

	s += fmt.Sprintln("Name: ", usrResp.Name)
	s += fmt.Sprintln("Job.: ", usrResp.Job)
	s += fmt.Sprintln("Id..: ", usrResp.ID)
	s += fmt.Sprintln("Date: ", usrResp.Date)

	return s
}

var baseURL = "https://reqres.in/"

// Headers ...
func Headers() map[string]string {
	return map[string]string {
		"Accept": "application/json",
		"Content-Type": "application/json",
	}
}

// ShowRes ...
func ShowRes(resp *resty.Response, err error, ti time.Duration) string {
	s := ""

	if(err != nil) {
		s += fmt.Sprintln("Error....: ", err)
	}

	s += fmt.Sprintln("Status...: ", resp.Status())
	s += fmt.Sprintln("Resp Time: ", ti)
	s += fmt.Sprintf("\n-----------------------------------------\n")

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
