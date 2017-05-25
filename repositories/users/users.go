package users

import (
    "database/sql"
    "github.com/tonymtz/almond/models"
)

type UsersRepository interface {
    Create(*models.User) (int64, error)
    GetOneById(int64) (*models.User, error)
    GetOneByGoogleId(string) (*models.User, error)
    Update(*models.User) (error)
}

type usersRepository struct {
    UsersRepository
    database *sql.DB
}

func (this *usersRepository) Create(newUser *models.User) (int64, error) {
    var lastInsertedId int64

    err := this.database.QueryRow(
        "INSERT INTO users (google_id, display_name, profile_picture, token, email) VALUES ($1, $2, $3, $4, '') RETURNING id",
        newUser.GoogleId,
        newUser.DisplayName,
        newUser.ProfilePicture,
        newUser.Token,
    ).Scan(&lastInsertedId)

    if err != nil {
        return -1, err
    }

    return lastInsertedId, nil
}

func (this *usersRepository) GetOneById(userId int64) (*models.User, error) {
    myUser := &models.User{}

    err := this.database.QueryRow(
        "SELECT id, google_id, created_at, display_name, profile_picture, token FROM users WHERE id = $1",
        userId,
    ).Scan(
        &myUser.Id,
        &myUser.GoogleId,
        &myUser.CreatedAt,
        &myUser.DisplayName,
        &myUser.ProfilePicture,
        &myUser.Token,
    )

    if err != nil {
        return nil, err
    }

    return myUser, nil
}

func (this *usersRepository) GetOneByGoogleId(googleId string) (*models.User, error) {
    myUser := &models.User{}

    err := this.database.QueryRow(
        "SELECT id, google_id, created_at, display_name, profile_picture, COALESCE(token, '') as token FROM users WHERE google_id = $1",
        googleId,
    ).Scan(
        &myUser.Id,
        &myUser.GoogleId,
        &myUser.CreatedAt,
        &myUser.DisplayName,
        &myUser.ProfilePicture,
        &myUser.Token,
    )

    if err != nil {
        return nil, err
    }

    return myUser, nil
}

func (this *usersRepository) Update(myUser *models.User) (error) {
    result, err := this.database.Exec(
        "UPDATE users SET token = $1 WHERE id = $2",
        myUser.Token,
        myUser.Id,
    )

    if err != nil {
        return err
    }

    _, err = result.RowsAffected()

    if err != nil {
        return err
    }

    return nil
}

func NewUsersRepository(database *sql.DB) UsersRepository {
    return &usersRepository{
        database: database,
    }
}
