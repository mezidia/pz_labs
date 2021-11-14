package common

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

//return class client
type Client struct {
	BaseURL string
}

//returns client
func client(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
	}
}

// client has methods get and post
func (s *Client) get(endpoint string) (string, error) {
	url := fmt.Sprintf("%s%s", s.BaseURL, endpoint)
	fmt.Println(url)

	//create get request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	client := &http.Client{}

	res, err := client.Do(req)
  if err != nil {
		return "", err
 }
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if 200 != res.StatusCode {
		return "", fmt.Errorf("%s", body)
	}

	resultString := string(body)
	fmt.Println(resultString)

	return resultString, nil
}

