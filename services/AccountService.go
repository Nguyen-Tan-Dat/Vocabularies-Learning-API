package services

import (
	"context"

	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/graph/model"
	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/repositories"
)

type AccountService struct {
	Repo repositories.AccountRepository
}

func (s AccountService) Create(ctx context.Context, input model.NewAccountInput, userID int32) (*model.Account, error) {
	return s.Repo.Create(ctx, input, userID)
}

func (s AccountService) All(ctx context.Context, userID int32) ([]*model.Account, error) {
	return s.Repo.GetByUserID(ctx, userID)
}

func (s AccountService) Get(ctx context.Context, userID int32, id int32) (*model.Account, error) {
	return s.Repo.GetByUserIdAndID(ctx, userID, id)
}

func (s AccountService) Delete(ctx context.Context, userID int32, id int32) (bool, error) {
	return s.Repo.Delete(ctx, userID, id)
}

func (s AccountService) Update(ctx context.Context, input model.UpdateAccountInput, userID int32) (*model.Account, error) {
	return s.Repo.Update(ctx, input, userID)
}
