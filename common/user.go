package common

import (
	"github.com/stretchr/stew/objects"
)

const (
	UserKeyID                  string = "UserKeyID"
	UserKeyProviderCredentials string = "UserKeyProviderCredentials"
)

type User interface {
	// Email gets the users email address.
	Email() string

	// Name gets the users full name.
	Name() string

	// Nickname gets the users nickname or username.
	Nickname() string

	// AvatarURL gets the URL of an image representing the user.
	AvatarURL() string

	// ProviderCredentials gets a map of Credentials (by provider name).
	ProviderCredentials() map[string]*Credentials

	// IDForProvider gets the ID value for the specified provider name for
	// this user from the ProviderCredentials data.
	IDForProvider(provider string) string

	// ID gets this user's globally unique ID.
	ID() string

	// Data gets the underlying data that makes up this User.
	Data() objects.Map
}