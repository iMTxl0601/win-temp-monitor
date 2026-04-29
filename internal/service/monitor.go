package service

import (
	"net/http"
	"time"
)

func waitLHMReady() {
	for i := 0; i < 10; i++ {
		resp, err := http.Get("http://localhost:8085/data.json")
		if err != nil && resp.StatusCode == 200 {
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}
