package api

import "github.com/ethanjmarchand/ezquote/internal/startup"

type GeckoWorker struct {
	Config *startup.Config
}

func (g GeckoWorker) GetAll() ([]byte, error) {
	return []byte("getall working"), nil
}

func (g GeckoWorker) GetQuote(s string) ([]byte, error) {
	return []byte("getquote working. Asked for: " + s), nil
}
