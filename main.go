package main

import (
	cmd "./src/commandHandler"
	f "./src/foundation"
	jsd "./src/jsonDecoder"
	"flag"
	"fmt"
	dsg "github.com/bwmarrin/discordgo"
	c "github.com/skilstak/go-colors"
	"os"
	"os/signal"
	"syscall"
)

var (
	FilePath string
)

func init() ***REMOVED***
	flag.StringVar(&FilePath, "f", "", "Path to json file containing bot info")
	flag.Parse()
***REMOVED***

func main() ***REMOVED***
	bot, err := jsd.GetBotInfo(FilePath)
	fmt.Println(c.B0 + "Reading bot prefrences file...")
	if err != nil ***REMOVED***
		fmt.Println(c.R + "Unable to read prefrences file. Exiting program." + c.X)
		panic(err)
	***REMOVED*** else ***REMOVED***
		fmt.Println(c.G + "Bot prefrences recived.")
	***REMOVED***
	fmt.Println(c.B0 + "Writing bot preferences to skilbot")
	f.MyBot = bot
	f.DG = runBot(f.MyBot.Auth.Username, f.MyBot.Auth.Secret, f.MyBot.Auth.ClientID, f.MyBot.Auth.Token)

	fmt.Println(c.B0 + "Bot is now running! Press CTRL+C to exit." + c.X)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	f.DG.Close()
***REMOVED***

func runBot(username string, secret string, id string, token string) *dsg.Session ***REMOVED***
	fmt.Println(c.B01 + "Info provided:")
	fmt.Println(c.B00 + "Username  : " + c.O + username)
	fmt.Println(c.B00 + "Secret    : " + c.O + secret)
	fmt.Println(c.B00 + "Client ID : " + c.O + id)
	fmt.Println(c.B00 + "Token     : " + c.O + token)

	fmt.Println(c.B0 + "Creating bot session...")
	dg, err := dsg.New("Bot " + token)

	if err != nil ***REMOVED***
		fmt.Println(c.R+"Error in creating discord session. Exiting program."+c.X, err)
		os.Exit(-1)
	***REMOVED*** else ***REMOVED***
		fmt.Println(c.G + "Session successfuly created.")
	***REMOVED***

	fmt.Println(c.B0 + "Adding message handlers...")
	dg.AddHandler(cmd.MessageCreate)

	fmt.Println(c.B0 + "Opening websocket to Discord...")
	err = dg.Open()
	if err != nil ***REMOVED***
		fmt.Println(c.R+"Error opening websocket to Discord. Exiting program."+c.X, err)
		os.Exit(-1)
	***REMOVED*** else ***REMOVED***
		fmt.Println(c.G + "Socket successfully opened.")
	***REMOVED***
	return dg
***REMOVED***
