package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLookupUbiAccountIDsByUsernameSubString(packet nex.PacketInterface) {
	if protocol.LookupUbiAccountIDsByUsernameSubString == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::LookupUbiAccountIDsByUsernameSubString not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	usernameSubString, err := parseLookupUbiAccountIDsByUsernameSubStringRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.LookupUbiAccountIDsByUsernameSubString(err, packet, callID, usernameSubString)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LookupUbiAccountIDsByUsernameSubString(nil, packet, callID, usernameSubString)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseLookupUbiAccountIDsByUsernameSubStringRequest decodes the parameters of Ubi Account Management::LookupUbiAccountIDsByUsernameSubString from the given stream
func parseLookupUbiAccountIDsByUsernameSubStringRequest(stream *nex.ByteStreamIn) (types.String, error) {
	var usernameSubString types.String

	if err := usernameSubString.ExtractFrom(stream); err != nil {
		return usernameSubString, fmt.Errorf("failed to read usernameSubString from parameters. %s", err.Error())
	}

	return usernameSubString, nil
}
