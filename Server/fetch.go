package groupie

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

const URL = "https://groupietrackers.herokuapp.com/api"

// !Data structure for the response from the API!
type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	Creation     int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	ID    int                 `json:"id"`
	Dates map[string][]string `json:"datesLocations"`
}

func fetchData(endpoint string, target interface{}) error {
	resp, err := http.Get(URL + endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}

	return nil
}

func fetch(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}

	return nil
}

func fetchArtists() ([]Artists, error) {
	var artists []Artists
	err := fetchData("/artists", &artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func fetchDates(ID int) ([]Date, error) {
	var dates []Date
	sID := strconv.Itoa(ID)
	err := fetchData("/dates/"+sID, &dates)
	if err != nil {
		return nil, err
	}
	return dates, nil
}

func fetchRelations(ID int) (Relations, error) {
	var relations Relations
	sID := strconv.Itoa(ID)
	err := fetchData("/relation/"+sID, &relations)
	if err != nil {
		return Relations{}, err
	}
	return relations, nil
}
