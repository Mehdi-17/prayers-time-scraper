package telegram_bot

import (
	"fmt"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"prayers-time-scraper/internal/scraping"
	"strconv"
)

func SetUpBotConfiguration(salatOfTheDay [5]scraping.Salat) {
	//TODO : voir pour sortir ça d'ici, on a pas besoin de récupérer ces données tous les jours, une fois suffit
	token, chatId := getApiTokenFromEnv()
	bot, err := tgBotApi.NewBotAPI(token)

	//TODO : gestion des erreurs
	if err != nil {
		log.Fatal("Erreur lors de la création du bot api : ", err)
	}
	bot.Debug = true

	c := cron.New()

	for _, salat := range salatOfTheDay {
		// We have to reassign prayer's data so that each anonymous function has its own data.
		salatName := salat.Name
		salatMinute := salat.Hour.Minute()
		salatHour := salat.Hour.Hour()

		_, errSetUpCron := c.AddFunc(fmt.Sprintf("CRON_TZ=Europe/Paris %d %d * * *", salatMinute, salatHour), func() {
			msg := tgBotApi.NewMessage(chatId, fmt.Sprintf("C'est l'heure de la prière : %s", salatName))
			_, errSending := bot.Send(msg)

			if errSending != nil {
				log.Fatal("Erreur lors de l'envoi du rappel : ", errSending)
			}
		})
		if errSetUpCron != nil {
			log.Fatal("Erreur lors du set up d'un cron : ", errSetUpCron)
		}
	}

	c.Start()
}

func getApiTokenFromEnv() (token string, chatId int64) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token = os.Getenv("TELEGRAM_API_TOKEN")
	chatIdStr := os.Getenv("TELEGRAM_CHAT_ID")
	chatIdInt, errConv := strconv.Atoi(chatIdStr)
	if errConv != nil {
		log.Fatal("Erreur lors de la conversion du chat id : ", errConv)
	}
	chatId = int64(chatIdInt)
	return
}
