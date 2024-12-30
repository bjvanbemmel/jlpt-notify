package main

import (
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/bjvanbemmel/jlpt-notify/notifier"
	"github.com/bjvanbemmel/jlpt-notify/scraper"
	"github.com/charmbracelet/log"
	"github.com/twilio/twilio-go"
)

var scraperAgent scraper.ScrapeAgent = scraper.ScrapeAgent{}
var signalChannel chan os.Signal = make(chan os.Signal, 1)

func main() {
	signal.Notify(signalChannel)

	go func() {
		for {
			signalHandler(<-signalChannel)
		}
	}()

	na := &notifier.NotifyAgent{}

	if err := na.SetSender(os.Getenv("SMS_SENDER")); err != nil {
		log.Fatal(err.Error())
	}

	if err := na.SetReceiver(os.Getenv("SMS_RECEIVER")); err != nil {
		log.Fatal(err.Error())
	}

	na.SetRestClient(twilio.NewRestClient())

	envInt := os.Getenv("SCRAPE_INTERVAL")
	interval, err := strconv.Atoi(envInt)
	if err != nil {
		log.Fatal(err.Error())
	}

	// na.SendMessage(fmt.Sprintf("JLPT-Notify has started scraping %s...", os.Getenv("SCRAPE_TARGET_URI")))

	if err := scraperAgent.SetNotifier(na); err != nil {
		log.Fatal(err.Error())
	}

	if err := scraperAgent.SetInterval(time.Minute * time.Duration(interval)); err != nil {
		log.Fatal(err.Error())
	}

	if err := scraperAgent.RunAgent(); err != nil {
		log.Fatal(err.Error())
	}
}

func signalHandler(signal os.Signal) {
	if signal == syscall.SIGTERM || signal == syscall.SIGINT || signal == syscall.SIGKILL {
		log.Warn("Gracefully exiting now...")

		backup, err := os.Create("tmp/page.backup")
		if err != nil {
			log.Fatal(err.Error())
		}

		if _, err := backup.WriteString(scraperAgent.Previous); err != nil {
			log.Fatal(err.Error())
		}

		log.Info("Successfully wrote backup to filesystem!")

		os.Exit(1)
	}
}
