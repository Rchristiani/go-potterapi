package potterapi

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetHouses(t *testing.T) {

	if !CheckExists("ENVIRONMENT") {
		err := godotenv.Load()

		if err != nil {
			t.Fatal(err)
		}
	}

	c := ClientInitialize(os.Getenv("POTTERAPIKEY"))

	_, err := c.GetHouses()

	if err != nil {
		t.Fatal(err)
	}

}

func TestGetHouseByID(t *testing.T) {
	if !CheckExists("ENVIRONMENT") {
		err := godotenv.Load()

		if err != nil {
			t.Fatal(err)
		}
	}

	c := ClientInitialize(os.Getenv("POTTERAPIKEY"))

	house, err := c.GetHouseByID("5a05e2b252f721a3cf2ea33f")

	if err != nil {
		t.Fatal(err)
	}

	if house[0].Name == "Gryffindor" {
		t.Log("Gryffindor!!")
	}
}
