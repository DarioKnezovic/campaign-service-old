package controllers

import (
	"encoding/json"
	"net/http"
)

func testHome(rw http.ResponseWriter, r *http.Request)  {
	response := map[string]string{
		"message": "Hello you jebem ti krmka",
	}
	json.NewEncoder(rw).Encode(response)
}