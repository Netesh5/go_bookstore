package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func parseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err != nil {
		if err := json.Unmarshal(body, x); err != nil {
			return
		}
	}

}
