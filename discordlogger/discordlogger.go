package discordlogger

import (
	"fmt"
	"github.com/jonas747/yagpdb/bot"
	"github.com/jonas747/yagpdb/bot/eventsystem"
	"github.com/jonas747/yagpdb/common"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var (
	// Send bot leaves joins to this disocrd channel
	BotLeavesJoins int64
)

func init() {
	BotLeavesJoins, _ = strconv.ParseInt(os.Getenv("YAGPDB_BOTLEAVESJOINS"), 10, 64)
}

func Register() {
	if BotLeavesJoins != 0 {
		logrus.Info("Listening for bot leaves and join")
		common.RegisterPlugin(&Plugin{})
	}
}

var _ bot.BotInitHandler = (*Plugin)(nil)

type Plugin struct{}

func (p *Plugin) BotInit() {
	eventsystem.AddHandler(EventHandler, eventsystem.EventNewGuild, eventsystem.EventGuildDelete)
}

func EventHandler(evt *eventsystem.EventData) {
	bot.State.RLock()
	count := len(bot.State.Guilds)
	bot.State.RUnlock()

	msg := ""
	switch evt.Type {
	case eventsystem.EventGuildDelete:
		if evt.GuildDelete().Unavailable {
			// Just a guild outage
			return
		}
		msg = fmt.Sprintf(":x: Left guild **%s** :(", evt.GuildDelete().Guild.Name)
	case eventsystem.EventNewGuild:
		msg = fmt.Sprintf(":white_check_mark: Joined guild **%s** :D", evt.GuildCreate().Guild.Name)
	}

	msg += fmt.Sprintf(" (now connected to %d servers)", count)
	common.BotSession.ChannelMessageSend(BotLeavesJoins, common.EscapeSpecialMentions(msg))
}

func (p *Plugin) Name() string {
	return "DiscordLogger"
}
