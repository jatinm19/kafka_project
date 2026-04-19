package tests

import (
	"testing"

	"kafka-pipeline/internal/model"
	"kafka-pipeline/internal/sorter"
)

func TestSortByID(t *testing.T) {
	data := []model.Record{
		{ID: 9},
		{ID: 1},
		{ID: 5},
	}

	sorter.ByID(data)

	if data[0].ID != 1 {
		t.Fail()
	}
}

func TestSortByName(t *testing.T) {
	data := []model.Record{
		{Name: "z"},
		{Name: "a"},
	}

	sorter.ByName(data)

	if data[0].Name != "a" {
		t.Fail()
	}
}
