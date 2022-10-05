package uptycs

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"net/http"
	"time"
)

type Config struct {
	Host       string
	APIKey     string
	APISecret  string
	CustomerID string
}

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

func CreateToken(apiKey string, apiSecret string) (string, error) {
	var err error
	currentTime := time.Now()
	atClaims := jwt.MapClaims{}
	atClaims["iss"] = apiKey
	atClaims["iat"] = currentTime
	atClaims["exp"] = currentTime.AddDate(0, 0, 1).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(apiSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateConfig(config Config) (bool, error) {
	configKeys := []string{config.Host, config.APIKey, config.APISecret, config.CustomerID}
	for _, configVal := range configKeys {
		if len(configVal) == 0 {
			return false, errors.New("required config value not found")
		}
	}
	return true, nil
}

func NewClient(config Config) (*Client, error) {
	ValidateConfig(config) //nolint:errcheck

	c := Client{
		HTTPClient: &http.Client{},
		HostURL:    fmt.Sprintf("%s/public/api/customers/%s", config.Host, config.CustomerID),
	}

	c.Token, _ = CreateToken(config.APIKey, config.APISecret)

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	//TODO Support pagination
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(time.Millisecond*30000),
	)

	defer cancel()
	req = req.WithContext(ctx)

	token := c.Token

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
