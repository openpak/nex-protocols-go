package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleForgetPassword(packet nex.PacketInterface) {
	if protocol.ForgetPassword == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::ForgetPassword not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	usernameOrEmail, err := parseForgetPasswordRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.ForgetPassword(err, packet, callID, usernameOrEmail)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ForgetPassword(nil, packet, callID, usernameOrEmail)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseForgetPasswordRequest decodes the parameters of Ubi Account Management::ForgetPassword from the given stream
func parseForgetPasswordRequest(stream *nex.ByteStreamIn) (types.String, error) {
	var usernameOrEmail types.String

	if err := usernameOrEmail.ExtractFrom(stream); err != nil {
		return usernameOrEmail, fmt.Errorf("failed to read usernameOrEmail from parameters. %s", err.Error())
	}

	return usernameOrEmail, nil
}
