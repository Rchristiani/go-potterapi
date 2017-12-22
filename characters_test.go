package potterapi

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetCharacters(t *testing.T) {
	if !CheckExists("ENVIRONMENT") {
		err := godotenv.Load()
		if err != nil {
			t.Fatal("Failed to load .env")
		}
	}
	c := ClientInitialize(os.Getenv("POTTERAPIKEY"))

	_, err := c.GetCharacters()

	if err != nil {
		t.Fatal(err)
	}

}

func TestGetCharactersByID(t *testing.T) {
	if !CheckExists("ENVIROMENT") {
		err := godotenv.Load()

		if err != nil {
			t.Fatal("Failed to load .env")
		}
	}

	c := ClientInitialize(os.Getenv("POTTERAPIKEY"))

	character, err := c.GetCharacterByID("5a12292a0f5ae10021650d7e")

	if err != nil {
		t.Fatal(err)
	}

	if character.Name == "Harry Potter" {
		t.Log("Yeerr a wizzard harry")
	} else {
		t.Fatal("That's not Harry!")
	}
}
