package copy

import (
	"mcolomerc/cc-tools/pkg/client"
	"mcolomerc/cc-tools/pkg/config"
)

type SchemaCopy struct {
	SourceClient client.RestClient
	DestClient   client.RestClient
	Conf         config.Config
}

func NewSchemaCopy(conf config.Config) (*SchemaCopy, error) {
	return nil, nil
}

func (s *SchemaCopy) Copy() {
	// Copy Schemas
	// s.copySchemas()
	// Copy Subjects
	// s.copySubjects()
}
