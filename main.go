package main

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"

	cron "github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
}

func compare() {
	prodResp, err := http.Get(os.Getenv("PROD_URL"))
	if err != nil {
		log.Error("Production URL value empty or null.")
		panic(err)
	}
	testResp, err := http.Get(os.Getenv("TEST_URL"))
	if err != nil {
		log.Error("Test URL value empty or null.")
		panic(err)
	}
	defer prodResp.Body.Close()
	defer testResp.Body.Close()

	var prodResponse map[string]interface{}
	err = json.NewDecoder(prodResp.Body).Decode(&prodResponse)
	if err != nil {
		log.Error("Production URL response parse error.")
		panic(err)
	}

	var testResponse map[string]interface{}
	err = json.NewDecoder(testResp.Body).Decode(&testResponse)
	if err != nil {
		log.Error("Test URL response parse error")
		panic(err)
	}
	if testResponse["core_version"] == prodResponse["core_version"] {
		log.Info("core_version match: ", prodResponse["core_version"])
	} else {
		log.Error("core_version mismatch !!! Below are the details")
		log.Error("Production core_version: ", prodResponse["core_version"])
		log.Error("Test core_version: ", testResponse["core_version"])
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	log.Info("Create new cron")
	log.Info("cron schedule:", os.Getenv("APP_CRON_SCHEDULE"))
	c := cron.New(cron.WithSeconds())
	c.AddFunc(os.Getenv("APP_CRON_SCHEDULE"), compare)

	// Start cron with scheduled job
	log.Info("Start cron")
	c.Start() // schedule jobs start
	wg.Wait() // waiting for schedule job to complete

}
