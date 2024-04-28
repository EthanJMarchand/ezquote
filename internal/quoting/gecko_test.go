package quoting_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ethanjmarchand/ezquote/internal/quoting"
	"github.com/ethanjmarchand/ezquote/internal/startup"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type mockDoer struct {
	err error
	q   quoting.Quote
	r   int
}

func (m *mockDoer) Do(req *http.Request) (*http.Response, error) {
	if m.err != nil {
		return nil, m.err
	}

	if m.r == 0 {
		m.r = http.StatusOK
	}
	b, _ := json.Marshal(m.q)
	resp := &http.Response{
		StatusCode: m.r,
		Body:       io.NopCloser(bytes.NewReader(b)),
		Header:     make(http.Header),
	}
	resp.Header.Set("Content-Type", "application/json")

	return resp, nil
}

func TestGeckoWorker_GetQuote(t *testing.T) {
	t.Run("given an error from /coins, get an error back.", func(t *testing.T) {

		g := quoting.GeckoWorker{
			NetClient: &mockDoer{err: errors.New("err_from_coins")},
			Config: &startup.Config{
				APIKey: "some-urk",
			},
		}

		_, err := g.GetQuote("some-coin")
		if err == nil {
			t.Fatal("expected an error, did not get one")
		}
		if !strings.Contains(err.Error(), "failed to send get request") {
			t.Fatalf("wrong error, got %s", err.Error())
		}
	})

	t.Run("get ErrInvalidCoin given a 400 from api", func(t *testing.T) {
		g := quoting.GeckoWorker{
			NetClient: &mockDoer{
				r: http.StatusBadRequest,
			},
			Config: &startup.Config{
				APIKey: "some-urk",
			},
		}

		_, err := g.GetQuote("some-coin")
		if err == nil {
			t.Fatal("expected an error, did not get one")
		}
		if !errors.Is(err, quoting.ErrInvalidCoinID) {
			t.Fatalf("expected error of type ErrInvalidCoin, got %s", reflect.TypeOf(err))
		}
	})

	t.Run("given a valid request, get a quote back", func(t *testing.T) {
		qID := "some-quote-id"
		g := quoting.GeckoWorker{
			NetClient: &mockDoer{
				q: quoting.Quote{
					ID: qID,
				},
			},
			Config: &startup.Config{
				APIKey: "some-urk",
			},
		}

		q, err := g.GetQuote("some-valid-coin")
		if err != nil {
			t.Fatalf("expected no error, but got %s", err.Error())
		}
		if q.ID != qID {
			t.Fatal("id not as expected")
		}
	})
}
