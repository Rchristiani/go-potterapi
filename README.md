# Potterapi Go Wrapper

[![CircleCI](https://img.shields.io/circleci/project/Rchristiani/go-potterapi.svg?style=flat-square)](https://circleci.com/Rchristiani/go-potterapi)

A Simple wrapper around the [potterapi](https://www.potterapi.com/) 

### Usage

```go
package main

import (
	"fmt"
	"github.com/rchristiani/go-potter" //import package
)

func main() {
    c := ClientInitialize("api key goes here")

	characters, err := c.GetCharacters()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(characters)
}
``` 


## Sorting Hat

### `SortingHat() string` 

Simple method to get what house you are in!

```go
c := ClientInitialize("api key")

fmt.Println(c.SortingHat())
```

## Characters

```go
type Character struct {
	ID                string 
	BloodStatus       string 
	DeathEater        bool   
	DumbledoresArmy   bool   
	House             string 
	MinistryOfMagic   bool   
	Name              string 
	OrderOfThePhoenix bool   
	Role              string 
	School            string 
	Species           string 
	Wand              string 
	Boggart           string 
	Alias             string 
	Animagus          string 
}
```


### `GetCharacters() ([]Character, error)` 

```go
c := ClientInitialize("api key goes here")

characters, err := c.GetCharacters()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(characters)
```

### `GetCharacterByID(id string) (Character, error)`

```go
c := ClientInitialize(os.Getenv("api key"))
character, err := c.GetCharacterByID("5a12292a0f5ae10021650d7e")

fmt.Println(character.Name) //Harry Potter!
```
