package ebzrest

import (
	"encoding/json"
	"net/http"

	"github.com/paulwizviz/lotterystat/internal/csvops"
	"github.com/paulwizviz/lotterystat/internal/lotto"
)

// LottoUploadCSV handles the upload of a Lotto CSV file and persists the draws.
func (r RESTFul) LottoUploadCSV(rw http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	recs := csvops.ExtractRec(req.Context(), file)
	drawChans := lotto.ProcessCSV(recs, 1)

	for _, dc := range drawChans {
		if dc.Err != nil {
			continue
		}
		_ = lotto.PersistsDraw(req.Context(), r.db, dc.Draw)
	}
	rw.WriteHeader(http.StatusAccepted)
}

// LottoDrawFrequencies returns the frequencies of Lotto draw balls.
func (r RESTFul) LottoDrawFrequencies(rw http.ResponseWriter, req *http.Request) {
	freqs, err := lotto.CalculateBallFreq(req.Context(), r.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(freqs)
}

// LottoBonusFrequencies returns the frequencies of Lotto bonus balls.
func (r RESTFul) LottoBonusFrequencies(rw http.ResponseWriter, req *http.Request) {
	freqs, err := lotto.CalculateBonusFreq(req.Context(), r.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(freqs)
}
