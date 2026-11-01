package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// ReportUnsuccessfulJoinSessionsRequest holds the request parameters of Game Session::ReportUnsuccessfulJoinSessions
type ReportUnsuccessfulJoinSessionsRequest struct {
	UnsuccessfulJoinSessions types.List[game_session_types.GameSessionUnsuccessfulJoinSession]
}

// WriteTo writes the ReportUnsuccessfulJoinSessionsRequest to the given writable
func (m ReportUnsuccessfulJoinSessionsRequest) WriteTo(writable types.Writable) {
	m.UnsuccessfulJoinSessions.WriteTo(writable)
}

// ExtractFrom extracts the ReportUnsuccessfulJoinSessionsRequest from the given readable
func (m *ReportUnsuccessfulJoinSessionsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.UnsuccessfulJoinSessions.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read ReportUnsuccessfulJoinSessions.unsuccessfulJoinSessions. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleReportUnsuccessfulJoinSessions(packet nex.PacketInterface) {
	if protocol.ReportUnsuccessfulJoinSessions == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::ReportUnsuccessfulJoinSessions not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params ReportUnsuccessfulJoinSessionsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.ReportUnsuccessfulJoinSessions(err, packet, callID, params.UnsuccessfulJoinSessions)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReportUnsuccessfulJoinSessions(nil, packet, callID, params.UnsuccessfulJoinSessions)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
