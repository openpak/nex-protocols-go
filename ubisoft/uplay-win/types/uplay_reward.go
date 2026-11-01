// Package types implements all the types used by the Uplay Win protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplayReward is a type within the Uplay Win protocol
type UplayReward struct {
	types.Structure
	Code           types.String
	Name           types.String
	Description    types.String
	Value          types.Int32
	RewardTypeName types.String
	GameCode       types.String
	Platforms      types.List[UplayRewardPlatform]
}

// WriteTo writes the UplayReward to the given writable
func (s UplayReward) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.Code.WriteTo(contentWritable)
	s.Name.WriteTo(contentWritable)
	s.Description.WriteTo(contentWritable)
	s.Value.WriteTo(contentWritable)
	s.RewardTypeName.WriteTo(contentWritable)
	s.GameCode.WriteTo(contentWritable)
	s.Platforms.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	s.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplayReward from the given readable
func (s *UplayReward) ExtractFrom(readable types.Readable) error {
	if err := s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward header. %s", err.Error())
	}

	if err := s.Code.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.Code. %s", err.Error())
	}

	if err := s.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.Name. %s", err.Error())
	}

	if err := s.Description.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.Description. %s", err.Error())
	}

	if err := s.Value.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.Value. %s", err.Error())
	}

	if err := s.RewardTypeName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.RewardTypeName. %s", err.Error())
	}

	if err := s.GameCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.GameCode. %s", err.Error())
	}

	if err := s.Platforms.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.Platforms. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplayReward
func (s UplayReward) Copy() types.RVType {
	copied := NewUplayReward()

	copied.StructureVersion = s.StructureVersion
	copied.Code = s.Code.Copy().(types.String)
	copied.Name = s.Name.Copy().(types.String)
	copied.Description = s.Description.Copy().(types.String)
	copied.Value = s.Value.Copy().(types.Int32)
	copied.RewardTypeName = s.RewardTypeName.Copy().(types.String)
	copied.GameCode = s.GameCode.Copy().(types.String)
	copied.Platforms = s.Platforms.Copy().(types.List[UplayRewardPlatform])

	return copied
}

// Equals checks if the given UplayReward contains the same data as the current UplayReward
func (s UplayReward) Equals(o types.RVType) bool {
	other, ok := o.(UplayReward)
	if !ok {
		return false
	}

	if s.StructureVersion != other.StructureVersion {
		return false
	}

	if !s.Code.Equals(other.Code) {
		return false
	}

	if !s.Name.Equals(other.Name) {
		return false
	}

	if !s.Description.Equals(other.Description) {
		return false
	}

	if !s.Value.Equals(other.Value) {
		return false
	}

	if !s.RewardTypeName.Equals(other.RewardTypeName) {
		return false
	}

	if !s.GameCode.Equals(other.GameCode) {
		return false
	}

	if !s.Platforms.Equals(other.Platforms) {
		return false
	}

	return true
}

// CopyRef copies the current value of the UplayReward
// and returns a pointer to the new copy
func (s UplayReward) CopyRef() types.RVTypePtr {
	copied := s.Copy().(UplayReward)
	return &copied
}

// Deref takes a pointer to the UplayReward
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (s *UplayReward) Deref() types.RVType {
	return *s
}

// String returns the string representation of the UplayReward
func (s UplayReward) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the UplayReward using the provided indentation level
func (s UplayReward) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UplayReward{\n")
	fmt.Fprintf(&b, "%sCode: %s,\n", indentationValues, s.Code)
	fmt.Fprintf(&b, "%sName: %s,\n", indentationValues, s.Name)
	fmt.Fprintf(&b, "%sDescription: %s,\n", indentationValues, s.Description)
	fmt.Fprintf(&b, "%sValue: %s,\n", indentationValues, s.Value)
	fmt.Fprintf(&b, "%sRewardTypeName: %s,\n", indentationValues, s.RewardTypeName)
	fmt.Fprintf(&b, "%sGameCode: %s,\n", indentationValues, s.GameCode)
	fmt.Fprintf(&b, "%sPlatforms: %s,\n", indentationValues, s.Platforms)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplayReward returns a new UplayReward
func NewUplayReward() UplayReward {
	return UplayReward{
		Code:           types.NewString(""),
		Name:           types.NewString(""),
		Description:    types.NewString(""),
		Value:          types.NewInt32(0),
		RewardTypeName: types.NewString(""),
		GameCode:       types.NewString(""),
		Platforms:      types.NewList[UplayRewardPlatform](),
	}
}
