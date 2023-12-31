package container

import (
	"net/http"
	"net/url"
	"slices"
	"strings"

	"github.com/bilus/microwire/turbo"
)

type Container struct {
	services []Service
}

func New(services ...Service) (Container, error) {
	return Container{
		services: services,
	}, nil // TODO(bilus): Validation of service configuration.
}

func (c Container) Services() []Service {
	return slices.Clone(c.services)
}

func (c Container) ProxyTurbo(w http.ResponseWriter, r *http.Request) bool {
	if !turbo.IsTurbo(r) {
		return false
	}
	service, ok := c.resolve(r.URL)
	if !ok {
		return false
	}
	service.proxy(w, r)
	return true
}

func (c Container) resolve(url *url.URL) (Service, bool) {
	for _, service := range c.services {
		if strings.HasPrefix(url.Path, service.Path) {
			return service, true
		}
	}
	return Service{}, false //nolint: exhaustruct
}
