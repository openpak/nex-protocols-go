package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLookupUbiAccountIDsByUsernamesGlobal(packet nex.PacketInterface) {
	if protocol.LookupUbiAccountIDsByUsernamesGlobal == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::LookupUbiAccountIDsByUsernamesGlobal not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	usernames, err := parseLookupUbiAccountIDsByUsernamesGlobalRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.LookupUbiAccountIDsByUsernamesGlobal(err, packet, callID, usernames)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LookupUbiAccountIDsByUsernamesGlobal(nil, packet, callID, usernames)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseLookupUbiAccountIDsByUsernamesGlobalRequest decodes the parameters of Ubi Account Management::LookupUbiAccountIDsByUsernamesGlobal from the given stream
func parseLookupUbiAccountIDsByUsernamesGlobalRequest(stream *nex.ByteStreamIn) (types.List[types.String], error) {
	var usernames types.List[types.String]

	if err := usernames.ExtractFrom(stream); err != nil {
		return usernames, fmt.Errorf("failed to read usernames from parameters. %s", err.Error())
	}

	return usernames, nil
}
