package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	player_stats_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/player-stats/types"
)

// ReadLeaderboardsNearPlayerRequest holds the request parameters of Player Stats::ReadLeaderboardsNearPlayer
type ReadLeaderboardsNearPlayerRequest struct {
	PlayerPID types.UInt32
	Count     types.UInt32
	Queries   types.List[player_stats_types.LeaderboardQuery]
}

// WriteTo writes the ReadLeaderboardsNearPlayerRequest to the given writable
func (m ReadLeaderboardsNearPlayerRequest) WriteTo(writable types.Writable) {
	m.PlayerPID.WriteTo(writable)
	m.Count.WriteTo(writable)
	m.Queries.WriteTo(writable)
}

// ExtractFrom extracts the ReadLeaderboardsNearPlayerRequest from the given readable
func (m *ReadLeaderboardsNearPlayerRequest) ExtractFrom(readable types.Readable) error {
	if err := m.PlayerPID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsNearPlayer.playerPID. %s", err.Error())
	}

	if err := m.Count.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsNearPlayer.count. %s", err.Error())
	}

	if err := m.Queries.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsNearPlayer.queries. %s", err.Error())
	}

	return nil
}

// ReadLeaderboardsNearPlayerResponse holds the response parameters of Player Stats::ReadLeaderboardsNearPlayer
type ReadLeaderboardsNearPlayerResponse struct {
	Results types.List[player_stats_types.LeaderboardResult]
}

// WriteTo writes the ReadLeaderboardsNearPlayerResponse to the given writable
func (m ReadLeaderboardsNearPlayerResponse) WriteTo(writable types.Writable) {
	m.Results.WriteTo(writable)
}

// ExtractFrom extracts the ReadLeaderboardsNearPlayerResponse from the given readable
func (m *ReadLeaderboardsNearPlayerResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Results.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsNearPlayer.results. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleReadLeaderboardsNearPlayer(packet nex.PacketInterface) {
	if protocol.ReadLeaderboardsNearPlayer == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Player Stats::ReadLeaderboardsNearPlayer not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params ReadLeaderboardsNearPlayerRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.ReadLeaderboardsNearPlayer(err, packet, callID, params.PlayerPID, params.Count, params.Queries)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReadLeaderboardsNearPlayer(nil, packet, callID, params.PlayerPID, params.Count, params.Queries)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
