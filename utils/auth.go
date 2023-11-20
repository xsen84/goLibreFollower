package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Body struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Auth(user string, password string, region string) (string, error) {

	client := &http.Client{}
	client.Timeout = 5 * time.Second
	posturl := "https://api-" + region + ".libreview.io/llu/auth/login"

	body := Body{
		Email:    user,
		Password: password,
	}

	marshalled, err := json.Marshal(body)
	if err != nil {
		log.Print("impossible to marshall the body:", err)
		return "", err
	}

	req, err := http.NewRequest("POST", posturl, bytes.NewReader(marshalled))
	if err != nil {
		log.Print("Error creating HTTP request:", err)
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("product", "llu.ios")
	req.Header.Add("version", "4.7.0")

	resp, err := client.Do(req)
	if err != nil {
		//fmt.Println("Error sending HTTP request")
		return "", err
	}

	rbody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print("Error reading HTTP response body")
		return "", err
	}

	// Print the response body
	//fmt.Println(string(rbody))
	return string(rbody), nil
}
