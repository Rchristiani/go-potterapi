package potterapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//House type used to create houses!
type House struct {
	ID          string   `json:"_id"`
	Name        string   `json:"name"`
	Mascot      string   `json:"mascot"`
	HeadOfHouse string   `json:"headOfHouse"`
	HouseGhost  string   `json:"houseGhost"`
	Founder     string   `json:"founder"`
	School      string   `json:"school"`
	Members     []string `json:"members"`
	Values      []string `json:"values"`
	Colors      []string `json:"colors"`
}

//Member is a single memeber
type Member struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

//HouseWithMembers used to create a house full of members
type HouseWithMembers struct {
	ID          string   `json:"_id"`
	Name        string   `json:"name"`
	Mascot      string   `json:"mascot"`
	HeadOfHouse string   `json:"headOfHouse"`
	HouseGhost  string   `json:"houseGhost"`
	Founder     string   `json:"founder"`
	School      string   `json:"school"`
	Members     []Member `json:"members"`
	Values      []string `json:"values"`
	Colors      []string `json:"colors"`
}

//GetHouses is used to get all the houses
func (c *Client) GetHouses() ([]House, error) {
	res, err := http.Get(BaseURL + "houses?key=" + c.apiKey)

	var houses []House

	if err != nil {
		return houses, err
	}

	defer res.Body.Close()

	houseBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return houses, err
	}

	err = json.Unmarshal(houseBytes, &houses)

	if err != nil {
		return houses, err
	}

	return houses, nil
}

//GetHouseByID is used to get a house by an ID
func (c *Client) GetHouseByID(id string) ([]HouseWithMembers, error) {
	res, err := http.Get(BaseURL + "houses/" + id + "?key=" + c.apiKey)

	var house []HouseWithMembers

	if err != nil {
		return house, err
	}

	defer res.Body.Close()

	houseBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return house, err
	}

	err = json.Unmarshal(houseBytes, &house)

	if err != nil {
		return house, err
	}

	return house, err
}
