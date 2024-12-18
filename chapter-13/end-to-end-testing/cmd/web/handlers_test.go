package main

import (
	"net/http"
	"testing"

	"github.com/ahnafms/learn-go/internal/assert"
)

func TestPing(t *testing.T) {
	app := newTestingApplication(t)

	ts := newTestServer(t, app.routes())

	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}
