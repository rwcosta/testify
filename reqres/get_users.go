package reqres

import (
    "fmt"
    "time"
    "math/rand"
    "github.com/go-resty/resty/v2"
)

// GetUsers ...
func GetUsers() (User, *resty.Response, time.Duration, string) {
    client := resty.New()

    resp, err := client.R().
	    EnableTrace().
	    SetQueryString("page=2").
	    SetHeaders(Headers()).
	    SetResult(&Users{}).
	    Get(baseURL + "api/users")

    ti      := resp.Request.TraceInfo()
    usrData := (*resp.Result().(*Users))
    rnd     := rand.New(rand.NewSource(time.Now().Unix()))
    usr     := usrData.Data[rnd.Intn(len(usrData.Data))]

    fmt.Printf("Get Users\n\n")
    fmt.Println(usr.toString())
    fmt.Println(ShowRes(resp, err, ti.ResponseTime))

    return usr, resp, ti.ResponseTime, "Get Users"
}
