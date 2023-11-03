package validators

import "time"

func ValidateDate(date string) error {
	layout := "2006-01-02"
	if _, err := time.Parse(layout, date); err != nil {
		return err
	}

	return nil
}
