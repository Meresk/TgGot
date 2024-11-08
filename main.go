package main

import (
	"flag"
	"log"
)

func main() {

	//tgClient = telegram.New(mustHost(), mustToken())

	// token = flags.Get(token)

	// tgClient = telegram.New(token)

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("bot-token",
		"",
		"token for access to tg bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is required")
	}

	return *token
}

func mustHost() string {
	host := flag.String("bot-host", "", "telegram bot host")
	flag.Parse()

	if *host == "" {
		log.Fatal("host is required")
	}

	return *host
}
