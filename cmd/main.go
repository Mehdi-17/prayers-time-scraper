package main

import (
	"prayers-time-scraper/internal/scraping"
	telegram_bot "prayers-time-scraper/internal/telegram-bot"
)

func main() {
	salatOfTheDay := scraping.GetSalatTime()
	telegram_bot.SetUpBotConfiguration(salatOfTheDay)
}
