package routers

import (
	"fmt"
	"net/http"
	"bluebird/controllers"
	"bluebird/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"bluebird/constants"

	"time"
	"strconv"

	"bluebird/models/dto"
)

func InitRouter() *gin.Engine {
    r := gin.New()

	fmt.Println(gin.IsDebugging())

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "authorization", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge: 86400,
	}))

	AuthController := new(controllers.AuthController)
	OrderController := new(controllers.OrderController)

	api := r.Group("/auth")
	api.POST("/login", AuthController.Login)

	api = r.Group("/order")
	api.POST("/", cekToken, OrderController.SendOrder)

	return r
}

func cekToken(c *gin.Context) {
	res := models.Response{}
	tokenString := c.Request.Header.Get("Authorization")

	if strings.HasPrefix(tokenString, "Bearer ") == false {
		res.ErrCode = constants.ERR_CODE_54
		res.ErrDesc = constants.ERR_CODE_54_MSG
		c.JSON(http.StatusUnauthorized, res)
		c.Abort()
		return
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			res.ErrCode = constants.ERR_CODE_54
			res.ErrDesc = constants.ERR_CODE_54_MSG
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
		}
		return []byte(constants.TokenSecretKey), nil
	})

	if token != nil && err == nil {
		claims := token.Claims.(jwt.MapClaims)

		unixNano := time.Now().UnixNano()
		timeNowInInt := unixNano / 1000000

		tokenCreated := (claims["tokenCreated"])
		dto.CurrUser = (claims["user"]).(string)

		tokenCreatedInString := tokenCreated.(string)
		tokenCreatedInInt, errTokenExpired := strconv.ParseInt(tokenCreatedInString, 10, 64)

		if errTokenExpired != nil {
			res.ErrCode = constants.ERR_CODE_55
			res.ErrDesc = constants.ERR_CODE_55_MSG
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		if ((timeNowInInt - tokenCreatedInInt) / 1000) > constants.TokenExpiredInMinutes {
			res.ErrCode = constants.ERR_CODE_55
			res.ErrDesc = constants.ERR_CODE_55_MSG
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}
	} else {
		res.ErrCode = constants.ERR_CODE_54
		res.ErrDesc = constants.ERR_CODE_54_MSG
		c.JSON(http.StatusUnauthorized, res)
		c.Abort()
		return
	}
}
