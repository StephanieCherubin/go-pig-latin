package slack

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
)

const helpMessage = "type in '@pig-latin <command_arg_1> <command_arg_2>'"

/*
   CreateSlackClient sets up the slack RTM (real-timemessaging) client library,
   initiating the socket connection and returning the client.
   DO NOT EDIT THIS FUNCTION. This is a fully complete implementation.
*/
func CreateSlackClient(apiKey string) *slack.RTM {
	api := slack.New(apiKey)
	rtm := api.NewRTM()
	go rtm.ManageConnection() // goroutine!
	return rtm
}

/*
   RespondToEvents waits for messages on the Slack client's incomingEvents channel,
   and sends a response when it detects the bot has been tagged in a message with @<botTag>.

   EDIT THIS FUNCTION IN THE SPACE INDICATED ONLY!
*/
func RespondToEvents(slackClient *slack.RTM) {
	for msg := range slackClient.IncomingEvents {
		fmt.Println("Event Received: ", msg.Type)
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			botTagString := fmt.Sprintf("<@%s> ", slackClient.GetInfo().User.ID)
			if !strings.Contains(ev.Msg.Text, botTagString) {
				continue
			}
			message := strings.Replace(ev.Msg.Text, botTagString, "", -1)

			fmt.Println(message)
			// TODO: Make your bot do more than respond to a help command. See notes below.
			// Make changes below this line and add additional funcs to support your bot's functionality.
			// sendHelp is provided as a simple example. Your team may want to call a free external API
			// in a function called sendResponse that you'd create below the definition of sendHelp,
			// and call in this context to ensure execution when the bot receives an event.

			// START SLACKBOT CUSTOM CODE
			// ===============================================================
			sendResponse(slackClient, message, ev.Channel)
			sendHelp(slackClient, message, ev.Channel)
			sendPigLatin(slackClient, message, ev.Channel)
			// ===============================================================
			// END SLACKBOT CUSTOM CODE
		default:

		}
	}
}

const (
	pigLatinSuffix             string = "ay"
	vowels                     string = "aeiou"
	firstLetterExceptionSuffix string = "d" + pigLatinSuffix
)

// sendPigLatin translates one or more english words into the PigLatin equlivent
func sendPigLatin(slackClient *slack.RTM, message, slackChannel string) {
	var pigLatinWords []string
	englishWords := strings.Split(message, " ")

	if len(message) == 0 && message == " " {
		return
	}

	for _, word := range englishWords {
		first := word[0:1]
		if strings.Contains(vowels, first) {
			pigLatinWords = append(pigLatinWords, word+firstLetterExceptionSuffix)
		} else {
			pigLatinWords = append(pigLatinWords, word[1:]+first+pigLatinSuffix)
		}
	}

	pigLatinString := strings.Join(pigLatinWords, " ")
	slackClient.SendMessage(slackClient.NewOutgoingMessage(pigLatinString, slackChannel))
}

// sendHelp is a working help message, for reference.
func sendHelp(slackClient *slack.RTM, message, slackChannel string) {
	if strings.ToLower(message) != "help" {
		return
	}
	slackClient.SendMessage(slackClient.NewOutgoingMessage(helpMessage, slackChannel))
}

// sendResponse is NOT unimplemented --- write code in the function body to complete!

func sendResponse(slackClient *slack.RTM, message, slackChannel string) {
	command := strings.ToLower(message)
	println("[RECEIVED] sendResponse:", command)

	// START SLACKBOT CUSTOM CODE
	// ===============================================================
	// TODO:
	//      1. Implement sendResponse for one or more of your custom Slackbot commands.
	//         You could call an external API here, or create your own string response. Anything goes!
	//      2. STRETCH: Write a goroutine that calls an external API based on the data received in this function.
	// ===============================================================
	// END SLACKBOT CUSTOM CODE
}
