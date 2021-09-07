package api

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/RedAFD/mega/internal/config"
	"github.com/RedAFD/mega/internal/core/context"
	"github.com/RedAFD/mega/internal/modules/user"
	"github.com/RedAFD/mega/internal/modules/user/model"
	"github.com/RedAFD/mega/internal/utils/i18n"
	"github.com/RedAFD/mega/internal/utils/logger"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/crypto/bcrypt"
)

type LoginReq struct {
	Email    string `json:"email" validate:"required" example:"radixholms@gmail.com"`
	Password string `json:"password" validate:"required" example:"123456"`
}

// Login godoc
// @Tags user
// @Accept json
// @Schemes https
// @Param body body LoginReq true "params"
// @Router /user/login [post]
// @Success 200 "成功"
// @Failure 400 "请求数据有误"
// @Failure 403 "验证失败"
// @Failure 429 "服务器繁忙，请稍后再试"
// @Failure 500 "服务器错误，请重新尝试/发生未知错误"
// @Header 200 {string} Set-Cookie "SessionID=54ac5448822e225a5a4656787074a4dc; expires=Sun, 05 Sep 2021 12:46:31 GMT"
func Login(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("login panic: %v", r)
			ctx.SetRespCode(http.StatusInternalServerError, i18n.Sprintf("发生未知错误"))
		}
	}()
	var param LoginReq
	if err := jsoniter.Unmarshal(ctx.GetReqBody(), &param); err != nil {
		logger.Error("failed to decode body: %v", err)
		ctx.SetRespCode(http.StatusBadRequest, i18n.Sprintf("请求数据有误"))
		return
	}
	mUser := &model.User{Email: param.Email}
	if ok, _ := mUser.FindUserByEmail(); !ok {
		ctx.SetRespCode(http.StatusForbidden, i18n.Sprintf("验证失败"))
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(mUser.Password), []byte(param.Password)); err != nil {
		ctx.SetRespCode(http.StatusForbidden, i18n.Sprintf("验证失败"))
		return
	}
	b := make([]byte, 16)
	binary.BigEndian.PutUint64(b, mUser.ID)
	binary.BigEndian.PutUint64(b[8:], uint64(time.Now().UnixNano()))
	newSessionID := fmt.Sprintf("%x", md5.Sum(b))
	cookieExpire := time.Now().Add(config.SessionExpiration)
	if err := user.SetSession(newSessionID, []byte(strconv.FormatUint(mUser.ID, 10)), config.SessionExpiration); err != nil {
		logger.Error("failed to set session: %v", err)
		ctx.SetRespCode(http.StatusInternalServerError, i18n.Sprintf("服务器错误，请重新尝试"))
		return
	}
	ctx.SetRespCookie("SessionID", newSessionID, cookieExpire)
}
