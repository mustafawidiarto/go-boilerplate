package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func HandleHealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		timeNow := time.Now()
		resp, _ := json.Marshal(map[string]interface{}{
			"status":   http.StatusOK,
			"message":  "Service is running",
			"location": fmt.Sprintf("%v", timeNow.Location()),
			"time":     fmt.Sprintf("%v", timeNow),
		})

		w.Write([]byte(resp))
	}
}
