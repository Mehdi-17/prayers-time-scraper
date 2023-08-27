package main

import (
	"prayers-time-scraper/internal/scraping"
	telegrambot "prayers-time-scraper/internal/telegram-bot"
)

func main() {
	salatOfTheDay := scraping.GetSalatTime()
	telegrambot.SetUpBotConfiguration(salatOfTheDay)
}
