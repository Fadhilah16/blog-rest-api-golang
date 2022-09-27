package utils

import (
	"encoding/json"
	"net/http"
)

func EncodeJson(w http.ResponseWriter, x interface{}, status int) {
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(status)

	jsonEncode := json.NewEncoder(w)
	jsonEncode.Encode(x)

}
