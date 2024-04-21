package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"translator/pkg/utils"

	"github.com/gin-gonic/gin"
)

type TranslateResponse struct {
	Text string `json:"text" swaggertype:"string" example:"test"`
}

// @BasePath /api

type YandexTranslateBody struct {
	TargetLanguageCode string   `json:"targetLanguageCode"`
	SourceLanguageCode string   `json:"sourceLanguageCode"`
	Texts              []string `json:"texts"`
	FolderId           string   `json:"folderId"`
}

type YandexTranslateResponse struct {
	Translations []struct {
		Text                 string `json:"text"`
		DetectedLanguageCode string `json:"detectedLanguageCode"`
	} `json:"translations"`
}

const yandexAPI = "https://translate.api.cloud.yandex.net/translate/v2/translate"

// Translate godoc
// @Summary translate something
// @Schemes
// @Param toLang path string true "Translate to language"
// @Param fromLang path string true "Translate from language"
// @Param index query int true "Index" default(0)
// @Description Translate input text `fromLang` `toLang`
// @Tags api
// @Accept json
// @Produce json
// @Success 200 {object} api.TranslateResponse
// @Router /translate/{name} [post]
func Translate(c *gin.Context) {
	//ctx := c.Request.Context()
	res := c.Writer
	toLang := c.Param("toLang")
	fromLang := c.Param("fromLang")
	//index := c.DefaultQuery("index", "0") // 0 - unlimited

	cfg, err := utils.GetConfig(c)
	if err != nil {
		http.Error(res, "Config load: "+err.Error(), http.StatusInternalServerError)
		return
	}

	text, err := io.ReadAll(c.Request.Body)
	if err != nil {
		http.Error(res, "Faulty input: "+err.Error(), http.StatusInternalServerError)
		return
	}

	body := YandexTranslateBody{
		FolderId:           cfg.FolderID,
		TargetLanguageCode: toLang,
		SourceLanguageCode: fromLang,
		Texts:              []string{string(text)},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		http.Error(res, "Cannot serialize request body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	r, err := http.NewRequest(http.MethodPost, yandexAPI, bytes.NewBuffer(jsonBody))
	if err != nil {
		http.Error(res, "Cannot build request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Api-Key %s", cfg.ApiKey))

	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		http.Error(res, "YCloud: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	var responseJson YandexTranslateResponse

	err = json.NewDecoder(response.Body).Decode(&responseJson)
	if err != nil {
		http.Error(res, "YCloud cant read body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(responseJson.Translations) == 0 {
		http.Error(res, "Empty response text", http.StatusInternalServerError)
		return
	}

	utils.Unisponse(c, responseJson.Translations[0].Text)
}
