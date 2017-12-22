package potterapi

import "os"

func CheckExists(value string) bool {
	_, ok := os.LookupEnv(value)
	if !ok {
		return false
	}
	return true
}
