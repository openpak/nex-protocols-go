// Package protocol implements the Player Stats protocol of the Quazal Rendez-Vous (Ubisoft) family.
//
// Method IDs and layouts follow the NintendoClients wiki page "Player Stats Protocol" (MIT).
//
// Layout provisional:
//   - ReadLeaderboardsByPlayers2 takes only the query list on the wiki; the player list is presumably carried in LeaderboardQuery2.m_estimatedPIDs.
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	player_stats_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/player-stats/types"
)

const (
	// ProtocolID is the protocol ID for the Player Stats protocol
	ProtocolID = 0x37

	// MethodWriteStats is the method ID for method WriteStats
	MethodWriteStats = 0x1

	// MethodReadStatsByPlayers is the method ID for method ReadStatsByPlayers
	MethodReadStatsByPlayers = 0x2

	// MethodReadLeaderboardsNearPlayer is the method ID for method ReadLeaderboardsNearPlayer
	MethodReadLeaderboardsNearPlayer = 0x3

	// MethodReadLeaderboardsByRank is the method ID for method ReadLeaderboardsByRank
	MethodReadLeaderboardsByRank = 0x4

	// MethodReadLeaderboardsByPlayers is the method ID for method ReadLeaderboardsByPlayers
	MethodReadLeaderboardsByPlayers = 0x5

	// MethodReadStatboardHistory is the method ID for method ReadStatboardHistory
	MethodReadStatboardHistory = 0x6

	// MethodReadLeaderboardHistory is the method ID for method ReadLeaderboardHistory
	MethodReadLeaderboardHistory = 0x7

	// MethodReadStatboardHistoryAggregated is the method ID for method ReadStatboardHistoryAggregated
	MethodReadStatboardHistoryAggregated = 0x8

	// MethodGetStatboardNextPurgeDate is the method ID for method GetStatboardNextPurgeDate
	MethodGetStatboardNextPurgeDate = 0x9

	// MethodReadLeaderboardsNearPlayer2 is the method ID for method ReadLeaderboardsNearPlayer2
	MethodReadLeaderboardsNearPlayer2 = 0xa

	// MethodReadLeaderboardsByRank2 is the method ID for method ReadLeaderboardsByRank2
	MethodReadLeaderboardsByRank2 = 0xb

	// MethodReadLeaderboardsByPlayers2 is the method ID for method ReadLeaderboardsByPlayers2
	MethodReadLeaderboardsByPlayers2 = 0xc

	// MethodReadPopulationStats is the method ID for method ReadPopulationStats
	MethodReadPopulationStats = 0xd
)

// Protocol handles the Player Stats protocol
type Protocol struct {
	endpoint                       nex.EndpointInterface
	WriteStats                     func(err error, packet nex.PacketInterface, callID uint32, playerStatUpdates types.List[player_stats_types.PlayerStatUpdate]) (*nex.RMCMessage, *nex.Error)
	ReadStatsByPlayers             func(err error, packet nex.PacketInterface, callID uint32, playerPIDs types.List[types.UInt32], queries types.List[player_stats_types.StatboardQuery]) (*nex.RMCMessage, *nex.Error)
	ReadLeaderboardsNearPlayer     func(err error, packet nex.PacketInterface, callID uint32, playerPID types.UInt32, count types.UInt32, queries types.List[player_stats_types.LeaderboardQuery]) (*nex.RMCMessage, *nex.Error)
	ReadLeaderboardsByRank         func(err error, packet nex.PacketInterface, callID uint32, startingRank types.UInt32, count types.UInt32, queries types.List[player_stats_types.LeaderboardQuery]) (*nex.RMCMessage, *nex.Error)
	ReadLeaderboardsByPlayers      func(err error, packet nex.PacketInterface, callID uint32, playerPIDs types.List[types.UInt32], queries types.List[player_stats_types.LeaderboardQuery]) (*nex.RMCMessage, *nex.Error)
	ReadStatboardHistory           func(err error, packet nex.PacketInterface, callID uint32, queries types.List[player_stats_types.StatboardHistoryQuery]) (*nex.RMCMessage, *nex.Error)
	ReadLeaderboardHistory         func(err error, packet nex.PacketInterface, callID uint32, queries types.List[player_stats_types.LeaderboardHistoryQuery]) (*nex.RMCMessage, *nex.Error)
	ReadStatboardHistoryAggregated func(err error, packet nex.PacketInterface, callID uint32, queries types.List[player_stats_types.StatboardHistoryAggregatedQuery]) (*nex.RMCMessage, *nex.Error)
	GetStatboardNextPurgeDate      func(err error, packet nex.PacketInterface, callID uint32, boardID types.UInt32, resetFrequency types.UInt32) (*nex.RMCMessage, *nex.Error)
	ReadLeaderboardsNearPlayer2    func(err error, packet nex.PacketInterface, callID uint32, playerPID types.UInt32, count types.UInt32, queries types.List[player_stats_types.LeaderboardQuery2]) (*nex.RMCMessage, *nex.Error)
	ReadLeaderboardsByRank2        func(err error, packet nex.PacketInterface, callID uint32, startingRank types.UInt32, count types.UInt32, queries types.List[player_stats_types.LeaderboardQuery2]) (*nex.RMCMessage, *nex.Error)
	ReadLeaderboardsByPlayers2     func(err error, packet nex.PacketInterface, callID uint32, queries types.List[player_stats_types.LeaderboardQuery2]) (*nex.RMCMessage, *nex.Error)
	ReadPopulationStats            func(err error, packet nex.PacketInterface, callID uint32, queries types.List[player_stats_types.PopulationStatQuery]) (*nex.RMCMessage, *nex.Error)
	Patches                        nex.ServiceProtocol
	PatchedMethods                 []uint32
}

// Interface implements the methods present on the Player Stats protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerWriteStats(handler func(err error, packet nex.PacketInterface, callID uint32, playerStatUpdates types.List[player_stats_types.PlayerStatUpdate]) (*nex.RMCMessage, *nex.Error))
	SetHandlerReadStatsByPlayers(handler func(err error, packet nex.PacketInterface, callID uint32, playerPIDs types.List[types.UInt32], queries types.List[player_stats_types.StatboardQuery]) (*nex.RMCMessage, *nex.Error))
	SetHandlerReadLeaderboardsNearPlayer(handler func(err error, packet nex.PacketInterface, callID uint32, playerPID types.UInt32, count types.UInt32, queries types.List[player_stats_types.LeaderboardQuery]) (*nex.RMCMessage, *nex.Error))
	SetHandlerReadLeaderboardsByRank(handler func(err error, packet nex.PacketInterface, callID uint32, startingRank types.UInt32, count types.UInt32, queries types.List[player_stats_types.LeaderboardQuery]) (*nex.RMCMessage, *nex.Error))
	SetHandlerReadLeaderboardsByPlayers(handler func(err error, packet nex.PacketInterface, callID uint32, playerPIDs types.List[types.UInt32], queries types.List[player_stats_types.LeaderboardQuery]) (*nex.RMCMessage, *nex.Error))
	SetHandlerReadStatboardHistory(handler func(err error, packet nex.PacketInterface, callID uint32, queries types.List[player_stats_types.StatboardHistoryQuery]) (*nex.RMCMessage, *nex.Error))
	SetHandlerReadLeaderboardHistory(handler func(err error, packet nex.PacketInterface, callID uint32, queries types.List[player_stats_types.LeaderboardHistoryQuery]) (*nex.RMCMessage, *nex.Error))
	SetHandlerReadStatboardHistoryAggregated(handler func(err error, packet nex.PacketInterface, callID uint32, queries types.List[player_stats_types.StatboardHistoryAggregatedQuery]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetStatboardNextPurgeDate(handler func(err error, packet nex.PacketInterface, callID uint32, boardID types.UInt32, resetFrequency types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerReadLeaderboardsNearPlayer2(handler func(err error, packet nex.PacketInterface, callID uint32, playerPID types.UInt32, count types.UInt32, queries types.List[player_stats_types.LeaderboardQuery2]) (*nex.RMCMessage, *nex.Error))
	SetHandlerReadLeaderboardsByRank2(handler func(err error, packet nex.PacketInterface, callID uint32, startingRank types.UInt32, count types.UInt32, queries types.List[player_stats_types.LeaderboardQuery2]) (*nex.RMCMessage, *nex.Error))
	SetHandlerReadLeaderboardsByPlayers2(handler func(err error, packet nex.PacketInterface, callID uint32, queries types.List[player_stats_types.LeaderboardQuery2]) (*nex.RMCMessage, *nex.Error))
	SetHandlerReadPopulationStats(handler func(err error, packet nex.PacketInterface, callID uint32, queries types.List[player_stats_types.PopulationStatQuery]) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerWriteStats sets the handler for the WriteStats method
func (protocol *Protocol) SetHandlerWriteStats(handler func(err error, packet nex.PacketInterface, callID uint32, playerStatUpdates types.List[player_stats_types.PlayerStatUpdate]) (*nex.RMCMessage, *nex.Error)) {
	protocol.WriteStats = handler
}

// SetHandlerReadStatsByPlayers sets the handler for the ReadStatsByPlayers method
func (protocol *Protocol) SetHandlerReadStatsByPlayers(handler func(err error, packet nex.PacketInterface, callID uint32, playerPIDs types.List[types.UInt32], queries types.List[player_stats_types.StatboardQuery]) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReadStatsByPlayers = handler
}

// SetHandlerReadLeaderboardsNearPlayer sets the handler for the ReadLeaderboardsNearPlayer method
func (protocol *Protocol) SetHandlerReadLeaderboardsNearPlayer(handler func(err error, packet nex.PacketInterface, callID uint32, playerPID types.UInt32, count types.UInt32, queries types.List[player_stats_types.LeaderboardQuery]) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReadLeaderboardsNearPlayer = handler
}

// SetHandlerReadLeaderboardsByRank sets the handler for the ReadLeaderboardsByRank method
func (protocol *Protocol) SetHandlerReadLeaderboardsByRank(handler func(err error, packet nex.PacketInterface, callID uint32, startingRank types.UInt32, count types.UInt32, queries types.List[player_stats_types.LeaderboardQuery]) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReadLeaderboardsByRank = handler
}

// SetHandlerReadLeaderboardsByPlayers sets the handler for the ReadLeaderboardsByPlayers method
func (protocol *Protocol) SetHandlerReadLeaderboardsByPlayers(handler func(err error, packet nex.PacketInterface, callID uint32, playerPIDs types.List[types.UInt32], queries types.List[player_stats_types.LeaderboardQuery]) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReadLeaderboardsByPlayers = handler
}

// SetHandlerReadStatboardHistory sets the handler for the ReadStatboardHistory method
func (protocol *Protocol) SetHandlerReadStatboardHistory(handler func(err error, packet nex.PacketInterface, callID uint32, queries types.List[player_stats_types.StatboardHistoryQuery]) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReadStatboardHistory = handler
}

// SetHandlerReadLeaderboardHistory sets the handler for the ReadLeaderboardHistory method
func (protocol *Protocol) SetHandlerReadLeaderboardHistory(handler func(err error, packet nex.PacketInterface, callID uint32, queries types.List[player_stats_types.LeaderboardHistoryQuery]) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReadLeaderboardHistory = handler
}

// SetHandlerReadStatboardHistoryAggregated sets the handler for the ReadStatboardHistoryAggregated method
func (protocol *Protocol) SetHandlerReadStatboardHistoryAggregated(handler func(err error, packet nex.PacketInterface, callID uint32, queries types.List[player_stats_types.StatboardHistoryAggregatedQuery]) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReadStatboardHistoryAggregated = handler
}

// SetHandlerGetStatboardNextPurgeDate sets the handler for the GetStatboardNextPurgeDate method
func (protocol *Protocol) SetHandlerGetStatboardNextPurgeDate(handler func(err error, packet nex.PacketInterface, callID uint32, boardID types.UInt32, resetFrequency types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetStatboardNextPurgeDate = handler
}

// SetHandlerReadLeaderboardsNearPlayer2 sets the handler for the ReadLeaderboardsNearPlayer2 method
func (protocol *Protocol) SetHandlerReadLeaderboardsNearPlayer2(handler func(err error, packet nex.PacketInterface, callID uint32, playerPID types.UInt32, count types.UInt32, queries types.List[player_stats_types.LeaderboardQuery2]) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReadLeaderboardsNearPlayer2 = handler
}

// SetHandlerReadLeaderboardsByRank2 sets the handler for the ReadLeaderboardsByRank2 method
func (protocol *Protocol) SetHandlerReadLeaderboardsByRank2(handler func(err error, packet nex.PacketInterface, callID uint32, startingRank types.UInt32, count types.UInt32, queries types.List[player_stats_types.LeaderboardQuery2]) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReadLeaderboardsByRank2 = handler
}

// SetHandlerReadLeaderboardsByPlayers2 sets the handler for the ReadLeaderboardsByPlayers2 method
func (protocol *Protocol) SetHandlerReadLeaderboardsByPlayers2(handler func(err error, packet nex.PacketInterface, callID uint32, queries types.List[player_stats_types.LeaderboardQuery2]) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReadLeaderboardsByPlayers2 = handler
}

// SetHandlerReadPopulationStats sets the handler for the ReadPopulationStats method
func (protocol *Protocol) SetHandlerReadPopulationStats(handler func(err error, packet nex.PacketInterface, callID uint32, queries types.List[player_stats_types.PopulationStatQuery]) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReadPopulationStats = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if protocol.Patches != nil && slices.Contains(protocol.PatchedMethods, message.MethodID) {
		protocol.Patches.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodWriteStats:
		protocol.handleWriteStats(packet)
	case MethodReadStatsByPlayers:
		protocol.handleReadStatsByPlayers(packet)
	case MethodReadLeaderboardsNearPlayer:
		protocol.handleReadLeaderboardsNearPlayer(packet)
	case MethodReadLeaderboardsByRank:
		protocol.handleReadLeaderboardsByRank(packet)
	case MethodReadLeaderboardsByPlayers:
		protocol.handleReadLeaderboardsByPlayers(packet)
	case MethodReadStatboardHistory:
		protocol.handleReadStatboardHistory(packet)
	case MethodReadLeaderboardHistory:
		protocol.handleReadLeaderboardHistory(packet)
	case MethodReadStatboardHistoryAggregated:
		protocol.handleReadStatboardHistoryAggregated(packet)
	case MethodGetStatboardNextPurgeDate:
		protocol.handleGetStatboardNextPurgeDate(packet)
	case MethodReadLeaderboardsNearPlayer2:
		protocol.handleReadLeaderboardsNearPlayer2(packet)
	case MethodReadLeaderboardsByRank2:
		protocol.handleReadLeaderboardsByRank2(packet)
	case MethodReadLeaderboardsByPlayers2:
		protocol.handleReadLeaderboardsByPlayers2(packet)
	case MethodReadPopulationStats:
		protocol.handleReadPopulationStats(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Player Stats method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Player Stats protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
