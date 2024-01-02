package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	// "telego/telego"
	// "goBot/go"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"

	"log"
	"net/http"
)

func sendGptRequest(message string) string {
	apiKey := os.Getenv("AITOKEN")
	url := "https://api.openai.com/v1/chat/completions"

	data := map[string]interface{}{
		"model":      "gpt-3.5-turbo",
		"messages": []map[string]string{{
			"role":    "user",
			"content": message,
		}},
	}

	payload, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		panic(err)
	}

	choices, ok := result["choices"].([]interface{})
	if !ok {
		panic("Invalid response format")
	}

	if len(choices) == 0 {
		return ""
	}

	messageMap, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	if !ok {
		panic("Invalid response format")
	}

	return messageMap["content"].(string)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("TOKEN")
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// botUser, err := bot.GetMe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPulling(nil)

	// Create bot handler and specify from where to get updates
	bh, _ := th.NewBotHandler(bot, updates)

	// Stop handling updates
	defer bh.Stop()

	// Stop getting updates
	defer bot.StopLongPulling()

	//request

	// Register new handler with match on command `/start`
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Send message
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Hello! How can I assist you today?"),
		))
	}, th.CommandEqual("start"))

	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		// Get chat ID from the message
		chatID := tu.ID(message.Chat.ID)

		//now make request to the open ai api

		_, _ = bot.SendMessage(tu.Message(
			chatID,
			sendGptRequest(message.Text),
		))

	})

	// Start handling updates
	bh.Start()
}
