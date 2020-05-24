package generate

import "remind-go/handlers"

func GetPhone(phone string) *handlers.Phone {
	return &handlers.Phone{Phone: phone}
}
