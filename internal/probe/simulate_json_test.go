package probe

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSimulateBatchResultMarshalsAndUnmarshals(t *testing.T) {
	t.Parallel()

	original := SimulateBatchResult{
		Items: []SimulateBatchItem{
			{
				Name: "transfer",
				Result: &SimulateResult{
					ContractID:  "C123",
					Function:    "transfer",
					Args:        []string{"addr:GABC", "i64:100"},
					Success:     true,
					ReturnValue: "ok",
					Cost: Cost{
						Instructions: 1234,
						ResourceFee:  56,
					},
					Footprint: Footprint{
						ReadOnly:  []string{"account:GABC"},
						ReadWrite: []string{"balance:GABC"},
					},
					LatestLedger: 987,
				},
			},
			{
				Name:  "withdraw",
				Error: "simulation failed",
			},
		},
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded SimulateBatchResult
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if !reflect.DeepEqual(decoded, original) {
		t.Fatalf("round trip mismatch:\noriginal: %#v\ndecoded:  %#v", original, decoded)
	}
}
