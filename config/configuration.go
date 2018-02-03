package config

import (
	"fmt"
	"strconv"

	fm "github.com/chentex/go-fm"
	"github.com/pkg/errors"
)

//Params parameters to initialize configuration
type Params struct {
	Force          string
	CfgFile        string
	VersionFile    string
	InitialVersion string
}

//Configure defines how to manage configuration
type Configure interface {
	InitConfig(p Params) error
}

//NewConfigure returns a new instance of a configure
func NewConfigure() Configure {
	return &Configurator{}
}

//Configurator implementation of a Configure
type Configurator struct {
}

//InitConfig initializes the configuration
func (c *Configurator) InitConfig(p Params) error {
	// Validate if repository already initialized
	fileManager := fm.NewFileManager()
	f, err := strconv.ParseBool(p.Force)
	if err != nil {
		return errors.Wrapf(err, "Error parsing flag.")
	}

	if exists, _ := fileManager.ExistsFile(p.VersionFile); exists && !f {
		return errors.Errorf("Repository already initialized with version file: %s", p.VersionFile)
	}

	yaml := fmt.Sprintf(`versionfile: %s`, p.VersionFile)
	err = fileManager.WriteFile(p.CfgFile, []byte(yaml), 0644)
	if err != nil {
		return errors.Wrapf(err, "Creating configuration file: %s", p.CfgFile)
	}

	err = fileManager.WriteFile(p.VersionFile, []byte(p.InitialVersion), 0644)
	if err != nil {
		return errors.Wrapf(err, "Creating version file: %s", p.VersionFile)
	}
	return nil
}
