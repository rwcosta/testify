package reqres

import (
	"fmt"
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func TestReqres(t *testing.T) {
	assert := assert.New(t)

	usr, resp, time := GetUsers()
	Status(assert, http.StatusOK, resp.StatusCode(), "Get Users")
	RespTime(assert, 350, time, "GetUsers")

	fmt.Println(usrToString(usr))
}
