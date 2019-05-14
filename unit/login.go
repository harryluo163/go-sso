package unit

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"go-sso/models"
	"strconv"
	"time"
)

func AuthenticateUserForLogin(loginname, password string) (*models.User, error) {
	if len(password) == 0 || len(loginname) == 0 {
		return nil, errors.New("Error: 用户名或密码为空")
	}
	//密码md5加密
	data := []byte(password)
	has := md5.Sum(data)
	password = fmt.Sprintf("%x", has) //将[]byte转成16进制
	var user *models.User
	o := orm.NewOrm()
	ids := []string{loginname, password}
	err := o.Raw("SELECT user_id,login_name FROM user where login_name= ? and user_password= ? ", ids).QueryRow(&user)
	if err != nil {
		return nil, errors.New("Error: 未找到该用户")
	} else {
		return user, nil
	}

}

func CreateToken(userName string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	//添加令牌关键信息
	Tokenexp, _ := strconv.Atoi(beego.AppConfig.String("Tokenexp"))
	//添加令牌期限
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(Tokenexp)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["userName"] = userName
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(beego.AppConfig.String("TokenSecrets")))
	return tokenString
}

// 校验token是否有效 返回参数
func CheckToken(tokenString string) string {
	userName := ""
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(beego.AppConfig.String("TokenSecrets")), nil
	})

	if token != nil && token.Valid {
		claims, _ := token.Claims.(jwt.MapClaims)
		userName = claims["userName"].(string)
	}
	return userName
}
