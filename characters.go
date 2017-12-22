package potterapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Character struct
type Character struct {
	ID                string `json:"_id"`
	BloodStatus       string `json:"bloodStatus"`
	DeathEater        bool   `json:"deathEater"`
	DumbledoresArmy   bool   `json:"dumbledoresArmy"`
	House             string `json:"house"`
	MinistryOfMagic   bool   `json:"ministryOfMagic"`
	Name              string `json:"name"`
	OrderOfThePhoenix bool   `json:"orderOfThePhoenix"`
	Role              string `json:"role"`
	School            string `json:"school"`
	Species           string `json:"species"`
	Wand              string `json:"wand"`
	Boggart           string `json:"boggart"`
	Alias             string `json:"alias"`
	Animagus          string `json:"animagus"`
}

//GetCharacters is used to get characters
func (c *Client) GetCharacters() ([]Character, error) {
	res, err := http.Get(BaseURL + "characters?key=" + c.apiKey)

	var characters []Character

	if err != nil {
		return characters, err
	}

	defer res.Body.Close()
	charBytes, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(charBytes, &characters)

	if err != nil {
		return characters, err
	}

	return characters, nil
}

//GetCharacterByID is used to get a character from the api by ID
func (c *Client) GetCharacterByID(id string) (Character, error) {
	res, err := http.Get(BaseURL + "characters/" + id + "?key=" + c.apiKey)

	var character Character

	if err != nil {
		return character, err
	}

	charBytes, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(charBytes, &character)

	if err != nil {
		return character, err
	}

	return character, nil
}
