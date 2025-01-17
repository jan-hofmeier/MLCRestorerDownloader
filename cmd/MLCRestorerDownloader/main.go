package main

import (
	"bufio"
	"fmt"
	"os"

	downloader "github.com/Xpl0itU/MLCRestorerDownloader"
)

func main() {
	fmt.Println("Menu:")
	fmt.Println("1. Download MLC titles")
	fmt.Println("2. Download SLC titles")
	fmt.Println("3. Exit")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Select an option: ")
	option, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	switch option {
	case "1\n", "1\r\n":
		showSubmenu("MLC")
	case "2\n", "2\r\n":
		showSubmenu("SLC")
	case "3\n", "3\r\n":
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println("Invalid option")
		return
	}
}

func showSubmenu(titleType string) {
	if titleType != "MLC" && titleType != "SLC" {
		fmt.Println("Invalid title type")
		return
	}
	titles, err := readTitleInfoFromFile("titles.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var chosenTitles map[string][]string
	switch titleType {
	case "MLC":
		chosenTitles = titles.MLC
	case "SLC":
		chosenTitles = titles.SLC
	default:
		fmt.Println("Invalid option")
		return
	}

	fmt.Println("Menu:")
	fmt.Printf("1. Download EUR %s titles\n", titleType)
	fmt.Printf("2. Download USA %s titles\n", titleType)
	fmt.Printf("3. Download JPN %s titles\n", titleType)
	fmt.Println("4. Back to main menu")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Select an option: ")
	option, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	switch option {
	case "1\n", "1\r\n":
		downloadTitles("EUR", chosenTitles, titleType)
	case "2\n", "2\r\n":
		downloadTitles("USA", chosenTitles, titleType)
	case "3\n", "3\r\n":
		downloadTitles("JPN", chosenTitles, titleType)
	case "4\n", "4\r\n":
		fmt.Println("Going back to the main menu...")
		main()
	default:
		fmt.Println("Invalid option")
		return
	}
}

func downloadTitles(region string, titles map[string][]string, titleType string) {
	selectedRegionTitles := titles[region]
	allRegionTitles := titles["All"]

	allTitles := append(selectedRegionTitles, allRegionTitles...)

	commonKey, err := getCommonKey()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, titleID := range allTitles {
		if titleID == "dummy" {
			continue
		}
		fmt.Printf("Downloading files for title %s on region %s for type %s\n", titleID, region, titleType)
		err := downloader.DownloadTitle(titleID, fmt.Sprintf("output/%s/%s/%s", titleType, region, titleID), commonKey)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Printf("Download files for title %s on region %s for type %s done\n", titleID, region, titleType)
	}
	fmt.Println("All done!")
}
