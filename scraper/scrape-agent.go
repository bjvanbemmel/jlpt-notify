package scraper

import (
	"bytes"
	"errors"
	"os"
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
}

var (
	ErrArgNil      error = errors.New("Argument may not be nil.")
	ErrPageEmpty   error = errors.New("Scraped page's body is empty.")
	ErrTargetEmpty error = errors.New("The scrape target may not be empty. Please provide a valid URI.")
)

func (s ScrapeAgent) requestCallback(r *colly.Request) {
	log.Printf("Scraping %s...", r.URL)
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
		s.Notifier.SendMessage("Contents have changed. Check https://jlpt-leiden.nl.")
		log.Info("Contents have changed!")
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

func (s *ScrapeAgent) CheckBackup() error {
	if _, err := os.Stat("./tmp/page.backup"); err != nil {
		return err
	}

	raw, err := os.ReadFile("./tmp/page.backup")
	if err != nil {
		return err
	}

	log.Info("Creating comparable page from backup.")
	s.Previous = bytes.NewBuffer(raw).String()

	return nil
}

func (s *ScrapeAgent) RunAgent() error {
	target := os.Getenv("SCRAPE_TARGET_URI")
	if target == "" {
		return ErrTargetEmpty
	}

	err := s.CheckBackup()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	for {
		c := colly.NewCollector()

		c.OnRequest(s.requestCallback)
		c.OnScraped(s.scrapedCallback)
		s.SetCollector(c)

		if err := s.Collector.Visit(target); err != nil {
			return err
		}

		time.Sleep(s.Interval)
	}
}
