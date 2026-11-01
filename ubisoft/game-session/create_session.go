package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// CreateSessionRequest holds the request parameters of Game Session::CreateSession
type CreateSessionRequest struct {
	GameSession game_session_types.GameSession
}

// WriteTo writes the CreateSessionRequest to the given writable
func (m CreateSessionRequest) WriteTo(writable types.Writable) {
	m.GameSession.WriteTo(writable)
}

// ExtractFrom extracts the CreateSessionRequest from the given readable
func (m *CreateSessionRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSession.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read CreateSession.gameSession. %s", err.Error())
	}

	return nil
}

// CreateSessionResponse holds the response parameters of Game Session::CreateSession
type CreateSessionResponse struct {
	GameSessionKey game_session_types.GameSessionKey
}

// WriteTo writes the CreateSessionResponse to the given writable
func (m CreateSessionResponse) WriteTo(writable types.Writable) {
	m.GameSessionKey.WriteTo(writable)
}

// ExtractFrom extracts the CreateSessionResponse from the given readable
func (m *CreateSessionResponse) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read CreateSession.gameSessionKey. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleCreateSession(packet nex.PacketInterface) {
	if protocol.CreateSession == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::CreateSession not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params CreateSessionRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.CreateSession(err, packet, callID, params.GameSession)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.CreateSession(nil, packet, callID, params.GameSession)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
