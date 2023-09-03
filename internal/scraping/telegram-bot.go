package scraping

import (
	"fmt"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"strconv"
)

func SetUpBotConfiguration(salatOfTheDay [5]Salat) {
	token, chatId := getApiTokenFromEnv()
	bot, err := tgBotApi.NewBotAPI(token)

	if err != nil {
		log.Fatal("Erreur lors de la création du bot api : ", err)
	}
	bot.Debug = true

	c := cron.New()

	for _, salat := range salatOfTheDay {
		salat.setUpAdhanAndReminder(c, bot, chatId)
	}

	c.Start()
}

func (salat Salat) setUpAdhanAndReminder(c *cron.Cron, bot *tgBotApi.BotAPI, chatId int64) {
	_, errSetUpCronReminder := c.AddFunc(fmt.Sprintf("CRON_TZ=Europe/Paris %d %d * * *", salat.reminderTime.Minute(), salat.reminderTime.Hour()), func() {
		sendMessageToTelegram(bot, chatId, fmt.Sprintf("RAPPEL : Salat %s à %02d:%02d.", salat.Name, salat.Time.Hour(), salat.Time.Minute()))
	})

	if errSetUpCronReminder != nil {
		log.Fatal("Erreur lors du set up d'un cron pour le rappel: ", errSetUpCronReminder)
	}

	_, errSetUpCron := c.AddFunc(fmt.Sprintf("CRON_TZ=Europe/Paris %d %d * * *", salat.Time.Minute(), salat.Time.Hour()), func() {
		sendMessageToTelegram(bot, chatId, fmt.Sprintf("C'est l'heure de la prière : %s", salat.Name))
	})

	if errSetUpCron != nil {
		log.Fatal("Erreur lors du set up d'un cron : ", errSetUpCron)
	}
}

func sendMessageToTelegram(bot *tgBotApi.BotAPI, chatId int64, text string) {
	msg := tgBotApi.NewMessage(chatId, text)
	_, errSending := bot.Send(msg)

	if errSending != nil {
		log.Fatal("Erreur lors de l'envoi du message : ", errSending)
	}
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
