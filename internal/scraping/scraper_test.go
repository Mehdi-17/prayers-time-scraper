package scraping

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"testing"
)

type MockWebScraper struct{}

func (m *MockWebScraper) ScrapeMawaqitWebsite(ctx context.Context, nodes []*cdp.Node) [5]Salat {
	return [5]Salat{}
}

func TestGetSalatTime(t *testing.T) {
	realScraper := &RealScraper{}

	got := getSalatTime(realScraper)
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
