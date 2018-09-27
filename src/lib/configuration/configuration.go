package configuration

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/4piso/okevents/src/lib/persistence/dblayer"
)

var (
	DBTypeDefault          = dblayer.DBType("mongodb")
	DBConnectionDefault    = "mongodb://127.0.0.1"
	RestFulEndPointDefault = "localhost:8181"
)

// ServiceConfig struct definition
type ServiceConfig struct {
	DataBaseType    dblayer.DBType `json:"databasetype"`
	DBConnection    string         `json:"dbconnection"`
	RestFulEndPoint string         `json:"restfulapi_endpoint"`
}

// ExtractConfiguration (filename) => (ServiceConfig, error)
func ExtractConfiguration(filename string) (ServiceConfig, error) {
	// add local properties just in case the file is not there
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestFulEndPointDefault,
	}

	// read the file from the files system
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found, continue with the default values")
		return conf, err
	}

	err = json.NewDecoder(file).Decode(&conf)

	return conf, err
}
