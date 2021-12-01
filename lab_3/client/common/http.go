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
func (c *Client) Get(endpoint string) ([]dto.ForumsResponse, error) {
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

	//decode answer if no error
	body := []dto.ForumsResponse{}
	err = json.NewDecoder(res.Body).Decode(&body)

	if err != nil {
		return nil, err
	}
	return body, nil
}

// composing post request
func (c *Client) Post(endpoint string, userInfo *dto.User) (int, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)
	j, err := json.Marshal(userInfo)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return 1, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 1, err
	}
	defer res.Body.Close()

	//decode answer if no error
	statusCode := res.StatusCode;

	if err != nil {
		return 1, err
	}
	return statusCode, nil
}
