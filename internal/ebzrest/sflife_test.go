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
	"github.com/paulwizviz/lotterystat/internal/sflife"
	"github.com/paulwizviz/lotterystat/internal/sqlops"
	"github.com/stretchr/testify/assert"
)

func TestSFLifeHandlers(t *testing.T) {
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err := sqlops.CreateTables(context.TODO(), db, sflife.CreateTableFn); err != nil {
		t.Fatal(err)
	}

	mux := http.NewServeMux()
	ebzrest.New(mux, db)

	// Test CSV Upload
	t.Run("Upload CSV", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", "sflife.csv")
		assert.NoError(t, err)

		csvContent := `DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Life Ball,Ball Set,Machine,DrawNumber
19-Feb-2026,5,9,13,34,45,8,SFL3,Excalibur6,724
`
		_, err = io.WriteString(part, csvContent)
		assert.NoError(t, err)
		writer.Close()

		req := httptest.NewRequest("POST", "/sflife/csv", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusAccepted, rr.Code)
	})

	// Test Ball Frequencies
	t.Run("Get Ball Frequencies", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/sflife/draw/frequency", nil)
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		var freqs []sflife.BallFrequency
		err := json.NewDecoder(rr.Body).Decode(&freqs)
		assert.NoError(t, err)
		assert.NotEmpty(t, freqs)
	})

	// Test LBall Frequencies
	t.Run("Get LBall Frequencies", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/sflife/lball/frequency", nil)
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		var freqs []sflife.LBallFrequency
		err := json.NewDecoder(rr.Body).Decode(&freqs)
		assert.NoError(t, err)
		assert.NotEmpty(t, freqs)
	})
}
