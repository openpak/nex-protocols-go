package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	player_stats_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/player-stats/types"
)

// ReadLeaderboardsByRank2Request holds the request parameters of Player Stats::ReadLeaderboardsByRank2
type ReadLeaderboardsByRank2Request struct {
	StartingRank types.UInt32
	Count        types.UInt32
	Queries      types.List[player_stats_types.LeaderboardQuery2]
}

// WriteTo writes the ReadLeaderboardsByRank2Request to the given writable
func (m ReadLeaderboardsByRank2Request) WriteTo(writable types.Writable) {
	m.StartingRank.WriteTo(writable)
	m.Count.WriteTo(writable)
	m.Queries.WriteTo(writable)
}

// ExtractFrom extracts the ReadLeaderboardsByRank2Request from the given readable
func (m *ReadLeaderboardsByRank2Request) ExtractFrom(readable types.Readable) error {
	if err := m.StartingRank.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsByRank2.startingRank. %s", err.Error())
	}

	if err := m.Count.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsByRank2.count. %s", err.Error())
	}

	if err := m.Queries.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsByRank2.queries. %s", err.Error())
	}

	return nil
}

// ReadLeaderboardsByRank2Response holds the response parameters of Player Stats::ReadLeaderboardsByRank2
type ReadLeaderboardsByRank2Response struct {
	Results types.List[player_stats_types.LeaderboardResult]
}

// WriteTo writes the ReadLeaderboardsByRank2Response to the given writable
func (m ReadLeaderboardsByRank2Response) WriteTo(writable types.Writable) {
	m.Results.WriteTo(writable)
}

// ExtractFrom extracts the ReadLeaderboardsByRank2Response from the given readable
func (m *ReadLeaderboardsByRank2Response) ExtractFrom(readable types.Readable) error {
	if err := m.Results.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsByRank2.results. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleReadLeaderboardsByRank2(packet nex.PacketInterface) {
	if protocol.ReadLeaderboardsByRank2 == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Player Stats::ReadLeaderboardsByRank2 not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params ReadLeaderboardsByRank2Request

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.ReadLeaderboardsByRank2(err, packet, callID, params.StartingRank, params.Count, params.Queries)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReadLeaderboardsByRank2(nil, packet, callID, params.StartingRank, params.Count, params.Queries)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
