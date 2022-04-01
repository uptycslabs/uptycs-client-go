package uptycs

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

var hmacSampleSecret []byte

func CreateToken(apiKey string, apiSecret string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["iss"] = apiKey
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(apiSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewClient() (*Client, error) {
	host := os.Getenv("UPTYCS_HOST")
	if len(host) == 0 {
		return &Client{}, errors.New("required env var UPTYCS_HOST not found")
	}

	customerId := os.Getenv("UPTYCS_CUSTOMER_ID")
	if len(customerId) == 0 {
		return &Client{}, errors.New("required env var UPTYCS_CUSTOMER_ID not found")
	}

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    fmt.Sprintf("https://%s/public/api/customers/%s", host, customerId),
	}

	apiKey := os.Getenv("UPTYCS_API_KEY")
	if len(apiKey) == 0 {
		return &Client{}, errors.New("required env var UPTYCS_API_KEY not found")
	}
	apiSecret := os.Getenv("UPTYCS_API_SECRET")
	if len(apiSecret) == 0 {
		return &Client{}, errors.New("required env var UPTYCS_API_SECRET not found")
	}

	c.Token, _ = CreateToken(apiKey, apiSecret)

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	token := c.Token

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
