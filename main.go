package main

import (
	"log"
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

type Result struct {
    Weather []struct {
        Main string
        Description string
        Icon string
    }
    Main struct {
        Temp float64
        Humidity int
    }
}

// API Key
const APIKEY = "" 

func getWeatherData(city string) []byte {
    // Prepare the http client for the request
    client := &http.Client{}
    // URL for the request
    url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + APIKEY + "&units=metric"

    // Get the response
    response, err := client.Get(url) 
    // Panic if error
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()

    // Read the response
    body, readingErr := io.ReadAll(response.Body)

    // If error occure during reading response body return fatal error 
    if readingErr != nil {
        log.Fatal("Error: ", readingErr)
    }

    return body
}

func main() {
    log.SetPrefix("weazher: ")

    // Get data returned by the API request
    data := getWeatherData("abidjan")

    var result Result
    if err := json.Unmarshal(data, &result); err != nil {
            log.Fatal(err)
    }

    fmt.Println(result)
}
