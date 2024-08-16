package ie2utilities

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func IE2GetEnv(envname string) (string, error) {

	log.Printf("Retrieving env variable %s", envname)
	res := os.Getenv(envname)

	if len(res) <= 0 {
		msg := fmt.Sprintf("missing environment variable: %s", envname)
		log.Print(msg)
		return "", errors.New(msg)
	}

	return res, nil
}
