package wa

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bio426/monchisss/datasource"
	"github.com/bio426/monchisss/internal/core"
)

type WaSvc core.Service

var HttpClient = http.Client{Timeout: time.Second * 5}

type SvcProcessMessagePayload struct {
	Text      string
	Latitude  float32
	Longitude float32
}
type SvcProcessMessageParams struct {
	From    string
	To      string
	Type    string
	Payload SvcProcessMessagePayload
}

func (s *WaSvc) ProccessMessage(c context.Context, params SvcProcessMessageParams) error {
	tx, err := datasource.Postgres.BeginTx(c, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// get store info
	var (
		id      int32
		waId    string
		waToken string
	)
	row := tx.QueryRowContext(c,
		"select id, wa_id, wa_token from stores where wa_phone = $1",
		params.To,
	)
	if err := row.Scan(&id, &waId, &waToken); err != nil {
		return err
	}

    // get last message stage


	var lastConversationStage string
	// get last message
	row = tx.QueryRowContext(c,
		`
        select 
          co.stage
        from 
          conversations co 
          join customers cu on cu.id = co.customer 
        where 
          cu.wa_phone = $1 
        order by 
          co.created_at desc 
        limit 
          1
        `,
		params.From,
	)
	if err = row.Scan(&lastConversationStage); err != nil {
		if err == sql.ErrNoRows {
			createCustomer()
		} else {
			return err
		}
	}

	err = AnswerWithText(AnswerWithTextParams{
		PhoneId: waId,
		Token:   waToken,
		To:      params.From,
	})

	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func createCustomer() {

}

type AnswerWithTextParams struct {
	PhoneId string
	Token   string
	To      string
}

func AnswerWithText(params AnswerWithTextParams) error {
	req, err := newTextAnswerRequest(params.PhoneId, params.To, params.Token)
	if err != nil {
		return err
	}
	res, err := HttpClient.Do(req)
	if err != nil {
		return err
	}
	log.Println(res.Status)

	return nil
}

type graphBody struct {
	MessagingProduct string        `json:"messaging_product"`
	RecipientType    string        `json:"recipient_type"`
	To               string        `json:"to"`
	Type             string        `json:"type"`
	Text             graphBodyText `json:"text"`
}
type graphBodyText struct {
	PreviewUrl bool   `json:"preview_url"`
	Body       string `json:"body"`
}

func newTextAnswerRequest(phoneId, to, token string) (*http.Request, error) {
	requestUrl := fmt.Sprintf("https://graph.facebook.com/v19.0/%s/messages", phoneId)
	body := graphBody{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               to,
		Type:             "text",
		Text:             graphBodyText{PreviewUrl: false, Body: "Pong"},
	}
	marshalled, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, requestUrl, bytes.NewReader(marshalled))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	return req, nil
}

var Service = &WaSvc{}
