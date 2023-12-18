package bp128

import (
	"fmt"

	"github.com/prometheus/prometheus/tsdb/encoding"
)

type PrometheusPostingsEncoder struct {
}

func (b *PrometheusPostingsEncoder) Encode(e *encoding.Encbuf, offs []uint32) {
	packedInts := DeltaPack(offs)
	encodedInts, err := packedInts.GobEncode()
	if err != nil {
		panic(fmt.Sprintf("failed to encode ints: %v", err))
	}
	e.PutBytes(encodedInts)
}
