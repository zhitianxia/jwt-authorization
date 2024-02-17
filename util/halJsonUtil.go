package util

import (
	"authorization/lang"
	"authorization/model"
	"encoding/json"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/hal+json")
	w.WriteHeader(status)
	w.Write(response)
}

func ResponseWithProblem(w http.ResponseWriter, status int, problemCode string) {
	detail := lang.Get(problemCode)
	w.Header().Add("Problem-Code",problemCode)
	problemModel := model.ProblemModel{Status:status,Title:problemCode,Detail:detail}
	ResponseWithJson(w, status, problemModel)
}
