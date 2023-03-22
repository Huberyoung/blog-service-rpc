package bapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	USERNAME = "hubery"
	PASSWORD = "123456"
)

type AccessToken struct {
	Token string `json:"token"`
}

type Api struct {
	URL string
}

func NewApi(url string) *Api {
	return &Api{url}
}

func (a *Api) httpGet(ctx context.Context, path string) ([]byte, error) {
	sprintf := fmt.Sprintf("%s/%s", a.URL, path)
	resp, err := http.Get(sprintf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (a Api) httpPost(url string, data url.Values) ([]byte, error) {
	resp, err := http.PostForm(a.URL+"/"+url, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (a *Api) getAccessToken() (string, error) {
	data := make(url.Values)
	data["username"] = []string{USERNAME}
	data["password"] = []string{PASSWORD}
	body, err := a.httpPost("user", data)
	if err != nil {
		return "", err
	}

	var accessToken AccessToken
	_ = json.Unmarshal(body, &accessToken)
	return accessToken.Token, nil
}

func (a *Api) GetTagList(ctx context.Context, name string) ([]byte, error) {
	token, err := a.getAccessToken()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("%s?token=%s&name=%s", "api/v1/tags", token, name)
	body, err := a.httpGet(ctx, path)
	if err != nil {
		return nil, err
	}
	return body, nil
}
