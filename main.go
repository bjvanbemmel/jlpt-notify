package main

import (
	"os"
	"os/signal"
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
	na.SetSender(os.Getenv("SMS_SENDER"))
	na.SetReceiver(os.Getenv("SMS_RECEIVER"))
	na.SetRestClient(twilio.NewRestClient())

	scraperAgent.SetNotifier(na)
	scraperAgent.SetInterval(time.Second * 5)
	scraperAgent.RunAgent()
}

func signalHandler(signal os.Signal) {
    log.Warn(signal.String())

    if signal == syscall.SIGTERM || signal == syscall.SIGINT || signal == syscall.SIGKILL {
        log.Warn("Graceful exiting now...")

        os.Exit(1)
    }
}
