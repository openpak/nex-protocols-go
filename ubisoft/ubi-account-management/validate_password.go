package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleValidatePassword(packet nex.PacketInterface) {
	if protocol.ValidatePassword == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::ValidatePassword not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	password, username, err := parseValidatePasswordRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.ValidatePassword(err, packet, callID, password, username)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ValidatePassword(nil, packet, callID, password, username)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseValidatePasswordRequest decodes the parameters of Ubi Account Management::ValidatePassword from the given stream
func parseValidatePasswordRequest(stream *nex.ByteStreamIn) (types.String, types.String, error) {
	var password types.String
	var username types.String

	if err := password.ExtractFrom(stream); err != nil {
		return password, username, fmt.Errorf("failed to read password from parameters. %s", err.Error())
	}

	if err := username.ExtractFrom(stream); err != nil {
		return password, username, fmt.Errorf("failed to read username from parameters. %s", err.Error())
	}

	return password, username, nil
}
