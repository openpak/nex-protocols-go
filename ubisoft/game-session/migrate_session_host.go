package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/game-session/types"
)

// MigrateSessionHostRequest holds the request parameters of Game Session::MigrateSessionHost
type MigrateSessionHostRequest struct {
	GameSessionKey game_session_types.GameSessionKey
}

// WriteTo writes the MigrateSessionHostRequest to the given writable
func (m MigrateSessionHostRequest) WriteTo(writable types.Writable) {
	m.GameSessionKey.WriteTo(writable)
}

// ExtractFrom extracts the MigrateSessionHostRequest from the given readable
func (m *MigrateSessionHostRequest) ExtractFrom(readable types.Readable) error {
	if err := m.GameSessionKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read MigrateSessionHost.gameSessionKey. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleMigrateSessionHost(packet nex.PacketInterface) {
	if protocol.MigrateSessionHost == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::MigrateSessionHost not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params MigrateSessionHostRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.MigrateSessionHost(err, packet, callID, params.GameSessionKey)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.MigrateSessionHost(nil, packet, callID, params.GameSessionKey)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
