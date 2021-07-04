package dice

import (
	"fmt"
	"math/rand"

	"github.com/gadget-bot/gadget/router"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func rollD6() *router.MentionRoute {
	var pluginRoute router.MentionRoute
	pluginRoute.Permissions = append(pluginRoute.Permissions, "*")
	pluginRoute.Name = "dice.rollD6"
	pluginRoute.Description = "Rolls two d6 dice"
	pluginRoute.Help = "roll some dice"
	pluginRoute.Pattern = `(?i)^(roll some dice|dice me)[!.]?$`
	pluginRoute.Plugin = func(api slack.Client, router router.Router, ev slackevents.AppMentionEvent, message string) {
		// Here's how we can react to the message
		msgRef := slack.NewRefToMessage(ev.Channel, ev.TimeStamp)
		api.AddReaction("game_die", msgRef)

		// Roll virtual dice
		dice := []int{1, 2, 3, 4, 5, 6}
		roll1 := rollDie(dice)
		roll2 := rollDie(dice)

		// Here's how we send a reply
		api.PostMessage(
			ev.Channel,
			slack.MsgOptionText(
				fmt.Sprintf("<@%s> rolled a %d and a %d", ev.User, roll1, roll2),
				false,
			),
		)
	}

	// We've got to return the MentionRoute
	return &pluginRoute
}

func rollDie(sides []int) int {
	index := rand.Intn(len(sides))
	return sides[index]
}

// GetMentionRoutes is used to retrieve all Mention Routes from this plugin
func GetMentionRoutes() []router.MentionRoute {
	return []router.MentionRoute{
		*rollD6(),
	}
}
