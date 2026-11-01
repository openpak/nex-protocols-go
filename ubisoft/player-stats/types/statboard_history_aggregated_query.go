package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// StatboardHistoryAggregatedQuery is a type within the Player Stats protocol
type StatboardHistoryAggregatedQuery struct {
	types.Structure
	PlayerPIDs    types.List[types.UInt32]
	BoardID       types.UInt32
	ContextIDs    types.List[types.UInt32]
	StatIDs       types.List[types.UInt32]
	ResultRange   types.ResultRange
	PlayerStats   types.List[PropertyVariant]
	DateRanges    types.List[DateRange]
	SortCriterias types.List[PlayerStatSortCriteria]
	FilterOption  types.UInt32
}

// WriteTo writes the StatboardHistoryAggregatedQuery to the given writable
func (shaq StatboardHistoryAggregatedQuery) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	shaq.PlayerPIDs.WriteTo(contentWritable)
	shaq.BoardID.WriteTo(contentWritable)
	shaq.ContextIDs.WriteTo(contentWritable)
	shaq.StatIDs.WriteTo(contentWritable)
	shaq.ResultRange.WriteTo(contentWritable)
	shaq.PlayerStats.WriteTo(contentWritable)
	shaq.DateRanges.WriteTo(contentWritable)
	shaq.SortCriterias.WriteTo(contentWritable)
	shaq.FilterOption.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	shaq.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the StatboardHistoryAggregatedQuery from the given readable
func (shaq *StatboardHistoryAggregatedQuery) ExtractFrom(readable types.Readable) error {
	if err := shaq.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryAggregatedQuery header. %s", err.Error())
	}

	if err := shaq.PlayerPIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryAggregatedQuery.PlayerPIDs. %s", err.Error())
	}

	if err := shaq.BoardID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryAggregatedQuery.BoardID. %s", err.Error())
	}

	if err := shaq.ContextIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryAggregatedQuery.ContextIDs. %s", err.Error())
	}

	if err := shaq.StatIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryAggregatedQuery.StatIDs. %s", err.Error())
	}

	if err := shaq.ResultRange.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryAggregatedQuery.ResultRange. %s", err.Error())
	}

	if err := shaq.PlayerStats.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryAggregatedQuery.PlayerStats. %s", err.Error())
	}

	if err := shaq.DateRanges.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryAggregatedQuery.DateRanges. %s", err.Error())
	}

	if err := shaq.SortCriterias.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryAggregatedQuery.SortCriterias. %s", err.Error())
	}

	if err := shaq.FilterOption.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract StatboardHistoryAggregatedQuery.FilterOption. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of StatboardHistoryAggregatedQuery
func (shaq StatboardHistoryAggregatedQuery) Copy() types.RVType {
	copied := NewStatboardHistoryAggregatedQuery()

	copied.StructureVersion = shaq.StructureVersion
	copied.PlayerPIDs = shaq.PlayerPIDs.Copy().(types.List[types.UInt32])
	copied.BoardID = shaq.BoardID.Copy().(types.UInt32)
	copied.ContextIDs = shaq.ContextIDs.Copy().(types.List[types.UInt32])
	copied.StatIDs = shaq.StatIDs.Copy().(types.List[types.UInt32])
	copied.ResultRange = shaq.ResultRange.Copy().(types.ResultRange)
	copied.PlayerStats = shaq.PlayerStats.Copy().(types.List[PropertyVariant])
	copied.DateRanges = shaq.DateRanges.Copy().(types.List[DateRange])
	copied.SortCriterias = shaq.SortCriterias.Copy().(types.List[PlayerStatSortCriteria])
	copied.FilterOption = shaq.FilterOption.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given StatboardHistoryAggregatedQuery contains the same data as the current StatboardHistoryAggregatedQuery
func (shaq StatboardHistoryAggregatedQuery) Equals(o types.RVType) bool {
	other, ok := o.(StatboardHistoryAggregatedQuery)
	if !ok {
		return false
	}

	if shaq.StructureVersion != other.StructureVersion {
		return false
	}

	if !(shaq.PlayerPIDs.Equals(other.PlayerPIDs)) {
		return false
	}

	if !(shaq.BoardID.Equals(other.BoardID)) {
		return false
	}

	if !(shaq.ContextIDs.Equals(other.ContextIDs)) {
		return false
	}

	if !(shaq.StatIDs.Equals(other.StatIDs)) {
		return false
	}

	if !(shaq.ResultRange == other.ResultRange) {
		return false
	}

	if !(shaq.PlayerStats.Equals(other.PlayerStats)) {
		return false
	}

	if !(shaq.DateRanges.Equals(other.DateRanges)) {
		return false
	}

	if !(shaq.SortCriterias.Equals(other.SortCriterias)) {
		return false
	}

	if !(shaq.FilterOption.Equals(other.FilterOption)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the StatboardHistoryAggregatedQuery
// and returns a pointer to the new copy
func (shaq StatboardHistoryAggregatedQuery) CopyRef() types.RVTypePtr {
	copied := shaq.Copy().(StatboardHistoryAggregatedQuery)
	return &copied
}

// Deref takes a pointer to the StatboardHistoryAggregatedQuery
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (shaq *StatboardHistoryAggregatedQuery) Deref() types.RVType {
	return *shaq
}

// String returns the string representation of the StatboardHistoryAggregatedQuery
func (shaq StatboardHistoryAggregatedQuery) String() string {
	return shaq.FormatToString(0)
}

// FormatToString pretty-prints the StatboardHistoryAggregatedQuery using the provided indentation level
func (shaq StatboardHistoryAggregatedQuery) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("StatboardHistoryAggregatedQuery{\n")
	fmt.Fprintf(&b, "%sPlayerPIDs: %s,\n", indentationValues, shaq.PlayerPIDs)
	fmt.Fprintf(&b, "%sBoardID: %s,\n", indentationValues, shaq.BoardID)
	fmt.Fprintf(&b, "%sContextIDs: %s,\n", indentationValues, shaq.ContextIDs)
	fmt.Fprintf(&b, "%sStatIDs: %s,\n", indentationValues, shaq.StatIDs)
	fmt.Fprintf(&b, "%sResultRange: %s,\n", indentationValues, shaq.ResultRange.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sPlayerStats: %s,\n", indentationValues, shaq.PlayerStats)
	fmt.Fprintf(&b, "%sDateRanges: %s,\n", indentationValues, shaq.DateRanges)
	fmt.Fprintf(&b, "%sSortCriterias: %s,\n", indentationValues, shaq.SortCriterias)
	fmt.Fprintf(&b, "%sFilterOption: %s,\n", indentationValues, shaq.FilterOption)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewStatboardHistoryAggregatedQuery returns a new StatboardHistoryAggregatedQuery
func NewStatboardHistoryAggregatedQuery() StatboardHistoryAggregatedQuery {
	return StatboardHistoryAggregatedQuery{
		PlayerPIDs:    types.NewList[types.UInt32](),
		BoardID:       types.NewUInt32(0),
		ContextIDs:    types.NewList[types.UInt32](),
		StatIDs:       types.NewList[types.UInt32](),
		ResultRange:   types.NewResultRange(),
		PlayerStats:   types.NewList[PropertyVariant](),
		DateRanges:    types.NewList[DateRange](),
		SortCriterias: types.NewList[PlayerStatSortCriteria](),
		FilterOption:  types.NewUInt32(0),
	}
}
