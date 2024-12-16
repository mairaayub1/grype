package v5

import (
	"io"

	"github.com/mairaayub1/grype/grype/match"
)

type ProviderStore struct {
	VulnerabilityProvider
	VulnerabilityMetadataProvider
	match.ExclusionProvider
	io.Closer
}
