package httpgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func Request(url, method string, body map[string]any, headers map[string]string) (*http.Response, int, error) {
	payload, err := json.Marshal(body)

	if err != nil {
		return nil, 0, fmt.Errorf("error marshalling request body: %v", err)
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		return nil, 0, fmt.Errorf("error creating request: %v", err)
	}
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
		fmt.Println("====================")
		fmt.Println("Request url:", url)
		fmt.Println("Request method:", method)

		if req.Body != nil {
			requestBody := req.Body
			var requestBodyBytes []byte
			if requestBody != nil {
				requestBodyBytes, _ = io.ReadAll(requestBody)
			}
			fmt.Println("Request body:", string(requestBodyBytes))
			req.Body = io.NopCloser(bytes.NewBuffer(requestBodyBytes))
		}
		fmt.Println("Request headers:")
		for key, value := range req.Header {
			for _, v := range value {
				fmt.Printf("%s:%s\n", key, v)
			}
		}
		fmt.Println("====================")
	}
	var resp *http.Response
	var retryErr error
	resp, retryErr = client.Do(req)
	if retryErr != nil {
		return nil, 0, fmt.Errorf("error sending request: %v", retryErr)
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, fmt.Errorf("received 401 status code: %d", resp.StatusCode)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, resp.StatusCode, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}
	return resp, resp.StatusCode, retryErr
}
