package otel

import (
	"fmt"
)

var globalRegistry = NewRegistry()

type Factory func(c *Conf, serviceName string, uuid string) error

// Registry is the interface for callers to get registered middleware.
type Registry interface {
	Register(name string, factory Factory)
	Create(c *Conf, serviceName string, uuid string) error
}

type discoveryRegistry struct {
	discovery map[string]Factory
}

// NewRegistry returns a new middleware registry.
func NewRegistry() Registry {
	return &discoveryRegistry{
		discovery: map[string]Factory{},
	}
}

func (d *discoveryRegistry) Register(name string, factory Factory) {
	d.discovery[name] = factory
}

func (d *discoveryRegistry) Create(c *Conf, serviceName string, uuid string) error {
	key := fmt.Sprintf("%s:%s", c.ProviderType.String(), c.Type.Type())
	factory, ok := d.discovery[key]
	if !ok {
		return fmt.Errorf("discovery %s has not been registered", key)
	}

	err := factory(c, serviceName, uuid)
	if err != nil {
		return fmt.Errorf("create discovery error: %s", err)
	}
	return nil
}

// Register registers one discovery.
func Register(name string, factory Factory) {
	globalRegistry.Register(name, factory)
}

// Create instantiates a discovery based on `discoveryDSN`.
func Create(c *Conf, serviceName string, uuid string) error {
	return globalRegistry.Create(c, serviceName, uuid)
}
