package encoding

import (
	"encoding/json"
	"fmt"
	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
	"os"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	jsonFile, err := os.ReadFile(j.FileInput)

	j.DockerCompose = &models.DockerCompose{}
	if err = json.Unmarshal(jsonFile, j.DockerCompose); err != nil {
		return fmt.Errorf("error unmarshal json: %s", err)
	}

	yamlData, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		return fmt.Errorf("error marshal yaml: %s", err)
	}

	err = os.WriteFile(j.FileOutput, yamlData, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %s", err)
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	yamlFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		return fmt.Errorf("error read file: %s", err)
	}

	y.DockerCompose = &models.DockerCompose{}
	if err = yaml.Unmarshal(yamlFile, y.DockerCompose); err != nil {
		return fmt.Errorf("error unmarshal yaml: %s", err)
	}

	jsonData, err := json.Marshal(y.DockerCompose)
	if err != nil {
		return fmt.Errorf("error marshal json: %s", err)
	}

	err = os.WriteFile(y.FileOutput, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %s", err)
	}

	return nil
}
