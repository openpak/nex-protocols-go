package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	player_stats_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/player-stats/types"
)

// ReadStatsByPlayersRequest holds the request parameters of Player Stats::ReadStatsByPlayers
type ReadStatsByPlayersRequest struct {
	PlayerPIDs types.List[types.UInt32]
	Queries    types.List[player_stats_types.StatboardQuery]
}

// WriteTo writes the ReadStatsByPlayersRequest to the given writable
func (m ReadStatsByPlayersRequest) WriteTo(writable types.Writable) {
	m.PlayerPIDs.WriteTo(writable)
	m.Queries.WriteTo(writable)
}

// ExtractFrom extracts the ReadStatsByPlayersRequest from the given readable
func (m *ReadStatsByPlayersRequest) ExtractFrom(readable types.Readable) error {
	if err := m.PlayerPIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadStatsByPlayers.playerPIDs. %s", err.Error())
	}

	if err := m.Queries.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadStatsByPlayers.queries. %s", err.Error())
	}

	return nil
}

// ReadStatsByPlayersResponse holds the response parameters of Player Stats::ReadStatsByPlayers
type ReadStatsByPlayersResponse struct {
	Results types.List[player_stats_types.StatboardResult]
}

// WriteTo writes the ReadStatsByPlayersResponse to the given writable
func (m ReadStatsByPlayersResponse) WriteTo(writable types.Writable) {
	m.Results.WriteTo(writable)
}

// ExtractFrom extracts the ReadStatsByPlayersResponse from the given readable
func (m *ReadStatsByPlayersResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Results.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadStatsByPlayers.results. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleReadStatsByPlayers(packet nex.PacketInterface) {
	if protocol.ReadStatsByPlayers == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Player Stats::ReadStatsByPlayers not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params ReadStatsByPlayersRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.ReadStatsByPlayers(err, packet, callID, params.PlayerPIDs, params.Queries)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReadStatsByPlayers(nil, packet, callID, params.PlayerPIDs, params.Queries)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
