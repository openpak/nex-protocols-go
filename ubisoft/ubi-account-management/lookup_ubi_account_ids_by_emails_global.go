package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLookupUbiAccountIDsByEmailsGlobal(packet nex.PacketInterface) {
	if protocol.LookupUbiAccountIDsByEmailsGlobal == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::LookupUbiAccountIDsByEmailsGlobal not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	emails, err := parseLookupUbiAccountIDsByEmailsGlobalRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.LookupUbiAccountIDsByEmailsGlobal(err, packet, callID, emails)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LookupUbiAccountIDsByEmailsGlobal(nil, packet, callID, emails)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseLookupUbiAccountIDsByEmailsGlobalRequest decodes the parameters of Ubi Account Management::LookupUbiAccountIDsByEmailsGlobal from the given stream
func parseLookupUbiAccountIDsByEmailsGlobalRequest(stream *nex.ByteStreamIn) (types.List[types.String], error) {
	var emails types.List[types.String]

	if err := emails.ExtractFrom(stream); err != nil {
		return emails, fmt.Errorf("failed to read emails from parameters. %s", err.Error())
	}

	return emails, nil
}
