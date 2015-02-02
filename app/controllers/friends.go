package controllers

import (
	"github.com/revel/revel"
	"i2med-server/app/models"
	"strconv"
	"time"
)

type FriendController struct {
	*revel.Controller
}

func (f FriendController) AddFriend() revel.Result {
	friendId, _ := strconv.ParseInt(f.Params.Get("friendId"), 10, 64)
	userId, _ := strconv.ParseInt(f.Params.Get("userId"), 10, 64)

	now := time.Now().Unix()
	friend := models.Friend{
		UserId:     userId,
		FriendId:   friendId,
		Status:     models.FriendRequest,
		CreateTime: now,
		ModifyTime: now,
	}
	id, err := models.CreateFriend(&friend)
	if err != nil {
		panic(err)
		return f.RenderJson(CommonResult{
			ErrorCode:    999,
			ErrorMessage: "错误啦啦啦啦",
		})
	} else {
		friend.Id = id
		return f.RenderJson(CommonResult{
			ErrorCode:    0,
			ErrorMessage: "",
			Result:       friend,
		})
	}
}

func (f FriendController) List() revel.Result {
	userId, _ := strconv.ParseInt(f.Params.Get("userId"), 10, 64)

	friends, err := models.GetFriendsByUserId(userId)

	if err != nil {
		panic(err)
		return f.RenderJson(CommonResult{
			ErrorCode:    999,
			ErrorMessage: "出错啦啦啦啦啦",
		})
	} else {
		return f.RenderJson(CommonResult{
			ErrorCode:    0,
			ErrorMessage: "",
			Result:       friends,
		})
	}
}

func (f FriendController) Get() revel.Result {
	userId, _ := strconv.ParseInt(f.Params.Get("userId"), 10, 64)
	friendId, _ := strconv.ParseInt(f.Params.Get("friendId"), 10, 64)

	friend, err := models.GetFriendByUserIdAndFriendId(userId, friendId)

	if err != nil {
		panic(err)
		return f.RenderJson(CommonResult{
			ErrorCode:    999,
			ErrorMessage: "出错啦啦啦啦",
		})
	} else {
		return f.RenderJson(CommonResult{
			ErrorCode:    0,
			ErrorMessage: "",
			Result:       friend,
		})
	}
}

func (f FriendController) Search() revel.Result {
	return f.RenderText("developing")
}
