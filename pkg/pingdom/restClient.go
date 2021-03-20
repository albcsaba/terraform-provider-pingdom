package pingdom

import (
	"io"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Body       []byte
	StatusCode int
}

type Client interface {
	Do(path, method string, requestBody io.Reader) (*Response, error)
}

type RestClient struct {
	baseURL string
	token   string
}

func (client *RestClient) New(baseURL, token string) {
	client.baseURL = baseURL
	client.token = token
}

func (client RestClient) Do(path, method string, requestBody io.Reader) (*Response, error) {
	request, err := http.NewRequest(method, client.baseURL+path, requestBody)
	if err != nil {
		return nil, err
	}

	request.Header.Add("accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+client.token)

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		Body:       responseBody,
		StatusCode: response.StatusCode,
	}, nil
}
