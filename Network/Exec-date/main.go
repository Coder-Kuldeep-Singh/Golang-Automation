package main

import (
	"log"
	"os/exec"
)

func main() {

	// Send two pings and send the output to STDOUT
	output1, err := exec.Command("ping", "-c", "2", "8.8.8.8").Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("The results are --> %s\n", output1)

	// Run the date command and store the output
	output2, err := exec.Command("date").Output()
	if err != nil {
		log.Print(err)
	}
	log.Printf("The date is --> %s ", output2)
}
