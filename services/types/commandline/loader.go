package commandline

import (
	"github.com/pkg/errors"
	"github.com/yext/edward/services"
)

type CommandLineLoader struct {
}

func (l *CommandLineLoader) New() services.ConfigType {
	return &ConfigCommandLine{}
}

func (l *CommandLineLoader) Handles(c services.ConfigType) bool {
	_, ok := c.(*ConfigCommandLine)
	return ok
}

func (l *CommandLineLoader) Builder(s *services.ServiceConfig) (services.Builder, error) {
	return l.buildandrun(s)
}

func (l *CommandLineLoader) Runner(s *services.ServiceConfig) (services.Runner, error) {
	return l.buildandrun(s)
}

func (l *CommandLineLoader) buildandrun(s *services.ServiceConfig) (*buildandrun, error) {
	if config, ok := s.TypeConfig.(*ConfigCommandLine); ok {
		return &buildandrun{
			Service: s,
			Config:  config,
		}, nil
	}
	return nil, errors.New("config was not of expected type")
}
