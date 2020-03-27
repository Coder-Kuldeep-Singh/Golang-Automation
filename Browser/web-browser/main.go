package main

import (
	"log"

	"github.com/sclevine/agouti"
)

func main() {

	driver := agouti.ChromeDriver()

	if err := driver.Start(); err != nil {
		log.Fatal("Failed to start driver:", err)
	}

	page, err := driver.NewPage()
	if err != nil {
		log.Fatal("Failed to open page:", err)
	}
	if err := page.Navigate("https://agouti.org"); err != nil {
		log.Fatal("Failed to navigate:", err)
	}

	sectionTitle, err := page.FindByID(`getting-agouti`).Text()
	log.Println(sectionTitle)

	if err := driver.Stop(); err != nil {
		log.Fatal("Failed to close pages and stop WebDriver:", err)
	}

}
