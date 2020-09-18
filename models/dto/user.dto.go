package dto

var CurrUser string

type LoginRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Response string `json:"response"`
}

type LoginResponseDto struct {
	ErrCode string `json:"errCode"`
	ErrDesc string `json:"errDesc"`
	Token   string `json:"token"`
}

type CheckAuthMenuRequestDto struct {
	MenuLink string `json:"menuLink"`
}