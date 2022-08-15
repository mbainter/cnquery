package azure

import (
	"errors"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/spf13/afero"
	"go.mondoo.io/mondoo/motor/providers"
	"go.mondoo.io/mondoo/motor/providers/fsutil"
)

var (
	_ providers.Transport                   = (*Provider)(nil)
	_ providers.TransportPlatformIdentifier = (*Provider)(nil)
)

func New(cfg *providers.Config) (*Provider, error) {
	if cfg.Backend != providers.ProviderType_AZURE {
		return nil, providers.ErrProviderTypeDoesNotMatch
	}

	if cfg.Options == nil || len(cfg.Options["subscriptionID"]) == 0 {
		return nil, errors.New("azure provider requires a subscriptionID")
	}

	if cfg.Options == nil || len(cfg.Options["tenantID"]) == 0 {
		return nil, errors.New("azure provider requires a tenantID")
	}

	return &Provider{
		subscriptionID: cfg.Options["subscriptionID"],
		tenantID:       cfg.Options["tenantID"],
		opts:           cfg.Options,
	}, nil
}

type Provider struct {
	subscriptionID string
	tenantID       string
	opts           map[string]string
}

func (p *Provider) RunCommand(command string) (*providers.Command, error) {
	return nil, providers.ErrRunCommandNotImplemented
}

func (p *Provider) FileInfo(path string) (providers.FileInfoDetails, error) {
	return providers.FileInfoDetails{}, providers.ErrFileInfoNotImplemented
}

func (p *Provider) FS() afero.Fs {
	return &fsutil.NoFs{}
}

func (p *Provider) Close() {}

func (p *Provider) Capabilities() providers.Capabilities {
	return providers.Capabilities{
		providers.Capability_Azure,
	}
}

func (p *Provider) Options() map[string]string {
	return p.opts
}

func (p *Provider) Kind() providers.Kind {
	return providers.Kind_KIND_API
}

func (p *Provider) Runtime() string {
	return providers.RUNTIME_AZ
}

func (p *Provider) PlatformIdDetectors() []providers.PlatformIdDetector {
	return []providers.PlatformIdDetector{
		providers.TransportPlatformIdentifierDetector,
	}
}

func GetAuthorizer() (autorest.Authorizer, error) {
	return auth.NewAuthorizerFromCLI()
}

func (p *Provider) Authorizer() (autorest.Authorizer, error) {
	return GetAuthorizer()
}

func (p *Provider) AuthorizerWithAudience(audience string) (autorest.Authorizer, error) {
	return auth.NewAuthorizerFromCLIWithResource(audience)
}

func (p *Provider) ParseResourceID(id string) (*ResourceID, error) {
	return ParseResourceID(id)
}
