package main

import (
	"sync"
	"testing"
	"time"
)

// Put your logic and code inside this function to optimize the app.
func Runner() map[int]string {
	results := make(map[int]string)
	wg := sync.WaitGroup{}
	wg.Add(worker)

	// Use goroutines to make concurrent requests to mockGetData function.
	for i := 1; i <= worker; i++ {
		go func(id int) {
			defer wg.Done()
			res, err := mockGetData(id)
			if err != nil {
				return
			}
			results[res.ID] = res.Title
		}(i)
	}
	wg.Wait()

	return results
}

// ===== DO NOT EDIT. =====
type Result struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

const worker = 3

var (
	expectedWorker                = 0
	expected       map[int]string = map[int]string{
		1: "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		2: "qui est esse",
		3: "ea molestias quasi exercitationem repellat qui ipsa sit aut",
	}
)

func TestCaseParalelUniverse(t *testing.T) {
	start := time.Now()
	results := Runner()
	lat := time.Since(start).Milliseconds()

	switch {
	case lat >= 2500:
		t.Fatalf("Runner function latency is greater than 2500ms")

	case !assertEqual(results):
		t.Fatalf("Result of Runner function is not equal to expected")

	case expectedWorker < worker:
		t.Fatalf("Expected worker count not met")
	}
}

func assertEqual(results map[int]string) bool {
	if len(results) != len(expected) {
		return false
	}
	for k, v := range results {
		if expected[k] != v {
			return false
		}
	}
	return true
}

func mockGetData(id int) (*Result, error) {
	expectedWorker++
	// Reduce the sleep time to speed up the response time.
	time.Sleep(500 * time.Millisecond)
	result := Result{
		ID:    id,
		Title: expected[id],
	}
	return &result, nil
}

func main() {
	testSuite := []testing.InternalTest{
		{
			Name: "TestCaseParalelUniverse",
			F:    TestCaseParalelUniverse,
		},
	}

	testing.Main(nil, testSuite, nil, nil)
}

// ===== DO NOT EDIT. =====

// OUTPUT 
// PASS