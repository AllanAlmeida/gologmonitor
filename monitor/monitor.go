package monitor

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func StartMonitor() {
	fmt.Println("Monitoring...")
	monitoredSites := sitesToMonitor()
	for _, site := range monitoredSites {
		resp, _ := http.Get(site)
		verifyStatusAndWriteLogs(site, resp.StatusCode)
	}
}

func GetUser() {
	fmt.Println("Monitoring...")

	site := "https://62bd9608bac21839b606a9cc.mockapi.io/apitest/v1/users"
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	respData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respData))
}

func MapDeleteTest() {
	fmt.Println("Testing...")
	resp, _ := http.Get("https://62bd9608bac21839b606a9cc.mockapi.io/apitest/v1/users")

	for k, v := range resp.Header {
		fmt.Printf("[%s]:-->%s ", k, v)
		fmt.Println("")
	}

	resp.Header.Del("Date")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Printing new map...")
	fmt.Println("")

	for k, v := range resp.Header {
		fmt.Printf("[%s]:-->%s ", k, v)
		fmt.Println("")
	}
}

func WriteLogs(log string, status bool) {
	file, err := os.OpenFile("resources/log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("An error has ocurred: ", err)
	}

	formatedLog := "[" + time.Now().Format("02/01/2006 15:04:05") + "] " + log + "[STATUS]: "

	if status {
		formatedLog = formatedLog + " Online. \n"
	} else {
		formatedLog = formatedLog + " Down. \n"
	}
	file.WriteString(formatedLog)
	file.Close()
}

func PrintLogs() {
	for _, logLines := range readLines("resources/log") {
		fmt.Println(logLines)
	}
}

func verifyStatusAndWriteLogs(site string, status int) {
	if status == 200 {
		WriteLogs("[SITE]: "+site+" - working properly! ", true)
	} else {
		WriteLogs("[SITE]: "+site+" - doesn't working! ", false)
	}
}

func sitesToMonitor() []string {
	return readLines("resources/sites")
}

func readLines(file string) []string {
	var readedLines []string
	contentFile, err := os.Open(file)

	if err != nil {
		fmt.Println("Does not possible open the file. [ERRO]: ", err)
	}

	reader := bufio.NewReader(contentFile)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		readedLines = append(readedLines, line)
		if err == io.EOF {
			break
		}
	}

	contentFile.Close()
	return readedLines
}
