package web_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethanjmarchand/ezquote/internal/quoting"
	"github.com/ethanjmarchand/ezquote/internal/web"
)

type MockWorker struct {
	Err error
}

func (mw MockWorker) GetAll() ([]quoting.Coin, error) {
	if mw.Err != nil {
		return nil, mw.Err
	}
	return nil, nil
}

func (mw MockWorker) GetQuote(s string) (*quoting.Quote, error) {
	return nil, nil
}

func TestHomehandler_ServeHTTP(t *testing.T) {
	t.Run("testing home route", func(t *testing.T) {
		// TODO: Build the worker so you can test for an error.
		// Great test name: Given an incorrect coin, I get a 400 response.
		mw := MockWorker{}
		service := quoting.NewAPIService(mw)
		handler := web.HomeHandler{
			Service: *service,
		}
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		handler.ServeHTTP(w, r)
		resp := w.Result()
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error("could not read the response body")
		}
		want := "GetAll()"
		if string(body) != want {
			t.Errorf("want %q got %q", want, string(body))
		}
	})
	t.Run("testing Quote route", func(t *testing.T) {
		mw := MockWorker{}
		service := quoting.NewAPIService(mw)
		handler := web.HomeHandler{
			Service: *service,
		}
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/quote/test", nil)
		handler.ServeHTTP(w, r)
		resp := w.Result()
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error("could not read the response body")
		}
		want := "test"
		if string(body) != want {
			t.Errorf("want %q got %q", want, string(body))
		}
	})

}
