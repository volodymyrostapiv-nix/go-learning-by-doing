package main

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ps := httprouter.Params{}

	Index(w, req, ps)
	resp := w.Result()

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %q", err)
	}
	if string(data) != INDEX {
		t.Errorf("expected %q got %q", INDEX, string(data))
	}
}

func TestTarget(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ps := httprouter.Params{{Key: "target", Value: "whatever"}}

	Target(w, req, ps)
	resp := w.Result()

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusSeeOther {
		t.Errorf("expected 303 got %q", resp.StatusCode)
	}

}
