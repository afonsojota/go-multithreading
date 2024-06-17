package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type ApiResponse struct {
	Data []byte
	API  string
	Time time.Duration
	Err  error
}

func getFromAPI(url string, apiName string, ch chan<- ApiResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	client := http.Client{
		Timeout: time.Second,
	}
	responseStart := time.Now()
	response, err := client.Get(url)
	duration := time.Since(responseStart)
	if err != nil {
		ch <- ApiResponse{Err: fmt.Errorf("%s request timeout or error: %w", apiName, err)}
		return
	}
	defer response.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		ch <- ApiResponse{Err: fmt.Errorf("%s JSON decoding error: %w", apiName, err)}
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		ch <- ApiResponse{Err: fmt.Errorf("%s JSON marshaling error: %w", apiName, err)}
		return
	}

	ch <- ApiResponse{Data: jsonData, API: apiName, Time: duration}
}

func main() {
	cep := "01153000"
	ch := make(chan ApiResponse, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	brasilAPIURL := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	viaCEPURL := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	go getFromAPI(viaCEPURL, "ViaCEP", ch, &wg)
	go getFromAPI(brasilAPIURL, "BrasilAPI", ch, &wg)

	wg.Wait()

	close(ch)

	var fastestResponse ApiResponse

	select {
	case fastestResponse = <-ch:
	case <-time.After(time.Second):
		fmt.Println("Timeout: No responses received within the time limit.")
		return
	}

	if fastestResponse.Err != nil {
		fmt.Printf("Error in fastest response: %v\n", fastestResponse.Err)
	} else {
		fmt.Printf("Fastest response from API %s (%v ms):\n%s\n", fastestResponse.API, fastestResponse.Time.Milliseconds(), string(fastestResponse.Data))
	}

}
