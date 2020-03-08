package spaceapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//Schema is the go representation of the Spaceapi schema
type Schema struct {
	API      string `json:"api"`
	Space    string `json:"space"`
	Logo     string `json:"logo"`
	URL      string `json:"url"`
	Location struct {
		Address string  `json:"address"`
		Lat     float64 `json:"lat"`
		Lon     float64 `json:"lon"`
	} `json:"location"`
	Contact struct {
		Phone    string `json:"phone"`
		Irc      string `json:"irc"`
		Twitter  string `json:"twitter"`
		Facebook string `json:"facebook"`
		Email    string `json:"email"`
	} `json:"contact"`
	IssueReportChannels []string `json:"issue_report_channels"`
	Cache               struct {
		Schedule string `json:"schedule"`
	} `json:"cache"`
	Projects []string `json:"projects"`
	Feeds    struct {
		Calendar struct {
			URL  string `json:"url"`
			Type string `json:"type"`
		} `json:"calendar"`
	} `json:"feeds"`
	State struct {
		ExtLockstate string `json:"ext_lockstate"`
		Open         bool   `json:"open"`
		Lastchange   int    `json:"lastchange"`
		Icon         struct {
			Open   string `json:"open"`
			Closed string `json:"closed"`
		} `json:"icon"`
	} `json:"state"`
	Sensors struct {
		Temperature []struct {
			Location string  `json:"location"`
			Unit     string  `json:"unit"`
			Value    float64 `json:"value"`
		} `json:"temperature"`
		Humidity []struct {
			Location string  `json:"location"`
			Unit     string  `json:"unit"`
			Value    float64 `json:"value"`
		} `json:"humidity"`
	} `json:"sensors"`
}

//GetSchemafromURL loads a SpaceAPI URL and delivers back a Schmea struct
func GetSchemafromURL(url string) (Schema, error) {
	var result Schema

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return result, errors.New("Error creating request:" + err.Error())
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, errors.New("Error executing request: " + err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		data, _ := ioutil.ReadAll(resp.Body)
		return result, errors.New("Negativ status code:" + string(data))
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return result, errors.New("Unable to decode on Schema: " + err.Error())
	}
	return result, nil
}
