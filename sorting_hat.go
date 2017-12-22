package potterapi

import (
	"io/ioutil"
	"net/http"
)

//SortingHat is used to get a random house from the API
func (c *Client) SortingHat() (string, error) {
	res, err := http.Get(BaseURL + "/sortinghat")

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
