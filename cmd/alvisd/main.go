package main

import (
	"fmt"
	dsg "github.com/bwmarrin/discordgo"
	c "github.com/skilstak/go-colors"
	"github.com/takama/daemon"
	f "github.com/whitman-colm/go-discord"
	"github.com/whitman-colm/go-discord/cmd/handler"
	"github.com/whitman-colm/go-discord/dat"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	name        = "godiscord"
	description = "go-discord bot"
	port        = "15212"
)

var option string

type Service struct {
	daemon.Daemon
}

func init() {
	flag.StringVar(&option, "s", "help", "Action to take")
	flag.Parse()
}

func (service *Service) Manage() (string, error) {
	usage := "Usage: " + name + " install | remove | start | stop | status"

	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}

	//TODO: set up the bot to exist here.

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		dat.Log.Println(err)
		return "Possibly was a problem with the port binding", err
	}

	listen := make(chan net.Conn, 100)
	go acceptConnection(listener, listen)

	for {
		select {
		case conn := <-listen:
			go handleClient(conn)
		case killSignal := <-interrupt:
			dat.Log.Println("Recived signal:", killSignal)
			dat.Log.Println("Stopping listening on ", listener.Addr())
			listener.Close()
			if killSignal == os.Interrupt {
				dat.Log.Println("Daemon was interrupted by system signal")
				return "Daemon was interrupted by system signa", nil
			}
			dat.Log.Println("Daemon was killed")
			return "Daemon was killed", nil
		}
	}

	return usage, nil
}

func acceptConnection(listener net.Listener, listen chan<- net.Conn) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		listen <- conn
	}
}

func handleClient(client net.Conn) {
	for {
		buf := make([]byte, 4096)
		numbytes, err := client.Read(buf)
		if numbytes == 0 || err != nil {
			return
		}
		client.Write(buf[:numbytes])
	}
}

func main() {
	srv, err := daemon.New(name, description, "")
	if err != nil {
		dat.Log.Println(err)
		os.Exit(1)
	}
	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		dat.Log.Println(err)
		os.Exit(1)
	}
	fmt.Println(status)

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
	dat.Log.Println("Creating bot session")
	dg, err := dsg.New("Bot " + token)

	if err != nil {
		dat.Log.Fatalln(err)
	} else {
		dat.Log.Println("Session successfully created.")
	}

	dg.AddHandler(cmd.MessageCreate)

	dat.Log.Println("Opening websocket to Discord")
	err = dg.Open()
	if err != nil {
		dat.Log.Fatalln(err.Error())
	} else {
		dat.Log.Println("Socket successfully opened.")
	}
	return dg
}

/*func installWizard(service *Service) (string, error) {
	fmt.Println("You will be taken through the process of installing an instance of")
	fmt.Println("the go-discord bot. Please have a bot user ready from")
	fmt.Println("https://discordapp.com/developers/applications/ before progressing.")
	clid, err := input.Prompt("Please enter the client ID:\n> ")
	if err != nil {
		dat.Log.Println(err)
		return "", err
	}
	clsc, err := input.Prompt("Please enter the client secret:\n> ")
	if err != nil {
		dat.Log.Println(err)
		return "", err
	}
	cltk, err := input.Prompt("Please enter the token:\n> ")
	if err != nil {
		dat.Log.Println(err)
		return "", err
	}
	pref, err := input.Prompt("Please decide on a prefix. This is used to call the bot (leave blank to only have an @ mention trigger)\nPrefix: ")
	if err != nil {
		dat.Log.Println(err)
		return "", err
	}
}*/
