package factory

import (
	"net/http"

	"github.com/maxwelbm/mdz/pkg/environment"
	"github.com/maxwelbm/mdz/pkg/iostreams"
)

type Factory struct {
	Token      string
	HTTPClient *http.Client
	IOStreams  *iostreams.IOStreams
	Env        *environment.Env
	Flags
}

type Flags struct {
	NoColor bool
}

func NewFactory(env *environment.Env) *Factory {
	return &Factory{
		HTTPClient: &http.Client{},
		IOStreams:  iostreams.System(),
		Env:        env,
	}
}
