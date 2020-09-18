package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"bluebird/models/dbmodels"
	"bluebird/models"
	"bluebird/services"

	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	DB *gorm.DB
}

func (h *OrderController) SendOrder(c *gin.Context) {
	var req dbmodels.OrderReq
	var res models.Response

	topic := ""
	var data interface{}

	body := c.Request.Body
	reqBody, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(reqBody, &req); err != nil {
		c.JSON(500, gin.H{"error": "failed get body struct"})
		c.Abort()
		return
	}

	topic = req.Topic
	data = req.Data

	var svc services.KafkaService
	if res := svc.AddToKafka(topic, data); res != nil {
		fmt.Println("failed save to kafka  ====> ", res)
		c.AbortWithStatusJSON(200, res)
	}

	res.ErrCode = "00"
	res.ErrDesc = "Success"
	c.JSON(200, res)
}