package authx

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
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
		return errors.New("没有该接口")
	}

	var roleId int64

	switch tokenData.TokenType {
	case common.UserTokenType:
		// 判断是否存在用户
		user, err := authenticationRpc.SysUserFindOne(r.Context(), &authenticationclient.SysUserFindOneReq{
			Id: tokenData.Uid, // 用户ID
		})
		if err != nil {
			return err
		}
		// 判断用户状态
		switch user.State {
		case 2:
			return fmt.Errorf("此令牌已停用,请联系管理员:%s", tokenData.NickName)
		case 3:
			return fmt.Errorf("此令牌已封禁,请联系管理员:%s", tokenData.NickName)
		}

		roleId = user.RoleId // 用户角色ID
		// 取用户角色ID
		if roleId == 0 {
			return fmt.Errorf("该用户没有角色")
		}

		// 角色信息
		role, err := authenticationRpc.SysRoleFindOne(r.Context(), &authenticationclient.SysRoleFindOneReq{
			Id: roleId, // 角色ID
		})
		if err != nil {
			return err
		}

		// 判断角色状态
		if role.RoleType != 1 && role.RoleType != 2 {
			return fmt.Errorf("该用户角色分配存在问题")
		}

	case common.AuthTokenType:
		// 查询第三方token 状态 和判断是否存在
		auth, err := authenticationRpc.SysAuthFindOne(r.Context(), &authenticationclient.SysAuthFindOneReq{
			Id: tokenData.Uid,
		})

		if err != nil {
			return err
		}

		roleId = auth.RoleId // 用户第三方角色ID

		// 取第三方角色ID
		if roleId == 0 {
			return fmt.Errorf("该第三方没有角色")
		}

		// 判断第三方用户状态
		switch auth.State {
		case 2:
			return fmt.Errorf("此令牌已停用,请联系管理员:%s", tokenData.NickName)
		case 3:
			return fmt.Errorf("此令牌已封禁,请联系管理员:%s", tokenData.NickName)
		}
		// 取第三方用户角色ID
		role, err := authenticationRpc.SysRoleFindOne(r.Context(), &authenticationclient.SysRoleFindOneReq{
			Id: roleId, // 角色ID
		})
		if err != nil {
			return err
		}
		// 判断角色状态
		if role.RoleType != 3 {
			return fmt.Errorf("该第三方角色分配存在问题")
		}
	}

	// 获取角色权限
	interfaceIds, err := authenticationRpc.SysInterfaceByRoleIdRespIDs(r.Context(), &authenticationclient.SysInterfaceByRoleIdReq{
		RoleId: roleId,
	})

	// 判断是否有权限
	is := common.IsAvailableInt64(reqRes.Id, interfaceIds.Ids)
	if !is {
		return errors.New("该用户没有该权限")
	}

	return nil

}
