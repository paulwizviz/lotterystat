package ebzrest

import (
	"encoding/json"
	"net/http"

	"github.com/paulwizviz/lotterystat/internal/csvops"
	"github.com/paulwizviz/lotterystat/internal/sflife"
)

// SFLifeUploadCSV handles the upload of a Set For Life CSV file and persists the draws.
func (r RESTFul) SFLifeUploadCSV(rw http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	recs := csvops.ExtractRec(req.Context(), file)
	drawChans := sflife.ProcessCSV(recs, 1)

	for _, dc := range drawChans {
		if dc.Err != nil {
			continue
		}
		_ = sflife.PersistsDraw(req.Context(), r.db, dc.Draw)
	}
	rw.WriteHeader(http.StatusAccepted)
}

// SFLifeDrawFrequencies returns the frequencies of Set For Life draw balls.
func (r RESTFul) SFLifeDrawFrequencies(rw http.ResponseWriter, req *http.Request) {
	freqs, err := sflife.CalculateBallFreq(req.Context(), r.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(freqs)
}

// SFLifeLBallFrequencies returns the frequencies of Set For Life life balls.
func (r RESTFul) SFLifeLBallFrequencies(rw http.ResponseWriter, req *http.Request) {
	freqs, err := sflife.CalculateLBallFreq(req.Context(), r.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(freqs)
}
