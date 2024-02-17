package handler

import (
	"authorization/util"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	util.ResponseWithProblem(w,200,"ok")
}
