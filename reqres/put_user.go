package reqres

import (
    "fmt"
    "time"
    "github.com/go-resty/resty/v2"
)

// PutUser ...
func PutUser(user PostUserResp, body UserBody) (UpdateUserResp, *resty.Response, time.Duration, string) {
    client := resty.New()

    resp, err := client.R().
        EnableTrace().
        SetHeaders(Headers()).
        SetResult(&UpdateUserResp{}).
        SetBody(body).
        Put(baseURL + "/users/" + fmt.Sprint(user.ID))

    ti  := resp.Request.TraceInfo()
    usr := (*resp.Result().(*UpdateUserResp))

    fmt.Printf("PUT /users/{userId}\n\n")
    fmt.Println(usr.toString())
    fmt.Println(ShowRes(resp, err, ti.ResponseTime))

    return usr, resp, ti.ResponseTime, "Put User"
}
