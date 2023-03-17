package ctx

import (
	"go/build"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wuhc1010/go-zero/core/stringx"
	"github.com/wuhc1010/go-zero/tools/goctl/util/pathx"
)

func TestProjectFromGoPath(t *testing.T) {
	dft := build.Default
	gp := dft.GOPATH
	if len(gp) == 0 {
		return
	}
	projectName := stringx.Rand()
	dir := filepath.Join(gp, "src", projectName)
	err := pathx.MkdirIfNotExist(dir)
	if err != nil {
		return
	}
	defer func() {
		_ = os.RemoveAll(dir)
	}()

	ctx, err := projectFromGoPath(dir)
	assert.Nil(t, err)
	assert.Equal(t, dir, ctx.Dir)
	assert.Equal(t, projectName, ctx.Path)
}

func TestProjectFromGoPathNotInGoSrc(t *testing.T) {
	dft := build.Default
	gp := dft.GOPATH
	if len(gp) == 0 {
		return
	}
	projectName := stringx.Rand()
	dir := filepath.Join(gp, "src", projectName)
	err := pathx.MkdirIfNotExist(dir)
	if err != nil {
		return
	}
	defer func() {
		_ = os.RemoveAll(dir)
	}()

	_, err = projectFromGoPath("testPath")
	assert.NotNil(t, err)
}
