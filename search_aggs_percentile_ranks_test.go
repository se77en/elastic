package elastic

import (
	"encoding/json"
	"testing"
)

func TestPercentileRanksAggregation(t *testing.T) {
	agg := NewPercentileRanksAggregation().Field("load_time")
	data, err := json.Marshal(agg.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"percentile_ranks":{"field":"load_time"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}

func TestPercentileRanksAggregationWithCustomValues(t *testing.T) {
	agg := NewPercentileRanksAggregation().Field("load_time").Values(15, 30)
	data, err := json.Marshal(agg.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"percentile_ranks":{"field":"load_time","values":[15,30]}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
