package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendDiscordNotification(webhookURL, message string) error {
    payload := map[string]string{"content": message}
    body, _ := json.Marshal(payload)

    resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(body))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
        return fmt.Errorf("discord webhook failed: %s", resp.Status)
    }
    return nil
}
