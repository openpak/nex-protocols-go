package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	player_stats_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/player-stats/types"
)

// WriteStatsRequest holds the request parameters of Player Stats::WriteStats
type WriteStatsRequest struct {
	PlayerStatUpdates types.List[player_stats_types.PlayerStatUpdate]
}

// WriteTo writes the WriteStatsRequest to the given writable
func (m WriteStatsRequest) WriteTo(writable types.Writable) {
	m.PlayerStatUpdates.WriteTo(writable)
}

// ExtractFrom extracts the WriteStatsRequest from the given readable
func (m *WriteStatsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.PlayerStatUpdates.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read WriteStats.playerStatUpdates. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleWriteStats(packet nex.PacketInterface) {
	if protocol.WriteStats == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Player Stats::WriteStats not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params WriteStatsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.WriteStats(err, packet, callID, params.PlayerStatUpdates)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.WriteStats(nil, packet, callID, params.PlayerStatUpdates)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
