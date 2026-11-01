package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// DeleteSessionRequest holds the request parameters of Game Session::DeleteSession
type DeleteSessionRequest struct {
	GameSessionKey game_session_types.GameSessionKey
}

// WriteTo writes the DeleteSessionRequest to the given writable
func (m DeleteSessionRequest) WriteTo(writable types.Writable) {
	m.GameSessionKey.WriteTo(writable)
}

// ExtractFrom extracts the DeleteSessionRequest from the given readable
func (m *DeleteSessionRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read DeleteSession.gameSessionKey. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleDeleteSession(packet nex.PacketInterface) {
	if protocol.DeleteSession == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::DeleteSession not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params DeleteSessionRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.DeleteSession(err, packet, callID, params.GameSessionKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DeleteSession(nil, packet, callID, params.GameSessionKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
