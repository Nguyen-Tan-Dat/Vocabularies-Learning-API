package repositories

import (
	"context"

	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/graph/model"
	"gorm.io/gorm"
)

type EnglishRepository struct {
	DB *gorm.DB
}

func NewEnglishRepository(db *gorm.DB) *EnglishRepository {
	return &EnglishRepository{
		DB: db,
	}
}

// CREATE
func (r *EnglishRepository) Create(
	ctx context.Context,
	word string,
	phonetic *string,
	audio *string,
) (*model.English, error) {

	english := &model.English{
		Word:     word,
		Phonetic: phonetic,
		Audio:    audio,
	}

	if err := r.DB.
		WithContext(ctx).
		Create(english).
		Error; err != nil {

		return nil, err
	}

	return english, nil
}

// READ - Get by ID
func (r *EnglishRepository) GetByID(
	ctx context.Context,
	id int,
) (*model.English, error) {

	var english model.English

	if err := r.DB.
		WithContext(ctx).
		First(&english, id).
		Error; err != nil {

		return nil, err
	}

	return &english, nil
}

// READ - Get all
func (r *EnglishRepository) GetAll(
	ctx context.Context,
) ([]*model.English, error) {

	var list []*model.English

	if err := r.DB.
		WithContext(ctx).
		Find(&list).
		Error; err != nil {

		return nil, err
	}

	return list, nil
}

// UPDATE (partial update – đúng GraphQL UpdateEnglishInput)
func (r *EnglishRepository) Update(
	ctx context.Context,
	input *model.UpdateEnglishInput,
) (*model.English, error) {

	updates := make(map[string]interface{})

	if input.Word != nil {
		updates["word"] = *input.Word
	}

	if input.Phonetic != nil {
		updates["phonetic"] = *input.Phonetic
	}

	if input.Audio != nil {
		updates["audio"] = *input.Audio
	}

	if err := r.DB.
		WithContext(ctx).
		Model(&model.English{}).
		Where("id = ?", input.ID).
		Updates(updates).
		Error; err != nil {

		return nil, err
	}

	return r.GetByID(ctx, int(input.ID))
}

// DELETE
func (r *EnglishRepository) Delete(
	ctx context.Context,
	id int,
) error {

	if err := r.DB.
		WithContext(ctx).
		Delete(&model.English{}, id).
		Error; err != nil {

		return err
	}

	return nil
}
