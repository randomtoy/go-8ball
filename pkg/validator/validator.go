package validator

import (
	"log"
)

func IsValidCode(code string) bool {
	log.Printf("code: %+v", code)

	switch code {
	case "en":
		return true
	case "en-alice":
		return true
	default:
		return false
	}
}
