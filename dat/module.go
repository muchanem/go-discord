package dat

import (
	"bytes"
	"encoding/json"
	"flag"
	dsg "github.com/bwmarrin/discordgo"
	f "github.com/skilstak/discord-public/lib"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

var (
	currentTime string
	path        string
	logger      *log.Logger
)

func init() {
	flag.StringVar(&path, "p", "./dat", "Path to directory where the bot can store and work with data")
	flag.Parse()

	currentTime = time.Now().Format("2006-01-02@15h04m05s")

	file, err := os.Create(path + "logs/system-logs@" + currentTime + ".log")
	if err != nil {
		panic(err)
	}

	logger = log.New(file, "", log.Ldate|log.Ltime|log.Llongfile|log.LUTC)
}

var lock sync.Mutex

func Save(fileName string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()

	file, err := os.Create(path + "botData/" + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}

	reader := bytes.NewReader(b)

	_, err = io.Copy(file, reader)

	return err
}

func Load(fileName string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	file, err := os.Open(path + "botData/" + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(v)
	return err
}

func GetBotInfo() (f.BotType, error) {
	raw, err := ioutil.ReadFile(path + "staticData/preferences.json")
	var b f.BotType

	if err != nil {
		return b, err
	}

	err = json.Unmarshal(raw, &b)

	if err != nil {
		return b, err
	}

	return b, nil
}

/* A Better Way To Panic™
* Panic() (not to be confused with the built in panic() is a function that
* helps handle and log errors that may occur from commands. This logs errors
* from commands without always killing the bot (but does if it comes to it) or
* forcing command creators to make their own logging systems that would be
* independant and confusing for a person running the bot.
*
* Parameters:
* - s (type *discordgo.Session) : You know the drill
* - m (type *discordgo.MessageCreate) : Just pass in the one your command is
*					running off of please.
* - err (type string) : The error to be logged. This is a string in case you
*			have your own error message you want to log, just put
*			in YOURERRORVAR.Error() if you don't.
* - fatal (type bool) : Notes if the entire bot needs to come down from the
*		        error. Please use sparingly
*
* Returns:
* Nothing. Please just put an empty return statment after your call.
 */
func Panic(s *dsg.Session, m *dsg.MessageCreate, err string, fatal bool) {
	s.ChannelMessageSend(m.ChannelID, "**ERROR ENCOUNTERED. DETAILS FOLLOW:**\n```"+err+"```\nThis incident will be reported.")
	if fatal {
		s.ChannelMessageSend(m.ChannelID, "The bot is now \"gracefully\" force quitting, however it might fail to close out of its session with discord and may still apear online.\n\n*Have a good day!*")
		s.Close()
		logger.Fatalln("FATAL! " + err)
	} else {
		logger.Println(err)
	}
}
