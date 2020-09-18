package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// "net/http"
	constants "bluebird/constants"
	models "bluebird/models"
	dto "bluebird/models/dto"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"bluebird/services"
)

type AuthController struct {
	DB *gorm.DB
}

var UserService = new(services.UserService)

func (h *AuthController) Login(c *gin.Context) {
	req := dto.LoginRequestDto{}
	res := models.Response{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		fmt.Println("Error, body Request ")
		res.ErrCode = constants.ERR_CODE_03
		res.ErrDesc = constants.ERR_CODE_03_MSG
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, UserService.AuthLogin(&req))
}