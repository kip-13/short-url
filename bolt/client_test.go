package bolt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDB(t *testing.T) {
	boltClient := NewBoltClient("short-url.db", 0600)
	db, err := boltClient.getDB()
	assert.Nil(t, err)
	assert.NotNil(t, db)
}

func TestSetAndGetValue(t *testing.T) {
	boltClient := NewBoltClient("short-url.db", 0600)
	boltClient.Set("short", "sag/2$#", []byte("1"))
	boltClient.Set("short", "!@##$", []byte("2"))
	boltClient.Set("short", "+_()*(", []byte("3"))
	v1, err := boltClient.Get("short", "sag/2$#")
	assert.Nil(t, err)
	assert.EqualValues(t, v1, "1")
	v2, err := boltClient.Get("short", "!@##$")
	assert.Nil(t, err)
	assert.EqualValues(t, v2, "2")
	v3, err := boltClient.Get("short", "+_()*(")
	assert.Nil(t, err)
	assert.EqualValues(t, v3, "3")

}
