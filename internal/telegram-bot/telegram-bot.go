package telegram_bot

import (
	"fmt"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"prayers-time-scraper/internal/scraping"
)

func SetUpBotConfiguration(salatOfTheDay [5]scraping.Salat) {
	token, chatId := getApiTokenFromEnv()
	bot, err := tgBotApi.NewBotAPI(token)

	if err != nil {
		log.Fatal("Erreur lors de la création du bot api : ", err)
	}
	bot.Debug = true

	fmt.Printf("Token : %s\n Chatid : %s", token, chatId)
	//TODO : paramétrage des envoies de notifications :
	// use to build msg => msg := tgBotApi.NewMessage(update.Message.Chat.ID, "corps du message")
	// send msg => bot.Send(msg)

}

func getApiTokenFromEnv() (token, chatId string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token = os.Getenv("TELEGRAM_API_TOKEN")
	chatId = os.Getenv("TELEGRAM_CHAT_ID")
	return
}
