package reqres

import (
    "fmt"
    "testing"
    "net/http"
    "github.com/go-resty/resty/v2"
    "github.com/stretchr/testify/suite"
    "github.com/stretchr/testify/assert"
)

type UserSuite struct {
    suite.Suite
    APIClient *resty.Client
}

func (usr *UserSuite) SetupTest() {
    usr.APIClient = resty.New()
}

func (suite *UserSuite) TestData() {
    assert := assert.New(suite.T())

    usr, resp, time, name := GetUsers(suite)
    Status(assert, http.StatusOK, resp.StatusCode(), name)
    RespTime(assert, 370, time, name)

    if(!assert.NotEqual(nil, usr, fmt.Sprint("User Object is nil)"))) {
        return
    }

    usr, resp, time, name = GetUser(suite, usr)
    Status(assert, http.StatusOK, resp.StatusCode(), name)
    RespTime(assert, 350, time, name)

    if(!assert.NotEqual(nil, usr, fmt.Sprint("User Object is nil)"))) {
        return
    }
}

func (suite *UserSuite) TestInsert() {
    assert := assert.New(suite.T())

    usrResp, resp, time, name := PostUser(suite, UserBody{"Josh", "Engineer"})
    Status(assert, http.StatusCreated, resp.StatusCode(), name)
    RespTime(assert, 375, time, name)

    if(!assert.NotEqual(nil, usrResp, fmt.Sprint("PostUserResp Object is nil"))) {
        return
    }

    uptResp := UpdateUserResp{}
    uptResp, resp, time, name = PutUser(suite, usrResp, UserBody{"Smith", "Driver"})
    Status(assert, http.StatusOK, resp.StatusCode(), name)
    RespTime(assert, 385, time, name)

    if(!assert.NotEqual(nil, uptResp, fmt.Sprint("UpdateUserResp Object is nil"))) {
        return
    }
}

func TestReqres(t *testing.T) {
    suite.Run(t, new(UserSuite))
}
