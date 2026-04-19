package tests

import (
	"kafka-pipeline/internal/model"
	"testing"
)

func TestRecordStructure(t *testing.T) {
	r := model.Record{
		ID:        1,
		Name:      "test",
		Address:   "addr",
		Continent: "Asia",
	}

	if r.ID != 1 {
		t.Error("invalid ID")
	}
}

func TestCSVFormat(t *testing.T) {
	// basic sanity test placeholder
	r := model.Record{ID: 1, Name: "a", Address: "b", Continent: "c"}

	if r.Name == "" {
		t.Error("name empty")
	}
}
