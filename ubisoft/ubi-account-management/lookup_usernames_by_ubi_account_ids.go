package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLookupUsernamesByUbiAccountIDs(packet nex.PacketInterface) {
	if protocol.LookupUsernamesByUbiAccountIDs == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::LookupUsernamesByUbiAccountIDs not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	ubiAccountIds, err := parseLookupUsernamesByUbiAccountIDsRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.LookupUsernamesByUbiAccountIDs(err, packet, callID, ubiAccountIds)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LookupUsernamesByUbiAccountIDs(nil, packet, callID, ubiAccountIds)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseLookupUsernamesByUbiAccountIDsRequest decodes the parameters of Ubi Account Management::LookupUsernamesByUbiAccountIDs from the given stream
func parseLookupUsernamesByUbiAccountIDsRequest(stream *nex.ByteStreamIn) (types.List[types.String], error) {
	var ubiAccountIds types.List[types.String]

	if err := ubiAccountIds.ExtractFrom(stream); err != nil {
		return ubiAccountIds, fmt.Errorf("failed to read ubiAccountIds from parameters. %s", err.Error())
	}

	return ubiAccountIds, nil
}
