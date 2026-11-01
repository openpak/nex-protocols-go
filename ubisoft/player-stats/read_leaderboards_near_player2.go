package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	player_stats_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/player-stats/types"
)

// ReadLeaderboardsNearPlayer2Request holds the request parameters of Player Stats::ReadLeaderboardsNearPlayer2
type ReadLeaderboardsNearPlayer2Request struct {
	PlayerPID types.UInt32
	Count     types.UInt32
	Queries   types.List[player_stats_types.LeaderboardQuery2]
}

// WriteTo writes the ReadLeaderboardsNearPlayer2Request to the given writable
func (m ReadLeaderboardsNearPlayer2Request) WriteTo(writable types.Writable) {
	m.PlayerPID.WriteTo(writable)
	m.Count.WriteTo(writable)
	m.Queries.WriteTo(writable)
}

// ExtractFrom extracts the ReadLeaderboardsNearPlayer2Request from the given readable
func (m *ReadLeaderboardsNearPlayer2Request) ExtractFrom(readable types.Readable) error {
	if err := m.PlayerPID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsNearPlayer2.playerPID. %s", err.Error())
	}

	if err := m.Count.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsNearPlayer2.count. %s", err.Error())
	}

	if err := m.Queries.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsNearPlayer2.queries. %s", err.Error())
	}

	return nil
}

// ReadLeaderboardsNearPlayer2Response holds the response parameters of Player Stats::ReadLeaderboardsNearPlayer2
type ReadLeaderboardsNearPlayer2Response struct {
	Results types.List[player_stats_types.LeaderboardResult]
}

// WriteTo writes the ReadLeaderboardsNearPlayer2Response to the given writable
func (m ReadLeaderboardsNearPlayer2Response) WriteTo(writable types.Writable) {
	m.Results.WriteTo(writable)
}

// ExtractFrom extracts the ReadLeaderboardsNearPlayer2Response from the given readable
func (m *ReadLeaderboardsNearPlayer2Response) ExtractFrom(readable types.Readable) error {
	if err := m.Results.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReadLeaderboardsNearPlayer2.results. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleReadLeaderboardsNearPlayer2(packet nex.PacketInterface) {
	if protocol.ReadLeaderboardsNearPlayer2 == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Player Stats::ReadLeaderboardsNearPlayer2 not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params ReadLeaderboardsNearPlayer2Request

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.ReadLeaderboardsNearPlayer2(err, packet, callID, params.PlayerPID, params.Count, params.Queries)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReadLeaderboardsNearPlayer2(nil, packet, callID, params.PlayerPID, params.Count, params.Queries)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
