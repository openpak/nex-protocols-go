package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PlayerRank is a type within the Player Stats protocol
type PlayerRank struct {
	types.Structure
	PlayerStatSet PlayerStatSet
	RankStatus    types.UInt32
	Rank          types.UInt32
	Score         types.Variant
}

// WriteTo writes the PlayerRank to the given writable
func (pr PlayerRank) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	pr.PlayerStatSet.WriteTo(contentWritable)
	pr.RankStatus.WriteTo(contentWritable)
	pr.Rank.WriteTo(contentWritable)
	pr.Score.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	pr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the PlayerRank from the given readable
func (pr *PlayerRank) ExtractFrom(readable types.Readable) error {
	if err := pr.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PlayerRank header. %s", err.Error())
	}

	if err := pr.PlayerStatSet.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PlayerRank.PlayerStatSet. %s", err.Error())
	}

	if err := pr.RankStatus.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PlayerRank.RankStatus. %s", err.Error())
	}

	if err := pr.Rank.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PlayerRank.Rank. %s", err.Error())
	}

	if err := pr.Score.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract PlayerRank.Score. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PlayerRank
func (pr PlayerRank) Copy() types.RVType {
	copied := NewPlayerRank()

	copied.StructureVersion = pr.StructureVersion
	copied.PlayerStatSet = pr.PlayerStatSet.Copy().(PlayerStatSet)
	copied.RankStatus = pr.RankStatus.Copy().(types.UInt32)
	copied.Rank = pr.Rank.Copy().(types.UInt32)
	copied.Score = pr.Score.Copy().(types.Variant)

	return copied
}

// Equals checks if the given PlayerRank contains the same data as the current PlayerRank
func (pr PlayerRank) Equals(o types.RVType) bool {
	other, ok := o.(PlayerRank)
	if !ok {
		return false
	}

	if pr.StructureVersion != other.StructureVersion {
		return false
	}

	if !(pr.PlayerStatSet.Equals(other.PlayerStatSet)) {
		return false
	}

	if !(pr.RankStatus.Equals(other.RankStatus)) {
		return false
	}

	if !(pr.Rank.Equals(other.Rank)) {
		return false
	}

	if !(pr.Score.Equals(other.Score)) {
		return false
	}

	return true
}

// CopyRef copies the current value of the PlayerRank
// and returns a pointer to the new copy
func (pr PlayerRank) CopyRef() types.RVTypePtr {
	copied := pr.Copy().(PlayerRank)
	return &copied
}

// Deref takes a pointer to the PlayerRank
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (pr *PlayerRank) Deref() types.RVType {
	return *pr
}

// String returns the string representation of the PlayerRank
func (pr PlayerRank) String() string {
	return pr.FormatToString(0)
}

// FormatToString pretty-prints the PlayerRank using the provided indentation level
func (pr PlayerRank) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PlayerRank{\n")
	fmt.Fprintf(&b, "%sPlayerStatSet: %s,\n", indentationValues, pr.PlayerStatSet.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%sRankStatus: %s,\n", indentationValues, pr.RankStatus)
	fmt.Fprintf(&b, "%sRank: %s,\n", indentationValues, pr.Rank)
	fmt.Fprintf(&b, "%sScore: %s,\n", indentationValues, pr.Score)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewPlayerRank returns a new PlayerRank
func NewPlayerRank() PlayerRank {
	return PlayerRank{
		PlayerStatSet: NewPlayerStatSet(),
		RankStatus:    types.NewUInt32(0),
		Rank:          types.NewUInt32(0),
		Score:         types.NewVariant(),
	}
}
