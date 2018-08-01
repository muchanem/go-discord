package main

import (
	"fmt"
	dsg "github.com/bwmarrin/discordgo"
	c "github.com/skilstak/go-colors"
	f "github.com/skilstak/go-discord"
	"github.com/skilstak/go-discord/cmd/handler"
	"github.com/skilstak/go-discord/dat"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	bot, err := dat.GetBotInfo()
	fmt.Println(c.B0 + "Reading bot prefrences file...")
	dat.Log.Println("Reading bot prefs file...")
	if err != nil {
		dat.Log.Fatalln(err.Error())
	} else {
		dat.Log.Println("Bot prefrences recived.")
		fmt.Println(c.G + "Bot prefrences recived.")
	}
	fmt.Println(c.B0 + "Writing bot preferences")
	f.MyBot = bot
	f.DG = runBot(f.MyBot.Auth.Username, f.MyBot.Auth.Secret, f.MyBot.Auth.ClientID, f.MyBot.Auth.Token)

	f.DG.UpdateStatus(0, f.MyBot.Prefs.Playing)

	fmt.Println(c.B0 + "Bot is now running! Press CTRL+C to exit." + c.X)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	f.DG.Close()
	dat.Log.Println("Escape for bot called. The system is now closing cleanly")
}

func runBot(username string, secret string, id string, token string) *dsg.Session {
	fmt.Println(c.B01 + "Info provided:")
	fmt.Println(c.B00 + "Username  : " + c.O + username)
	fmt.Println(c.B00 + "Secret    : " + c.O + secret)
	fmt.Println(c.B00 + "Client ID : " + c.O + id)
	fmt.Println(c.B00 + "Token     : " + c.O + token)

	fmt.Println(c.B0 + "Creating bot session...")
	dat.Log.Println("Creating bot session")
	dg, err := dsg.New("Bot " + token)

	if err != nil {
		dat.Log.Fatalln(err.Error())
	} else {
		dat.Log.Println("Session successfully created.")
		fmt.Println(c.G + "Session successfully created.")
	}

	fmt.Println(c.B0 + "Adding message handlers...")
	dat.Log.Println("Adding message handlers")
	dg.AddHandler(cmd.MessageCreate)

	fmt.Println(c.B0 + "Opening websocket to Discord...")
	dat.Log.Println("Opening websocket to Discord")
	err = dg.Open()
	if err != nil {
		dat.Log.Fatalln(err.Error())
	} else {
		dat.Log.Println("Socket successfully opened.")
		fmt.Println(c.G + "Socket successfully opened.")
	}
	return dg
}
