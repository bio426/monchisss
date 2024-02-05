package wa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/bio426/monchisss/internal/core"
)

type WaSvc core.Service

type ReqBody struct {
	MessagingProduct string   `json:"messaging_product"`
	RecipientType    string   `json:"recipient_type"`
	To               string   `json:"to"`
	Type             string   `json:"type"`
	Text             TextBody `json:"text"`
}
type TextBody struct {
	PreviewUrl bool   `json:"preview_url"`
	Body       string `json:"body"`
}

var HttpClient = http.Client{Timeout: time.Second * 10}

func (svc *WaSvc) Answer(c context.Context) error {
	senderId := "225734773951886"

	accessToken := "EAAY7kYjNhu4BOxEcFolDl2bO4kWHKDVni5D0L49ISe7uB3vGt0mWyxXBmVDygfcPp4P7P7wU9jeMDwHS5ZB1ZBEpymR7wg0RBp04HIZCz7JYlZCDPFxKBRVLfHdB5xVcKunCd8UK9hltqte5GfZCgM9ZCYSXj2zxUZAviGnj4UNoxMJQ24kDFqSbfFpR9XttGiDkVewPkiaRDkIDbhI5mNXS5CoG8gGCLuh0mUZD"

	log.Println("start answer")
	requestUrl := fmt.Sprintf("https://graph.facebook.com/v19.0/%s/messages", senderId)
	body := ReqBody{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               "51910831282",
		Type:             "text",
		Text:             TextBody{PreviewUrl: false, Body: "Answer from monchisss"},
	}
	marshalled, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewReader(marshalled))
	authBearer := fmt.Sprintf("Bearer %s", accessToken)
	req.Header.Add("Authorization", authBearer)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return err
	}
	res, err := HttpClient.Do(req)
	if err != nil {
		return err
	}
	// print res body
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	log.Println(res.StatusCode)
	log.Println(string(resBytes))
	return nil
}

var Service = &WaSvc{}
