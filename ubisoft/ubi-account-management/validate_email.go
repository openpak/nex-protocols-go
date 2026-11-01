package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleValidateEmail(packet nex.PacketInterface) {
	if protocol.ValidateEmail == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::ValidateEmail not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	email, err := parseValidateEmailRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.ValidateEmail(err, packet, callID, email)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ValidateEmail(nil, packet, callID, email)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseValidateEmailRequest decodes the parameters of Ubi Account Management::ValidateEmail from the given stream
func parseValidateEmailRequest(stream *nex.ByteStreamIn) (types.String, error) {
	var email types.String

	if err := email.ExtractFrom(stream); err != nil {
		return email, fmt.Errorf("failed to read email from parameters. %s", err.Error())
	}

	return email, nil
}
