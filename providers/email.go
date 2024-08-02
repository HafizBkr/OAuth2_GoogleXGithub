package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type EmailProvider struct {
	key             string
	emailSenderAddr string
}

func NewEmailProvider(key, emailSenderAddr string) *EmailProvider {
	return &EmailProvider{
		key:             key,
		emailSenderAddr: emailSenderAddr,
	}
}

func (p EmailProvider) SendSampleEmail(recipient, content string) error {
	payload := map[string]interface{}{
		"Recipients": map[string]interface{}{
			"To": []string{recipient},
		},
		"Content": map[string]interface{}{
			"From": p.emailSenderAddr,
			"Body": []map[string]interface{}{
				{
					"ContentType": "HTML",
					"Content":     content,
					"Charset":     "utf-8",
				},
			},
		},
	}
	bytesData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("Error while marshalling data: %w", err)
	}
	req, err := http.NewRequest("POST", "https://api.elasticemail.com/v4/emails/transactional", bytes.NewBuffer(bytesData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-ElasticEmail-ApiKey", p.key)
	if err != nil {
		return fmt.Errorf("Error while sending mail: constructing request failed: %w", err)
	}
	netClient := http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := netClient.Do(req)
	if err != nil {
		return fmt.Errorf("Error while sending mail: sending HTTP request failed: %w", err)
	}
	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Error while sending mail: reading response body failed: %w", err)
		}
		fmt.Println(string(body))
		return fmt.Errorf("Error while sending mail: Got HTTP %d from Elastic Email", resp.StatusCode)
	}
	return nil
}
