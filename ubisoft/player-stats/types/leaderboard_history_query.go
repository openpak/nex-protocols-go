package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// LeaderboardHistoryQuery is a type within the Player Stats protocol
type LeaderboardHistoryQuery struct {
	types.Structure
	PlayerPIDs   types.List[types.UInt32]
	BoardID      types.UInt32
	ContextID    types.UInt32
	ResultRange  types.ResultRange
	DateRanges   types.List[DateRange]
	FilterOption types.UInt32
}

// WriteTo writes the LeaderboardHistoryQuery to the given writable
func (lhq LeaderboardHistoryQuery) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	lhq.PlayerPIDs.WriteTo(contentWritable)
	lhq.BoardID.WriteTo(contentWritable)
	lhq.ContextID.WriteTo(contentWritable)
	lhq.ResultRange.WriteTo(contentWritable)
	lhq.DateRanges.WriteTo(contentWritable)
	lhq.FilterOption.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	lhq.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the LeaderboardHistoryQuery from the given readable
func (lhq *LeaderboardHistoryQuery) ExtractFrom(readable types.Readable) error {
	if err := lhq.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardHistoryQuery header. %s", err.Error())
	}

	if err := lhq.PlayerPIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardHistoryQuery.PlayerPIDs. %s", err.Error())
	}

	if err := lhq.BoardID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardHistoryQuery.BoardID. %s", err.Error())
	}

	if err := lhq.ContextID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardHistoryQuery.ContextID. %s", err.Error())
	}

	if err := lhq.ResultRange.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardHistoryQuery.ResultRange. %s", err.Error())
	}

	if err := lhq.DateRanges.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardHistoryQuery.DateRanges. %s", err.Error())
	}

	if err := lhq.FilterOption.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardHistoryQuery.FilterOption. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of LeaderboardHistoryQuery
func (lhq LeaderboardHistoryQuery) Copy() types.RVType {
	copied := NewLeaderboardHistoryQuery()

	copied.StructureVersion = lhq.StructureVersion
	copied.PlayerPIDs = lhq.PlayerPIDs.Copy().(types.List[types.UInt32])
	copied.BoardID = lhq.BoardID.Copy().(types.UInt32)
	copied.ContextID = lhq.ContextID.Copy().(types.UInt32)
	copied.ResultRange = lhq.ResultRange.Copy().(types.ResultRange)
	copied.DateRanges = lhq.DateRanges.Copy().(types.List[DateRange])
	copied.FilterOption = lhq.FilterOption.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given LeaderboardHistoryQuery contains the same data as the current LeaderboardHistoryQuery
func (lhq LeaderboardHistoryQuery) Equals(o types.RVType) bool {
	other, ok := o.(LeaderboardHistoryQuery)
	if !ok {
		return false
	}

	if lhq.StructureVersion != other.StructureVersion {
		return false
	}

	if !(lhq.PlayerPIDs.Equals(other.PlayerPIDs)) {
		return false
	}

	if !(lhq.BoardID.Equals(other.BoardID)) {
		return false
	}

	if !(lhq.ContextID.Equals(other.ContextID)) {
		return false
	}

	if !(lhq.ResultRange == other.ResultRange) {
		return false
	}

	if !(lhq.DateRanges.Equals(other.DateRanges)) {
		return false
	}

	if !(lhq.FilterOption.Equals(other.FilterOption)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the LeaderboardHistoryQuery
// and returns a pointer to the new copy
func (lhq LeaderboardHistoryQuery) CopyRef() types.RVTypePtr {
	copied := lhq.Copy().(LeaderboardHistoryQuery)
	return &copied
}

// Deref takes a pointer to the LeaderboardHistoryQuery
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (lhq *LeaderboardHistoryQuery) Deref() types.RVType {
	return *lhq
}

// String returns the string representation of the LeaderboardHistoryQuery
func (lhq LeaderboardHistoryQuery) String() string {
	return lhq.FormatToString(0)
}

// FormatToString pretty-prints the LeaderboardHistoryQuery using the provided indentation level
func (lhq LeaderboardHistoryQuery) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("LeaderboardHistoryQuery{\n")
	fmt.Fprintf(&b, "%sPlayerPIDs: %s,\n", indentationValues, lhq.PlayerPIDs)
	fmt.Fprintf(&b, "%sBoardID: %s,\n", indentationValues, lhq.BoardID)
	fmt.Fprintf(&b, "%sContextID: %s,\n", indentationValues, lhq.ContextID)
	fmt.Fprintf(&b, "%sResultRange: %s,\n", indentationValues, lhq.ResultRange.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sDateRanges: %s,\n", indentationValues, lhq.DateRanges)
	fmt.Fprintf(&b, "%sFilterOption: %s,\n", indentationValues, lhq.FilterOption)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewLeaderboardHistoryQuery returns a new LeaderboardHistoryQuery
func NewLeaderboardHistoryQuery() LeaderboardHistoryQuery {
	return LeaderboardHistoryQuery{
		PlayerPIDs:   types.NewList[types.UInt32](),
		BoardID:      types.NewUInt32(0),
		ContextID:    types.NewUInt32(0),
		ResultRange:  types.NewResultRange(),
		DateRanges:   types.NewList[DateRange](),
		FilterOption: types.NewUInt32(0),
	}
}
