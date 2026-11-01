package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleValidateUsername(packet nex.PacketInterface) {
	if protocol.ValidateUsername == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::ValidateUsername not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	username, err := parseValidateUsernameRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.ValidateUsername(err, packet, callID, username)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ValidateUsername(nil, packet, callID, username)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseValidateUsernameRequest decodes the parameters of Ubi Account Management::ValidateUsername from the given stream
func parseValidateUsernameRequest(stream *nex.ByteStreamIn) (types.String, error) {
	var username types.String

	if err := username.ExtractFrom(stream); err != nil {
		return username, fmt.Errorf("failed to read username from parameters. %s", err.Error())
	}

	return username, nil
}
