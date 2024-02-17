package jwt

import (
	"authorization/conf"
	"authorization/logger"
	"authorization/util"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/thedevsaddam/gojsonq/v2"
)

const (
	AuthorizationHeader = "authorization"
	AuthorizationPrefix = "Bearer "
	Authorities         = "authorities"
)

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get(AuthorizationHeader)
		if tokenStr == "" || !strings.HasPrefix(tokenStr, AuthorizationPrefix) {
			logger.Error.Printf("未携带token。url[%s]\n", r.RequestURI)
			util.ResponseWithProblem(w, http.StatusUnauthorized, "Unauthorized")
		} else {
			tokenStr = tokenStr[7:]
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					logger.Error.Printf("无效token[%s]。url[%s]\n", tokenStr, r.RequestURI)
					util.ResponseWithProblem(w, http.StatusUnauthorized, "Unauthorized")
					return nil, fmt.Errorf("not authorization")
				}
				return []byte(conf.AuthConfig.SecretKey), nil
			})
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				usernameNode := conf.AuthConfig.AuthProxyUsernameJsonNode
				if usernameNode == "" {
					logger.Access.Printf("放行。url[%s]\n", r.RequestURI)
					next.ServeHTTP(w, r)
				} else {
					subject := claims["sub"]
					content, _ := io.ReadAll(r.Body)
					username := gojsonq.New().FromString(string(content)).Find(usernameNode)
					if username == nil || strings.Compare(username.(string), subject.(string)) == 0 {
						logger.Access.Printf("用户%s验证通过，放行。url[%s]\n", subject, r.RequestURI)
						next.ServeHTTP(w, r)
					} else {
						logger.Error.Printf("参数中的用户名[%s]和令牌中[%s]不一致。url[%s]\n", username, subject, r.RequestURI)
						util.ResponseWithProblem(w, http.StatusUnauthorized, "ValidationErrorUser")
					}
				}
			} else if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					logger.Error.Printf("格式错误。url[%s],token[%s]\n", r.RequestURI, tokenStr)
					util.ResponseWithProblem(w, http.StatusUnauthorized, "ValidationErrorMalformed")
				} else if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
					logger.Error.Printf("已过期[%v]。url[%s]\n", token.Claims, r.RequestURI)
					util.ResponseWithProblem(w, http.StatusUnauthorized, "ValidationErrorExpired")
				} else if ve.Errors&(jwt.ValidationErrorNotValidYet|jwt.ValidationErrorIssuedAt) != 0 {
					logger.Error.Printf("未生效[%v]。url[%s]\n", token.Claims, r.RequestURI)
					util.ResponseWithProblem(w, http.StatusUnauthorized, "ValidationErrorNotValidYet")
				} else if ve.Errors&jwt.ValidationErrorIssuer != 0 {
					logger.Error.Printf("不信任的签发机构：%v。url[%s]\n", token.Claims, r.RequestURI)
					util.ResponseWithProblem(w, http.StatusUnauthorized, "ValidationErrorIssuer")
				} else {
					logger.Error.Printf("其他异常:%v。url[%s]\n", err, r.RequestURI)
					util.ResponseWithProblem(w, http.StatusUnauthorized, "Unauthorized")
				}
			} else {
				logger.Error.Printf("验证失败:%s。url[%s]\n", tokenStr, r.RequestURI)
				util.ResponseWithProblem(w, http.StatusUnauthorized, "Unauthorized")
			}
		}
	})
}

func GenerateToken(username string) (string, error) {
	type CustomClaims struct {
		Authorities string `json:"authorities"`
		jwt.StandardClaims
	}
	t := time.Now().Unix()

	claims := CustomClaims{
		"",
		jwt.StandardClaims{
			Subject:   username,
			IssuedAt:  t,
			ExpiresAt: t + conf.AuthConfig.ValidityInMilliseconds,
			Issuer:    conf.AuthConfig.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.AuthConfig.SecretKey))
}

func WriteToken(w http.ResponseWriter, username string) {
	token, _ := GenerateToken(username)
	w.Header().Add(AuthorizationHeader, AuthorizationPrefix+token)
}

func SetAuthorizationHeader(response *http.Response, username string) {
	token, _ := GenerateToken(username)
	response.Header.Add(AuthorizationHeader, AuthorizationPrefix+token)
}
