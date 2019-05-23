package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const name string = "sunmoon20190522"
const email string = "holyreaper28@gmail.com"

type jwtSunMoonClaim struct {
	jwt.StandardClaims
	Name  string `json:"name"`
	Email string `json:"email"`
}

var claim jwtSunMoonClaim
var token *jwt.Token

func init() {
	claim = jwtSunMoonClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 24 * 7).Unix()),
			Issuer:    "sunmoon",
		},
		Name:  name,
		Email: email,
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
}

//CreateToken 生成token
func CreateToken(str string) (string, error) {
	strtoken, err := token.SignedString([]byte(str))
	return strtoken, err
}

//CheckToken 检查token
func CheckToken(ctoken string, str string) (bool, error) {
	_token, err := jwt.Parse(ctoken, func(*jwt.Token) (interface{}, error) {
		return []byte(str), nil
	})
	if err != nil {
		return false, err
	}
	return _token.Claims.Valid() == nil, nil
}
