// Package types implements the structures used by the Player Stats protocol

package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PlayerStatUpdate is a type within the Player Stats protocol
type PlayerStatUpdate struct {
	types.Structure
	BoardID    types.UInt32
	ContextIDs types.List[types.UInt32]
	Stats      types.List[PropertyVariant]
}

// WriteTo writes the PlayerStatUpdate to the given writable
func (psu PlayerStatUpdate) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	psu.BoardID.WriteTo(contentWritable)
	psu.ContextIDs.WriteTo(contentWritable)
	psu.Stats.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	psu.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PlayerStatUpdate from the given readable
func (psu *PlayerStatUpdate) ExtractFrom(readable types.Readable) error {
	if err := psu.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PlayerStatUpdate header. %s", err.Error())
	}

	if err := psu.BoardID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PlayerStatUpdate.BoardID. %s", err.Error())
	}

	if err := psu.ContextIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PlayerStatUpdate.ContextIDs. %s", err.Error())
	}

	if err := psu.Stats.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PlayerStatUpdate.Stats. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PlayerStatUpdate
func (psu PlayerStatUpdate) Copy() types.RVType {
	copied := NewPlayerStatUpdate()

	copied.StructureVersion = psu.StructureVersion
	copied.BoardID = psu.BoardID.Copy().(types.UInt32)
	copied.ContextIDs = psu.ContextIDs.Copy().(types.List[types.UInt32])
	copied.Stats = psu.Stats.Copy().(types.List[PropertyVariant])

	return copied
}

// Equals checks if the given PlayerStatUpdate contains the same data as the current PlayerStatUpdate
func (psu PlayerStatUpdate) Equals(o types.RVType) bool {
	other, ok := o.(PlayerStatUpdate)
	if !ok {
		return false
	}

	if psu.StructureVersion != other.StructureVersion {
		return false
	}

	if !(psu.BoardID.Equals(other.BoardID)) {
		return false
	}

	if !(psu.ContextIDs.Equals(other.ContextIDs)) {
		return false
	}

	if !(psu.Stats.Equals(other.Stats)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the PlayerStatUpdate
// and returns a pointer to the new copy
func (psu PlayerStatUpdate) CopyRef() types.RVTypePtr {
	copied := psu.Copy().(PlayerStatUpdate)
	return &copied
}

// Deref takes a pointer to the PlayerStatUpdate
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (psu *PlayerStatUpdate) Deref() types.RVType {
	return *psu
}

// String returns the string representation of the PlayerStatUpdate
func (psu PlayerStatUpdate) String() string {
	return psu.FormatToString(0)
}

// FormatToString pretty-prints the PlayerStatUpdate using the provided indentation level
func (psu PlayerStatUpdate) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PlayerStatUpdate{\n")
	fmt.Fprintf(&b, "%sBoardID: %s,\n", indentationValues, psu.BoardID)
	fmt.Fprintf(&b, "%sContextIDs: %s,\n", indentationValues, psu.ContextIDs)
	fmt.Fprintf(&b, "%sStats: %s,\n", indentationValues, psu.Stats)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewPlayerStatUpdate returns a new PlayerStatUpdate
func NewPlayerStatUpdate() PlayerStatUpdate {
	return PlayerStatUpdate{
		BoardID:    types.NewUInt32(0),
		ContextIDs: types.NewList[types.UInt32](),
		Stats:      types.NewList[PropertyVariant](),
	}
}
