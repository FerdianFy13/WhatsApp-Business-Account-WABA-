package usecase

import (
	"waba/internal/domain/model"
	"waba/internal/domain/repository"
)

type WhatsAppUsecase interface {
	CreateTemplate(input model.Template) error
	SendTemplate(input model.TemplateMessage) error
}

type whatsappUsecase struct {
	repo repository.WhatsAppRepository
}

// NewWhatsAppUsecase returns a WhatsAppUsecase object with the given repo.
func NewWhatsAppUsecase(repo repository.WhatsAppRepository) WhatsAppUsecase {
	return &whatsappUsecase{repo: repo}
}

// Create a template on WhatsApp Business Cloud API
func (u *whatsappUsecase) CreateTemplate(input model.Template) error {
	return u.repo.CreateTemplate(input)
}

// Send a template message to a WhatsApp phone number.
func (u *whatsappUsecase) SendTemplate(input model.TemplateMessage) error {
	return u.repo.SendTemplateMessage(input)
}