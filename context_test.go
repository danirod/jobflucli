package main

import (
	"testing"
	"time"
)

var offers []Offer = []Offer{
	Offer{
		id:           1000,
		URL:          "http://localhost:8080/offer/1000",
		Description:  "This is offer number 1",
		Company:      "ActionSoft Inc",
		Position:     "Rockstar Ninja",
		CreationDate: time.Now(),
		Tags:         []string{"alpha", "one"},
	},
	Offer{
		id:           2000,
		URL:          "http://localhost:8080/offer/2000",
		Description:  "This is offer number 2",
		Company:      "ActionRocks Inc",
		Position:     "Code Ops",
		CreationDate: time.Now(),
		Tags:         []string{"beta", "one", "two"},
	},
	Offer{
		id:           3000,
		URL:          "http://localhost:8080/offer/3000",
		Description:  "This is offer number 3",
		Company:      "ActionWare Inc",
		Position:     "Security Lead",
		CreationDate: time.Now(),
		Tags:         []string{"gamma", "two", "three"},
	},
}

func TestContextProperties(t *testing.T) {
	context := NewContext(offers)
	if context.CountOffers() != 3 {
		t.Errorf("CountOffers() did not return valid number")
	}

	cases := []int{1000, 2000, 3000}
	for _, c := range cases {
		if context.GetOffer(c).id != c {
			t.Errorf("GetOffer() did not return valid offer")
		}
	}
}

func TestContextIdIndex(t *testing.T) {
	cases := []struct {
		id   int
		want string
	}{
		{1000, "This is offer number 1"},
		{2000, "This is offer number 2"},
		{3000, "This is offer number 3"},
	}

	context := NewContext(offers)

	// Test that IDs do not appear by themselves in the index.
	if _, ok := context.idIndex[500]; ok {
		t.Errorf("Expected missing ID not to be found in the index")
	}
	// Test that offers are found.
	for _, c := range cases {
		expected, ok := context.idIndex[c.id]
		if !ok {
			t.Errorf("Expected an existing ID to be found in the index")
		}
		if expected.Description != c.want {
			t.Errorf("Offer found in the index is not valid: %s do not match %s", expected.Description, c.want)
		}
	}
}

func TestContextTagIndex(t *testing.T) {
	cases := []struct {
		tag  string
		want []int
	}{
		{"alpha", []int{1000}},
		{"beta", []int{2000}},
		{"gamma", []int{3000}},
		{"one", []int{1000, 2000}},
		{"two", []int{2000, 3000}},
	}

	context := NewContext(offers)

	for _, c := range cases {
		expected, ok := context.tagIndex[c.tag]
		if !ok {
			t.Errorf("Expected tag entry to be found in the index")
		}
		if len(c.want) != len(expected) {
			t.Errorf("Mismatching number of offers for tag index entry")
		}

	wantedTagLoop:
		for _, wantedID := range c.want {
			for _, expectedID := range expected {
				if expectedID == wantedID {
					continue wantedTagLoop
				}
			}
			t.Errorf("Expected offer ID to be found in the tag index")
		}
	}
}
