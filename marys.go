package marys

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

const (
	timeout  = 5 * time.Second
	tmplName = "Notification"
	tmpl     = `
Dari: {{.Sender}}

{{.Message}}
`
)

var (
	bot             *TelegramBot
	recipientID     int
	messageTemplate *template.Template
)

// TelegramBot acts as telegram bot.
type TelegramBot struct {
	client *http.Client
	url    string
	token  string
}

// Notification represents a notification data.
type Notification struct {
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

// TelegramMessage represents a telegram message.
type TelegramMessage struct {
	ChatID                int    `json:"chat_id"`
	Text                  string `json:"text"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
}

func init() {
	tgToken := os.Getenv("TELEGRAM_TOKEN")
	if tgToken == "" {
		log.Fatalln("TELEGRAM_TOKEN is not set")
	}

	tgURL := os.Getenv("TELEGRAM_URL")
	if tgURL == "" {
		log.Fatalln("TELEGRAM_URL is not set")
	}

	var err error
	recipientID, err = strconv.Atoi(os.Getenv("TELEGRAM_RECIPIENT_ID"))
	if err != nil || recipientID == 0 {
		log.Fatalln("TELEGRAM_RECIPIENT_ID is not set properly")
	}

	messageTemplate = template.Must(template.New(tmplName).Parse(tmpl))

	client := &http.Client{
		Timeout: timeout,
	}

	bot = &TelegramBot{
		url:    tgURL,
		token:  tgToken,
		client: client,
	}
}

// ReceiveNotification receives notification.
func ReceiveNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`Only POST will be processed. Otherwise, it returns status OK`))
		return
	}

	var notif Notification
	if err := json.NewDecoder(r.Body).Decode(&notif); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	if notif.Sender == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`Sender can't be empty`))
		return
	}

	if err := sendNotifToTelegram(&notif); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`Success`))
}

func sendNotifToTelegram(notif *Notification) error {
	var out bytes.Buffer
	if err := messageTemplate.Execute(&out, notif); err != nil {
		return err
	}

	message := &TelegramMessage{
		ChatID:                recipientID,
		Text:                  out.String(),
		DisableWebPagePreview: true,
	}

	url := fmt.Sprintf("%s%s/sendMessage", bot.url, bot.token)
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := bot.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		log.Println(err)
		log.Println(string(body))
		return fmt.Errorf("response status code is %d", resp.StatusCode)
	}
	return nil
}
