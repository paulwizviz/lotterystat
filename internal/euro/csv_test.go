package euro

import (
	"bytes"
	"context"
	"fmt"
)

func Example_processEuroCSV() {
	input := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,DrawNumber
29-Sep-2023,9,11,13,21,32,2,7,"HQSB24670",1672
26-Sep-2023,2,6,14,19,23,5,7,"VPRC26636",1671
22-Sep-2023,3,23,24,34,35,5,8,"HNRB16622",1670
19-Sep-2023,10,15,31,41,42,2,5,"JMQP30657",1669
15-Sep-2023,12,14,21,45,48,8,11,"HLQH38434,HLQJ62979,HLQK11974,TKPP96754,VKPN30889,VLQB24044,XKPN60194,XLPX34097,ZLPX51278,ZLPZ96812",1668`)

	ecd := processCSV(context.TODO(), bytes.NewReader(input))
	for d := range ecd {
		fmt.Println(d) // All draws will be displayed
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	ecd = processCSV(ctx, bytes.NewReader(input))
	// The following step will not be called
	for d := range ecd {
		fmt.Println(d)
	}

	// Output:
	// {{2023-09-29 00:00:00 +0000 UTC Friday 9 11 13 21 32 2 7 HQSB24670 1672} <nil>}
	// {{2023-09-26 00:00:00 +0000 UTC Tuesday 2 6 14 19 23 5 7 VPRC26636 1671} <nil>}
	// {{2023-09-22 00:00:00 +0000 UTC Friday 3 23 24 34 35 5 8 HNRB16622 1670} <nil>}
	// {{2023-09-19 00:00:00 +0000 UTC Tuesday 10 15 31 41 42 2 5 JMQP30657 1669} <nil>}
	// {{2023-09-15 00:00:00 +0000 UTC Friday 12 14 21 45 48 8 11 HLQH38434,HLQJ62979,HLQK11974,TKPP96754,VKPN30889,VLQB24044,XKPN60194,XLPX34097,ZLPX51278,ZLPZ96812 1668} <nil>}

}
