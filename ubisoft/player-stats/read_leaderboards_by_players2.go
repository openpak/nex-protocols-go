package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	player_stats_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/player-stats/types"
)

// ReadLeaderboardsByPlayers2Request holds the request parameters of Player Stats::ReadLeaderboardsByPlayers2
type ReadLeaderboardsByPlayers2Request struct {
	Queries types.List[player_stats_types.LeaderboardQuery2]
}

// WriteTo writes the ReadLeaderboardsByPlayers2Request to the given writable
func (m ReadLeaderboardsByPlayers2Request) WriteTo(writable types.Writable) {
	m.Queries.WriteTo(writable)
}

// ExtractFrom extracts the ReadLeaderboardsByPlayers2Request from the given readable
func (m *ReadLeaderboardsByPlayers2Request) ExtractFrom(readable types.Readable) error {
	if err := m.Queries.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsByPlayers2.queries. %s", err.Error())
	}

	return nil
}

// ReadLeaderboardsByPlayers2Response holds the response parameters of Player Stats::ReadLeaderboardsByPlayers2
type ReadLeaderboardsByPlayers2Response struct {
	Results types.List[player_stats_types.LeaderboardResult]
}

// WriteTo writes the ReadLeaderboardsByPlayers2Response to the given writable
func (m ReadLeaderboardsByPlayers2Response) WriteTo(writable types.Writable) {
	m.Results.WriteTo(writable)
}

// ExtractFrom extracts the ReadLeaderboardsByPlayers2Response from the given readable
func (m *ReadLeaderboardsByPlayers2Response) ExtractFrom(readable types.Readable) error {
	if err := m.Results.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsByPlayers2.results. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleReadLeaderboardsByPlayers2(packet nex.PacketInterface) {
	if protocol.ReadLeaderboardsByPlayers2 == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Player Stats::ReadLeaderboardsByPlayers2 not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params ReadLeaderboardsByPlayers2Request

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.ReadLeaderboardsByPlayers2(err, packet, callID, params.Queries)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReadLeaderboardsByPlayers2(nil, packet, callID, params.Queries)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
