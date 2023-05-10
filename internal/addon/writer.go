package addon

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/stoewer/go-strcase"
	"gopkg.in/yaml.v3"
)

func readSqlcConf(baseDir string) (SqlcConf, error) {
	filename, _ := filepath.Abs(fmt.Sprintf("%s/sqlc.yaml", baseDir))
	file, err := os.ReadFile(filename)
	if err != nil {
		return SqlcConf{}, err
	}

	var conf SqlcConf

	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		return SqlcConf{}, err
	}

	return conf, nil
}

func readSqlcAddonConf(baseDir string) (SqlcAddonConf, error) {
	filename, _ := filepath.Abs(fmt.Sprintf("%s/sqlc-addon.yaml", baseDir))
	file, err := os.ReadFile(filename)
	if err != nil {
		return SqlcAddonConf{}, err
	}

	var conf SqlcAddonConf

	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		return SqlcAddonConf{}, err
	}

	return conf, nil
}

func extractVersionFromMigrationFileName(filename string) (int, error) {
	r := regexp.MustCompile(`v(?P<version>\d)_.*`)

	subMatch := r.FindStringSubmatch(filename)
	if len(subMatch) == 0 {
		return 0, fmt.Errorf("migration file does not have a valid name : %s", filename)
	}

	version, err := strconv.Atoi(subMatch[1])
	if err != nil {
		return 0, err
	}

	return version, nil
}

func extractNameFromMigrationFileName(filename string) string {
	camelCaseName := strcase.LowerCamelCase(filename)

	return camelCaseName[:len(camelCaseName)-4]
}

func listMigrationFilesInFolder(folder string) ([]MigrationFile, error) {
	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	filenames := make([]MigrationFile, len(files))
	for i, file := range files {
		version, err := extractVersionFromMigrationFileName(file.Name())
		if err != nil {
			return nil, err
		}

		content, err := os.ReadFile(fmt.Sprintf("%s/%s", folder, file.Name()))
		if err != nil {
			return nil, err
		}

		filenames[i] = MigrationFile{
			Name:    extractNameFromMigrationFileName(file.Name()),
			Version: version,
			Content: strings.Trim(string(content), " \n"),
			Last:    i == len(files)-1,
		}
	}

	return filenames, nil
}

func WriteTemplate(baseDir string, templateName string, sqlcAddonVersion string) error {
	conf, err := readSqlcConf(baseDir)
	if err != nil {
		return err
	}
	addonConf, err := readSqlcAddonConf(baseDir)
	if err != nil {
		return err
	}

	migrationFiles, err := listMigrationFilesInFolder(fmt.Sprintf("%s/%s", baseDir, conf.Sql[0].Schema))
	if err != nil {
		return err
	}

	err = os.MkdirAll(fmt.Sprintf("%s/%s", baseDir, addonConf.AddonOut), os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("%s/%s/%s", baseDir, addonConf.AddonOut, templateName[:len(templateName)-7]))
	if err != nil {
		return err
	}
	defer file.Close()

	values := TemplateValues{
		DatabaseName:     addonConf.DatabaseName,
		MigrationFiles:   migrationFiles,
		SqlcAddonVersion: sqlcAddonVersion,
	}

	tmplFS, _ := fs.Sub(templates, "templates")
	t := template.Must(template.ParseFS(tmplFS, templateName))
	err = t.Execute(file, values)
	if err != nil {
		return err
	}

	return nil
}
