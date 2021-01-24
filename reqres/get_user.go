package reqres

import (
    "fmt"
    "time"
    "github.com/go-resty/resty/v2"
)

// GetUser ...
func GetUser(user User) (User, *resty.Response, time.Duration, string) {
    client := resty.New()

    resp, err := client.R().
        EnableTrace().
        SetHeaders(Headers()).
        SetResult(&User{}).
        Get(baseURL + "/users/" + fmt.Sprint(user.Data.ID))

    ti  := resp.Request.TraceInfo()
    usr := (*resp.Result().(*User))

    fmt.Printf("GET /users/{userId}\n\n")
    fmt.Println(usr.toString())
    fmt.Println(ShowRes(resp, err, ti.ResponseTime))

    return usr, resp, ti.ResponseTime, "Get User"
}
