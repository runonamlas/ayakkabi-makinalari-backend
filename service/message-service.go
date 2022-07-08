package service

import (
	"github.com/mashingan/smapping"
	"github.com/runonamlas/ayakkabi-makinalari-backend/dto"
	"github.com/runonamlas/ayakkabi-makinalari-backend/entity"
	"github.com/runonamlas/ayakkabi-makinalari-backend/repository"
	"log"
)

type MessageService interface {
	Insert(p dto.MessageCreateDTO) entity.Message
	Update(p dto.MessageUpdateDTO) entity.Message
	Delete(p entity.Message)
	All(cityID uint64) []entity.Message
	AllMessages() []entity.Message
	FindByID(messageID uint64) entity.Message
	IsAllowedToEdit(cityID string, messageID uint64) bool
}

type messageService struct {
	messageRepository repository.MessageRepository
}

func NewMessageService(messageRepo repository.MessageRepository) MessageService {
	return &messageService{
		messageRepository: messageRepo,
	}
}

func (service *messageService) Insert(p dto.MessageCreateDTO) entity.Message {
	message := entity.Message{}
	err := smapping.FillStruct(&message, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.messageRepository.InsertMessage(message)
	return res
}

func (service *messageService) Update(p dto.MessageUpdateDTO) entity.Message {
	message := entity.Message{}
	err := smapping.FillStruct(&message, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.messageRepository.UpdateMessage(message)
	return res
}

func (service *messageService) Delete(p entity.Message) {
	service.messageRepository.DeleteMessage(p)
}

func (service *messageService) All(cityID uint64) []entity.Message {
	return service.messageRepository.AllMessage(cityID)
}

func (service *messageService) AllMessages() []entity.Message {
	return service.messageRepository.AllMessages()
}

func (service *messageService) FindByID(messageID uint64) entity.Message {
	return service.messageRepository.FindMessageByID(messageID)
}

func (service *messageService) IsAllowedToEdit(cityID string, messageID uint64) bool {
	//p := service.messageRepository.FindMessageByID(messageID)
	//id := fmt.Sprintf("%v", p.CityID)
	//return cityID == id
	return false
}
