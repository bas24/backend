package main

import (
	"os"
	"strconv"
	"strings"
	"time"
)

var ginPort = os.Getenv("GIN_PORT")
var dbCockroachPathConst = os.Getenv("COCKROACH_PATH") // cockroach db
var NewsApiKey = os.Getenv("NEWSAPI_ORG_KEY")          // www.newsapi.org
var googleApiKey = os.Getenv("GOOGLE_TRANSLATOR_API_KEY")
var msFreeApiKeyLimit = 2000000 // limit of characters for free Microsoft text translator
var msPaidApiKeyLimit = 0       // limit for paid translator api

var msBingSpellCheckApiKey = os.Getenv("AZURE_BING_SPELL_CHECK_API_KEY")
var msBingSpellCheckApiKeyLimit = 10000 // api calls Miscrosoft Bing Spell Check Api
var meaningcloudApiKey string

func initMeaningcloudKeys() {
	key := strings.Split(os.Getenv("MEANINGCLOUD_API_KEY_1"), ";")
	mc := &meaningcloudApiKey
	*mc = key[0]
}

func initMicrosoftApiKeys() error {
	t := time.Now()
	month := t.Month().String()
	for i := 1; i <= 9999; i++ {
		k := strconv.Itoa(i)
		keyFree := "AZURE_TRANSLATOR_API_KEY_" + k
		keyPaid := "AZURE_TRANSLATOR_API_KEY_PAID_" + k
		apiKeyFree := os.Getenv(keyFree)
		apiKeyPaid := os.Getenv(keyPaid)
		if len(apiKeyFree) > 10 {
			data := strings.Split(apiKeyFree, ";")
			key := data[0]
			billDay, err := strconv.Atoi(data[1])
			if err != nil {
				return err
			}
			msK := msKey{
				Id:      int64(0), // db will set correct id
				Key:     keyFree,
				ApiKey:  key,
				Usage:   0,
				FreeKey: true,
				Month:   month,
				BillDay: billDay,
			}
			loadMicrosoftApiKeyToDb(msK)
		}
		if len(apiKeyPaid) > 10 {
			data := strings.Split(apiKeyPaid, ";")
			key := data[0]
			billDay, err := strconv.Atoi(data[1])
			if err != nil {
				return err
			}
			msK := msKey{
				Id:      int64(0), // db will set correct id
				Key:     keyPaid,
				ApiKey:  key,
				Usage:   0,
				FreeKey: false,
				Month:   month,
				BillDay: billDay,
			}
			loadMicrosoftApiKeyToDb(msK)
		}
		if len(apiKeyFree) < 10 && len(apiKeyPaid) < 10 {
			break
		}
	}
	return nil
}
