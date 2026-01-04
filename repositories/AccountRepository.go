package repositories

import (
	"context"

	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/graph/model"
	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
}

func (r AccountRepository) Create(ctx context.Context, input model.NewAccountInput, userID int32) (*model.Account, error) {
	account := &model.Account{
		UserName: input.UserName,
		Password: input.Password,
		Allocate: input.Allocate,
		OfUser:   userID,
	}
	if err := r.DB.WithContext(ctx).Create(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (r AccountRepository) GetByUserID(ctx context.Context, userId int32) ([]*model.Account, error) {
	var acounts []*model.Account
	if err := r.DB.WithContext(ctx).Where("of_user = ?", userId).Find(&acounts).Error; err != nil {
		return nil, err
	}
	return acounts, nil
}

func (r AccountRepository) GetByUserIdAndID(ctx context.Context, userID int32, id int32) (*model.Account, error) {
	var acounts *model.Account
	if err := r.DB.WithContext(ctx).Where("of_user = ? and id= ?", userID, id).Find(&acounts).Error; err != nil {
		return nil, err
	}
	return acounts, nil
}

func (r AccountRepository) Delete(ctx context.Context, userID int32, id int32) (bool, error) {
	topic := &model.Account{
		ID:     id,
		OfUser: userID,
	}
	if err := r.DB.WithContext(ctx).Delete(topic).Error; err != nil {
		return false, nil
	}
	return true, nil
}

func (r AccountRepository) Update(ctx context.Context, input model.UpdateAccountInput, userID int32) (*model.Account, error) {
	account := &model.Account{
		ID:       input.ID,
		UserName: input.UserName,
		Password: input.Password,
		Allocate: input.Allocate,
		OfUser:   userID,
	}
	if err := r.DB.WithContext(ctx).Save(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}
