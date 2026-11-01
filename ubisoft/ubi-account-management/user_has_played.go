package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleUserHasPlayed(packet nex.PacketInterface) {
	if protocol.UserHasPlayed == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::UserHasPlayed not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	ubiAccountIDs, err := parseUserHasPlayedRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.UserHasPlayed(err, packet, callID, ubiAccountIDs)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UserHasPlayed(nil, packet, callID, ubiAccountIDs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseUserHasPlayedRequest decodes the parameters of Ubi Account Management::UserHasPlayed from the given stream
func parseUserHasPlayedRequest(stream *nex.ByteStreamIn) (types.List[types.String], error) {
	var ubiAccountIDs types.List[types.String]

	if err := ubiAccountIDs.ExtractFrom(stream); err != nil {
		return ubiAccountIDs, fmt.Errorf("failed to read ubiAccountIDs from parameters. %s", err.Error())
	}

	return ubiAccountIDs, nil
}
