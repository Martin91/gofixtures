package fixtures

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func (f *Fixtures) loadWalker(path string, d os.FileInfo, _ error) error {
	if d.IsDir() {
		return nil
	}

	if filepath.Ext(path) != ".yaml" && filepath.Ext(path) != ".yml" {
		log.Printf("skip loading %s because it has no a .yaml or .yml extention name", path)
		return nil
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("load yaml file failed: %s", err.Error())
		return err
	}

	return f.Parse(content)
}

func (f *Fixtures) Load() error {
	return filepath.Walk(f.path, f.loadWalker)
}

func (f *Fixtures) Parse(content []byte) error {
	definitions := map[string]interface{}{}
	err := yaml.Unmarshal(content, definitions)
	if err != nil {
		log.Panicf("invalid yaml error: %s", err.Error())
	}

	for collectionName, collectionDefinitionsIface := range definitions {
		collectionDefinitions := collectionDefinitionsIface.(map[interface{}]interface{})
		collection := &Collection{
			DbName:    collectionDefinitions["db"].(string),
			TableName: collectionDefinitions["table_name"].(string),
			Rows:      map[string]*Fixture{},
		}
		rows := collectionDefinitions["rows"].(map[interface{}]interface{})
		for rowName, rowDef := range rows {
			fixture := &Fixture{
				Columns: map[string]interface{}{},
			}
			columns := rowDef.(map[interface{}]interface{})
			for columnName, columnValue := range columns {
				fixture.Columns[columnName.(string)] = columnValue
			}
			collection.Rows[rowName.(string)] = fixture
		}
		f.collections[collectionName] = collection
	}

	return nil
}
