package graph

import (
	"context"
	"fmt"

	"github.com/Nguyen-Tan-Dat/Vocabualries-Learning-API/models"
	"gorm.io/gorm"
)

// Resolver struct contains the database connection (e.g., GORM) and any other dependencies you might need
type Resolver struct {
	DB *gorm.DB
}

// Resolver for the Query type: vocabularies
func (r *Resolver) Vocabularies(ctx context.Context, topic string) ([]*models.Vocabulary, error) {
	var vocabularies []*models.Vocabulary

	// If a topic is provided, fetch vocabularies related to that topic
	if topic != "" {
		// Find vocabularies that are related to the specified topic
		err := r.DB.Preload("Topics").Where("topics.name = ?", topic).Find(&vocabularies).Error
		if err != nil {
			return nil, fmt.Errorf("failed to fetch vocabularies by topic: %w", err)
		}
	} else {
		// If no topic is provided, fetch all vocabularies
		err := r.DB.Preload("Topics").Find(&vocabularies).Error
		if err != nil {
			return nil, fmt.Errorf("failed to fetch vocabularies: %w", err)
		}
	}
	return vocabularies, nil
}

// Resolver for the Query type: topics
func (r *Resolver) Topics(ctx context.Context) ([]*models.Topic, error) {
	var topics []*models.Topic
	err := r.DB.Find(&topics).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch topics: %w", err)
	}
	return topics, nil
}

// Resolver for the Mutation type: createVocabulary
func (r *Resolver) CreateVocabulary(ctx context.Context, word string, phonetic string, meaning string, topicIds []int) (*models.Vocabulary, error) {
	vocabulary := &models.Vocabulary{
		Word:     word,
		Phonetic: phonetic,
		Meaning:  meaning,
	}

	// Start a transaction
	tx := r.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Create the vocabulary
	if err := tx.Create(vocabulary).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create vocabulary: %w", err)
	}

	// Link vocabulary with topics (many-to-many relationship)
	for _, topicId := range topicIds {
		// Add a record to the vocabularies_topics table
		if err := tx.Exec("INSERT INTO vocabularies_topics (topic, vocabulary) VALUES (?, ?)", topicId, vocabulary.ID).Error; err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to associate vocabulary with topic: %w", err)
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return vocabulary, nil
}

// Resolver for the Mutation type: createTopic
func (r *Resolver) CreateTopic(ctx context.Context, name string, userId string) (*models.Topic, error) {
	topic := &models.Topic{
		Name:   name,
		UserID: userId,
	}

	err := r.DB.Create(topic).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create topic: %w", err)
	}

	return topic, nil
}

// Resolver for the Mutation type: updateVocabulary
func (r *Resolver) UpdateVocabulary(ctx context.Context, id int, word *string, phonetic *string, meaning *string, topicIds []int) (*models.Vocabulary, error) {
	var vocabulary models.Vocabulary
	if err := r.DB.First(&vocabulary, id).Error; err != nil {
		return nil, fmt.Errorf("vocabulary not found: %w", err)
	}

	// Update fields if they are provided
	if word != nil {
		vocabulary.Word = *word
	}
	if phonetic != nil {
		vocabulary.Phonetic = *phonetic
	}
	if meaning != nil {
		vocabulary.Meaning = *meaning
	}

	// Save the updated vocabulary
	if err := r.DB.Save(&vocabulary).Error; err != nil {
		return nil, fmt.Errorf("failed to update vocabulary: %w", err)
	}

	// Update the vocabulary-topic relationships
	// First, clear old associations
	if err := r.DB.Exec("DELETE FROM vocabularies_topics WHERE vocabulary = ?", vocabulary.ID).Error; err != nil {
		return nil, fmt.Errorf("failed to clear old vocabulary-topic associations: %w", err)
	}

	// Add new associations
	for _, topicId := range topicIds {
		if err := r.DB.Exec("INSERT INTO vocabularies_topics (topic, vocabulary) VALUES (?, ?)", topicId, vocabulary.ID).Error; err != nil {
			return nil, fmt.Errorf("failed to associate vocabulary with topic: %w", err)
		}
	}

	return &vocabulary, nil
}

// Resolver for the Mutation type: deleteVocabulary
func (r *Resolver) DeleteVocabulary(ctx context.Context, id int) (bool, error) {
	var vocabulary models.Vocabulary
	if err := r.DB.First(&vocabulary, id).Error; err != nil {
		return false, fmt.Errorf("vocabulary not found: %w", err)
	}

	if err := r.DB.Delete(&vocabulary).Error; err != nil {
		return false, fmt.Errorf("failed to delete vocabulary: %w", err)
	}

	return true, nil
}
