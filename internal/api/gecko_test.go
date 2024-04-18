package api_test

import (
	"testing"

	"github.com/ethanjmarchand/ezquote/internal/api"
)

func TestGeckoWorker_GetQuote(t *testing.T) {
	w := api.GeckoWorker{}
	arg := "test"
	got, err := w.GetQuote(arg)
	want := "getquote working. Asked for: " + arg
	if err != nil {
		t.Errorf("got err: %s, did not want an error", err.Error())
	}
	if string(got) != want {
		t.Errorf("Got %q, want %q", string(got), want)
	}
}
