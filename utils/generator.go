package utils

import (
	"strconv"
	"time"
)

func GenerateTrxID() string {
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	return timestamp
}
