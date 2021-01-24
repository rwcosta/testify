package reqres

import (
    "fmt"
    "testing"
    "net/http"
    "github.com/stretchr/testify/assert"
)

func TestData(t *testing.T) {
    assert := assert.New(t)

    usr, resp, time, name := GetUsers()
    Status(assert, http.StatusOK, resp.StatusCode(), name)
    RespTime(assert, 370, time, name)

    if(!assert.NotEqual(nil, usr, fmt.Sprint("User Object is nil)"))) {
        return
    }

    usr, resp, time, name = GetUser(usr)
    Status(assert, http.StatusOK, resp.StatusCode(), name)
    RespTime(assert, 350, time, name)

    if(!assert.NotEqual(nil, usr, fmt.Sprint("User Object is nil)"))) {
        return
    }
}

func TestInsert(t *testing.T) {
    assert := assert.New(t)

    usrResp, resp, time, name := PostUser(UserBody{"Josh", "Engineer"})
    Status(assert, http.StatusCreated, resp.StatusCode(), name)
    RespTime(assert, 375, time, name)

    if(!assert.NotEqual(nil, usrResp, fmt.Sprint("PostUserResp Object is nil"))) {
        return
    }
}
