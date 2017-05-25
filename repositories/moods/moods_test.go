package moods

import (
    "testing"
    "database/sql"
    "github.com/stretchr/testify/assert"
)

func TestNewMoodsRepository(t *testing.T) {
    mockDB := &sql.DB{}

    // execute our method
    result := NewMoodsRepository(mockDB)

    // expectations
    assert.NotNil(t, result)
}
