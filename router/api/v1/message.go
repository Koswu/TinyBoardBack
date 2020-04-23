package v1

import (
	"boarderbackend/models"
	"boarderbackend/pkgs/e"
	"boarderbackend/pkgs/logging"
	"boarderbackend/pkgs/setting"
	"boarderbackend/pkgs/utils"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"html"
	"net/http"
)

type post struct {
	Title string `json: "title" valid:"Required MaxSize(100)"`
	Content string `json: "content" valid:"Required MaxSize(1000)"`
	Token string `json: "token" valid:"Required MaxSize(10000)"`
}

func GetComments(c *gin.Context) {
	data :=make(map[string]interface{})
	maps :=make(map[string]interface{})
	validator := validation.Validation{}
	pageNum := utils.GetPage(c)

	validator.Range(pageNum, 1, 1000, "page_num").Message("页码超出范围")
	code := e.InvalidParams
	if ! validator.HasErrors(){
		count := models.GetCommentCount(maps)
		pageSize := setting.App.PageSize
		code = e.Success
		data["lists"] =models.GetComments(pageNum, pageSize, maps)
		data["count"] = count
		data["page_count"] = (count + pageSize -1) / pageSize
	} else {
		for _, err:= range validator.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data":data,
	})
}

func GetComment(c *gin.Context) {
	validator := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	validator.Range(id, 1, 1000000, "id").Message("id不在有效范围内")
	var data interface{}
	code := e.InvalidParams
	if !validator.HasErrors(){
		if models.ExistCommentByID(id){
			code = e.Success
			data =models.GetComment(id)
		} else{
			code = e.ErrorNotExistComment
		}
	} else {
		for _, err :=range validator.Errors{
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(code,
		gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
		})
}

func PutComment(c *gin.Context) {
	p := post{}
	p.Title = c.PostForm("title")
	p.Content = c.PostForm("content")
	p.Token = c.PostForm("token")
	p.Title = html.EscapeString(p.Title)
	p.Content = html.EscapeString(p.Content)

	validator := validation.Validation{}
	ok, _ := validator.Valid(p)

	data := make(map[string]interface{})
	code := e.InvalidParams

	if ok{
		data["title"] = p.Title
		data["content"] = p.Content
		claim, _ := utils.ParseToken(p.Token)
		data["posted_by"] = claim.Username
		models.AddComment(data)
		code = e.Success
	} else {
		for _, err := range validator.Errors{
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(code,
		gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data": make(map[string]interface{}),
		})
}

