package compressor

import (
	"testing"

	"github.com/4nkitd/gobackup/helper"
	"github.com/stretchr/testify/assert"
)

func TestTar_options(t *testing.T) {
	ctx := &Tar{}
	opts := ctx.options()
	if helper.IsGnuTar {
		assert.Equal(t, opts[0], "--ignore-failed-read")
		assert.Equal(t, opts[1], "-cf")
	} else {
		assert.Equal(t, opts[0], "-cf")
	}

}
