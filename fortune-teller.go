package main

import (
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const TOKEN = "7081723927:AAGLnRbovX3DeJ5fXEx3WY4LPKtZbTy4SkA"

var bot *tgbotapi.BotAPI
var fortuneTellerNames = [2]string{"лев", "лев толстой"}
var chatId int64
var answers = []string{
	"Да",
	"Нет",
	"Конечно",
	"Несомненно",
	"Даже не думай",
	"Иди поспи",
	"Самая важная учеба, какая есть на свете, это научиться любить и научиться различать свою любовь. (\"Анна Каренина\")",
	"Если вы хотите изменить мир, начните с себя. (\"Цитата из разных произведений\")",
	"Лучше всего начинать с того, что лежит у ваших ног. (\"Анна Каренина\")",
	"Любовь своего ближнего — это то, ради чего один лишь человек может назвать свою жизнь смысловой. (\"Воскресение\")",
	"Нет ничего более убедительного и сложного, чем простое правда. (\"Воскресение\")",
	"Делайте то, что считаете правильным, и дайте вам быть высмеянными. (\"Хаджи-Мурат\")",
	"Нельзя надеяться на справедливость в мире, где существует сила. (\"Война и мир\")",
	"Мы живем, чтобы делать добро. Что же еще остается делать? (\"Анна Каренина\")",
	"Не суди других людей, и тогда ты не будешь судим. (\"Анна Каренина\")",
	"Никогда не делай ничего против своей совести, даже если целый мир в тебя упрется. (\"Война и мир\")",
	"Правда, неважно, сколько ей лет; она всегда остается правдой. (\"Война и мир\")",
	"Будь собой, но стремись к лучшему. (\"Анна Каренина\")",
	"Судьба не столько связана с тем, что случилось с нами, сколько с тем, как мы реагируем на то, что случилось с нами. (\"Анна Каренина\")",
	// Добавьте остальные цитаты здесь
}

func connectWithTelegram() {
	var err error
	bot, err = tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		panic("Cannot connect to telegram")
	}
}

func sendMessage(msg string) {
	msgConfig := tgbotapi.NewMessage(chatId, msg)
	bot.Send(msgConfig)
}

func getFortuneTellersAnswer() string {
	index := rand.Intn(len(answers))
	return answers[index]
}

func isMessageForFortuneTeller(update *tgbotapi.Update) bool {
	if update.Message == nil || update.Message.Text == "" {
		return false
	}

	msgInLowerCase := strings.ToLower(update.Message.Text)
	for _, name := range fortuneTellerNames {
		if strings.Contains(msgInLowerCase, name) {
			return true
		}
	}
	return false
}

func sendAnswer(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(chatId, getFortuneTellersAnswer())
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func main() {
	connectWithTelegram()

	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatId = update.Message.Chat.ID
			sendMessage("Не соизволите ли вы задать мне вопрос, на который я мог бы ответить лишь одним словом, да или нет, обращаясь ко мне по имени Лев?" +
				"Сие дело облегчит мою участь и позволит мне сосредоточиться на истинном и важном." +
				"Пример: \"Лев, стоит ли мне сегодня изюмничать?\"")
		}
		if isMessageForFortuneTeller(&update) {
			sendAnswer(&update)
		}
	}
}
