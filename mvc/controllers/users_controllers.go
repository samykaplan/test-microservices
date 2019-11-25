package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/samykaplan/test-microservices/mvc/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		if _, err := resp.Write([]byte("user_id must be a an integer\n")); err != nil {
			// log.Fatalln(err)
			panic(err)
		}

	}

	user, err := services.GetUser(userId)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		if _, err := resp.Write([]byte(err.Error())); err != nil {
			// log.Fatalln(err)
			// panic(err)
			return
		}
	}

	jsonValue, _ := json.Marshal(user)
	if _, err := resp.Write(jsonValue); err != nil {
		// log.Fatalln(err)
		panic(err)
	}
}
