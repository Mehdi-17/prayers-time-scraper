package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"prayers-time-scraper/internal/scraping"
	telegrambot "prayers-time-scraper/internal/telegram-bot"
)

func main() {
	c := cron.New()

	_, err := c.AddFunc("CRON_TZ=Europe/Paris 0 3 * * *", func() {
		salatOfTheDay := scraping.GetSalatTime()
		telegrambot.SetUpBotConfiguration(salatOfTheDay)
	})
	c.Start()
	if err != nil {
		log.Fatal("Erreur lors du paramétrage du cron pour scrapper les données : ", err)
	}
	select {}
}
