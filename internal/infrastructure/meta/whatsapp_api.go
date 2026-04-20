package meta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"waba/internal/domain/model"
	"waba/internal/domain/repository"
	"waba/internal/infrastructure/config"
)

type whatsappAPI struct {
	cfg    *config.Config
	client *http.Client
}

// NewWhatsAppAPI is a struct that takes a config and a client as arguments.
func NewWhatsAppAPI(cfg *config.Config) repository.WhatsAppRepository {
	return &whatsappAPI{
		cfg:    cfg,
		client: &http.Client{},
	}
}

// Create a docstring for create template and send template message functions
func (w *whatsappAPI) CreateTemplate(template model.Template) error {
	url := fmt.Sprintf("%s/%s/message_templates", w.cfg.BaseURL, w.cfg.BusinessID)

	payload := map[string]interface{}{
		"name":     template.Name,
		"language": template.Language,
		"category": template.Category,
		"components": []map[string]interface{}{
			{
				"type": "BODY",
				"text": template.BodyText,
			},
		},
	}

	return w.doRequest(url, payload)
}

// sends a template message to a WhatsApp phone number.
func (w *whatsappAPI) SendTemplateMessage(message model.TemplateMessage) error {
	url := fmt.Sprintf("%s/%s/messages", w.cfg.BaseURL, w.cfg.PhoneNumberID)

	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                message.To,
		"type":              "template",
		"template": map[string]interface{}{
			"name": message.TemplateName,
			"language": map[string]string{
				"code": message.Language,
			},
		},
	}

	return w.doRequest(url, payload)
}

// doRequest makes a POST request to the WhatsApp Business Cloud API.
func (w *whatsappAPI) doRequest(url string, body interface{}) error {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+w.cfg.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := w.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("meta api error: %s", resp.Status)
	}

	return nil
}
