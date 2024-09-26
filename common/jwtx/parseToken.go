package jwtx

import (
	"context"
	"fmt"
)

func ParseToken(ctx context.Context) *TokenData {
	// 用户登录信息
	var uidInterface = ctx.Value("uid")
	var nickNameInterface = ctx.Value("nickName")
	var tokenTypeInterface = ctx.Value("tokenType")
	uid := fmt.Sprintf("%v", uidInterface)
	nickName := fmt.Sprintf("%v", nickNameInterface)
	tokenType := fmt.Sprintf(fmt.Sprintf("%v", tokenTypeInterface))
	return &TokenData{
		Uid:       uid,
		NickName:  nickName,
		TokenType: tokenType,
	}
}

func ParseTokenMap(data map[string]interface{}) *TokenData {
	// 用户登录信息
	var uidInterface = data["uid"]
	var nickNameInterface = data["nickName"]
	var tokenTypeInterface = data["tokenType"]
	uid := fmt.Sprintf("%v", uidInterface)
	nickName := fmt.Sprintf("%v", nickNameInterface)
	tokenType := fmt.Sprintf(fmt.Sprintf("%v", tokenTypeInterface))
	return &TokenData{
		Uid:       uid,
		NickName:  nickName,
		TokenType: tokenType,
	}
}

type TokenData struct {
	Uid       string `json:"uid"`
	NickName  string `json:"nickName"`  // 昵称
	TokenType string `json:"tokenType"` // 类型
}
