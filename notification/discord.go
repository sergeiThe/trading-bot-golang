package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Discord struct {
	WebhookURL string
}



func (n Discord) Notify(message string) error {
	payload := map[string]string{"content": message}
	body, _ := json.Marshal(payload)

	resp, err := http.Post(n.WebhookURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("discord webhook failed: %s", resp.Status)
	}

	log.Println("Discord notified")
	return nil
}
