package ebzrest

import (
	"encoding/json"
	"net/http"

	"github.com/paulwizviz/lotterystat/internal/csvops"
	"github.com/paulwizviz/lotterystat/internal/euro"
)

// EuroUploadCSV handles the upload of a EuroMillions CSV file and persists the draws.
func (r RESTFul) EuroUploadCSV(rw http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	recs := csvops.ExtractRec(req.Context(), file)
	drawChans := euro.ProcessCSV(recs, 1)

	for _, dc := range drawChans {
		if dc.Err != nil {
			continue
		}
		_ = euro.PersistsDraw(req.Context(), r.db, dc.Draw)
	}
	rw.WriteHeader(http.StatusAccepted)
}

// EuroDrawFrequencies returns the frequencies of EuroMillions draw balls.
func (r RESTFul) EuroDrawFrequencies(rw http.ResponseWriter, req *http.Request) {
	freqs, err := euro.CalculateBallFreq(req.Context(), r.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(freqs)
}

// EuroStarFrequencies returns the frequencies of EuroMillions lucky stars.
func (r RESTFul) EuroStarFrequencies(rw http.ResponseWriter, req *http.Request) {
	freqs, err := euro.CalculateStarFreq(req.Context(), r.db)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(freqs)
}
