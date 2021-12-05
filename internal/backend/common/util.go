package common

import (
	"fmt"
	t "github.com/pedroalbanese/kp/internal/backend/types"
	"time"
)

func CompareUUIDs(me t.UUIDer, them t.UUIDer) (bool, error) {
	myUUID, err := me.UUIDString()
	if err != nil {
		return false, err
	}

	theirUUID, err := them.UUIDString()
	if err != nil {
		return false, err
	}

	return theirUUID == myUUID, nil
}

func FormatTime(t time.Time) (formatted string) {
	timeFormat := "Mon Jan 2 15:04:05 MST 2006"
	if (t == time.Time{}) {
		formatted = "unknown"
	} else {
		since := time.Since(t).Round(time.Duration(1) * time.Second)
		sinceString := since.String()

		// greater than or equal to 1 day
		if since.Hours() >= 24 {
			sinceString = fmt.Sprintf("%d day(s) ago", int(since.Hours()/24))
		}

		// greater than or equal to ~1 month
		if since.Hours() >= 720 {
			// rough estimate, not accounting for non-30-day months
			months := int(since.Hours() / 720)
			sinceString = fmt.Sprintf("about %d month(s) ago", months)
		}

		// greater or equal to 1 year
		if since.Hours() >= 8760 {
			// yes yes yes, leap years aren't 365 days long
			years := int(since.Hours() / 8760)
			sinceString = fmt.Sprintf("about %d year(s) ago", years)
		}

		// less than a second
		if since.Seconds() < 1.0 {
			sinceString = "less than a second ago"
		}

		formatted = fmt.Sprintf("%s (%s)", t.Local().Format(timeFormat), sinceString)
	}
	return
}
