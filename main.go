package main

import (
	"flag"
	"fmt"
	"regexp"
	"sync"
)

const (
	SeriesUrl = "http://seasonvar.ru/serial-14498-Popadetc.html"
	IPUrl     = "http://ifconfig.co"
)

func PrintSecureMark(url string) {
	body, err := HttpGet(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	dataSecure := regexp.MustCompile(`\'secureMark\'\:\ \'([a-z0-9]+)\'`).FindStringSubmatch(body)
	if len(dataSecure) == 0 {
		fmt.Println("Cant find secure mark")
		return
	}

	fmt.Printf("Your secure mark: %s\n", dataSecure[1])

	dataComments := regexp.MustCompile(`.ru\/([a-z0-9]+)\/comments\/`).FindStringSubmatch(body)
	if len(dataComments) == 0 {
		fmt.Println("Cant find comments hash")
		return
	}

	fmt.Printf("Your comments hash: %s\n", dataComments[1])
}

func PrintIPAddress(url string) {
	body, err := HttpGet(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Your IP address: %s\n", body)
}

func main() {
	var input string
	var wg sync.WaitGroup

	seriesUrl := flag.String("url", SeriesUrl, "series url")
	flag.Parse()

	wg.Add(2)

	go func() {
		PrintIPAddress(IPUrl)
		wg.Done()
	}()

	go func() {
		PrintSecureMark(*seriesUrl)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Press Ctrl+C to exit.")
	fmt.Scanln(&input)
}
