package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// LeaderboardQuery2 is a type within the Player Stats protocol
type LeaderboardQuery2 struct {
	types.Structure
	BoardID        types.UInt32
	ContextID      types.UInt32
	ResetFrequency types.UInt32
	StatIDs        types.List[types.UInt32]
	EstimatedPIDs  types.List[types.UInt32]
}

// WriteTo writes the LeaderboardQuery2 to the given writable
func (lq LeaderboardQuery2) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	lq.BoardID.WriteTo(contentWritable)
	lq.ContextID.WriteTo(contentWritable)
	lq.ResetFrequency.WriteTo(contentWritable)
	lq.StatIDs.WriteTo(contentWritable)
	lq.EstimatedPIDs.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	lq.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the LeaderboardQuery2 from the given readable
func (lq *LeaderboardQuery2) ExtractFrom(readable types.Readable) error {
	if err := lq.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardQuery2 header. %s", err.Error())
	}

	if err := lq.BoardID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardQuery2.BoardID. %s", err.Error())
	}

	if err := lq.ContextID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardQuery2.ContextID. %s", err.Error())
	}

	if err := lq.ResetFrequency.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardQuery2.ResetFrequency. %s", err.Error())
	}

	if err := lq.StatIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardQuery2.StatIDs. %s", err.Error())
	}

	if err := lq.EstimatedPIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardQuery2.EstimatedPIDs. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of LeaderboardQuery2
func (lq LeaderboardQuery2) Copy() types.RVType {
	copied := NewLeaderboardQuery2()

	copied.StructureVersion = lq.StructureVersion
	copied.BoardID = lq.BoardID.Copy().(types.UInt32)
	copied.ContextID = lq.ContextID.Copy().(types.UInt32)
	copied.ResetFrequency = lq.ResetFrequency.Copy().(types.UInt32)
	copied.StatIDs = lq.StatIDs.Copy().(types.List[types.UInt32])
	copied.EstimatedPIDs = lq.EstimatedPIDs.Copy().(types.List[types.UInt32])

	return copied
}

// Equals checks if the given LeaderboardQuery2 contains the same data as the current LeaderboardQuery2
func (lq LeaderboardQuery2) Equals(o types.RVType) bool {
	other, ok := o.(LeaderboardQuery2)
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

	if !(lq.EstimatedPIDs.Equals(other.EstimatedPIDs)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the LeaderboardQuery2
// and returns a pointer to the new copy
func (lq LeaderboardQuery2) CopyRef() types.RVTypePtr {
	copied := lq.Copy().(LeaderboardQuery2)
	return &copied
}

// Deref takes a pointer to the LeaderboardQuery2
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (lq *LeaderboardQuery2) Deref() types.RVType {
	return *lq
}

// String returns the string representation of the LeaderboardQuery2
func (lq LeaderboardQuery2) String() string {
	return lq.FormatToString(0)
}

// FormatToString pretty-prints the LeaderboardQuery2 using the provided indentation level
func (lq LeaderboardQuery2) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("LeaderboardQuery2{\n")
	fmt.Fprintf(&b, "%sBoardID: %s,\n", indentationValues, lq.BoardID)
	fmt.Fprintf(&b, "%sContextID: %s,\n", indentationValues, lq.ContextID)
	fmt.Fprintf(&b, "%sResetFrequency: %s,\n", indentationValues, lq.ResetFrequency)
	fmt.Fprintf(&b, "%sStatIDs: %s,\n", indentationValues, lq.StatIDs)
	fmt.Fprintf(&b, "%sEstimatedPIDs: %s,\n", indentationValues, lq.EstimatedPIDs)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewLeaderboardQuery2 returns a new LeaderboardQuery2
func NewLeaderboardQuery2() LeaderboardQuery2 {
	return LeaderboardQuery2{
		BoardID:        types.NewUInt32(0),
		ContextID:      types.NewUInt32(0),
		ResetFrequency: types.NewUInt32(0),
		StatIDs:        types.NewList[types.UInt32](),
		EstimatedPIDs:  types.NewList[types.UInt32](),
	}
}
