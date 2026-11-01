package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UserContentKey is a type within the User Storage protocol
type UserContentKey struct {
	types.Structure
	TypeID    types.UInt32
	ContentID types.UInt64
}

// WriteTo writes the UserContentKey to the given writable
func (uck UserContentKey) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	uck.TypeID.WriteTo(contentWritable)
	uck.ContentID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	uck.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UserContentKey from the given readable
func (uck *UserContentKey) ExtractFrom(readable types.Readable) error {
	if err := uck.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserContentKey header. %s", err.Error())
	}

	if err := uck.TypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserContentKey.TypeID. %s", err.Error())
	}

	if err := uck.ContentID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserContentKey.ContentID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UserContentKey
func (uck UserContentKey) Copy() types.RVType {
	copied := NewUserContentKey()

	copied.StructureVersion = uck.StructureVersion
	copied.TypeID = uck.TypeID.Copy().(types.UInt32)
	copied.ContentID = uck.ContentID.Copy().(types.UInt64)

	return copied
}

// Equals checks if the given UserContentKey contains the same data as the current UserContentKey
func (uck UserContentKey) Equals(o types.RVType) bool {
	other, ok := o.(UserContentKey)
	if !ok {
		return false
	}

	if uck.StructureVersion != other.StructureVersion {
		return false
	}

	if !(uck.TypeID.Equals(other.TypeID)) {
		return false
	}

	if !(uck.ContentID.Equals(other.ContentID)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the UserContentKey
// and returns a pointer to the new copy
func (uck UserContentKey) CopyRef() types.RVTypePtr {
	copied := uck.Copy().(UserContentKey)
	return &copied
}

// Deref takes a pointer to the UserContentKey
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (uck *UserContentKey) Deref() types.RVType {
	return *uck
}

// String returns the string representation of the UserContentKey
func (uck UserContentKey) String() string {
	return uck.FormatToString(0)
}

// FormatToString pretty-prints the UserContentKey using the provided indentation level
func (uck UserContentKey) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UserContentKey{\n")
	fmt.Fprintf(&b, "%sTypeID: %s,\n", indentationValues, uck.TypeID)
	fmt.Fprintf(&b, "%sContentID: %s,\n", indentationValues, uck.ContentID)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUserContentKey returns a new UserContentKey
func NewUserContentKey() UserContentKey {
	return UserContentKey{
		TypeID:    types.NewUInt32(0),
		ContentID: types.NewUInt64(0),
	}
}
