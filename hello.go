package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

func main() {

	for {
		showIntrodution()

		command := readCommand()

		switch command {
		case 1:
			StartMonitoring()
		case 2:
			fmt.Println("Showing logns...")
		case 0:
			fmt.Println("Bye Bye :((")
			os.Exit(0)
		default:
			fmt.Println("Invalid!")
			os.Exit(-1)
		}
	}
}

func showIntrodution() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func StartMonitoring() {
	fmt.Println("Monitoring...")
	//sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}

	sites := readFileSites()
	for i := 0; i < monitoring; i++ {
		for _, site := range sites {
			testSite(site)
			time.Sleep(delay * time.Second)
		}
		fmt.Println("")
	}

}

func exibeNomes() {
	nomes := []string{"Daniel", "Douglas", "Bernardo"}
	fmt.Println(nomes)
}

func testSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "Loaded sucess")
	} else {
		fmt.Println("Status code error", response.StatusCode, "on", site)
	}
}

func readFileSites() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err != io.EOF {
			break
		}
	}

	file.Close()
	return sites
}
