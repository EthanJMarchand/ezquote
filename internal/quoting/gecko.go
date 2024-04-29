package quoting

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ethanjmarchand/ezquote/internal/startup"
)

var (
	ErrInvalidCoinID = errors.New("invalid coin id")
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type GeckoWorker struct {
	Config    *startup.Config
	NetClient Doer
}

type Coin struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// TODO: Add more context to failed messages
func (g GeckoWorker) GetAll() ([]Coin, error) {
	var cs []Coin
	req, err := http.NewRequest(http.MethodGet, g.Config.APIURL+"coins/list", nil)
	if err != nil {
		return nil, fmt.Errorf("getall(): %w", err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-cg-demo-api-key", g.Config.APIKey)
	res, err := g.NetClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("getall(): %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("getall(): %w", err)
	}
	err = json.Unmarshal(body, &cs)
	if err != nil {
		return nil, fmt.Errorf("getall(): %w", err)
	}
	return cs, nil
}

func (g GeckoWorker) GetQuote(s string) (*Quote, error) {
	var Quote Quote
	// TODO: utilize sprint instead of building the string
	req, err := http.NewRequest(http.MethodGet, g.Config.APIURL+"coins/"+s, nil)
	if err != nil {
		return nil, fmt.Errorf("getquote(): %w", err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-cg-demo-api-key", g.Config.APIKey)
	res, err := g.NetClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("getquote() failed to send get request: %w", err)
	}
	// TODO:  switch on the status code < 400 = Success, >= 400, try again in 2 minutes.
	if res.StatusCode != http.StatusOK {
		return nil, ErrInvalidCoinID
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("getquote() failed to readall: %w", err)
	}
	err = json.Unmarshal(body, &Quote)
	if err != nil {
		return nil, fmt.Errorf("getquote() failed to unmarshal: %w", err)
	}
	return &Quote, nil
}
