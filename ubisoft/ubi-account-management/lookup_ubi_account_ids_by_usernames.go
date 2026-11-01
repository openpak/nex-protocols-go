package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLookupUbiAccountIDsByUsernames(packet nex.PacketInterface) {
	if protocol.LookupUbiAccountIDsByUsernames == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::LookupUbiAccountIDsByUsernames not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	usernames, err := parseLookupUbiAccountIDsByUsernamesRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.LookupUbiAccountIDsByUsernames(err, packet, callID, usernames)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LookupUbiAccountIDsByUsernames(nil, packet, callID, usernames)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseLookupUbiAccountIDsByUsernamesRequest decodes the parameters of Ubi Account Management::LookupUbiAccountIDsByUsernames from the given stream
func parseLookupUbiAccountIDsByUsernamesRequest(stream *nex.ByteStreamIn) (types.List[types.String], error) {
	var usernames types.List[types.String]

	if err := usernames.ExtractFrom(stream); err != nil {
		return usernames, fmt.Errorf("failed to read usernames from parameters. %s", err.Error())
	}

	return usernames, nil
}
