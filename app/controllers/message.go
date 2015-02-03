package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"i2med-server/app/models"
	"strconv"
	"time"
)

type MessageController struct {
	*revel.Controller
}

func (c MessageController) CreateMessage() revel.Result {
	fmt.Println("hello")
	messageBody := c.Params.Get("message")
	userId, _ := strconv.ParseInt(c.Params.Get("userId"), 10, 64)

	fmt.Println("id=>?", 1)
	message := models.Message{
		UserId:     userId,
		Body:       messageBody,
		CommitTime: time.Now().Unix(),
	}
	id := models.CreateMessage(&message)

	fmt.Println("id=>?", id)
	if id > 0 {
		result := CommonResult{
			ErrorMessage: "ok",
			ErrorCode:    0,
		}

		return c.RenderJson(result)
	} else {
		result := CommonResult{
			ErrorMessage: "错误啦啦啦",
			ErrorCode:    999,
		}

		return c.RenderJson(result)
	}
}
