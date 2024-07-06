package euro

import (
	"fmt"
	"paulwizviz/lotterystat/internal/csvutil"
)

// func processCSV(ctx context.Context, r io.Reader) <-chan DrawChan {
// 	c := make(chan DrawChan)
// 	go func() {
// 		cr := csv.NewReader(r)
// 		cr.Read() // remove titles
// 		ln := 1
// 		defer close(c)

// 	loop:
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				break loop
// 			default:
// 				ln++
// 				rec, err := cr.Read()
// 				if errors.Is(err, io.EOF) {
// 					break loop
// 				}
// 				if err != nil {
// 					c <- DrawChan{
// 						Draw: Draw{},
// 						Err:  err,
// 					}
// 					continue loop
// 				}
// 				drawDate, err := csvutil.ParseDateTime(rec[0])
// 				if err != nil {
// 					c <- DrawChan{
// 						Draw: Draw{},
// 						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
// 					}
// 					continue loop
// 				}
// 				b1, err := csvutil.ParseDrawNum(rec[1], 50)
// 				if err != nil {
// 					c <- DrawChan{
// 						Draw: Draw{},
// 						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
// 					}
// 					continue loop
// 				}
// 				b2, err := csvutil.ParseDrawNum(rec[2], 50)
// 				if err != nil {
// 					c <- DrawChan{
// 						Draw: Draw{},
// 						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
// 					}
// 					continue loop
// 				}
// 				b3, err := csvutil.ParseDrawNum(rec[3], 50)
// 				if err != nil {
// 					c <- DrawChan{
// 						Draw: Draw{},
// 						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
// 					}
// 					continue loop
// 				}
// 				b4, err := csvutil.ParseDrawNum(rec[4], 50)
// 				if err != nil {
// 					c <- DrawChan{
// 						Draw: Draw{},
// 						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
// 					}
// 					continue loop
// 				}
// 				b5, err := csvutil.ParseDrawNum(rec[5], 50)
// 				if err != nil {
// 					c <- DrawChan{
// 						Draw: Draw{},
// 						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
// 					}
// 					continue loop
// 				}
// 				ls1, err := csvutil.ParseDrawNum(rec[6], 12)
// 				if err != nil {
// 					c <- DrawChan{
// 						Draw: Draw{},
// 						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
// 					}
// 					continue loop
// 				}
// 				ls2, err := csvutil.ParseDrawNum(rec[7], 12)
// 				if err != nil {
// 					c <- DrawChan{
// 						Draw: Draw{},
// 						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
// 					}
// 					continue loop
// 				}
// 				dn, err := csvutil.ParseDrawSeq(rec[9])
// 				if err != nil {
// 					c <- DrawChan{
// 						Draw: Draw{},
// 						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
// 					}
// 					continue loop
// 				}
// 				c <- DrawChan{
// 					Draw: Draw{
// 						DrawDate:  drawDate,
// 						DayOfWeek: drawDate.Weekday(),
// 						Ball1:     uint8(b1),
// 						Ball2:     uint8(b2),
// 						Ball3:     uint8(b3),
// 						Ball4:     uint8(b4),
// 						Ball5:     uint8(b5),
// 						LS1:       uint8(ls1),
// 						LS2:       uint8(ls2),
// 						UKMarker:  rec[8],
// 						DrawNo:    dn,
// 					},
// 					Err: nil,
// 				}
// 			}
// 		}
// 	}()
// 	return c
// }

// func persistsCSV(ctx context.Context, db *sql.DB, nworkers int) error {
// 	r, err := csvutil.DownloadFrom(CSVUrl)
// 	if err != nil {
// 		return err
// 	}
// 	ch := processCSV(ctx, r)
// 	var wg sync.WaitGroup
// 	wg.Add(nworkers)
// 	for i := 0; i < nworkers; i++ {
// 		go func() {
// 			defer wg.Done()
// 			err := persistsDrawChan(ctx, db, ch)
// 			if err != nil {
// 				log.Println(err)
// 			}

// 		}()
// 	}
// 	wg.Wait()
// 	return nil
// }

func processCSV(recs chan csvutil.CSVRec) chan DrawChan {
	c := make(chan DrawChan)
	go func(ch chan DrawChan) {
		defer close(ch)
		for rec := range recs {
			dt, err := csvutil.ParseDateTime(rec.Record[0])
			if err != nil {
				ch <- DrawChan{
					Draw: Draw{},
					Err:  err,
				}
				continue
			}
			b1, err := csvutil.ParseDrawNum(rec.Record[1], 50)
			if err != nil {
				ch <- DrawChan{
					Draw: Draw{},
					Err:  err,
				}
				continue
			}
			b2, err := csvutil.ParseDrawNum(rec.Record[2], 50)
			if err != nil {
				ch <- DrawChan{
					Draw: Draw{},
					Err:  err,
				}
				continue
			}
			b3, err := csvutil.ParseDrawNum(rec.Record[3], 50)
			if err != nil {
				ch <- DrawChan{
					Draw: Draw{},
					Err:  err,
				}
				continue
			}
			b4, err := csvutil.ParseDrawNum(rec.Record[4], 50)
			if err != nil {
				ch <- DrawChan{
					Draw: Draw{},
					Err:  err,
				}
				continue
			}
			b5, err := csvutil.ParseDrawNum(rec.Record[5], 50)
			if err != nil {
				ch <- DrawChan{
					Draw: Draw{},
					Err:  err,
				}
				continue
			}
			s1, err := csvutil.ParseDrawNum(rec.Record[6], 12)
			if err != nil {
				ch <- DrawChan{
					Draw: Draw{},
					Err:  err,
				}
				continue
			}
			s2, err := csvutil.ParseDrawNum(rec.Record[7], 12)
			if err != nil {
				ch <- DrawChan{
					Draw: Draw{},
					Err:  err,
				}
				continue
			}
			ukMarker := rec.Record[8]
			dn, err := csvutil.ParseDrawSeq(rec.Record[9])
			if err != nil {
				ch <- DrawChan{
					Draw: Draw{},
					Err:  fmt.Errorf("%w-%s", csvutil.ErrCSVInvalidDrawDigit, err.Error()),
				}
				continue
			}
			ch <- DrawChan{
				Draw: Draw{
					DrawDate:  dt,
					DayOfWeek: dt.Weekday(),
					Ball1:     b1,
					Ball2:     b2,
					Ball3:     b3,
					Ball4:     b4,
					Ball5:     b5,
					LS1:       s1,
					LS2:       s2,
					UKMarker:  ukMarker,
					DrawNo:    dn,
				},
			}
		}
	}(c)
	return c
}
