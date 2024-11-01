package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	accessToken = "vk1.a.token"
	itRollID    = -218375169
	apiVersion  = "5.199"
	baseURL     = "https://api.vk.com/method"
)

type VKResponse struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID       int `json:"id"`
			Comments struct {
				Count int `json:"count"`
			} `json:"comments"`
		} `json:"items"`
	} `json:"response"`
}

type VKCommentsResponse struct {
	Response struct {
		Items []struct {
			Text string `json:"text"`
		} `json:"items"`
	} `json:"response"`
}

func makeRequest(ctx context.Context, method string, params map[string]string) (*http.Response, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/%s?", baseURL, method)
	params["access_token"] = accessToken
	params["v"] = apiVersion

	queryParams := make([]string, 0)
	for k, v := range params {
		queryParams = append(queryParams, fmt.Sprintf("%s=%v", k, v))
	}
	url += strings.Join(queryParams, "&")

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}

func getPostsCount(ctx context.Context) (int, error) {
	params := map[string]string{
		"owner_id": fmt.Sprintf("%d", itRollID),
		"count":    "1",
	}

	resp, err := makeRequest(ctx, "wall.get", params)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var result VKResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}

	return result.Response.Count, nil
}

func getPostsWithComments(ctx context.Context, count int, offset int) ([]int, error) {
	params := map[string]string{
		"owner_id": fmt.Sprintf("%d", itRollID),
		"count":    fmt.Sprintf("%d", count),
		"offset":   fmt.Sprintf("%d", offset),
	}

	resp, err := makeRequest(ctx, "wall.get", params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result VKResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	var posts []int
	for _, item := range result.Response.Items {
		if item.Comments.Count > 0 {
			posts = append(posts, item.ID)
		}
	}

	return posts, nil
}

func getPostEnvelopeCount(ctx context.Context, postID int) (int, error) {
	params := map[string]string{
		"owner_id": fmt.Sprintf("%d", itRollID),
		"post_id":  fmt.Sprintf("%d", postID),
	}

	resp, err := makeRequest(ctx, "wall.getComments", params)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var result VKCommentsResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}

	count := 0
	for _, comment := range result.Response.Items {
		count += strings.Count(strings.ToLower(comment.Text), "энвилоуп")
	}

	return count, nil
}

func main() {
	start := time.Now()
	ctx := context.Background()
	step := 20

	postsCount, err := getPostsCount(ctx)
	if err != nil {
		return
	}

	var wg sync.WaitGroup
	postsChan := make(chan int, 1000)

	for offset := 0; offset < postsCount; offset += step {
		wg.Add(1)
		go func(offset int) {
			defer wg.Done()
			posts, err := getPostsWithComments(ctx, step, offset)
			if err != nil {
				return
			}
			for _, post := range posts {
				postsChan <- post
			}
		}(offset)
	}

	go func() {
		wg.Wait()
		close(postsChan)
	}()

	var totalCount int
	var countWg sync.WaitGroup
	countChan := make(chan int)

	go func() {
		for count := range countChan {
			totalCount += count
		}
	}()

	for postID := range postsChan {
		countWg.Add(1)
		go func(postID int) {
			defer countWg.Done()
			count, err := getPostEnvelopeCount(ctx, postID)
			if err != nil {
				return
			}
			countChan <- count
		}(postID)
	}

	countWg.Wait()
	close(countChan)

	duration := time.Since(start)
	fmt.Printf("Количество \"энвилоуп\" в комментариях: %d\n", totalCount)
	fmt.Printf("Время, затраченное на подсчёт энвилоупов: %.3f с\n", duration.Seconds())
}
