package startup_test

import (
	"testing"

	"github.com/ethanjmarchand/ezquote/internal/startup"
)

func TestLoadEnv(t *testing.T) {
	cfg, err := startup.LoadEnv("../testdata/.env.test")
	if err != nil {
		t.Error("There was an error loading the env file")
	}
	if cfg.APIKey == "" {
		t.Error("APIKey cannot be blank")
	}
	if cfg.APIURL == "" {
		t.Error("APIURL cannot be blank")
	}
	if cfg.Server == "" {
		t.Error("server address cannot be blank")
	}
}
