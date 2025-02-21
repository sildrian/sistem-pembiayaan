package library

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SetResponse(rsp Response, message string, emptyArr interface{}) Response {
	rsp.Message = message
	rsp.Data = emptyArr
	return rsp
}

func Res_400(w http.ResponseWriter, msg string) {
	var err Response
	err = SetResponse(err, msg, []string{})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(err)
	return
}

func Res_500(w http.ResponseWriter, msg string) {
	var err Response
	err = SetResponse(err, msg, []string{})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(err)
	return
}

func Res_Unknown(w http.ResponseWriter, msg string) {
	var err Response
	err = SetResponse(err, msg, []string{})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
	return
}

func Res_200(w http.ResponseWriter, msg string, data interface{}) {
	var rsp Response
	rsp = SetResponse(rsp, msg, data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rsp)
	return
}
