package tests

import (
	"testing"

	"kafka-pipeline/internal/model"
	"kafka-pipeline/internal/sorter"
)

func TestSortByID(t *testing.T) {
	data := []model.Record{
		{ID: 3},
		{ID: 1},
		{ID: 2},
	}

	sorter.ByID(data)

	if data[0].ID != 1 {
		t.Error("sorting by ID failed")
	}
}

func TestSortByName(t *testing.T) {
	data := []model.Record{
		{Name: "c"},
		{Name: "a"},
		{Name: "b"},
	}

	sorter.ByName(data)

	if data[0].Name != "a" {
		t.Error("sorting by name failed")
	}
}

func TestSortByContinent(t *testing.T) {
	data := []model.Record{
		{Continent: "Europe"},
		{Continent: "Asia"},
		{Continent: "Africa"},
	}

	sorter.ByContinent(data)

	if data[0].Continent != "Africa" {
		t.Error("sorting by continent failed")
	}
}
