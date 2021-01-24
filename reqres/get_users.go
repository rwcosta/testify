package reqres

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/go-resty/resty/v2"
)

// GetUsers ...
func GetUsers() (User, *resty.Response, time.Duration) {
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		SetQueryString("page=2").
		SetHeader(Header()).
		SetResult(&Users{}).
		Get(baseURL + "api/users")

	ti      := resp.Request.TraceInfo()
	usrData := (*resp.Result().(*Users))
	rnd     := rand.New(rand.NewSource(time.Now().Unix()))
	usr     := User{usrData.Data[rnd.Intn(len(usrData.Data))]}

	fmt.Println(usrToString(usr))
	fmt.Println(showRes(resp, err, ti.ResponseTime))

	return usr, resp, ti.ResponseTime
}
