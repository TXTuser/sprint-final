package api

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	start, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", fmt.Errorf("invalid date")
	}

	parts := strings.Split(repeat, " ")
	rule := parts[0]

	switch rule {

	case "y":
		d := start
		for {
			d = d.AddDate(1, 0, 0)
			if d.After(now) {
				return d.Format(DateFormat), nil
			}
		}

	case "d":
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid repeat")
		}

		days, err := strconv.Atoi(parts[1])
		if err != nil || days <= 0 || days > 400 {
			return "", fmt.Errorf("invalid repeat")
		}

		d := start
		for {
			d = d.AddDate(0, 0, days)
			if d.After(now) {
				return d.Format(DateFormat), nil
			}
		}
	}

	return "", fmt.Errorf("unsupported repeat")
}
