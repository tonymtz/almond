package moods

import (
    "github.com/tonymtz/almond/models"
    "database/sql"
    "errors"
)

type MoodsRepository interface {
    Create(mood *models.Mood) (int64, error)
}

type moodsRepository struct {
    MoodsRepository
    database *sql.DB
}

func (this *moodsRepository) Create(newMood *models.Mood) (int64, error) {
    return -1, errors.New("Not Implemented")
}

func NewMoodsRepository(database *sql.DB) MoodsRepository {
    return &moodsRepository{database: database}
}
