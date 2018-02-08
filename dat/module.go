package nil

import (
	"encoding/json"
	"flag"
	dsg "github.com/bwmarrin/discordgo"
	f "github.com/skilstak/discord-public/lib"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	time   string
	path   string
	logger *log.Logger
)

func init() {
	flag.StringVar(&path, "p", "./dat", "Path to directory where the bot can store and work with data")
	flag.Parse()

	time = time.Now().Format("2006-01-02@15:04:05")

	file, err := os.Create(logpath + "logs/system-logs@" + time + ".log")
	if err != nil {
		panic(err)
	}

	logger = log.New(file, "", Ldate|Ltime|Llongfile|LUTC)
}

func GetBotInfo() (f.BotType, error) {
	raw, err0 := ioutil.ReadFile(path)
	var b f.BotType

	if err0 != nil {
		return b, err0
	}

	err1 := json.Unmarshal(raw, &b)

	if err1 != nil {
		return b, err1
	}

	return b, nil
}

func Panic(s *dsg.Session, m *dsg.MessageCreate, err error, fatal bool) {
	s.ChannelMessageSend(m.ChannelID, "**ERROR ENCOUNTERED. DETAILS FOLLOW:**\n```"+err.Error()+"```\nThis incident will be reported.")
	if fatal {
		s.ChannelMessageSend(m.ChannelID, "The bot is now \"gracefully\" force quitting, however it will fail to close out of its session with discord and may still apear online.\n*Have a good day!*")
		logger.Fatalln(err.Error())
	} else {
		logger.Println(err.Error())
	}
}
