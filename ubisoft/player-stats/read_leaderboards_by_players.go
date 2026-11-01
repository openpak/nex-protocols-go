package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	player_stats_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/player-stats/types"
)

// ReadLeaderboardsByPlayersRequest holds the request parameters of Player Stats::ReadLeaderboardsByPlayers
type ReadLeaderboardsByPlayersRequest struct {
	PlayerPIDs types.List[types.UInt32]
	Queries    types.List[player_stats_types.LeaderboardQuery]
}

// WriteTo writes the ReadLeaderboardsByPlayersRequest to the given writable
func (m ReadLeaderboardsByPlayersRequest) WriteTo(writable types.Writable) {
	m.PlayerPIDs.WriteTo(writable)
	m.Queries.WriteTo(writable)
}

// ExtractFrom extracts the ReadLeaderboardsByPlayersRequest from the given readable
func (m *ReadLeaderboardsByPlayersRequest) ExtractFrom(readable types.Readable) error {
	if err := m.PlayerPIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsByPlayers.playerPIDs. %s", err.Error())
	}

	if err := m.Queries.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsByPlayers.queries. %s", err.Error())
	}

	return nil
}

// ReadLeaderboardsByPlayersResponse holds the response parameters of Player Stats::ReadLeaderboardsByPlayers
type ReadLeaderboardsByPlayersResponse struct {
	Results types.List[player_stats_types.LeaderboardResult]
}

// WriteTo writes the ReadLeaderboardsByPlayersResponse to the given writable
func (m ReadLeaderboardsByPlayersResponse) WriteTo(writable types.Writable) {
	m.Results.WriteTo(writable)
}

// ExtractFrom extracts the ReadLeaderboardsByPlayersResponse from the given readable
func (m *ReadLeaderboardsByPlayersResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Results.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsByPlayers.results. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleReadLeaderboardsByPlayers(packet nex.PacketInterface) {
	if protocol.ReadLeaderboardsByPlayers == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Player Stats::ReadLeaderboardsByPlayers not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params ReadLeaderboardsByPlayersRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.ReadLeaderboardsByPlayers(err, packet, callID, params.PlayerPIDs, params.Queries)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReadLeaderboardsByPlayers(nil, packet, callID, params.PlayerPIDs, params.Queries)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
