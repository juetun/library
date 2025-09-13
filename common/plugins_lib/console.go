package plugins_lib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/app/app_start"
	"github.com/juetun/base-wrapper/lib/common"
	"github.com/juetun/library/common/app_param"
	"net/http"
	"strconv"
	"strings"
)

//客服后台接口权限验证
func AdminImportAuth(arg *app_start.PluginsOperate) (err error) {
	app_start.AdminNetHandlerFunc = append(app_start.AdminNetHandlerFunc, func(c *gin.Context) {
		//如果是内网访问接口,则无须验证
		if strings.HasPrefix(c.Request.URL.Path, fmt.Sprintf("/%v/%v", app_obj.App.AppName, app_obj.App.AppRouterPrefix.Intranet)) {
			c.Next()
			return
		}
		if c.Request != nil && c.Request.Header.Get("Connection") == "Upgrade" {
			c.Next()
			return
		}
		switch c.Request.Method {
		case http.MethodHead, http.MethodOptions, http.MethodConnect, http.MethodTrace, http.MethodPatch, "":
			c.Next()
			return
		}
		var (
			arg       = &app_param.ArgParamsUserHaveConsoleImport{ImportKey: strings.ReplaceAll(c.Request.URL.Path, "/", "#"), User: &app_param.RequestUser{}}
			resPermit *app_param.ResultConsoleHaveImportPermit
			uid       int64
			reqUser   *app_param.ResultUser
			e         error
			uidString = c.Request.Header.Get(app_obj.HttpUserHid)
		)

		if uid, e = strconv.ParseInt(uidString, 10, 64); uid <= 0 {
			c.JSON(http.StatusOK, common.NewHttpResult().
				SetCode(http.StatusUnauthorized).
				SetMessage("未登录"))
			c.Abort()
			return
		}
		if user, exists := c.Get(app_param.GinContentUser); exists {
			arg.User = user.(*app_param.RequestUser)
		}
		if arg.User.UUserHid <= 0 {
			if reqUser, e = app_param.GetResultUserByUid(fmt.Sprintf("%d", uid), nil); e != nil {
				return
			}
			arg.User.UUserHid = uid
			if resultUserItem, ok := reqUser.List[uid]; ok {
				arg.User.UPortrait = resultUserItem.Portrait
				arg.User.UNickName = resultUserItem.NickName
				arg.User.UUserName = resultUserItem.UserName
				arg.User.UGender = resultUserItem.Gender
				arg.User.UStatus = resultUserItem.Status
				arg.User.UAuthStatus = resultUserItem.AuthStatus
				arg.User.UShopId = resultUserItem.ShopId
				arg.User.UHaveDashboard = resultUserItem.HaveDashboard
				arg.User.UUserMobileIndex = resultUserItem.UserMobileIndex
				arg.User.UUserEmailIndex = resultUserItem.UserEmailIndex
				arg.User.URememberToken = resultUserItem.RememberToken
				arg.User.UMsgReadTimeCursor = resultUserItem.MsgReadTimeCursor
				arg.User.UHaveDashboard = resultUserItem.HaveDashboard
				arg.User.UExists = resultUserItem.Exists
				arg.User.UDialog = resultUserItem.MsgInfo
				arg.User.UIsMocking = resultUserItem.IsMocking
			}
			c.Set(app_param.GinContentUser, arg.User)
		}
		if resPermit, e = app_param.GetUserHaveConsoleImportPermit(arg); e != nil {
			c.AbortWithStatusJSON(http.StatusOK, common.NewHttpResult().
				SetCode(http.StatusForbidden).
				SetMessage(fmt.Sprintf("%v-%v(err:%v,uid:%v)", c.Request.URL.Path, c.Request.Method, e.Error(), arg.User.UUserHid)))
			return
		}
		if resPermit == nil {
			c.AbortWithStatusJSON(http.StatusOK, common.NewHttpResult().
				SetCode(http.StatusForbidden).
				SetMessage(fmt.Sprintf("%v-%v(err:响应异常,uid:%v)", c.Request.URL.Path, c.Request.Method, arg.User.UUserHid)))

			return
		}
		if resPermit.IsSuper { //如果是超级管理员
			c.Next()
			return
		}
		if resPermit.NotHavePermit {
			c.AbortWithStatusJSON(http.StatusOK, common.NewHttpResult().
				SetCode(resPermit.StatusCode).
				SetMessage(fmt.Sprintf("%v-%v(err:%v,uid:%v)", c.Request.URL.Path, c.Request.Method, resPermit.ErrorMsg, arg.User.UUserHid)))
			return
		}
		c.Next()
		return
	})
	return
}
