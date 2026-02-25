package ebzrest

import (
	"database/sql"
	"net/http"
)

type RESTFul struct {
	db *sql.DB
}

func New(mux *http.ServeMux, db *sql.DB) *http.ServeMux {
	rest := RESTFul{
		db: db,
	}

	mux.HandleFunc("POST /tball/csv", rest.TBallUploadCSV)
	mux.HandleFunc("GET /tball/draw/frequency", rest.TBallDrawFrequencies)
	mux.HandleFunc("GET /tball/tball/frequency", rest.TBallFrequencies)

	mux.HandleFunc("POST /euro/csv", rest.EuroUploadCSV)
	mux.HandleFunc("GET /euro/draw/frequency", rest.EuroDrawFrequencies)
	mux.HandleFunc("GET /euro/star/frequency", rest.EuroStarFrequencies)

	mux.HandleFunc("POST /lotto/csv", rest.LottoUploadCSV)
	mux.HandleFunc("GET /lotto/draw/frequency", rest.LottoDrawFrequencies)
	mux.HandleFunc("GET /lotto/bonus/frequency", rest.LottoBonusFrequencies)

	mux.HandleFunc("POST /sflife/csv", rest.SFLifeUploadCSV)
	mux.HandleFunc("GET /sflife/draw/frequency", rest.SFLifeDrawFrequencies)
	mux.HandleFunc("GET /sflife/lball/frequency", rest.SFLifeLBallFrequencies)

	return mux
}
