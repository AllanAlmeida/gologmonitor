package main

import "os"
import "fmt"
import "logmonitor/banner"
import "logmonitor/monitor"



func main() {
	banner.Banner()
 	choosedOpt := getChoosedOption()
 	executeFlow(choosedOpt)
}

func getChoosedOption() int {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")

	option := 0
	fmt.Scan(&option)
	return option
}

func executeFlow(opt int) {
	switch opt {
		case 1: 
			monitor.StartMonitor()
		case 2: 
			fmt.Println("Showing logs...")
			monitor.PrintLogs()
		case 0: 
			fmt.Println("Exit")
			os.Exit(0)
		default: 
			fmt.Println("Invalid option!")
			os.Exit(-1)
	}
}
