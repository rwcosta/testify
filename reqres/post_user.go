package reqres

import (
    "fmt"
    "time"
    "github.com/go-resty/resty/v2"
)

// PostUser ...
func PostUser(user UserBody) (PostUserResp, *resty.Response, time.Duration, string) {
    client := resty.New()

    resp, err := client.R().
        EnableTrace().
        SetHeaders(Headers()).
        SetResult(&PostUserResp{}).
        SetBody(user).
        Post(baseURL + "/users")

    ti      := resp.Request.TraceInfo()
    usrResp := (*resp.Result().(*PostUserResp))

    fmt.Printf("POST /users\n\n")
    fmt.Println(usrResp.toString())
    fmt.Println(ShowRes(resp, err, ti.ResponseTime))

    return usrResp, resp, ti.ResponseTime, "Post User"
}
