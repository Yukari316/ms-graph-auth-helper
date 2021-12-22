package main

import (
	"encoding/json"
	"os"
)

type TokenResult struct {
	ClientId     string `json:"client-id"`
	ClientCode   string `json:"client-code"`
	Secret       string `json:"secret"`
	AccessToken  string `json:"access-token"`
	RefreshToken string `json:"refresh-token"`
}

func (res *TokenResult) Save() (*[]byte, error) {
	file, err := os.Create("result.json")
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	if err != nil {
		return nil, err
	}
	jsonBytes, err := json.MarshalIndent(res, "","	")
	if err != nil {
		return nil, err
	}
 	_, err = file.Write(jsonBytes)
	return &jsonBytes, err
}
