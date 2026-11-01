package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	player_stats_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/player-stats/types"
)

// ReadStatboardHistoryAggregatedRequest holds the request parameters of Player Stats::ReadStatboardHistoryAggregated
type ReadStatboardHistoryAggregatedRequest struct {
	Queries types.List[player_stats_types.StatboardHistoryAggregatedQuery]
}

// WriteTo writes the ReadStatboardHistoryAggregatedRequest to the given writable
func (m ReadStatboardHistoryAggregatedRequest) WriteTo(writable types.Writable) {
	m.Queries.WriteTo(writable)
}

// ExtractFrom extracts the ReadStatboardHistoryAggregatedRequest from the given readable
func (m *ReadStatboardHistoryAggregatedRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Queries.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadStatboardHistoryAggregated.queries. %s", err.Error())
	}

	return nil
}

// ReadStatboardHistoryAggregatedResponse holds the response parameters of Player Stats::ReadStatboardHistoryAggregated
type ReadStatboardHistoryAggregatedResponse struct {
	Results types.List[player_stats_types.StatboardResult]
}

// WriteTo writes the ReadStatboardHistoryAggregatedResponse to the given writable
func (m ReadStatboardHistoryAggregatedResponse) WriteTo(writable types.Writable) {
	m.Results.WriteTo(writable)
}

// ExtractFrom extracts the ReadStatboardHistoryAggregatedResponse from the given readable
func (m *ReadStatboardHistoryAggregatedResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Results.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadStatboardHistoryAggregated.results. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleReadStatboardHistoryAggregated(packet nex.PacketInterface) {
	if protocol.ReadStatboardHistoryAggregated == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Player Stats::ReadStatboardHistoryAggregated not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params ReadStatboardHistoryAggregatedRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.ReadStatboardHistoryAggregated(err, packet, callID, params.Queries)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReadStatboardHistoryAggregated(nil, packet, callID, params.Queries)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
