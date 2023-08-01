package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

var jwtSecret = []byte("lgb24kcs")

type TokenInfo struct {
	Id int64
}

type MyClaims struct {
	TokenInfo TokenInfo `json:"token_info"`
	jwt.RegisteredClaims
}

/*生成token*/
func CreateToken(userId int64) (string, error) {

	tokeninfo := TokenInfo{
		Id: userId,
	}
	myClaims := MyClaims{
		TokenInfo: tokeninfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        strconv.Itoa(int(userId)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))),
			Issuer:    "local",
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	signedString, err := claims.SignedString(jwtSecret)
	if err != nil {
		//global.APP.Log.Error(err.Error())
		return "", err
	} else {

		// nowUnix := time.Now().Unix()

		// timer := time.Duration(myClaims.RegisteredClaims.ExpiresAt.Time.Unix()-nowUnix) * time.Second
		// fmt.Printf("timer: %v\n", timer)
		// // if b := RedisSet(strconv.Itoa(int(CtxAccountId)), nowUnix, timer); !b {
		// // 	return "", errors.New("token存储异常")
		// // }
		// err := consts.App.Redis.SetNX(context.Background(), strconv.Itoa(int(CtxAccountId)), nowUnix, timer).Err()
		// if err != nil {
		// 	fmt.Printf("err: %v\n", err)
		// 	return "", errors.New("token存储异常")
		// } else {

		// }

	}
	return signedString, nil
}

/*解析token*/
func ParseToken(token string) (*MyClaims, error) {
	tokenString, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
	// 要传入指针，项目中结构体都是用指针传递，节省空间。
	if claims, ok := tokenString.Claims.(*MyClaims); ok && tokenString.Valid {
		return claims, nil
	}
	return nil, err
}

// 更新token
//func RefreshToken(tokenStr string) string {
//	token, _ := ParseToken(tokenStr)
//	if token.ExpiresAt.Time.Unix() > time.Now().Unix() {
//		admin := model.AdminInfo{
//			ID:       token.TokenInfo.Id,
//			UserName: token.TokenInfo.UserName,
//			IsAdmin:  token.TokenInfo.IsAdmin,
//			RoleIds:  token.TokenInfo.RoleIds,
//		}
//		s, _ := CreateToken(admin)
//		return s
//	}
//	return ""
//}
