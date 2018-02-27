package controllers;
import (
	"strings"
	"errors"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)
type BaseController struct {
	beego.Controller;
}


func (c *BaseController) ParseToken()(t *jwt.Token, e *error){
	authString := c.Ctx.Input.Header("Authorization")
	errInputData := errors.New("error input data")
	errExpired := errors.New("err Expired")
	kv := strings.Split(authString," ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		return nil,&errInputData
	}
	tokenString := kv[1]
	token,err := jwt.Parse(tokenString,func(token *jwt.Token) (interface{},error){
		return []byte("vito"),nil
	})
	if err != nil {
		beego.Error("Parse token:",err)
		if ve,ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0{
				return nil ,&errInputData
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet)!=0{
				return nil,&errExpired
			} else {
				return nil, &errInputData
			}
		} else {
			return nil,&errInputData
		}
	}
	if !token.Valid {
		beego.Error("Token invalid:",tokenString)
		return nil,&errInputData
	}
	beego.Debug("Token:",token)
	return token,nil
}