package identifier

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Identifier represents a namespaced identifier in the format namespace:value
type Identifier struct {
	Namespace string
	Value     string
}

const (
	DefaultNamespace = "minecraft"
)

var (
	namespaceRegex = regexp.MustCompile("^[a-z0-9.\\-_]+$")
	valueRegex     = regexp.MustCompile("^[a-z0-9.\\-_/]+$")

	ErrInvalidNamespace = errors.New("invalid namespace: must contain only lowercase alphanumeric characters, dots, dashes, and underscores")
	ErrInvalidValue     = errors.New("invalid value: must contain only lowercase alphanumeric characters, dots, dashes, underscores, and slashes")
	ErrEmptyValue       = errors.New("value cannot be empty")
)

// NewIdentifier creates a new identifier with validation
func NewIdentifier(namespace, value string) (*Identifier, error) {
	if namespace == "" {
		namespace = DefaultNamespace
	}

	if value == "" {
		return nil, ErrEmptyValue
	}

	if !namespaceRegex.MatchString(namespace) {
		return nil, ErrInvalidNamespace
	}

	if !valueRegex.MatchString(value) {
		return nil, ErrInvalidValue
	}

	return &Identifier{
		Namespace: namespace,
		Value:     value,
	}, nil
}

// ParseIdentifier parses a string in the format "namespace:value" or "value"
func ParseIdentifier(input string) (*Identifier, error) {
	if input == "" {
		return nil, ErrEmptyValue
	}

	parts := strings.SplitN(input, ":", 2)

	if len(parts) == 1 {
		return NewIdentifier("", parts[0])
	}

	return NewIdentifier(parts[0], parts[1])
}

// String returns the string representation of the identifier
func (id *Identifier) String() string {
	return fmt.Sprintf("%s:%s", id.Namespace, id.Value)
}

// IsDefault returns true if the identifier uses the default namespace
func (id *Identifier) IsDefault() bool {
	return id.Namespace == DefaultNamespace
}

// IsValid checks if the identifier is valid
func (id *Identifier) IsValid() bool {
	return id.Value != "" &&
		namespaceRegex.MatchString(id.Namespace) &&
		valueRegex.MatchString(id.Value)
}

// ValidateNamespace checks if a namespace string is valid
func ValidateNamespace(namespace string) bool {
	return namespace != "" && namespaceRegex.MatchString(namespace)
}

// ValidateValue checks if a value string is valid
func ValidateValue(value string) bool {
	return value != "" && valueRegex.MatchString(value)
}

// ValidateIdentifierString validates a full identifier string without parsing
func ValidateIdentifierString(input string) bool {
	if input == "" {
		return false
	}

	parts := strings.SplitN(input, ":", 2)

	if len(parts) == 1 {
		return ValidateValue(parts[0])
	}

	return ValidateNamespace(parts[0]) && ValidateValue(parts[1])
}
