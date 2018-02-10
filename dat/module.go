package dat

import (
	"bytes"
	"encoding/json"
	"flag"
	//dsg "github.com/bwmarrin/discordgo"
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
	stdlog      *log.Logger
	errlog      *log.Logger
	mfulog      *log.Logger
)

func init() {
	flag.StringVar(&path, "p", "./dat", "Path to directory where the bot can store and work with data")
	flag.Parse()

	currentTime = time.Now().Format("2006-01-02@15h04m")

	file, err := os.Create(path + "logs/general-logs@" + currentTime + ".log")
	if err != nil {
		panic(err)
	}
	stdlog = log.New(file, "", log.Ldate|log.Ltime|log.Llongfile|log.LUTC)

	file, err = os.Create(path + "logs/error-logs@" + currentTime + ".log")
	if err != nil {
		panic(err)
	}
	errlog = log.New(file, "", log.Ldate|log.Ltime|log.Llongfile|log.LUTC)

	file, err = os.Create(path + "logs/mfu-logs@" + currentTime + ".log")
	if err != nil {
		panic(err)
	}
	mfulog = log.New(file, "", log.Ldate|log.Ltime|log.Llongfile|log.LUTC)
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

/* Universal error logger
* Log() is a logger wrapper for the whole bot to use for logging events that
* happen. This can range from minor command permission failures to massive json
* data extraction issues. This helps the people running the bot understand
* whats going wrong and who they need to write to about it.
*
* Parameters:
* - err (type error) : The event to be logged. This directly calls err.Error()
* - priority (type int) : The priority of the error. While all events should be
*			  logged, some are actually important while others can
*			  be treated as "this certainly did happen". More
*			  important errors are also logged in their own file
*
* priority > 0 : Nothing out of the ordinary. Will not be noted apart from
*		 being put in the standard-log.log file (stdlog)
* priority = 0 : A minor error that forces a command to halt but does not
*		 affect the rest of the system. Will be put in error-log.log
*		 and given a "priority" prefix in stdlog.
* priority < 0 : A major failure that could be recovered from but might also
*		 force the system to crash. These will be put in mfu-log.log
*		 and given a "DANGER" prefix in stdlog
*
 */
func Log(err error, priority int) {
	if priority > 0 {
		stdlog.Println(err.Error())
	} else if priority == 0 {
		stdlog.Println("Priority: " + err.Error())
		errlog.Println(err.Error())
	} else if priority < 0 {
		stdlog.Println("DANGER: " + err.Error())
		mfulog.Println(err.Error())
	} else {
		stdlog.Println("DANGER: Log.priority (type int) index out of range.")
		mfulog.Println("Log.priority (type int) index out of range.")
	}
}
