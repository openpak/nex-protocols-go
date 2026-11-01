package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// LeaderboardQuery is a type within the Player Stats protocol
type LeaderboardQuery struct {
	types.Structure
	BoardID        types.UInt32
	ContextID      types.UInt32
	ResetFrequency types.UInt32
	StatIDs        types.List[types.UInt32]
}

// WriteTo writes the LeaderboardQuery to the given writable
func (lq LeaderboardQuery) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	lq.BoardID.WriteTo(contentWritable)
	lq.ContextID.WriteTo(contentWritable)
	lq.ResetFrequency.WriteTo(contentWritable)
	lq.StatIDs.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	lq.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the LeaderboardQuery from the given readable
func (lq *LeaderboardQuery) ExtractFrom(readable types.Readable) error {
	if err := lq.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardQuery header. %s", err.Error())
	}

	if err := lq.BoardID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardQuery.BoardID. %s", err.Error())
	}

	if err := lq.ContextID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardQuery.ContextID. %s", err.Error())
	}

	if err := lq.ResetFrequency.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardQuery.ResetFrequency. %s", err.Error())
	}

	if err := lq.StatIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardQuery.StatIDs. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of LeaderboardQuery
func (lq LeaderboardQuery) Copy() types.RVType {
	copied := NewLeaderboardQuery()

	copied.StructureVersion = lq.StructureVersion
	copied.BoardID = lq.BoardID.Copy().(types.UInt32)
	copied.ContextID = lq.ContextID.Copy().(types.UInt32)
	copied.ResetFrequency = lq.ResetFrequency.Copy().(types.UInt32)
	copied.StatIDs = lq.StatIDs.Copy().(types.List[types.UInt32])

	return copied
}

// Equals checks if the given LeaderboardQuery contains the same data as the current LeaderboardQuery
func (lq LeaderboardQuery) Equals(o types.RVType) bool {
	other, ok := o.(LeaderboardQuery)
	if !ok {
		return false
	}

	if lq.StructureVersion != other.StructureVersion {
		return false
	}

	if !(lq.BoardID.Equals(other.BoardID)) {
		return false
	}

	if !(lq.ContextID.Equals(other.ContextID)) {
		return false
	}

	if !(lq.ResetFrequency.Equals(other.ResetFrequency)) {
		return false
	}

	if !(lq.StatIDs.Equals(other.StatIDs)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the LeaderboardQuery
// and returns a pointer to the new copy
func (lq LeaderboardQuery) CopyRef() types.RVTypePtr {
	copied := lq.Copy().(LeaderboardQuery)
	return &copied
}

// Deref takes a pointer to the LeaderboardQuery
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (lq *LeaderboardQuery) Deref() types.RVType {
	return *lq
}

// String returns the string representation of the LeaderboardQuery
func (lq LeaderboardQuery) String() string {
	return lq.FormatToString(0)
}

// FormatToString pretty-prints the LeaderboardQuery using the provided indentation level
func (lq LeaderboardQuery) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("LeaderboardQuery{\n")
	fmt.Fprintf(&b, "%sBoardID: %s,\n", indentationValues, lq.BoardID)
	fmt.Fprintf(&b, "%sContextID: %s,\n", indentationValues, lq.ContextID)
	fmt.Fprintf(&b, "%sResetFrequency: %s,\n", indentationValues, lq.ResetFrequency)
	fmt.Fprintf(&b, "%sStatIDs: %s,\n", indentationValues, lq.StatIDs)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewLeaderboardQuery returns a new LeaderboardQuery
func NewLeaderboardQuery() LeaderboardQuery {
	return LeaderboardQuery{
		BoardID:        types.NewUInt32(0),
		ContextID:      types.NewUInt32(0),
		ResetFrequency: types.NewUInt32(0),
		StatIDs:        types.NewList[types.UInt32](),
	}
}
