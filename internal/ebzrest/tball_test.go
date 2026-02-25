package ebzrest_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/paulwizviz/lotterystat/internal/ebzrest"
	"github.com/paulwizviz/lotterystat/internal/sqlops"
	"github.com/paulwizviz/lotterystat/internal/tball"
	"github.com/stretchr/testify/assert"
)

func TestTBallHandlers(t *testing.T) {
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err := sqlops.CreateTables(context.TODO(), db, tball.CreateTableFn); err != nil {
		t.Fatal(err)
	}

	mux := http.NewServeMux()
	ebzrest.New(mux, db)

	// Test CSV Upload
	t.Run("Upload CSV", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", "tball.csv")
		assert.NoError(t, err)

		csvContent := `DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Thunderball,Ball Set,Machine,DrawNumber
28-Aug-2024,16,4,6,13,28,3,T6,Excalibur 1,3547
`
		_, err = io.WriteString(part, csvContent)
		assert.NoError(t, err)
		writer.Close()

		req := httptest.NewRequest("POST", "/tball/csv", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusAccepted, rr.Code)
	})

	// Test Ball Frequencies
	t.Run("Get Ball Frequencies", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/tball/draw/frequency", nil)
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		var freqs []tball.BallFrequency
		err := json.NewDecoder(rr.Body).Decode(&freqs)
		assert.NoError(t, err)
		assert.NotEmpty(t, freqs)
	})

	// Test TBall Frequencies
	t.Run("Get TBall Frequencies", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/tball/tball/frequency", nil)
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		var freqs []tball.TBallFrequency
		err := json.NewDecoder(rr.Body).Decode(&freqs)
		assert.NoError(t, err)
		assert.NotEmpty(t, freqs)
	})
}
