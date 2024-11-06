package authx

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
	"tpmt-zt/common"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/authentication/authentication"
	"tpmt-zt/service/authentication/authenticationclient"
)

func Auth(r *http.Request, authenticationRpc authentication.Authentication, accessSecret string) error {
	// 解析token
	token := r.Header.Get("Authorization")
	token = strings.TrimSpace(strings.Replace(token, "Bearer", "", -10))
	var mapClaims jwt.MapClaims
	_, err := jwt.ParseWithClaims(token, &mapClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return errors.New(msg.TokenError)
	}
	// jwt转换
	tokenData := jwtx.ParseTokenMap(mapClaims)

	// 获取请求路径
	path := r.URL.Path

	// 查询请求路径信息
	reqRes, err := authenticationRpc.FindOneInterfaceByPathAndInterfaceType(r.Context(), &authenticationclient.FindOneInterfaceByPathAndInterfaceTypeReq{
		Path:          path,
		InterfaceType: r.Method,
	})

	if err != nil {
		return err
	}
	fmt.Println(reqRes)

	switch tokenData.TokenType {
	case common.UserTokenType:
		// 判断是否存在用户

		// 判断用户状态

		// 取用户角色ID

	case common.AuthTokenType:
		// 查询第三方token 状态 和判断是否存在
		auth, err := authenticationRpc.SysAuthFindOne(r.Context(), &authenticationclient.SysAuthFindOneReq{
			Id: tokenData.Uid,
		})

		if err != nil {
			return err
		}

		// 判断第三方用户状态
		switch auth.State {
		case 2:
			return fmt.Errorf("此令牌已停用,请联系管理员:%s", tokenData.NickName)
		case 3:
			return fmt.Errorf("此令牌已封禁,请联系管理员:%s", tokenData.NickName)
		}

		// 取第三方用户角色ID

	}

	// 获取角色权限

	// 判断是否有权限

	return nil

}
