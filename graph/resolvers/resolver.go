package resolvers

import (
	"errors"

	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/services"
	"golang.org/x/net/context"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TopicService   services.TopicService
	EnglishService services.EnglishService
}

func getUserIDFromContext(ctx context.Context) (int32, error) {
	userID, ok := ctx.Value("userID").(int32) // Ensure the userID is of type int32
	if !ok {
		return 0, errors.New("unauthorized: userID not found in context")
	}
	return userID, nil
}
