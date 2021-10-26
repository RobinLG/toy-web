package v1

import "fmt"

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func SignUp(c *Context) {
	req := &signUpReq{}
	err := c.ReadJson(req)
	if err != nil {
		_ = c.BadRequestJson(&commonResponse{
			BizCode: 4,
			Msg:     fmt.Sprintf("invalid request: %v", err),
		})
		return
	}
	_ = c.OkJson(&commonResponse{
		Data: 123,
	})
}
