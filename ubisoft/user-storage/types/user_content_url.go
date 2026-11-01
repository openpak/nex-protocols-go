package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UserContentURL is a type within the User Storage protocol
type UserContentURL struct {
	types.Structure
	Protocol types.String
	Host     types.String
	Path     types.String
}

// WriteTo writes the UserContentURL to the given writable
func (ucurl UserContentURL) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ucurl.Protocol.WriteTo(contentWritable)
	ucurl.Host.WriteTo(contentWritable)
	ucurl.Path.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ucurl.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UserContentURL from the given readable
func (ucurl *UserContentURL) ExtractFrom(readable types.Readable) error {
	if err := ucurl.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserContentURL header. %s", err.Error())
	}

	if err := ucurl.Protocol.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserContentURL.Protocol. %s", err.Error())
	}

	if err := ucurl.Host.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserContentURL.Host. %s", err.Error())
	}

	if err := ucurl.Path.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserContentURL.Path. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UserContentURL
func (ucurl UserContentURL) Copy() types.RVType {
	copied := NewUserContentURL()

	copied.StructureVersion = ucurl.StructureVersion
	copied.Protocol = ucurl.Protocol.Copy().(types.String)
	copied.Host = ucurl.Host.Copy().(types.String)
	copied.Path = ucurl.Path.Copy().(types.String)

	return copied
}

// Equals checks if the given UserContentURL contains the same data as the current UserContentURL
func (ucurl UserContentURL) Equals(o types.RVType) bool {
	other, ok := o.(UserContentURL)
	if !ok {
		return false
	}

	if ucurl.StructureVersion != other.StructureVersion {
		return false
	}

	if !(ucurl.Protocol.Equals(other.Protocol)) {
		return false
	}

	if !(ucurl.Host.Equals(other.Host)) {
		return false
	}

	if !(ucurl.Path.Equals(other.Path)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the UserContentURL
// and returns a pointer to the new copy
func (ucurl UserContentURL) CopyRef() types.RVTypePtr {
	copied := ucurl.Copy().(UserContentURL)
	return &copied
}

// Deref takes a pointer to the UserContentURL
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (ucurl *UserContentURL) Deref() types.RVType {
	return *ucurl
}

// String returns the string representation of the UserContentURL
func (ucurl UserContentURL) String() string {
	return ucurl.FormatToString(0)
}

// FormatToString pretty-prints the UserContentURL using the provided indentation level
func (ucurl UserContentURL) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UserContentURL{\n")
	fmt.Fprintf(&b, "%sProtocol: %s,\n", indentationValues, ucurl.Protocol)
	fmt.Fprintf(&b, "%sHost: %s,\n", indentationValues, ucurl.Host)
	fmt.Fprintf(&b, "%sPath: %s,\n", indentationValues, ucurl.Path)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUserContentURL returns a new UserContentURL
func NewUserContentURL() UserContentURL {
	return UserContentURL{
		Protocol: types.NewString(""),
		Host:     types.NewString(""),
		Path:     types.NewString(""),
	}
}
