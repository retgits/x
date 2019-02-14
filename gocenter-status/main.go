package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// GoCenter status API constants.
const (
	GoCenterStatusAPIRoot = `https://status.gocenter.io`
	StatusPath            = `/api/v2/status.json`
)

// Welcome is the human-readable form of the status API.
type Welcome struct {
	Page   Page   `json:"page"`
	Status Status `json:"status"`
}

// Page is the human-readable page element of the status API.
type Page struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	TimeZone  string `json:"time_zone"`
	UpdatedAt string `json:"updated_at"`
}

// Status is the human-readable status element of the status API.
type Status struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}

// LastStatus returns a request to the most recent status for GoCenter.
func LastStatus() *http.Request {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", GoCenterStatusAPIRoot, StatusPath), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")

	return req
}

func main() {

	req := LastStatus()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var w Welcome
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&w)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status for %s:\n%s\n", w.Page.Name, w.Status.Description)

	if w.Status.Indicator != "none" {
		os.Exit(1)
	}
}
