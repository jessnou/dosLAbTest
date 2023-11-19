package usecase

import (
	"dosLAbTest/internal/entity"
	"dosLAbTest/pkg/postgres"
	"dosLAbTest/pkg/postgres/models"
	"dosLAbTest/pkg/postgres/query"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func UpdateStatistics(pg *postgres.Postgres) {

	for {
		var result []models.MaxWord
		posts := getData()
		jobs := make(chan entity.Post, len(posts))
		resultChan := make(chan []models.MaxWord, len(posts))
		var wg sync.WaitGroup
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go worker(jobs, resultChan, &wg)

		}

		for _, post := range posts {
			jobs <- post
		}

		select {
		case <-time.After(time.Second):
			close(jobs)
			close(resultChan)
			wg.Wait()
			for words := range resultChan {
				for _, word := range words {
					result = append(result, word)
				}

			}
			if err := query.Insert(*pg, result); err != nil {
				log.Fatalf("cannot insert or update data %v", err)
			}
		}

		time.Sleep(5 * time.Minute)
	}
}
func worker(posts <-chan entity.Post, results chan<- []models.MaxWord, wg *sync.WaitGroup) {
	defer wg.Done()

	for post := range posts {
		maxWord := counter(post)

		results <- maxWord

	}

}

func counter(post entity.Post) []models.MaxWord {
	var wordsss []models.MaxWord
	wordCount := make(map[string]int)

	words := strings.Fields(strings.ToLower(post.Body))
	for _, word := range words {
		wordCount[word]++
	}

	for word, key := range wordCount {
		mxW := models.MaxWord{
			PostID: post.ID,
			Count:  key,
			Word:   word,
		}
		wordsss = append(wordsss, mxW)
	}

	return wordsss
}

func getData() (posts []entity.Post) {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalf("cannot fetch to server %v", err)
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Fatalf("cannot close body response %v", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Fatalf("cannot unmarshal %v", err)
	}
	return
}
