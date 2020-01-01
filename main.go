package main

import (
	_ "net/http"

	log "github.com/sirupsen/logrus"
	_ "github.com/spf13/cobra"
	_ "github.com/spf13/viper"
)

func main() {
	log.Println("test")
}
