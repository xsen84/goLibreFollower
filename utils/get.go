package utils

import (
	"io"
	"log"
	"net/http"
	"time"
)

func GetReadings(authToken string) (string, error) {
	client := &http.Client{}
	client.Timeout = 10 * time.Second

	getUrl := "https://api-de.libreview.io/llu/connections"

	req, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		log.Print("Error creating HTTP request:", err)
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("product", "llu.ios")
	req.Header.Add("version", "4.7.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Print("Error sending HTTP request:", err)
		return "", err
	}

	rbody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print("Error reading HTTP response body:", err)
		return "", err
	}

	// Print the response body
	return string(rbody), nil
}
