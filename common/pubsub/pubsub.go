// The event system is used to propegate events from different yagpdb instances
// For example when you change the streamer settings, and event gets fired
// Telling the streamer plugin to recheck everyones streaming status

package pubsub

import (
	"encoding/json"
	"fmt"
	"github.com/jonas747/yagpdb/common"
	"github.com/mediocregopher/radix.v3"
	"github.com/sirupsen/logrus"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
)

type Event struct {
	TargetGuild    string // The guild this event was meant for, or * for all
	TargetGuildInt int64
	EventName      string
	Data           interface{}
}

type eventHandler struct {
	evt     string
	handler func(*Event)
}

var (
	eventHandlers = make([]*eventHandler, 0)
	eventTypes    = make(map[string]reflect.Type)
)

// AddEventHandler adds a event handler
// For the specified event, should only be done during startup
func AddHandler(evt string, cb func(*Event), t interface{}) {
	handler := &eventHandler{
		evt:     evt,
		handler: cb,
	}

	if t != nil {
		eventTypes[evt] = reflect.TypeOf(t)
	}

	eventHandlers = append(eventHandlers, handler)
	logrus.WithField("evt", evt).Debug("Added event handler")
}

// PublishEvent publishes the specified event
func Publish(evt string, target int64, data interface{}) error {
	dataStr := ""
	if data != nil {
		encoded, err := json.Marshal(data)
		if err != nil {
			return err
		}
		dataStr = string(encoded)
	}

	value := fmt.Sprintf("%d,%s,%s", target, evt, dataStr)
	return common.RedisPool.Do(radix.Cmd(nil, "PUBLISH", "events", value))
}

func PollEvents() {
	// Create a new client for pubsub
	// radix.PersistentPubSub("tcp", common.Conf.Redis, radix.Dial)

	client, err := radix.Dial("tcp", common.Conf.Redis)
	if err != nil {
		panic(err)
	}

	pubsubClient := radix.PubSub(client)

	msgChan := make(chan radix.PubSubMessage)
	if err := pubsubClient.Subscribe(msgChan, "events"); err != nil {
		logrus.WithError(err).Fatal("Failed subscribing to events")
		return
	}

	for msg := range msgChan {
		if len(msg.Message) > 0 {
			logrus.WithField("evt", string(msg.Message)).Info("Handling PubSub event")
			handleEvent(string(msg.Message))
		}
	}

	for {
	}
}

func handleEvent(evt string) {
	split := strings.SplitN(evt, ",", 3)

	if len(split) < 3 {
		logrus.WithField("evt", evt).Error("Invalid event")
		return
	}

	target := split[0]
	name := split[1]
	data := split[2]

	t, ok := eventTypes[name]
	if !ok && data != "" {
		// No handler for this event
		logrus.WithField("evt", name).Info("No handler for pubsub event")
		return
	}

	var decoded interface{}
	if data != "" && t != nil {
		decoded = reflect.New(t).Interface()
		err := json.Unmarshal([]byte(data), decoded)
		if err != nil {
			logrus.WithError(err).Error("Failed unmarshaling event")
			return
		}
	} else if t != nil {
		logrus.Error("No data provided for event that requires data")
		return
	}

	event := &Event{
		TargetGuild: target,
		EventName:   name,
		Data:        decoded,
	}

	parsedTarget, _ := strconv.ParseInt(target, 10, 64)
	event.TargetGuildInt = parsedTarget

	defer func() {
		if r := recover(); r != nil {
			stack := string(debug.Stack())
			logrus.Error("Recovered from panic in pubsub event handler", r, "\n", stack)
		}
	}()

	for _, handler := range eventHandlers {
		if handler.evt != name {
			continue
		}

		handler.handler(event)
	}
}
