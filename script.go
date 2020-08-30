package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// endpoint for the API
var endpoint = "http://localhost:9000/api/v1?user_id="

var numberGoRoutines = 50

func getUserInfo(wg *sync.WaitGroup, userID int) {

	var responseData []byte
	var err error
	var url string

	defer wg.Done()

	if userID == 31 {
		time.Sleep(time.Second * 3)
	}

	if userID >= 15 && userID <= 30 {
		url = fmt.Sprintf("%s%s", endpoint, "bad-request")
	} else {
		url = fmt.Sprintf("%s%d", endpoint, userID)
	}

	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		log.Printf("Error while getting response from url : %v", err)
	} else {
		if response.StatusCode != http.StatusOK {
			log.Printf("Service returned status-code : %v and respone is : %v", response.StatusCode, response)
		} else {
			responseData, err = ioutil.ReadAll(response.Body)
			if err != nil {
				log.Printf("error while reading response body : %v", err)
			}
		}
	}

	log.Printf("response for %d is : %v", userID, string(responseData))

}

func main() {

	var (
		wg sync.WaitGroup
	)

	for id := 1; id <= numberGoRoutines; id++ {

		wg.Add(1)

		getUserInfo(&wg, id)

		wg.Wait()
	}
}
