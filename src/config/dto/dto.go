package dto

type Res struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type LogBody struct {
	SpendTime string      `json:"spend_time"`
	Path      string      `json:"path"`
	Method    string      `json:"method"`
	Status    int         `json:"status"`
	Proto     string      `json:"proto"`
	Ip        string      `json:"ip"`
	Body      string      `json:"body"`
	Query     string      `json:"query"`
	Message   interface{} `json:"message"`
}
type Captcha struct {
	Id      string `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
}

type TokenClaims struct {
	Id    uint   `json:"id,omitempty" `
	Name  string `json:"name,omitempty" binding:"-"`
	Phone string `json:"phone,omitempty" binding:"-"`
	Email string `json:"email"`
	Role  int32  `json:"role"`
}
type TokenAndExp struct {
	Token   string `json:"token,omitempty"`
	ExpTime string `json:"exp_time,omitempty"`
}
