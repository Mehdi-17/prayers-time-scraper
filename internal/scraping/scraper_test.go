package scraping

import (
	"testing"
)

func TestGetSalatTime(t *testing.T) {
	got := getSalatTime()
	wantedSalatNamed := []string{"Fajr", "Dhuhr", "Asr", "Maghrib", "Isha"}

	for i, salatName := range wantedSalatNamed {
		if got[i].Name != salatName {
			t.Fatalf("The name of on salat is wrong ! want : %s; got : %s", salatName, got[i].Name)
		}

		if i != len(got)-1 && got[i+1].Time.Before(got[i].Time) {
			t.Fatalf("The time of %s is before the time of %s", got[i+1].Name, got[i].Name)
		}
	}
}
