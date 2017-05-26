package moods

import (
    "github.com/tonymtz/almond/models"
    "database/sql"
    "errors"
)

type MoodsRepository interface {
    Create(mood *models.Mood) (int64, error)
    GetAllByOwnerId(int64) ([]*models.Mood, error)
    Update(*models.Mood) (error)
}

type moodsRepository struct {
    MoodsRepository
    database *sql.DB
}

func (this *moodsRepository) Create(newMood *models.Mood) (int64, error) {
    var lastInsertedId int64

    err := this.database.QueryRow(
        "INSERT INTO moods (type, title, owner_id) VALUES ($1, $2, $3) RETURNING id",
        newMood.Type,
        newMood.Title,
        newMood.Owner.Id,
    ).Scan(&lastInsertedId)

    if err != nil {
        return -1, err
    }

    return lastInsertedId, nil
}

func (this *moodsRepository) GetAllByOwnerId(ownerId int64) ([]*models.Mood, error) {
    return nil, errors.New("Not Implemented")
}

func (this *moodsRepository) Update(newMood *models.Mood) error {
    return errors.New("Not Implemented")
}

func NewMoodsRepository(database *sql.DB) MoodsRepository {
    return &moodsRepository{database: database}
}
