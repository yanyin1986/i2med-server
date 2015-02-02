package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"i2med-server/app/models"
	"strconv"
)

type UserController struct {
	*revel.Controller
}

type CommonResult struct {
	ErrorMessage string
	ErrorCode    int32
	Result       interface{}
}

type InnerLoginResult struct {
	Id int64
}

func (c UserController) Login() revel.Result {
	user := c.Params.Get("user")
	pwd := c.Params.Get("pwd")

	if user == "admin" && pwd == "admin" {
		result := CommonResult{
			ErrorMessage: "",
			ErrorCode:    0,
			Result: InnerLoginResult{
				Id: 1986,
			},
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

func (c UserController) Detail() revel.Result {
	id, _ := strconv.ParseInt(c.Params.Get("id"), 10, 64)
	user := models.GetUserById(id)
	if user == nil {
		return c.RenderJson(CommonResult{
			ErrorCode:    999,
			ErrorMessage: "错误啦啦啦啦",
		})
	} else {
		return c.RenderJson(CommonResult{
			ErrorCode:    0,
			ErrorMessage: "",
			Result:       user,
		})
	}
}

func (c UserController) Register() revel.Result {
	name := c.Params.Get("name")
	age, _ := strconv.ParseInt(c.Params.Get("age"), 10, 32)
	title := c.Params.Get("title")
	department := c.Params.Get("department")
	province := c.Params.Get("province")
	city := c.Params.Get("city")

	user := models.User{
		Name:       name,
		Age:        int32(age),
		Title:      title,
		Department: department,
		Province:   province,
		City:       city,
	}

	id, err := models.CreateUser(&user)

	if err != nil {
		return c.RenderJson(CommonResult{
			ErrorCode:    999,
			ErrorMessage: "错误啦啦啦啦",
		})
	} else {
		user.Id = id
		return c.RenderJson(CommonResult{
			ErrorCode:    0,
			ErrorMessage: "",
			Result:       user,
		})
	}

}

func (c UserController) Create() revel.Result {
	fmt.Println("yeyeyeyeye")
	user := models.User{
		Name: "朱胜超",
		Age:  14,
		Hospital: models.Hospital{
			Name: "南京人民医院",
		},
		Title:      "主任医师",
		Department: "肿瘤科",
		Province:   "江苏",
		City:       "南京",
		Education: []models.Education{
			models.Education{SchoolName: "南京大学", Collage: "医学院"},
		},
	}
	id, err := models.CreateUser(&user)
	if err != nil {
		return c.RenderJson(CommonResult{
			ErrorCode:    999,
			ErrorMessage: "错误啦啦啦啦",
		})
	} else {
		user.Id = id
		return c.RenderJson(CommonResult{
			ErrorCode:    0,
			ErrorMessage: "",
			Result:       user,
		})
	}
}
