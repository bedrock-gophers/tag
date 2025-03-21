package main

import (
	"fmt"
	"github.com/bedrock-gophers/tag/tag"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/sirupsen/logrus"
)

func main() {
	err := tag.Load("assets/tags/")
	if err != nil {
		panic(err)
	}
	fmt.Println(tag.All())
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.InfoLevel

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	conf, err := server.DefaultConfig().Config(log)
	if err != nil {
		log.Fatalln(err)
	}

	srv := conf.New()
	srv.CloseOnProgramEnd()

	srv.Listen()

	for srv.Accept(func(p *player.Player) {
		p.Handle(&handler{p: p})
	}) {

	}
}

type handler struct {
	player.NopHandler
	p *player.Player
}

func (h *handler) HandleChat(ctx *event.Context, message *string) {
	ctx.Cancel()

	canada, ok := tag.ByName("canada")
	if !ok {
		return
	}
	format := text.Colourf("%s %s: %s", canada.Format(), h.p.Name(), *message)
	_, _ = chat.Global.WriteString(format)
}
