// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tFriendController struct {}
var FriendController tFriendController


func (_ tFriendController) AddFriend(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("FriendController.AddFriend", args).Url
}

func (_ tFriendController) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("FriendController.List", args).Url
}

func (_ tFriendController) Get(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("FriendController.Get", args).Url
}

func (_ tFriendController) Search(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("FriendController.Search", args).Url
}


type tMessageController struct {}
var MessageController tMessageController


func (_ tMessageController) CreateMessage(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MessageController.CreateMessage", args).Url
}


type tUserController struct {}
var UserController tUserController


func (_ tUserController) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UserController.Login", args).Url
}

func (_ tUserController) Detail(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UserController.Detail", args).Url
}

func (_ tUserController) Register(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UserController.Register", args).Url
}

func (_ tUserController) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UserController.Create", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


