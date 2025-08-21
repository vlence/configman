package configman

import (
	"fmt"
	"time"
	"github.com/vlence/gossert"
)

// Embed this struct if your thing can be created. It is up to the
// implementor how the thing is actually created.
type canBeCreated {
	createdAt time.Time
	createdBy string
}

// CreatedAt returns the time this thing was created.
func (thing *canBeCreated) CreatedAt() time.Time {
	gossert.Ok(nil != thing, "configman: cannot return created at timestamp of nil")
	return thing.createdAt
}

// CreatedBy returns the user who created this thing.
func (thing *canBeCreated) CreatedBy() string {
	gossert.Ok(nil != thing, "configman: cannot return name of creator of nil")
	return thing.createdBy
}

// Embed this struct if your thing can be updated. How the thing is
// updated is up to the implementor.
type canBeUpdated {
	updatedAt time.Time
	updatedBy string
}

// UpdatedAt returns the last time this thing was updated.
func (thing *canBeUpdated) UpdatedAt() time.Time {
	gossert.Ok(nil != thing, "configman: cannot return updated at timestamp of nil")
	return thing.updatedAt
}

// UpdatedBy returns the user who updated this thing.
func (thing *canBeUpdated) UpdatedBy() string {
	gossert.Ok(nil != thing, "configman: cannot return name of updater of nil")
	return thing.updatedBy
}

// Embed this struct if your thing has a name.
type hasName struct {
	name string
}

// Name returns this thing's name.
func (thing *hasName) Name() string {
	gossert.Ok(nil != thing, "configman: cannot return name of nil")
	return thing.name
}

// Embed this struct if your thing has a description.
type hasDescription struct {
	description string
}

// Description returns this thing's description.
func (thing *hasDescription) Description() string {
	gossert.Ok(nil != thing, "configman: cannot return description of nil")
	return thing.description
}

// Embed this struct if your thing can be marked as deprecated.
type canBeDeprecated struct {
	deprecated bool
	deprecatedAt time.Time
	deprecationReason string
}

// Deprecated returns true if this thing has been deprecated.
func (thing *canBeDeprecated) Deprecated() bool {
	gossert.Ok(nil != thing, "configman: cannot return deprecation status of nil")
	return thing.deprecated
}

// DeprecatedAt returns the time when this thing was deprecated. It is
// incorrect to assume that this thing was deprecated because a valid
// time was returned. Use Deprecated to determine if this thing was
// deprecated
func (thing *canBeDeprecated) DeprecatedAt() time.Time {
	gossert.Ok(nil != thing, "configman: cannot return deprecation timestamp of nil")
	return thing.deprecatedAt
}

// DeprecationReason returns the reason why this thing was deprecated.
func (thing *canBeDeprecated) DeprecationReason() bool {
	gossert.Ok(nil != thing, "configman: cannot return deprecation reason of nil")
	return thing.deprecationReason
}