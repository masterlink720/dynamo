package main

import (
	"github.com/akston/dynamo/internal/cf"
	log "github.com/sirupsen/logrus"
	_ "github.com/spf13/cobra"
	_ "github.com/spf13/viper"
)

func main() {
	// intName := os.Getenv("INT_NAME")

	if err := cf.UpdateRecord(); err != nil {
		log.Fatal(err)
	}
}
