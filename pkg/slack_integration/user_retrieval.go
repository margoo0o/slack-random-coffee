package slack_integration

import (
	"github.com/slack-go/slack"
	"log"
)

var CoffeeBotUserID = ""

// GetUsers Get all the users in the channel to allow the pairs to be generated/*
func GetUsers(api *slack.Client, channelId string) []string {
	// https://slack.com/api/conversations.info
	// users, cursor, err
	users, _, err := api.GetUsersInConversation(&slack.GetUsersInConversationParameters{
		ChannelID: channelId,
		Limit:     100,
	},
	)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	setBotUserId(api)

	return cleanData(api, users)
}

func setBotUserId(api *slack.Client) {
	authResult, err := api.AuthTest()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	CoffeeBotUserID = authResult.UserID
}

func cleanData(api *slack.Client, users []string) []string {
	// Remove the coffee bot user - should not be included in list of users
	removeUser(users, CoffeeBotUserID)

	// Remove duplicates
	return unique(users)
}

func unique(users []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range users {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func removeUser(users []string, user string) []string {
	for i, v := range users {
		if v == user {
			return append(users[:i], users[i+1:]...)
		}
	}
	return users
}
