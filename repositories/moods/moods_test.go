package moods

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

var testMood = &models.Mood{
    Id:    13,
    Type:  1,
    Title: "Bono jugoso!",
    Owner: testUser,
}

func TestMoodsRepository_Create_NoErrorResult(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    mock.ExpectQuery("INSERT INTO moods").WithArgs(
        testMood.Type,
        testMood.Title,
        testMood.Owner.Id,
    ).WillReturnRows(sqlmock.NewRows([]string{"lastInsertedId"}).AddRow(13))

    myMoodsRepository := &moodsRepository{database: db}

    // execute our method
    if id, err := myMoodsRepository.Create(testMood); err != nil || id != 13 {
        t.Errorf("error was not expected while updating stats: %s", err)
    }

    // expectations
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expections: %s", err)
    }
}

func TestMoodsRepository_Create_ErrorResult(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    mock.ExpectQuery("INSERT INTO moods").WithArgs(
        testMood.Type,
        testMood.Title,
        testMood.Owner.Id,
    ).WillReturnError(errors.New("Expected Error"))

    myMoodsRepository := &moodsRepository{database: db}

    // execute our method
    if id, err := myMoodsRepository.Create(testMood); err == nil || id != -1 {
        t.Errorf("error was not expected while updating stats: %s", err)
    }

    // expectations
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expections: %s", err)
    }
}

func TestNewMoodsRepository(t *testing.T) {
    mockDB := &sql.DB{}

    // execute our method
    result := NewMoodsRepository(mockDB)

    // expectations
    assert.NotNil(t, result)
}
