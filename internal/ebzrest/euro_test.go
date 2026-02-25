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
	"github.com/paulwizviz/lotterystat/internal/euro"
	"github.com/paulwizviz/lotterystat/internal/sqlops"
	"github.com/stretchr/testify/assert"
)

func TestEuroHandlers(t *testing.T) {
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err := sqlops.CreateTables(context.TODO(), db, euro.CreateTableFn); err != nil {
		t.Fatal(err)
	}

	mux := http.NewServeMux()
	ebzrest.New(mux, db)

	// Test CSV Upload
	t.Run("Upload CSV", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", "euro.csv")
		assert.NoError(t, err)

		csvContent := `DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,Ball Set,Machine,DrawNumber
20-Feb-2026,13,24,28,33,35,5,9,ZDTF34718,,21,13,1922
`
		_, err = io.WriteString(part, csvContent)
		assert.NoError(t, err)
		writer.Close()

		req := httptest.NewRequest("POST", "/euro/csv", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusAccepted, rr.Code)
	})

	// Test Ball Frequencies
	t.Run("Get Ball Frequencies", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/euro/draw/frequency", nil)
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		var freqs []euro.BallFrequency
		err := json.NewDecoder(rr.Body).Decode(&freqs)
		assert.NoError(t, err)
		assert.NotEmpty(t, freqs)
	})

	// Test Star Frequencies
	t.Run("Get Star Frequencies", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/euro/star/frequency", nil)
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		var freqs []euro.StarFrequency
		err := json.NewDecoder(rr.Body).Decode(&freqs)
		assert.NoError(t, err)
		assert.NotEmpty(t, freqs)
	})
}
