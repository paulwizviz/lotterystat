package sforl

import (
	"bytes"
	"context"
	"fmt"
)

func Example_processEuroCSV() {
	input := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Life Ball,Ball Set,Machine,DrawNumber
19-Jan-2023,42,30,47,40,15,8,SFL1,Excalibur 5,402
16-Jan-2023,36,10,23,40,32,10,SFL1,Excalibur 5,401
12-Jan-2023,23,30,31,25,24,9,SFL3,Excalibur 5,400
`)

	ecd := ProcessCSV(context.TODO(), bytes.NewReader(input))
	for d := range ecd {
		fmt.Println(d) // All draws will be displayed
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	ecd = ProcessCSV(ctx, bytes.NewReader(input))
	// The following step will not be called
	for d := range ecd {
		fmt.Println(d)
	}

	// Output:
	// {{2023-01-19 00:00:00 +0000 UTC Thursday 42 30 47 40 15 8 SFL1 Excalibur 5 402} <nil>}
	// {{2023-01-16 00:00:00 +0000 UTC Monday 36 10 23 40 32 10 SFL1 Excalibur 5 401} <nil>}
	// {{2023-01-12 00:00:00 +0000 UTC Thursday 23 30 31 25 24 9 SFL3 Excalibur 5 400} <nil>}

}
