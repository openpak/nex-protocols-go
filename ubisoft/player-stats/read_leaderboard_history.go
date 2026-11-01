package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	player_stats_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/player-stats/types"
)

// ReadLeaderboardHistoryRequest holds the request parameters of Player Stats::ReadLeaderboardHistory
type ReadLeaderboardHistoryRequest struct {
	Queries types.List[player_stats_types.LeaderboardHistoryQuery]
}

// WriteTo writes the ReadLeaderboardHistoryRequest to the given writable
func (m ReadLeaderboardHistoryRequest) WriteTo(writable types.Writable) {
	m.Queries.WriteTo(writable)
}

// ExtractFrom extracts the ReadLeaderboardHistoryRequest from the given readable
func (m *ReadLeaderboardHistoryRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Queries.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardHistory.queries. %s", err.Error())
	}

	return nil
}

// ReadLeaderboardHistoryResponse holds the response parameters of Player Stats::ReadLeaderboardHistory
type ReadLeaderboardHistoryResponse struct {
	Results types.List[player_stats_types.LeaderboardResult]
}

// WriteTo writes the ReadLeaderboardHistoryResponse to the given writable
func (m ReadLeaderboardHistoryResponse) WriteTo(writable types.Writable) {
	m.Results.WriteTo(writable)
}

// ExtractFrom extracts the ReadLeaderboardHistoryResponse from the given readable
func (m *ReadLeaderboardHistoryResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Results.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardHistory.results. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleReadLeaderboardHistory(packet nex.PacketInterface) {
	if protocol.ReadLeaderboardHistory == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Player Stats::ReadLeaderboardHistory not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params ReadLeaderboardHistoryRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.ReadLeaderboardHistory(err, packet, callID, params.Queries)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReadLeaderboardHistory(nil, packet, callID, params.Queries)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
