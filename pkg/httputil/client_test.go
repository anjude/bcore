package httputil

import (
	"fmt"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	baseURL := "https://api.example.com"
	timeout := 10 * time.Second
	userAgent := "MyApp/1.0"

	client := NewHTTPClient(baseURL, timeout, userAgent)

	// Example GET request
	var getResp map[string]interface{}
	getHeaders := map[string]string{"Authorization": "Bearer token"}
	err := client.Request("GET", "/endpoint", getHeaders, nil, &getResp)
	if err != nil {
		fmt.Println("GET Error:", err)
		return
	}
	fmt.Println("GET Response:", getResp)

	// Example POST request with JSON payload
	postPayload := map[string]interface{}{"key": "value"}
	var postResp map[string]interface{}
	postHeaders := map[string]string{"Content-Type": "application/json"}
	err = client.Request("POST", "/endpoint", postHeaders, postPayload, &postResp)
	if err != nil {
		fmt.Println("POST Error:", err)
		return
	}
	fmt.Println("POST Response:", postResp)

}
