package main

import (
	"flag"
	"log"

	"github.com/Kwynto/read-adviser-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())
	// fetcher := fetcher.New(tgClient)
	// processor := processor.New(tgClient)

	// consumer.Start(fetcher, processor)
}
