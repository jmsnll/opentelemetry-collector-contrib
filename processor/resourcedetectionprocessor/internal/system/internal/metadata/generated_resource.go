// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ResourceBuilder is a helper struct to build resources predefined in metadata.yaml.
// The ResourceBuilder is not thread-safe and must not to be used in multiple goroutines.
type ResourceBuilder struct {
	config ResourceAttributesConfig
	res    pcommon.Resource
}

// NewResourceBuilder creates a new ResourceBuilder. This method should be called on the start of the application.
func NewResourceBuilder(rac ResourceAttributesConfig) *ResourceBuilder {
	return &ResourceBuilder{
		config: rac,
		res:    pcommon.NewResource(),
	}
}

// SetHostArch sets provided value as "host.arch" attribute.
func (rb *ResourceBuilder) SetHostArch(val string) {
	if rb.config.HostArch.Enabled {
		rb.res.Attributes().PutStr("host.arch", val)
	}
}

// SetHostCPUCacheL2Size sets provided value as "host.cpu.cache.l2.size" attribute.
func (rb *ResourceBuilder) SetHostCPUCacheL2Size(val int64) {
	if rb.config.HostCPUCacheL2Size.Enabled {
		rb.res.Attributes().PutInt("host.cpu.cache.l2.size", val)
	}
}

// SetHostCPUFamily sets provided value as "host.cpu.family" attribute.
func (rb *ResourceBuilder) SetHostCPUFamily(val int64) {
	if rb.config.HostCPUFamily.Enabled {
		rb.res.Attributes().PutInt("host.cpu.family", val)
	}
}

// SetHostCPUModelID sets provided value as "host.cpu.model.id" attribute.
func (rb *ResourceBuilder) SetHostCPUModelID(val int64) {
	if rb.config.HostCPUModelID.Enabled {
		rb.res.Attributes().PutInt("host.cpu.model.id", val)
	}
}

// SetHostCPUModelName sets provided value as "host.cpu.model.name" attribute.
func (rb *ResourceBuilder) SetHostCPUModelName(val string) {
	if rb.config.HostCPUModelName.Enabled {
		rb.res.Attributes().PutStr("host.cpu.model.name", val)
	}
}

// SetHostCPUStepping sets provided value as "host.cpu.stepping" attribute.
func (rb *ResourceBuilder) SetHostCPUStepping(val int64) {
	if rb.config.HostCPUStepping.Enabled {
		rb.res.Attributes().PutInt("host.cpu.stepping", val)
	}
}

// SetHostCPUVendorID sets provided value as "host.cpu.vendor.id" attribute.
func (rb *ResourceBuilder) SetHostCPUVendorID(val string) {
	if rb.config.HostCPUVendorID.Enabled {
		rb.res.Attributes().PutStr("host.cpu.vendor.id", val)
	}
}

// SetHostID sets provided value as "host.id" attribute.
func (rb *ResourceBuilder) SetHostID(val string) {
	if rb.config.HostID.Enabled {
		rb.res.Attributes().PutStr("host.id", val)
	}
}

// SetHostName sets provided value as "host.name" attribute.
func (rb *ResourceBuilder) SetHostName(val string) {
	if rb.config.HostName.Enabled {
		rb.res.Attributes().PutStr("host.name", val)
	}
}

// SetOsDescription sets provided value as "os.description" attribute.
func (rb *ResourceBuilder) SetOsDescription(val string) {
	if rb.config.OsDescription.Enabled {
		rb.res.Attributes().PutStr("os.description", val)
	}
}

// SetOsType sets provided value as "os.type" attribute.
func (rb *ResourceBuilder) SetOsType(val string) {
	if rb.config.OsType.Enabled {
		rb.res.Attributes().PutStr("os.type", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}
