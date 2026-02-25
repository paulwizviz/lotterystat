package ebzrest

import (
	"encoding/json"
	"net/http"

	"github.com/paulwizviz/lotterystat/internal/csvops"
	"github.com/paulwizviz/lotterystat/internal/tball"
)

// TBallUploadCSV handles the upload of a Thunderball CSV file and persists the draws.
func (r RESTFul) TBallUploadCSV(rw http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	recs := csvops.ExtractRec(req.Context(), file)
	drawChans := tball.ProcessCSV(recs, 1)

	for _, dc := range drawChans {
		if dc.Err != nil {
			continue
		}
		_ = tball.PersistsDraw(req.Context(), r.db, dc.Draw)
	}
	rw.WriteHeader(http.StatusAccepted)
}

// TBallDrawFrequencies returns the frequencies of Thunderball draw balls.
func (r RESTFul) TBallDrawFrequencies(rw http.ResponseWriter, req *http.Request) {
	freqs, err := tball.CalculateBallFreq(req.Context(), r.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(freqs)
}

// TBallFrequencies returns the frequencies of Thunderball thunderballs.
func (r RESTFul) TBallFrequencies(rw http.ResponseWriter, req *http.Request) {
	freqs, err := tball.CalculateTBallFreq(req.Context(), r.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(freqs)
}
