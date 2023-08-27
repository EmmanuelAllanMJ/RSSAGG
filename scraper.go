package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/EmmanuelAllanMJ/rssagg/internal/database"
)

func startScraper(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetFeeds(context.Background())

		if err != nil {
			log.Println("error getting feeds", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()

	}

}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("error marking feed as fetched", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("error fetching feed", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		// _, err := db.CreateItem(context.Background(), database.CreateItemParams{
		// 	FeedID:      feed.ID,
		// 	Title:       item.Title,
		// 	Description: item.Description,
		// 	Link:        item.Link,
		// 	PubDate:     item.PubDate,
		// })
		// if err != nil {
		// 	log.Println("error creating item", err)
		// 	return
		// }
		log.Println("Found Post", item.Title, "in feed", feed.Url)
	}
	log.Printf("Found %d posts in %s", len(rssFeed.Channel.Item), feed.Url)
}
