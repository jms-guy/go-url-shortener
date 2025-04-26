package store

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.client != nil)
}

func TestInsertAndRetrieval(t *testing.T) {
	short := "shorturl.com"
	long := "longlonglongurl.com"

	saveErr := SaveUrlMap(short, long)

	retrievedUrl, err := GetInitialUrl(short)

	assert.Equal(t, saveErr, nil)
	assert.Equal(t, long, retrievedUrl)
	assert.Equal(t, err, nil)
}