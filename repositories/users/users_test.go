package users

import (
    "testing"
    "database/sql"
    "gopkg.in/DATA-DOG/go-sqlmock.v1"
    "github.com/stretchr/testify/assert"
    "github.com/tonymtz/almond/models"
    "errors"
)

var testUser = &models.User{
    Id:             9,
    GoogleId:       "uid:111111",
    DisplayName:    "Mr Test User",
    ProfilePicture: "http://image.url/happy.png",
    Token:          "t:11111",
}

func TestUsersRepository_Create_NoErrorResult(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    mock.ExpectQuery("INSERT INTO users").WithArgs(
        testUser.GoogleId,
        testUser.DisplayName,
        testUser.ProfilePicture,
        testUser.Token,
    ).WillReturnRows(sqlmock.NewRows([]string{"lastInsertedId"}).AddRow(5))

    myUsersRepository := &usersRepository{database: db}

    // execute our method
    if id, err := myUsersRepository.Create(testUser); err != nil || id != 5 {
        t.Errorf("error was not expected while updating stats: %s", err)
    }

    // expectations
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expections: %s", err)
    }
}

func TestUsersRepository_Create_ErrorResult(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    mock.ExpectQuery("INSERT INTO users").WithArgs(
        testUser.GoogleId,
        testUser.DisplayName,
        testUser.ProfilePicture,
        testUser.Token,
    ).WillReturnError(errors.New("Expected Error"))

    myUsersRepository := &usersRepository{database: db}

    // execute our method
    if id, err := myUsersRepository.Create(testUser); err == nil || id != -1 {
        t.Errorf("error was not expected while updating stats: %s", err)
    }

    // expectations
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expections: %s", err)
    }
}

func TestUsersRepository_GetOneById_NoErrorResult(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    mock.ExpectQuery("SELECT (.+) FROM users WHERE").WithArgs(
        testUser.Id,
    ).WillReturnRows(
        sqlmock.NewRows([]string{
            "Id",
            "GoogleId",
            "CreatedAt",
            "DisplayName",
            "ProfilePicture",
            "Token",
        }).AddRow(
            testUser.Id,
            testUser.GoogleId,
            testUser.CreatedAt,
            testUser.DisplayName,
            testUser.ProfilePicture,
            testUser.Token,
        ))

    myUsersRepository := &usersRepository{database: db}

    // execute our method
    if user, err := myUsersRepository.GetOneById(testUser.Id); err != nil || user == nil {
        t.Errorf("error was not expected while updating stats: %s", err)
    }

    // expectations
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expections: %s", err)
    }
}

func TestUsersRepository_GetOneById_ErrorResult(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    mock.ExpectQuery("SELECT (.+) FROM users WHERE").WithArgs(
        testUser.Id,
    ).WillReturnError(errors.New("Expected Error"))

    myUsersRepository := &usersRepository{database: db}

    // execute our method
    if user, err := myUsersRepository.GetOneById(testUser.Id); err == nil || user != nil {
        t.Errorf("error was not expected while updating stats: %s", err)
    }

    // expectations
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expections: %s", err)
    }
}

func TestNewUsersRepository(t *testing.T) {
    mockDB := &sql.DB{}

    // execute our method
    result := NewUsersRepository(mockDB)

    // expectations
    assert.NotNil(t, result)
}
