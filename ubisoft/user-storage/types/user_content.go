package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UserContent is a type within the User Storage protocol
type UserContent struct {
	types.Structure
	Key        UserContentKey
	Pid        types.UInt32
	Properties types.List[ContentProperty]
}

// WriteTo writes the UserContent to the given writable
func (uc UserContent) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	uc.Key.WriteTo(contentWritable)
	uc.Pid.WriteTo(contentWritable)
	uc.Properties.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	uc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UserContent from the given readable
func (uc *UserContent) ExtractFrom(readable types.Readable) error {
	if err := uc.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserContent header. %s", err.Error())
	}

	if err := uc.Key.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserContent.Key. %s", err.Error())
	}

	if err := uc.Pid.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserContent.Pid. %s", err.Error())
	}

	if err := uc.Properties.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UserContent.Properties. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UserContent
func (uc UserContent) Copy() types.RVType {
	copied := NewUserContent()

	copied.StructureVersion = uc.StructureVersion
	copied.Key = uc.Key.Copy().(UserContentKey)
	copied.Pid = uc.Pid.Copy().(types.UInt32)
	copied.Properties = uc.Properties.Copy().(types.List[ContentProperty])

	return copied
}

// Equals checks if the given UserContent contains the same data as the current UserContent
func (uc UserContent) Equals(o types.RVType) bool {
	other, ok := o.(UserContent)
	if !ok {
		return false
	}

	if uc.StructureVersion != other.StructureVersion {
		return false
	}

	if !(uc.Key.Equals(other.Key)) {
		return false
	}

	if !(uc.Pid.Equals(other.Pid)) {
		return false
	}

	if !(uc.Properties.Equals(other.Properties)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the UserContent
// and returns a pointer to the new copy
func (uc UserContent) CopyRef() types.RVTypePtr {
	copied := uc.Copy().(UserContent)
	return &copied
}

// Deref takes a pointer to the UserContent
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (uc *UserContent) Deref() types.RVType {
	return *uc
}

// String returns the string representation of the UserContent
func (uc UserContent) String() string {
	return uc.FormatToString(0)
}

// FormatToString pretty-prints the UserContent using the provided indentation level
func (uc UserContent) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UserContent{\n")
	fmt.Fprintf(&b, "%sKey: %s,\n", indentationValues, uc.Key.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sPid: %s,\n", indentationValues, uc.Pid)
	fmt.Fprintf(&b, "%sProperties: %s,\n", indentationValues, uc.Properties)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUserContent returns a new UserContent
func NewUserContent() UserContent {
	return UserContent{
		Key:        NewUserContentKey(),
		Pid:        types.NewUInt32(0),
		Properties: types.NewList[ContentProperty](),
	}
}
