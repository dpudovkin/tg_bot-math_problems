package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	models "p134d/tg_math_bot/models"
	"strconv"
)

func Start(botUrl string) {
	offset := -1
	for {
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			log.Println("Unexpected error", err.Error())
		}
		for _, update := range updates {
			go respond(botUrl, update)
			offset = update.UpdateId + 1
		}
		if len(updates) > 0 {
			fmt.Println(updates)
		}

	}
}

func getUpdates(botUrl string, offset int) ([]models.Update, error) {
	response, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var restResponse models.Response
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}
