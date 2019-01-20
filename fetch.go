package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Location is an alias to the location type.
type Location int

// LocationData holds visible information about locations.
type LocationData struct {
	// Title to present to the user.
	title string
	// Slug to use when fetching.
	Slug string
}

// The list of locations that have available feeds.
const (
	LocationAmsterdam Location = iota
	LocationBarcelona
	LocationBerlin
	LocationLondon
	LocationMadrid
	LocationParis
	LocationRemote
)

// Locations hold the metadata for each location that has an available feed.
var Locations = map[Location]LocationData{
	LocationAmsterdam: {"Amsterdam", "amsterdam"},
	LocationBarcelona: {"Barcelona", "barcelona"},
	LocationBerlin:    {"Berlin", "berlin"},
	LocationLondon:    {"London", "london"},
	LocationMadrid:    {"Madrid", "madrid"},
	LocationParis:     {"Paris", "paris"},
	LocationRemote:    {"Remote", "remoto"},
}

// TargetServer points to the HTTP server to use for fetching offers.
const TargetServer = "https://www.jobfluent.com"

// UserAgent keeps the user agent to be used in HTTP requests.
const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:66.0) Gecko/20100101 Firefox/66.0"

// unmarshalResponse will convert the bytearray content into offers.
func unmarshalResponse(resp []byte) ([]Offer, error) {
	var offers []Offer
	err := json.Unmarshal(resp, &offers)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse response offers: %s", err)
	}
	return offers, nil
}

func executeHTTPRequest(url string) ([]byte, error) {
	// Let's be honest and use a real client so we can set UA header.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Cannot prepare HTTP request: %s", err)
	}
	req.Header.Set("User-Agent", UserAgent)

	// Execute HTTP request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Cannot execute HTTP request: %s", err)
	}
	defer resp.Body.Close()

	// Extract data from the URL response.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot read HTTP response: %s", err)
	}

	return body, nil
}

// FetchOffers will send an HTTP request to the feed URL and retrieve the
// available offers found in the feed.  This function will retrieve every
// offer in the list.  It is up to the application to filter this list
// before presenting the results to the user.  There are multiple feeds,
func FetchOffers(location Location) ([]Offer, error) {
	// Get the real slug to use in the URL behind this location.
	if _, ok := Locations[location]; !ok {
		return nil, fmt.Errorf("Invalid location")
	}
	locData := Locations[location]

	// Fetch the jobs.
	url := fmt.Sprintf("%s/es/feeds/jobs-%s.json", TargetServer, locData.Slug)
	content, err := executeHTTPRequest(url)
	if err != nil {
		// err is properly wrapped by executeHttpRequest().
		return nil, err
	}

	// Offload the remains to the unmarshal process.
	return unmarshalResponse(content)
}
