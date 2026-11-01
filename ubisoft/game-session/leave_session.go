package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// LeaveSessionRequest holds the request parameters of Game Session::LeaveSession
type LeaveSessionRequest struct {
	GameSessionKey game_session_types.GameSessionKey
}

// WriteTo writes the LeaveSessionRequest to the given writable
func (m LeaveSessionRequest) WriteTo(writable types.Writable) {
	m.GameSessionKey.WriteTo(writable)
}

// ExtractFrom extracts the LeaveSessionRequest from the given readable
func (m *LeaveSessionRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read LeaveSession.gameSessionKey. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleLeaveSession(packet nex.PacketInterface) {
	if protocol.LeaveSession == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::LeaveSession not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params LeaveSessionRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.LeaveSession(err, packet, callID, params.GameSessionKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LeaveSession(nil, packet, callID, params.GameSessionKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
