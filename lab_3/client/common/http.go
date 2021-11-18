package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/mezidia/pz_labs/lab_3/client/dto"
)

type Client struct {
	BaseURL string
}

// composing get request
func (c *Client) Get(endpoint string) ([]dto.Forum, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if 200 != res.StatusCode {
		return nil, fmt.Errorf("Status code not equal to 200")
	}

	//decode answer if no error
	body := []dto.Forum{}
	err = json.NewDecoder(res.Body).Decode(&body)

	if err != nil {
		return nil, err
	}
	return body, nil
}

// composing post request
func (c *Client) Post(endpoint string, userInfo *dto.User) (*dto.RegistrateUserResponse, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)
	j, err := json.Marshal(userInfo)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if 200 != res.StatusCode {
		return nil, fmt.Errorf("Status code not equal to 200")
	}

	//decode answer if no error
	body := &dto.RegistrateUserResponse{};
	err = json.NewDecoder(res.Body).Decode(&body)

	if err != nil {
		return nil, err
	}
	return body, nil
}
