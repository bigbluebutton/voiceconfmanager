package main

import (
	"fmt"
	"log"
	//"strconv"
	//regexp"
	"github.com/fiorix/go-eventsocket/eventsocket"
	//"github.com/bigbluebutton/voiceconfmanager/fseslclient"
)

func main() {
	fmt.Println("***** Starting....................***********")
	c, err := eventsocket.Dial("localhost:8021", "ClueCon")
	if err != nil {
		log.Fatal(err)
	}
	c.Send("events plain ALL")

	for {
		ev, err := c.ReadEvent()
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("\nNew event")
		//ev.PrettyPrint()
		confClass := ev.Get("Event-Subclass")
		if confClass == "conference::maintenance" {
			callingFunction := ev.Get("Event-Calling-Function")
			//uniqueId := ev.Get("Caller-Unique-ID")

			switch callingFunction {
			case "conference_add_member":
				fmt.Printf("\n=====%s=====\n", callingFunction)
				ev.PrettyPrint()
				fmt.Printf("\n***************\n")
				//handleUseJoinedEvent(ev)
			case "conference_del_member":
				fmt.Printf("\n=====%s=====\n", callingFunction)
				ev.PrettyPrint()
				fmt.Printf("\n***************\n")
			case "conf_api_sub_mute":
				fmt.Printf("\n=====%s=====\n", callingFunction)
				ev.PrettyPrint()
				fmt.Printf("\n***************\n")
			case "conf_api_sub_unmute":
				fmt.Printf("\n=====%s=====\n", callingFunction)
				ev.PrettyPrint()
				fmt.Printf("\n***************\n")
			case "conference_record_thread_run":
				fmt.Printf("\n=====%s=====\n", callingFunction)
				ev.PrettyPrint()
				fmt.Printf("\n***************\n")
				//action := ev.Get("Action")
			case "conference_loop_input":
				fmt.Printf("\n=====%s=====\n", callingFunction)
				ev.PrettyPrint()
				fmt.Printf("\n***************\n")
			default:
				fmt.Printf("\ndefault: %s\n", callingFunction)
			}
		}
	}
	c.Close()
}
