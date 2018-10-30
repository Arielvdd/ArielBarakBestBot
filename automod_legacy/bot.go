package automod_legacy

import (
	"github.com/google/safebrowsing"
	"github.com/jonas747/discordgo"
	"github.com/jonas747/yagpdb/bot"
	"github.com/jonas747/yagpdb/bot/eventsystem"
	"github.com/jonas747/yagpdb/commands"
	"github.com/jonas747/yagpdb/common"
	"github.com/jonas747/yagpdb/common/pubsub"
	"github.com/jonas747/yagpdb/moderation"
	"github.com/karlseguin/ccache"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

var _ bot.BotInitHandler = (*Plugin)(nil)
var _ bot.BotStopperHandler = (*Plugin)(nil)

var (
	// cache configs because they are used often
	confCache   *ccache.Cache
	SafeBrowser *safebrowsing.SafeBrowser
)

func (p *Plugin) BotInit() {
	commands.MessageFilterFuncs = append(commands.MessageFilterFuncs, CommandsMessageFilterFunc)

	eventsystem.AddHandler(HandleMessageUpdate, eventsystem.EventMessageUpdate)

	pubsub.AddHandler("update_automod_legacy_rules", HandleUpdateAutomodRules, nil)
	confCache = ccache.New(ccache.Configure().MaxSize(1000))

	safeBrosingAPIKey := os.Getenv("YAGPDB_GOOGLE_SAFEBROWSING_API_KEY")
	if safeBrosingAPIKey != "" {
		conf := safebrowsing.Config{
			APIKey: safeBrosingAPIKey,
			DBPath: "safebrowsing_db",
			Logger: logrus.StandardLogger().Writer(),
		}
		sb, err := safebrowsing.NewSafeBrowser(conf)
		if err != nil {
			logrus.WithError(err).Error("Failed initializing google safebrowsing client, integration will be disabled")
		} else {
			SafeBrowser = sb
		}
	}
}

func (p *Plugin) StopBot(wg *sync.WaitGroup) {
	if SafeBrowser != nil {
		err := SafeBrowser.Close()
		if err != nil {
			logrus.WithError(err).Error("Failed closing safebrowser")
		}
	}

	wg.Done()
}

// Invalidate the cache when the rules have changed
func HandleUpdateAutomodRules(event *pubsub.Event) {
	confCache.Delete(KeyConfig(event.TargetGuildInt))
}

// CachedGetConfig either retrieves from local application cache or redis
func CachedGetConfig(gID int64) (*Config, error) {
	confItem, err := confCache.Fetch(KeyConfig(gID), time.Minute*5, func() (interface{}, error) {
		c, err := GetConfig(gID)
		if err != nil {
			return nil, err
		}

		// Compile sites and words
		c.Sites.GetCompiled()
		c.Words.GetCompiled()

		return c, nil
	})

	if err != nil {
		return nil, err
	}

	return confItem.Value().(*Config), nil
}

func CommandsMessageFilterFunc(msg *discordgo.Message) bool {
	return !CheckMessage(msg)
}

func HandleMessageUpdate(evt *eventsystem.EventData) {
	CheckMessage(evt.MessageUpdate().Message)
}

func CheckMessage(m *discordgo.Message) bool {

	if m.Author == nil || m.Author.ID == common.BotUser.ID {
		return false // Pls no panicerinos or banerinos self
	}

	if m.Author.Bot {
		return false
	}

	cs := bot.State.Channel(true, m.ChannelID)
	if cs == nil {
		logrus.WithField("channel", m.ChannelID).Error("Channel not found in state")
		return false
	}

	if cs.IsPrivate() {
		return false
	}

	config, err := CachedGetConfig(cs.Guild.ID)
	if err != nil {
		logrus.WithError(err).Error("Failed retrieving config")
		return false
	}

	if !config.Enabled {
		return false
	}

	member, err := bot.GetMember(cs.Guild.ID, m.Author.ID)
	if err != nil {
		logrus.WithError(err).WithField("guild", cs.Guild.ID).Warn("Member not found in state, automod ignoring")
		return false
	}

	locked := true
	cs.Owner.RLock()
	defer func() {
		if locked {
			cs.Owner.RUnlock()
		}
	}()

	del := false // Set if a rule triggered a message delete
	punishMsg := ""
	highestPunish := PunishNone
	muteDuration := 0

	rules := []Rule{config.Spam, config.Invite, config.Mention, config.Links, config.Words, config.Sites}

	// We gonna need to have this locked while we check
	for _, r := range rules {
		if r.ShouldIgnore(m, member) {
			continue
		}

		d, punishment, msg, err := r.Check(m, cs)
		if d {
			del = true
		}
		if err != nil {
			logrus.WithError(err).WithField("guild", cs.Guild.ID).Error("Failed checking aumod rule:", err)
			continue
		}

		// If the rule did not trigger a deletion there wasn't any violation
		if !d {
			continue
		}

		punishMsg += msg + "\n"

		if punishment > highestPunish {
			highestPunish = punishment
			muteDuration = r.GetMuteDuration()
		}
	}

	if !del {
		return false
	}

	if punishMsg != "" {
		// Strip last newline
		punishMsg = punishMsg[:len(punishMsg)-1]
	}

	cs.Owner.RUnlock()
	locked = false

	go func() {
		switch highestPunish {
		case PunishNone:
			err = moderation.WarnUser(nil, cs.Guild.ID, cs.ID, common.BotUser, member.DGoUser(), "Automoderator: "+punishMsg)
		case PunishMute:
			err = moderation.MuteUnmuteUser(nil, true, cs.Guild.ID, cs.ID, common.BotUser, "Automoderator: "+punishMsg, member, muteDuration)
		case PunishKick:
			err = moderation.KickUser(nil, cs.Guild.ID, cs.ID, common.BotUser, "Automoderator: "+punishMsg, member.DGoUser())
		case PunishBan:
			err = moderation.BanUser(nil, cs.Guild.ID, cs.ID, common.BotUser, "Automoderator: "+punishMsg, member.DGoUser())
		}

		// Execute the punishment before removing the message to make sure it's included in logs
		common.BotSession.ChannelMessageDelete(m.ChannelID, m.ID)

		if err != nil && err != moderation.ErrNoMuteRole && !common.IsDiscordErr(err, discordgo.ErrCodeMissingPermissions, discordgo.ErrCodeMissingAccess) {
			logrus.WithError(err).Error("Error carrying out punishment")
		}
	}()

	return true

}
