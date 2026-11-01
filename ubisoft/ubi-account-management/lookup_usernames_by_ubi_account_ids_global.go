package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLookupUsernamesByUbiAccountIDsGlobal(packet nex.PacketInterface) {
	if protocol.LookupUsernamesByUbiAccountIDsGlobal == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::LookupUsernamesByUbiAccountIDsGlobal not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	ubiAccountIds, err := parseLookupUsernamesByUbiAccountIDsGlobalRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.LookupUsernamesByUbiAccountIDsGlobal(err, packet, callID, ubiAccountIds)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LookupUsernamesByUbiAccountIDsGlobal(nil, packet, callID, ubiAccountIds)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseLookupUsernamesByUbiAccountIDsGlobalRequest decodes the parameters of Ubi Account Management::LookupUsernamesByUbiAccountIDsGlobal from the given stream
func parseLookupUsernamesByUbiAccountIDsGlobalRequest(stream *nex.ByteStreamIn) (types.List[types.String], error) {
	var ubiAccountIds types.List[types.String]

	if err := ubiAccountIds.ExtractFrom(stream); err != nil {
		return ubiAccountIds, fmt.Errorf("failed to read ubiAccountIds from parameters. %s", err.Error())
	}

	return ubiAccountIds, nil
}
