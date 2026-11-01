package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// StatboardHistoryQuery is a type within the Player Stats protocol
type StatboardHistoryQuery struct {
	types.Structure
	PlayerPIDs    types.List[types.UInt32]
	BoardID       types.UInt32
	ContextIDs    types.List[types.UInt32]
	StatIDs       types.List[types.UInt32]
	ResultRange   types.ResultRange
	PlayerStats   types.List[PropertyVariant]
	DateRanges    types.List[DateRange]
	SortCriterias types.List[PlayerStatSortCriteria]
}

// WriteTo writes the StatboardHistoryQuery to the given writable
func (shq StatboardHistoryQuery) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	shq.PlayerPIDs.WriteTo(contentWritable)
	shq.BoardID.WriteTo(contentWritable)
	shq.ContextIDs.WriteTo(contentWritable)
	shq.StatIDs.WriteTo(contentWritable)
	shq.ResultRange.WriteTo(contentWritable)
	shq.PlayerStats.WriteTo(contentWritable)
	shq.DateRanges.WriteTo(contentWritable)
	shq.SortCriterias.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	shq.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the StatboardHistoryQuery from the given readable
func (shq *StatboardHistoryQuery) ExtractFrom(readable types.Readable) error {
	if err := shq.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryQuery header. %s", err.Error())
	}

	if err := shq.PlayerPIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryQuery.PlayerPIDs. %s", err.Error())
	}

	if err := shq.BoardID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryQuery.BoardID. %s", err.Error())
	}

	if err := shq.ContextIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryQuery.ContextIDs. %s", err.Error())
	}

	if err := shq.StatIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryQuery.StatIDs. %s", err.Error())
	}

	if err := shq.ResultRange.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryQuery.ResultRange. %s", err.Error())
	}

	if err := shq.PlayerStats.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryQuery.PlayerStats. %s", err.Error())
	}

	if err := shq.DateRanges.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryQuery.DateRanges. %s", err.Error())
	}

	if err := shq.SortCriterias.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryQuery.SortCriterias. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of StatboardHistoryQuery
func (shq StatboardHistoryQuery) Copy() types.RVType {
	copied := NewStatboardHistoryQuery()

	copied.StructureVersion = shq.StructureVersion
	copied.PlayerPIDs = shq.PlayerPIDs.Copy().(types.List[types.UInt32])
	copied.BoardID = shq.BoardID.Copy().(types.UInt32)
	copied.ContextIDs = shq.ContextIDs.Copy().(types.List[types.UInt32])
	copied.StatIDs = shq.StatIDs.Copy().(types.List[types.UInt32])
	copied.ResultRange = shq.ResultRange.Copy().(types.ResultRange)
	copied.PlayerStats = shq.PlayerStats.Copy().(types.List[PropertyVariant])
	copied.DateRanges = shq.DateRanges.Copy().(types.List[DateRange])
	copied.SortCriterias = shq.SortCriterias.Copy().(types.List[PlayerStatSortCriteria])

	return copied
}

// Equals checks if the given StatboardHistoryQuery contains the same data as the current StatboardHistoryQuery
func (shq StatboardHistoryQuery) Equals(o types.RVType) bool {
	other, ok := o.(StatboardHistoryQuery)
	if !ok {
		return false
	}

	if shq.StructureVersion != other.StructureVersion {
		return false
	}

	if !(shq.PlayerPIDs.Equals(other.PlayerPIDs)) {
		return false
	}

	if !(shq.BoardID.Equals(other.BoardID)) {
		return false
	}

	if !(shq.ContextIDs.Equals(other.ContextIDs)) {
		return false
	}

	if !(shq.StatIDs.Equals(other.StatIDs)) {
		return false
	}

	if !(shq.ResultRange == other.ResultRange) {
		return false
	}

	if !(shq.PlayerStats.Equals(other.PlayerStats)) {
		return false
	}

	if !(shq.DateRanges.Equals(other.DateRanges)) {
		return false
	}

	if !(shq.SortCriterias.Equals(other.SortCriterias)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the StatboardHistoryQuery
// and returns a pointer to the new copy
func (shq StatboardHistoryQuery) CopyRef() types.RVTypePtr {
	copied := shq.Copy().(StatboardHistoryQuery)
	return &copied
}

// Deref takes a pointer to the StatboardHistoryQuery
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (shq *StatboardHistoryQuery) Deref() types.RVType {
	return *shq
}

// String returns the string representation of the StatboardHistoryQuery
func (shq StatboardHistoryQuery) String() string {
	return shq.FormatToString(0)
}

// FormatToString pretty-prints the StatboardHistoryQuery using the provided indentation level
func (shq StatboardHistoryQuery) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("StatboardHistoryQuery{\n")
	fmt.Fprintf(&b, "%sPlayerPIDs: %s,\n", indentationValues, shq.PlayerPIDs)
	fmt.Fprintf(&b, "%sBoardID: %s,\n", indentationValues, shq.BoardID)
	fmt.Fprintf(&b, "%sContextIDs: %s,\n", indentationValues, shq.ContextIDs)
	fmt.Fprintf(&b, "%sStatIDs: %s,\n", indentationValues, shq.StatIDs)
	fmt.Fprintf(&b, "%sResultRange: %s,\n", indentationValues, shq.ResultRange.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sPlayerStats: %s,\n", indentationValues, shq.PlayerStats)
	fmt.Fprintf(&b, "%sDateRanges: %s,\n", indentationValues, shq.DateRanges)
	fmt.Fprintf(&b, "%sSortCriterias: %s,\n", indentationValues, shq.SortCriterias)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewStatboardHistoryQuery returns a new StatboardHistoryQuery
func NewStatboardHistoryQuery() StatboardHistoryQuery {
	return StatboardHistoryQuery{
		PlayerPIDs:    types.NewList[types.UInt32](),
		BoardID:       types.NewUInt32(0),
		ContextIDs:    types.NewList[types.UInt32](),
		StatIDs:       types.NewList[types.UInt32](),
		ResultRange:   types.NewResultRange(),
		PlayerStats:   types.NewList[PropertyVariant](),
		DateRanges:    types.NewList[DateRange](),
		SortCriterias: types.NewList[PlayerStatSortCriteria](),
	}
}
