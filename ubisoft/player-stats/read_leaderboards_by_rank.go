package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	player_stats_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/player-stats/types"
)

// ReadLeaderboardsByRankRequest holds the request parameters of Player Stats::ReadLeaderboardsByRank
type ReadLeaderboardsByRankRequest struct {
	StartingRank types.UInt32
	Count        types.UInt32
	Queries      types.List[player_stats_types.LeaderboardQuery]
}

// WriteTo writes the ReadLeaderboardsByRankRequest to the given writable
func (m ReadLeaderboardsByRankRequest) WriteTo(writable types.Writable) {
	m.StartingRank.WriteTo(writable)
	m.Count.WriteTo(writable)
	m.Queries.WriteTo(writable)
}

// ExtractFrom extracts the ReadLeaderboardsByRankRequest from the given readable
func (m *ReadLeaderboardsByRankRequest) ExtractFrom(readable types.Readable) error {
	if err := m.StartingRank.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsByRank.startingRank. %s", err.Error())
	}

	if err := m.Count.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsByRank.count. %s", err.Error())
	}

	if err := m.Queries.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsByRank.queries. %s", err.Error())
	}

	return nil
}

// ReadLeaderboardsByRankResponse holds the response parameters of Player Stats::ReadLeaderboardsByRank
type ReadLeaderboardsByRankResponse struct {
	Results types.List[player_stats_types.LeaderboardResult]
}

// WriteTo writes the ReadLeaderboardsByRankResponse to the given writable
func (m ReadLeaderboardsByRankResponse) WriteTo(writable types.Writable) {
	m.Results.WriteTo(writable)
}

// ExtractFrom extracts the ReadLeaderboardsByRankResponse from the given readable
func (m *ReadLeaderboardsByRankResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Results.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsByRank.results. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleReadLeaderboardsByRank(packet nex.PacketInterface) {
	if protocol.ReadLeaderboardsByRank == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Player Stats::ReadLeaderboardsByRank not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params ReadLeaderboardsByRankRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.ReadLeaderboardsByRank(err, packet, callID, params.StartingRank, params.Count, params.Queries)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReadLeaderboardsByRank(nil, packet, callID, params.StartingRank, params.Count, params.Queries)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
