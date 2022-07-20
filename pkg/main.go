package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"log"
	"os"
	"random-coffee-groups/pkg/group_generation"
	"random-coffee-groups/pkg/slack_integration"
)

const GROUP_SIZE = 4

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	OAuthToken := os.Getenv("SLACK_BOT_TOKEN")
	channelId := os.Getenv("CHANNEL_ID")

	api := slack.New(OAuthToken)

	users := slack_integration.GetUsers(api, channelId)
	groups := group_generation.GenerateGroups(users, GROUP_SIZE)

	attachment := slack.Attachment{
		Text: groups,
	}

	message := fmt.Sprintf("Here are the random coffees for the week: \n")
	channelId, timestamp, err := api.PostMessage(
		channelId,
		slack.MsgOptionText(message, true),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Fatalf("%s\n", err)
	}

	log.Printf("Message successfully sent to Channel %s at %s\n", channelId, timestamp)
}
