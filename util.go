package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func BuildMsAuthorizeUrl(clientId, permission string) string {
	u, _ := url.Parse("https://login.microsoftonline.com/common/oauth2/v2.0/authorize")
	urlQuery := u.Query()
	urlQuery.Set("client_id", clientId)
	urlQuery.Set("redirect_uri", "http://localhost:11451/auth")
	urlQuery.Set("response_type", "code")
	urlQuery.Set("scope", permission)
	u.RawQuery = urlQuery.Encode()
	return u.String()
}

func MsTokenRequest(clientId, clientSecret, clientCode string) (*TokenResp, error) {
	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("redirect_uri", "http://localhost:11451/auth")
	data.Set("client_secret", clientSecret)
	data.Set("code", clientCode)
	data.Set("grant_type", "authorization_code")
	r, err := http.NewRequest(http.MethodPost,
		"https://login.microsoftonline.com/common/oauth2/v2.0/token",
		strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(r)
	if resp.StatusCode != http.StatusOK {
		errInfo := fmt.Sprintf("response http code %v", resp.StatusCode)
		return nil, errors.New(errInfo)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	tokenData := TokenResp{}
	err = json.Unmarshal(body, &tokenData)
	if err != nil {
		return nil, err
	}
	return &tokenData, nil
}
