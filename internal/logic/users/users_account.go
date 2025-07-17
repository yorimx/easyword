package users

import (
	"context"
	"easyword/internal/consts"
	"easyword/internal/dao"
	"easyword/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
)

type jwtClaims struct {
	Id       uint
	Username string
	jwt.RegisteredClaims
}

// 单词表中保存有uid字段，我们需要在logic/users包中封装一个GetUid函数提供uid。
func (u *Users) GetUid(ctx context.Context) (uint, error) {
	user, err := u.Info(ctx)
	if err != nil {
		return 0, err
	}
	return user.Id, nil

}

func (u *Users) Login(ctx context.Context, username, password string) (tokenString string, err error) {

	var user entity.Users //本页的user是数据库里的user的信息
	err = dao.Users.Ctx(ctx).Where("username", username).Scan(&user)
	if err != nil {
		return "", gerror.New("用户或密码错误")
	}

	if user.Id == 0 {
		return "", gerror.New("用户不存在")
	}

	// 将密码加密后与数据库中的密码进行对比
	if user.Password != u.encryptPassword(password) {
		return "", gerror.New("用户或密码错误")
	}
	//生成token
	uc := &jwtClaims{
		Id:       user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
			//限制token的时限
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	return token.SignedString([]byte(consts.JwtKey))
}

func (u *Users) Info(ctx context.Context) (user *entity.Users, err error) {
	tokenString := g.RequestFromCtx(ctx).Request.Header.Get("Authorization")
	tokenClaims, _ := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtKey), nil
	})

	if claims, ok := tokenClaims.Claims.(*jwtClaims); ok && tokenClaims.Valid {
		err = dao.Users.Ctx(ctx).Where("id", claims.Id).Scan(&user)
	}
	return
}
