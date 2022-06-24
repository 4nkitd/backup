package storage

import (
	"testing"

	"github.com/4nkitd/gobackup/config"
	"github.com/stretchr/testify/assert"
)

func TestBase_newBase(t *testing.T) {
	model := config.ModelConfig{}
	archivePath := "/tmp/gobackup/test-storeage/foo.zip"
	base := newBase(model, archivePath)

	assert.Equal(t, base.archivePath, archivePath)
	assert.Equal(t, base.model, model)
	assert.Equal(t, base.viper, model.Viper)
	assert.Equal(t, base.keep, 0)
}
