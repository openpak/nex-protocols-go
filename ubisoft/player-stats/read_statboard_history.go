package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	player_stats_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/player-stats/types"
)

// ReadStatboardHistoryRequest holds the request parameters of Player Stats::ReadStatboardHistory
type ReadStatboardHistoryRequest struct {
	Queries types.List[player_stats_types.StatboardHistoryQuery]
}

// WriteTo writes the ReadStatboardHistoryRequest to the given writable
func (m ReadStatboardHistoryRequest) WriteTo(writable types.Writable) {
	m.Queries.WriteTo(writable)
}

// ExtractFrom extracts the ReadStatboardHistoryRequest from the given readable
func (m *ReadStatboardHistoryRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Queries.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadStatboardHistory.queries. %s", err.Error())
	}

	return nil
}

// ReadStatboardHistoryResponse holds the response parameters of Player Stats::ReadStatboardHistory
type ReadStatboardHistoryResponse struct {
	Results types.List[player_stats_types.StatboardResult]
}

// WriteTo writes the ReadStatboardHistoryResponse to the given writable
func (m ReadStatboardHistoryResponse) WriteTo(writable types.Writable) {
	m.Results.WriteTo(writable)
}

// ExtractFrom extracts the ReadStatboardHistoryResponse from the given readable
func (m *ReadStatboardHistoryResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Results.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadStatboardHistory.results. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleReadStatboardHistory(packet nex.PacketInterface) {
	if protocol.ReadStatboardHistory == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Player Stats::ReadStatboardHistory not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params ReadStatboardHistoryRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.ReadStatboardHistory(err, packet, callID, params.Queries)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReadStatboardHistory(nil, packet, callID, params.Queries)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
