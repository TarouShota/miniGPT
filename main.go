package main

import (
	"fmt"
	"os"
	// "telego/telego"
	// "goBot/go"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
    tu "github.com/mymmrac/telego/telegoutil"
	"github.com/joho/godotenv"

	"io/ioutil"
	"log"
	"net/http"
)



func getHolidays() string{
    resp, err := http.Get("https://calendarific.com/api/v2/holidays?&api_key="+os.Getenv("CALENDAR_API")+"&country=CZ&year=2022&month=10")
    if err != nil {
       log.Fatalln(err)
    }
 //We Read the response body on the line below.
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
       log.Fatalln(err)
    }
 //Convert the body to type string
    
    sb := string(body)
    return sb

}


func main() {
	err := godotenv.Load(".env")
	// os.Setenv("TOKEN","1")
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
			fmt.Sprintf("Hello %s!", getHolidays()),
		))
	}, th.CommandEqual("start"))

	// Register new handler with match on any command
	// Handlers will match only once and in order of registration, 
	// so this handler will be called on any command except `/start` command
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		// Send message
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Unknown command, use /start",
		))
	}, th.AnyCommand())

	// Start handling updates
    bh.Start()
}

