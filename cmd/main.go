package main

import (
	"log"
	"waba/internal/domain/model"
	"waba/internal/infrastructure/config"
	"waba/internal/infrastructure/meta"
	"waba/internal/usecase"
)

// main is an entry point for WhatsApp CRM SaaS application.
// It initializes config, then initializes WhatsApp API repository.
func main() {
	cfg := config.Load()
	repo := meta.NewWhatsAppAPI(cfg)
	uc := usecase.NewWhatsAppUsecase(repo)

	template := model.Template{
		Name:     "order_confirmation",
		Language: "en_US",
		Category: "TRANSACTIONAL",
		BodyText: "Your order has been confirmed",
	}

	err := uc.CreateTemplate(template)
	if err != nil {
		log.Fatal(err)
	}

	message := model.TemplateMessage{
		To:           "6281234567890",
		TemplateName: "order_confirmation",
		Language:     "en_US",
	}

	err = uc.SendTemplate(message)
	if err != nil {
		log.Fatal(err)
	}
}