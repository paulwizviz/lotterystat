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
	"github.com/paulwizviz/lotterystat/internal/lotto"
	"github.com/paulwizviz/lotterystat/internal/sqlops"
	"github.com/stretchr/testify/assert"
)

func TestLottoHandlers(t *testing.T) {
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err := sqlops.CreateTables(context.TODO(), db, lotto.CreateTableFn); err != nil {
		t.Fatal(err)
	}

	mux := http.NewServeMux()
	ebzrest.New(mux, db)

	// Test CSV Upload
	t.Run("Upload CSV", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", "lotto.csv")
		assert.NoError(t, err)

		csvContent := `DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Ball 6,Bonus Ball,Ball Set,Machine,DrawNumber
18-Feb-2026,1,11,12,13,18,49,33,L10,Lotto4,3147
`
		_, err = io.WriteString(part, csvContent)
		assert.NoError(t, err)
		writer.Close()

		req := httptest.NewRequest("POST", "/lotto/csv", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusAccepted, rr.Code)
	})

	// Test Ball Frequencies
	t.Run("Get Ball Frequencies", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/lotto/draw/frequency", nil)
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		var freqs []lotto.BallFrequency
		err := json.NewDecoder(rr.Body).Decode(&freqs)
		assert.NoError(t, err)
		assert.NotEmpty(t, freqs)
	})

	// Test Bonus Frequencies
	t.Run("Get Bonus Frequencies", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/lotto/bonus/frequency", nil)
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		var freqs []lotto.BonusFrequency
		err := json.NewDecoder(rr.Body).Decode(&freqs)
		assert.NoError(t, err)
		assert.NotEmpty(t, freqs)
	})
}
