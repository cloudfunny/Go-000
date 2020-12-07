package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohuishou/go-training/Week02/work/domain"
	"github.com/mohuishou/go-training/Week02/work/pkg/errcode"
	"github.com/pkg/errors"
)

type api struct {
	usecase domain.IUserUsecase
	r       gin.IRouter
}

// NewUserHandler NewUserHandler
func NewUserHandler(r gin.IRouter, usecase domain.IUserUsecase) {
	api := &api{usecase: usecase}
	r.POST("/login", api.Login)
}

type loginParam struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Login Login
// 有很多可以优化的点，这里就不写那么完善了
func (api *api) Login(ctx *gin.Context) {
	var params loginParam
	if err := ctx.ShouldBind(&params); err != nil {
		errResp(ctx, errors.Wrapf(errcode.ErrParams, "should bind err: %+v", err))
		return
	}

	user, err := api.usecase.Login(params.UserName, params.Password)
	if err != nil {
		errResp(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": "success", "message": "登录成功", "data": user.ID})
}

func errResp(ctx *gin.Context, err error) {

	var code *errcode.ErrorCode
	if !errors.As(err, &code) {
		code = errcode.ErrUnKnown
	}

	ctx.Error(err)
	ctx.JSON(code.Status, code)
}
