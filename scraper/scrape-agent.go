package scraper

import (
	"bytes"
	"errors"
	"time"

	"github.com/bjvanbemmel/jlpt-notify/notifier"
	"github.com/charmbracelet/log"
	"github.com/gocolly/colly"
)

type ScrapeAgent struct {
	Interval  time.Duration
	Collector *colly.Collector
	Notifier  *notifier.NotifyAgent
	Previous  string
	Current   string
}

var (
	ErrArgNil    error = errors.New("Argument may not be nil.")
	ErrPageEmpty error = errors.New("Scraped page's body is empty.")
)

func (s ScrapeAgent) requestCallback(r *colly.Request) {
	log.Infof("Scraping %s...", r.URL)
}

func (s *ScrapeAgent) scrapedCallback(r *colly.Response) {
	page := bytes.NewBuffer(r.Body).String()

	if page == "" {
		log.Error(ErrPageEmpty)
		return
	}

	if s.Previous == "" {
		s.Previous = page
		return
	}

	if s.Previous != page {
		s.Notifier.SendMessage("Contents have changed.")
		log.Warn("Contents have changed.")
	}

	s.Previous = page
}

func (s *ScrapeAgent) SetPrevious(p []byte) error {
	if p == nil {
		return ErrArgNil
	}
	s.Previous = bytes.NewBuffer(p).String()

	return nil
}

func (s *ScrapeAgent) SetCurrent(p []byte) error {
	if p == nil {
		return ErrArgNil
	}
	s.Previous = bytes.NewBuffer(p).String()

	return nil
}

func (s *ScrapeAgent) SetCollector(c *colly.Collector) error {
	if c == nil {
		return ErrArgNil
	}
	s.Collector = c

	return nil
}

func (s *ScrapeAgent) SetNotifier(n *notifier.NotifyAgent) error {
	if n == nil {
		return ErrArgNil
	}
	s.Notifier = n

	return nil
}

func (s *ScrapeAgent) SetInterval(i time.Duration) error {
	if i == 0 {
		return ErrArgNil
	}
	s.Interval = i

	return nil
}

func (s *ScrapeAgent) RunAgent() {

	for {
		log.Info("New iteration started")

		c := colly.NewCollector()

		c.OnRequest(s.requestCallback)
		c.OnScraped(s.scrapedCallback)
		s.SetCollector(c)

		s.Collector.Visit("http://192.168.2.15:8080")

		time.Sleep(s.Interval)
	}
}
