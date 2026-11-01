package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	player_stats_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/player-stats/types"
)

// ReadPopulationStatsRequest holds the request parameters of Player Stats::ReadPopulationStats
type ReadPopulationStatsRequest struct {
	Queries types.List[player_stats_types.PopulationStatQuery]
}

// WriteTo writes the ReadPopulationStatsRequest to the given writable
func (m ReadPopulationStatsRequest) WriteTo(writable types.Writable) {
	m.Queries.WriteTo(writable)
}

// ExtractFrom extracts the ReadPopulationStatsRequest from the given readable
func (m *ReadPopulationStatsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.Queries.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadPopulationStats.queries. %s", err.Error())
	}

	return nil
}

// ReadPopulationStatsResponse holds the response parameters of Player Stats::ReadPopulationStats
type ReadPopulationStatsResponse struct {
	Results types.List[player_stats_types.PopulationStatResult]
}

// WriteTo writes the ReadPopulationStatsResponse to the given writable
func (m ReadPopulationStatsResponse) WriteTo(writable types.Writable) {
	m.Results.WriteTo(writable)
}

// ExtractFrom extracts the ReadPopulationStatsResponse from the given readable
func (m *ReadPopulationStatsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Results.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadPopulationStats.results. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleReadPopulationStats(packet nex.PacketInterface) {
	if protocol.ReadPopulationStats == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Player Stats::ReadPopulationStats not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params ReadPopulationStatsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.ReadPopulationStats(err, packet, callID, params.Queries)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReadPopulationStats(nil, packet, callID, params.Queries)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
