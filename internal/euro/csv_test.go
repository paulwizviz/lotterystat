package euro

import (
	"bytes"
	"context"
	"fmt"
)

func Example_processEuroCSV() {
	input := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,DrawNumber
04-Apr-2023,10,16,31,33,50,3,8,"XCRG53171","",1621
a-Apr-2023,10,16,31,33,50,3,8,"XCRG53171","",1622
06-Apr-2023,b,18,28,34,47,5,10,"JBQS10867","",1623
07-Apr-2023,18,28,34,47,5,10,"JBQS10867","",1624
08-Apr-2023,16,18,28,34,47,5,10,"JBQS10867","",1625`)

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
	// {{2023-04-04 00:00:00 +0000 UTC Tuesday 10 16 31 33 50 3 8 XCRG53171  1621} <nil>}
	// {{0001-01-01 00:00:00 +0000 UTC Sunday 0 0 0 0 0 0 0   0} record on line: 3: invalid day format: improper day format}
	// {{0001-01-01 00:00:00 +0000 UTC Sunday 0 0 0 0 0 0 0   0} record on line: 4: invalid draw digit: strconv.Atoi: parsing "b": invalid syntax}
	// {{0001-01-01 00:00:00 +0000 UTC Sunday 0 0 0 0 0 0 0   0} record on line 5: wrong number of fields}
	// {{2023-04-08 00:00:00 +0000 UTC Saturday 16 18 28 34 47 5 10 JBQS10867  1625} <nil>}

}
