package addon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_readConf(t *testing.T) {
	conf, err := readSqlcConf("../../test")
	if err != nil {
		assert.Fail(t, "Can't read conf", "%v", err)
	}

	assert.Len(t, conf.Sql, 1)
	assert.Equal(t, "db/migrations", conf.Sql[0].Schema)
}

func TestWriteTemplate(t *testing.T) {
	err := WriteTemplate("../../test", "migration.go.gotmpl", "v0.0.0")
	if err != nil {
		assert.Fail(t, "Can't write", "%v", err)
	}

	err = WriteTemplate("../../test", "connect.go.gotmpl", "v0.0.0")
	if err != nil {
		assert.Fail(t, "Can't write", "%v", err)
	}
}

func Test_listFilesInFolder(t *testing.T) {

	folder := "../../test/db/migrations"

	filesInFolder, err := listMigrationFilesInFolder(folder)
	if err != nil {
		assert.Fail(t, "Can't read files", "%v", err.Error())
	}

	assert.Len(t, filesInFolder, 2)
	assert.Equal(t, "v1Init", filesInFolder[0].Name)
	assert.Equal(t, 1, filesInFolder[0].Version)
	assert.Equal(t, "v2AddThumbnails", filesInFolder[1].Name)
	assert.Equal(t, 2, filesInFolder[1].Version)

}
