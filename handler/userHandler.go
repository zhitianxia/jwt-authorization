package handler

import (
	"authorization/jwt"
	"authorization/util"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	token, _ := jwt.GenerateToken("testUserName")
	jwt.WriteToken(w, "testUserName")
	util.ResponseWithProblem(w, 200, token)
}
