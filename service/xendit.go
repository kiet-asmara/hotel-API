package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func createInvoice(input XenditRequest) (XenditResponse, error) {
	url := "https://api.xendit.co/v2/invoices"

	postBody, _ := json.Marshal(input)
	reqBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest(http.MethodPost, url, reqBody)
	if err != nil {
		return XenditResponse{}, err
	}

	req.Header.Add("Authorization", "Basic "+basicAuth(os.Getenv("API_KEY_XENDIT"), "")) // password kosong
	req.Header.Set("Content-Type", "application/json")

	// send HTTP req w/default client
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return XenditResponse{}, err
	}

	defer res.Body.Close()

	// read body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return XenditResponse{}, err
	}

	var respBody XenditResponse

	json.Unmarshal(body, &respBody)
	if respBody.Invoice_url == "" {
		return XenditResponse{}, errors.New(string(body))
	}

	return respBody, nil
}

func getInvoice(invoiceID string) (XenditResponse, error) {
	url := fmt.Sprintf("https://api.xendit.co/v2/invoices/%s", invoiceID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return XenditResponse{}, err
	}

	req.Header.Add("Authorization", "Basic "+basicAuth(os.Getenv("API_KEY_XENDIT"), "")) // password kosong
	req.Header.Set("Content-Type", "application/json")

	// send HTTP req w/default client
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return XenditResponse{}, err
	}

	defer res.Body.Close()

	// read body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return XenditResponse{}, err
	}

	var respBody XenditResponse

	json.Unmarshal(body, &respBody)
	if respBody.Invoice_url == "" {
		return XenditResponse{}, errors.New(string(body))
	}

	return respBody, nil
}
