package repository

import "waba/internal/domain/model"

type WhatsAppRepository interface {
	CreateTemplate(template model.Template) error
	SendTemplateMessage(message model.TemplateMessage) error
}
