package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// MigrateSessionRequest holds the request parameters of Game Session::MigrateSession
type MigrateSessionRequest struct {
	GameSessionKey game_session_types.GameSessionKey
}

// WriteTo writes the MigrateSessionRequest to the given writable
func (m MigrateSessionRequest) WriteTo(writable types.Writable) {
	m.GameSessionKey.WriteTo(writable)
}

// ExtractFrom extracts the MigrateSessionRequest from the given readable
func (m *MigrateSessionRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read MigrateSession.gameSessionKey. %s", err.Error())
	}

	return nil
}

// MigrateSessionResponse holds the response parameters of Game Session::MigrateSession
type MigrateSessionResponse struct {
	GameSessionKeyMigrated game_session_types.GameSessionKey
}

// WriteTo writes the MigrateSessionResponse to the given writable
func (m MigrateSessionResponse) WriteTo(writable types.Writable) {
	m.GameSessionKeyMigrated.WriteTo(writable)
}

// ExtractFrom extracts the MigrateSessionResponse from the given readable
func (m *MigrateSessionResponse) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionKeyMigrated.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read MigrateSession.gameSessionKeyMigrated. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleMigrateSession(packet nex.PacketInterface) {
	if protocol.MigrateSession == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::MigrateSession not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params MigrateSessionRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.MigrateSession(err, packet, callID, params.GameSessionKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.MigrateSession(nil, packet, callID, params.GameSessionKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
