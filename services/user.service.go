package services

import (
	"fmt"
	repository "bluebird/database"
	dto "bluebird/models/dto"
	"time"
	constants "bluebird/constants"
	jwt "github.com/dgrijalva/jwt-go"
)

type UserService struct {
}

func (h UserService) AuthLogin(userDto *dto.LoginRequestDto) dto.LoginResponseDto {
	var res dto.LoginResponseDto
	if userDto.Username == "" {
		res.ErrCode = constants.ERR_CODE_50
		res.ErrDesc = constants.ERR_CODE_50_MSG
		return res
	}

	if userDto.Password == "" {
		res.ErrCode = constants.ERR_CODE_50
		res.ErrDesc = constants.ERR_CODE_50_MSG
		return res
	}

	user, err := repository.GetUserByName(userDto.Username)
	if err != nil {
		res.ErrCode = constants.ERR_CODE_51
		res.ErrDesc = constants.ERR_CODE_51_MSG
		return res
	}

	if user.ID == 0 {
		res.ErrCode = constants.ERR_CODE_50
		res.ErrDesc = constants.ERR_CODE_50_MSG
		return res
	}

	if user.Password != userDto.Password {
		res.ErrCode = constants.ERR_CODE_50
		res.ErrDesc = constants.ERR_CODE_50_MSG
		return res
	}

	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := sign.Claims.(jwt.MapClaims)
	claims["user"] = user.UserName

	unixNano := time.Now().UnixNano()
	umillisec := unixNano / 1000000
	timeToString := fmt.Sprintf("%v", umillisec)
	fmt.Println("token Created ", timeToString)
	claims["tokenCreated"] = timeToString

	token, err := sign.SignedString([]byte(constants.TokenSecretKey))
	if err != nil {
		res.ErrCode = constants.ERR_CODE_53
		res.ErrDesc = constants.ERR_CODE_53_MSG
		res.Token = ""
		return res
	}

	res.ErrCode = constants.ERR_CODE_00
	res.ErrDesc = constants.ERR_CODE_00_MSG
	res.Token = token

	return res
}