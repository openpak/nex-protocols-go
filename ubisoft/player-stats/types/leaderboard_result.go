package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// LeaderboardResult is a type within the Player Stats protocol
type LeaderboardResult struct {
	types.Structure
	BoardID                     types.UInt32
	ContextID                   types.UInt32
	ResetFrequency              types.UInt32
	LeaderboardTotalPlayerCount types.UInt32
	PlayerRanks                 types.List[PlayerRank]
}

// WriteTo writes the LeaderboardResult to the given writable
func (lr LeaderboardResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	lr.BoardID.WriteTo(contentWritable)
	lr.ContextID.WriteTo(contentWritable)
	lr.ResetFrequency.WriteTo(contentWritable)
	lr.LeaderboardTotalPlayerCount.WriteTo(contentWritable)
	lr.PlayerRanks.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	lr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the LeaderboardResult from the given readable
func (lr *LeaderboardResult) ExtractFrom(readable types.Readable) error {
	if err := lr.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardResult header. %s", err.Error())
	}

	if err := lr.BoardID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardResult.BoardID. %s", err.Error())
	}

	if err := lr.ContextID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardResult.ContextID. %s", err.Error())
	}

	if err := lr.ResetFrequency.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardResult.ResetFrequency. %s", err.Error())
	}

	if err := lr.LeaderboardTotalPlayerCount.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardResult.LeaderboardTotalPlayerCount. %s", err.Error())
	}

	if err := lr.PlayerRanks.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract LeaderboardResult.PlayerRanks. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of LeaderboardResult
func (lr LeaderboardResult) Copy() types.RVType {
	copied := NewLeaderboardResult()

	copied.StructureVersion = lr.StructureVersion
	copied.BoardID = lr.BoardID.Copy().(types.UInt32)
	copied.ContextID = lr.ContextID.Copy().(types.UInt32)
	copied.ResetFrequency = lr.ResetFrequency.Copy().(types.UInt32)
	copied.LeaderboardTotalPlayerCount = lr.LeaderboardTotalPlayerCount.Copy().(types.UInt32)
	copied.PlayerRanks = lr.PlayerRanks.Copy().(types.List[PlayerRank])

	return copied
}

// Equals checks if the given LeaderboardResult contains the same data as the current LeaderboardResult
func (lr LeaderboardResult) Equals(o types.RVType) bool {
	other, ok := o.(LeaderboardResult)
	if !ok {
		return false
	}

	if lr.StructureVersion != other.StructureVersion {
		return false
	}

	if !(lr.BoardID.Equals(other.BoardID)) {
		return false
	}

	if !(lr.ContextID.Equals(other.ContextID)) {
		return false
	}

	if !(lr.ResetFrequency.Equals(other.ResetFrequency)) {
		return false
	}

	if !(lr.LeaderboardTotalPlayerCount.Equals(other.LeaderboardTotalPlayerCount)) {
		return false
	}

	if !(lr.PlayerRanks.Equals(other.PlayerRanks)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the LeaderboardResult
// and returns a pointer to the new copy
func (lr LeaderboardResult) CopyRef() types.RVTypePtr {
	copied := lr.Copy().(LeaderboardResult)
	return &copied
}

// Deref takes a pointer to the LeaderboardResult
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (lr *LeaderboardResult) Deref() types.RVType {
	return *lr
}

// String returns the string representation of the LeaderboardResult
func (lr LeaderboardResult) String() string {
	return lr.FormatToString(0)
}

// FormatToString pretty-prints the LeaderboardResult using the provided indentation level
func (lr LeaderboardResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("LeaderboardResult{\n")
	fmt.Fprintf(&b, "%sBoardID: %s,\n", indentationValues, lr.BoardID)
	fmt.Fprintf(&b, "%sContextID: %s,\n", indentationValues, lr.ContextID)
	fmt.Fprintf(&b, "%sResetFrequency: %s,\n", indentationValues, lr.ResetFrequency)
	fmt.Fprintf(&b, "%sLeaderboardTotalPlayerCount: %s,\n", indentationValues, lr.LeaderboardTotalPlayerCount)
	fmt.Fprintf(&b, "%sPlayerRanks: %s,\n", indentationValues, lr.PlayerRanks)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewLeaderboardResult returns a new LeaderboardResult
func NewLeaderboardResult() LeaderboardResult {
	return LeaderboardResult{
		BoardID:                     types.NewUInt32(0),
		ContextID:                   types.NewUInt32(0),
		ResetFrequency:              types.NewUInt32(0),
		LeaderboardTotalPlayerCount: types.NewUInt32(0),
		PlayerRanks:                 types.NewList[PlayerRank](),
	}
}
