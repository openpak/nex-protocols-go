package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PlayerStatSet is a type within the Player Stats protocol
type PlayerStatSet struct {
	types.Structure
	PlayerPID     types.UInt32
	PlayerName    types.String
	SubmittedTime types.DateTime
	Stats         types.List[PropertyVariant]
}

// WriteTo writes the PlayerStatSet to the given writable
func (pss PlayerStatSet) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	pss.PlayerPID.WriteTo(contentWritable)
	pss.PlayerName.WriteTo(contentWritable)
	pss.SubmittedTime.WriteTo(contentWritable)
	pss.Stats.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	pss.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PlayerStatSet from the given readable
func (pss *PlayerStatSet) ExtractFrom(readable types.Readable) error {
	if err := pss.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PlayerStatSet header. %s", err.Error())
	}

	if err := pss.PlayerPID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PlayerStatSet.PlayerPID. %s", err.Error())
	}

	if err := pss.PlayerName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PlayerStatSet.PlayerName. %s", err.Error())
	}

	if err := pss.SubmittedTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PlayerStatSet.SubmittedTime. %s", err.Error())
	}

	if err := pss.Stats.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PlayerStatSet.Stats. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PlayerStatSet
func (pss PlayerStatSet) Copy() types.RVType {
	copied := NewPlayerStatSet()

	copied.StructureVersion = pss.StructureVersion
	copied.PlayerPID = pss.PlayerPID.Copy().(types.UInt32)
	copied.PlayerName = pss.PlayerName.Copy().(types.String)
	copied.SubmittedTime = pss.SubmittedTime.Copy().(types.DateTime)
	copied.Stats = pss.Stats.Copy().(types.List[PropertyVariant])

	return copied
}

// Equals checks if the given PlayerStatSet contains the same data as the current PlayerStatSet
func (pss PlayerStatSet) Equals(o types.RVType) bool {
	other, ok := o.(PlayerStatSet)
	if !ok {
		return false
	}

	if pss.StructureVersion != other.StructureVersion {
		return false
	}

	if !(pss.PlayerPID.Equals(other.PlayerPID)) {
		return false
	}

	if !(pss.PlayerName.Equals(other.PlayerName)) {
		return false
	}

	if !(pss.SubmittedTime.Equals(other.SubmittedTime)) {
		return false
	}

	if !(pss.Stats.Equals(other.Stats)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the PlayerStatSet
// and returns a pointer to the new copy
func (pss PlayerStatSet) CopyRef() types.RVTypePtr {
	copied := pss.Copy().(PlayerStatSet)
	return &copied
}

// Deref takes a pointer to the PlayerStatSet
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (pss *PlayerStatSet) Deref() types.RVType {
	return *pss
}

// String returns the string representation of the PlayerStatSet
func (pss PlayerStatSet) String() string {
	return pss.FormatToString(0)
}

// FormatToString pretty-prints the PlayerStatSet using the provided indentation level
func (pss PlayerStatSet) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PlayerStatSet{\n")
	fmt.Fprintf(&b, "%sPlayerPID: %s,\n", indentationValues, pss.PlayerPID)
	fmt.Fprintf(&b, "%sPlayerName: %s,\n", indentationValues, pss.PlayerName)
	fmt.Fprintf(&b, "%sSubmittedTime: %s,\n", indentationValues, pss.SubmittedTime.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sStats: %s,\n", indentationValues, pss.Stats)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewPlayerStatSet returns a new PlayerStatSet
func NewPlayerStatSet() PlayerStatSet {
	return PlayerStatSet{
		PlayerPID:     types.NewUInt32(0),
		PlayerName:    types.NewString(""),
		SubmittedTime: types.NewDateTime(0),
		Stats:         types.NewList[PropertyVariant](),
	}
}
