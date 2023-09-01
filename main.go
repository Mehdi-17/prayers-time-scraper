package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"prayers-time-scraper/internal/scraping"
)

func main() {
	c := cron.New()

	_, err := c.AddFunc("CRON_TZ=Europe/Paris 0 3 * * *", scraping.ScrapeAndNotify)

	c.Start()
	if err != nil {
		log.Fatal("Erreur lors du paramétrage du cron pour scrapper les données : ", err)
	}
	select {}
}
