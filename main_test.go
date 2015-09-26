package main

import (
	"net/http/httptest"
	"testing"
)

func TestHandlerResponse(t *testing.T) {
	rec := httptest.NewRecorder()
	handler(rec, nil)
	if rec.Code > 299 {
		t.Errorf("request failed with code: %d", rec.Code)
	}
}
