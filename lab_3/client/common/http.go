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
func (c *Client) Get(endpoint string) (string, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	return c.doRequest(req)
}

// composing post request
func (c *Client) Post(endpoint string, userInfo *dto.User) (string, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)
	j, err := json.Marshal(userInfo)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return "", err
	}

	return c.doRequest(req)
}

//sending and processing request
func (c *Client) doRequest(req *http.Request) (string, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if 200 != res.StatusCode {
		return "", fmt.Errorf("Status code not equal to 200")
	}

	//decode answer if no error
	var body string
	// b := []dto.Forum{}
	err = json.NewDecoder(res.Body).Decode(&body) //.Decode(&b)

	if err != nil {
		return "", err
	}
	return body, nil
}
