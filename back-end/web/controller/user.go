package controller

import (
	"data-manager/exterror"
	middlewares "data-manager/middleware"
	"data-manager/model"
	webCtx "data-manager/web/context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserManagerRequest struct {
	//Name        string `json:"name" bson:"name"`
	//Password    string `json:"password" bson:"password"`
	//Role        string `json:"role" bson:"role"`
	//NewPassword string `json:"newpassword" bson:"newpassword"`
	//Passphrase  string `json:"passphrase" bson:"passphrase"`
}

var UserController = &UserManagerRequest{}

func (u UserManagerRequest) SetSuperAdmin(ctx *webCtx.Context) {
	var initInfo model.InitRequest
	if err := ctx.ShouldBind(&initInfo); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "绑定输入参数失败！")
		return
	}
	if initInfo.Name == "" || initInfo.Password == "" || initInfo.Passphrase == "" || initInfo.Address==""{
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "必填字段不能为空！")
		return
	}
	err := model.DefaultUserModel.SetSuperAdmin(ctx, initInfo.Name, initInfo.Password, initInfo.Address, initInfo.Passphrase)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "注册失败！")
		return

	}

	middlewares.ResponseSuccess(ctx, "注册成功", nil)
}

func (u *UserManagerRequest) Login(ctx *webCtx.Context) {

	var loginInfo model.LoginRequest


	if err := ctx.ShouldBind(&loginInfo); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败! ")
		return
	}
	if loginInfo.Name == "" || loginInfo.Password == "" {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "用户名或密码不能为空！")
		return
	}
	output, err := model.DefaultUserModel.UserLogin(ctx, loginInfo.Name, loginInfo.Password)

	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "用户名不存在或密码错误，请重新登陆！")
		return
	}

	middlewares.ResponseSuccess(ctx, "登陆成功", output)
	return
}

func (u *UserManagerRequest) NotFound(ctx *webCtx.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404 ,page not exists!",
	})
}

func LoginValid(name, password string) {
	//todo
}
func (u UserManagerRequest) AddUser(ctx *webCtx.Context) {

	var addUserInfo model.AddUserRequest
	if err := ctx.ShouldBind(&addUserInfo); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}
	if addUserInfo.Name == "" || addUserInfo.Password == "" || addUserInfo.Role == "" || addUserInfo.Passphrase == "" {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "必填字段不能为空！")
		return
	}
	res, err := model.DefaultUserModel.AddUser(ctx, addUserInfo.Name, addUserInfo.Password, addUserInfo.Role, addUserInfo.Passphrase)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "增加用户失败！")
		return

	}

	middlewares.ResponseSuccess(ctx, "添加用户成功", res)
}

func (u UserManagerRequest) DeleteUser(ctx *webCtx.Context) {
	//var user UserStruct
	var deleteUserInfo model.DeleteUserRequest
	if err := ctx.ShouldBind(&deleteUserInfo); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败! ")
		return
	}
	if deleteUserInfo.Name == "" {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "用户名不能为空！")
		return
	}
	err := model.DefaultUserModel.DeleteUser(ctx, deleteUserInfo.Name)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "删除用户失败！")
		return

	}

	middlewares.ResponseSuccess(ctx, "删除用户成功", nil)
}

func (u UserManagerRequest) Update(ctx *webCtx.Context) {
	var updateUserInfo model.UpdateUserRequest
	if err := ctx.ShouldBind(&updateUserInfo); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}
	if updateUserInfo.Name == "" {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "用户名不能为空！")
		return

	}
	res, err := model.DefaultUserModel.UserUpdate(ctx, updateUserInfo.Name, updateUserInfo.NewPassword, updateUserInfo.Role)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "操作失败！")
		return

	}

	middlewares.ResponseSuccess(ctx, "操作成功！", res)
}

func (u UserManagerRequest) GetUser(ctx *webCtx.Context) {
	//var user UserStruct
	var p model.Page
	if err := ctx.ShouldBind(&p); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}
	if p.PageSize == 0 || p.PageIndex == 0 {
		err := exterror.ErrParameterInvalid
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "未设置页码以每页大小！")
		return

	}
	res, err := model.DefaultUserModel.GetUser(ctx, p.PageIndex, p.PageSize)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "增加用户失败！")
		return

	}

	middlewares.ResponseSuccess(ctx, "添加用户成功", res)
}

func (u UserManagerRequest) GetAccount(ctx *webCtx.Context) {
	//var user UserStruct
	//var p model.Page
	var p string
	if err := ctx.ShouldBind(&p); nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "参数绑定失败！")
		return
	}
	_, err := model.DefaultUserModel.GetAccount(ctx, p)
	if nil != err {
		middlewares.ResponseError(ctx, http.StatusBadRequest, err, "生成账户失败！")
		return

	}

	middlewares.ResponseSuccess(ctx, "添加用户成功", nil)
}
