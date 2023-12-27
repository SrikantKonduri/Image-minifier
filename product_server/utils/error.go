package utils

import "log"

func FailOnError(err error, msg string) bool {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
		return true
	}
	return false
}
