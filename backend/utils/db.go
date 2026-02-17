package utils

import "log"

func SaveToDB(table string, item any) error {
	log.Printf("Saving to %s: %+v", table, item)
	return nil
}
