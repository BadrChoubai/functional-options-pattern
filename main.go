package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO ", log.LstdFlags)
	errorLog := log.New(os.Stdout, "ERROR ", log.LstdFlags)

	UberAPIClient, err := New(
		WithHealthcheckURLOption("https://api.uber.com/health"),
		WithErrorLogOption(errorLog),
		WithInfoLogOption(infoLog),
	)

	if err != nil {
		fmt.Print("failed to create client: baseClient")
	}

	fmt.Println(UberAPIClient.String())

	CallHealthcheck(UberAPIClient)
}
