package utils

import (
	"encoding/json"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	jsonDecode := json.NewDecoder(r.Body)
	jsonDecode.Decode(x)
}
