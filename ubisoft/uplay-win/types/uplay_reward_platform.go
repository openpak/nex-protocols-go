// Package types implements all the types used by the Uplay Win protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplayRewardPlatform is a type within the Uplay Win protocol
type UplayRewardPlatform struct {
	types.Structure
	PlatformCode types.String
	Purchased    types.Bool
}

// WriteTo writes the UplayRewardPlatform to the given writable
func (s UplayRewardPlatform) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.PlatformCode.WriteTo(contentWritable)
	s.Purchased.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplayRewardPlatform from the given readable
func (s *UplayRewardPlatform) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayRewardPlatform header. %s", err.Error())
	}

	if err := s.PlatformCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayRewardPlatform.PlatformCode. %s", err.Error())
	}

	if err := s.Purchased.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayRewardPlatform.Purchased. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplayRewardPlatform
func (s UplayRewardPlatform) Copy() types.RVType {
	copied := NewUplayRewardPlatform()

	copied.StructureVersion = s.StructureVersion
	copied.PlatformCode = s.PlatformCode.Copy().(types.String)
	copied.Purchased = s.Purchased.Copy().(types.Bool)

	return copied
}

// Equals checks if the given UplayRewardPlatform contains the same data as the current UplayRewardPlatform
func (s UplayRewardPlatform) Equals(o types.RVType) bool {
	other, ok := o.(UplayRewardPlatform)
	if !ok {
		return false
	}

	if s.StructureVersion != other.StructureVersion {
		return false
	}

	if !s.PlatformCode.Equals(other.PlatformCode) {
		return false
	}

	if !s.Purchased.Equals(other.Purchased) {
		return false
	}

	return true
}

// CopyRef copies the current value of the UplayRewardPlatform
// and returns a pointer to the new copy
func (s UplayRewardPlatform) CopyRef() types.RVTypePtr {
	copied := s.Copy().(UplayRewardPlatform)
	return &copied
}

// Deref takes a pointer to the UplayRewardPlatform
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *UplayRewardPlatform) Deref() types.RVType {
	return *s
}

// String returns the string representation of the UplayRewardPlatform
func (s UplayRewardPlatform) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the UplayRewardPlatform using the provided indentation level
func (s UplayRewardPlatform) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UplayRewardPlatform{\n")
	fmt.Fprintf(&b, "%sPlatformCode: %s,\n", indentationValues, s.PlatformCode)
	fmt.Fprintf(&b, "%sPurchased: %s,\n", indentationValues, s.Purchased)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplayRewardPlatform returns a new UplayRewardPlatform
func NewUplayRewardPlatform() UplayRewardPlatform {
	return UplayRewardPlatform{
		PlatformCode: types.NewString(""),
		Purchased:    types.NewBool(false),
	}
}
