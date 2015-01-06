package main

import (
	"fmt"
	"github.com/juju/loggo"
	"github.com/mipearson/rfw"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

var logger = loggo.GetLogger("main")
var rootLogger = loggo.GetLogger("")

func main() {
	args := os.Args
	if len(args) > 1 {
		loggo.ConfigureLoggers(args[1])
	} else {
		fmt.Println("Add a parameter to configure the logging:")
		fmt.Println("E.g. \"<root>=INFO;first=TRACE\"")
	}
	fmt.Println("\nCurrent logging levels:")
	fmt.Println(loggo.LoggerInfo())
	fmt.Println("")

	rootLogger.Infof("Start of test.")

	FirstCritical("first critical")
	FirstError("first error")
	FirstWarning("first warning")
	FirstInfo("first info")
	FirstTrace("first trace")

	SecondCritical("first critical")
	SecondError("first error")
	SecondWarning("first warning")
	SecondInfo("first info")
	SecondTrace("first trace")

	writer, err := rfw.Open("./myprogram", 0644)
	if err != nil {
		log.Fatalln("Could not open '/var/log/myprogram': ", err)
	}

	log := log.New(writer, "[myprogram] ", log.LstdFlags)
	log.Println("Logging as normal")

	//defaultWriter, _, err := loggo.RemoveWriter("default")

	// err is non-nil if and only if the name isn't found.
	//loggo.RegisterWriter("default", writer, loggo.TRACE)
	newWriter := loggo.NewSimpleWriter(writer, &loggo.DefaultFormatter{})
	err = loggo.RegisterWriter("testfile", newWriter, loggo.TRACE)

	now := time.Now().Nanosecond()
	fmt.Printf("CurrentTime: %d\n", now)
	FirstCritical("first critical")
	FirstError("first error")
	FirstWarning("first warning")
	FirstInfo("first info")
	FirstTrace("first trace")

	SecondCritical("first critical")
	SecondError("first error")
	SecondWarning("first warning")
	SecondInfo("first info")
	SecondTrace("first trace")

	viper.SetConfigName("config")        // name of config file (without extension)
	viper.AddConfigPath("/etc/appname/") // path to look for the config file in
	viper.AddConfigPath("$HOME/")        // call multiple times to add many search paths
	viper.ReadInConfig()                 // Find and read the config file

	title := viper.GetString("title")
	fmt.Printf("TITLE=%s\n", title)
}
