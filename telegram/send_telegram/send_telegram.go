package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	Token  string
	ChatId string
)

func getEnvVariable(env string) (string, error) {
	recomandation := fmt.Sprintf("Run the following command: 'export %s=\"<%s>\"'", env, env)
	envContent, ok := os.LookupEnv(env)
	if !ok {

		return "", fmt.Errorf("%s doesn't exists. %s", env, recomandation)
	}
	if len(strings.TrimSpace(envContent)) == 0 {
		return "", fmt.Errorf("%s is empty. %s", env, recomandation)
	}
	return envContent, nil
}

func getUrl() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", Token)
}

func SendMessage(destination string, text string, silent bool) (bool, error) {
	// Global variables
	var err error
	var response *http.Response

	// Send the message
	url := fmt.Sprintf("%s/sendMessage", getUrl())
	body, _ := json.Marshal(map[string]interface{}{
		"chat_id": destination,
		"text": text,
		"disable_notification":	silent,
	})
	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return false, err
	}

	// Close the request at the end
	defer response.Body.Close()

	// Body
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	// Log
	log.Infof("Message '%s' was sent", text)
	log.Infof("Response JSON: %s", string(body))

	// Return
	return true, nil
}

func main() {
	var (
		destination = flag.String("destination", "", "[MANDATORY] ID of the destination user/group")
		message 	= flag.String("message", "Hello There!", "Message that will be send")
		silent		= flag.Bool("silent", false, "If is set, the message won't be notified")
	)
	flag.Parse()
	if *destination == "" {
		log.Fatalf("%s", "'destination' flag is required. Execute '-h' for help")
		os.Exit(1)
	}

	// Get the TOKEN
	var err error
	Token, err = getEnvVariable("TOKEN")
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Send a message
	_, err = SendMessage(*destination, *message, *silent)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
