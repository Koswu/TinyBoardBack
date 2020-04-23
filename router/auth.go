package router

import (
	"boarderbackend/models"
	"boarderbackend/pkgs/e"
	"boarderbackend/pkgs/logging"
	"boarderbackend/pkgs/utils"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"html"
	"io/ioutil"
	"net/http"
)

type account struct {

	Username string `json: "username" valid:"Required; MaxSize(50)"`
	Password string `json: "password" valid:"Required; MaxSize(50)"`
}

func printBody(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	logging.Info(string(body))
}

func GetAuth(c *gin.Context) {
	data := make(map[string]interface{})

	validator := validation.Validation{}
	acc := account{}
	acc.Username = html.EscapeString(c.PostForm("username"))
	acc.Password = html.EscapeString(c.PostForm("password"))
	ok, _ := validator.Valid(&acc)

	if !ok {
		for _, err :=range validator.Errors{
			logging.Info(err.Key, err.Message)
		}
		retFailed(c, e.InvalidParams)
		return
	}
	isAuthSuccess := models.CheckAuth(acc.Username, acc.Password)
	if !isAuthSuccess {
		retFailed(c, e.ErrorLogin)
		return
	}
	token, err := utils.GenerateToken(acc.Username)
	if err != nil{
		retFailed(c, e.ErrorAuthToken)
		return
	}
	code := e.Success
	data["username"] = acc.Username
	data["token"] = token
	c.JSON(http.StatusOK,
		gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
		})
}

func retFailed(c *gin.Context, code int) {
	c.JSON(http.StatusOK,
		gin.H{
			"code": code,
			"msg": e.GetMsg(code),
			"data": make(map[string]interface{}),
		})
	c.Abort()
}

func RegisterUser(c *gin.Context) {
	username := html.EscapeString(c.PostForm("username"))
	password := html.EscapeString(c.PostForm("password"))
	validator := validation.Validation{}
	data := make(map[string]interface{})
	a := account{
		Username: username,
		Password: password,
	}
	ok, _ := validator.Valid(&a)

	if !ok {
		for _, err :=range validator.Errors{
			logging.Info(err.Key, err.Message)
		}
		retFailed(c, e.InvalidParams)
		return
	}
	isExist := models.IsExistUser(username)
	if isExist {
		retFailed(c, e.ErrorRegisterUserExist)
		return
	}
	models.RegisterUser(username, password)
	code := e.Success
	c.JSON(http.StatusOK, gin.H{
		"code":code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}
