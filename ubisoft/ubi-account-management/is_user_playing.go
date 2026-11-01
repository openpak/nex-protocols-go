package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleIsUserPlaying(packet nex.PacketInterface) {
	if protocol.IsUserPlaying == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UbiAccountManagement::IsUserPlaying not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	ubiAccountIDs, err := parseIsUserPlayingRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.IsUserPlaying(err, packet, callID, ubiAccountIDs)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.IsUserPlaying(nil, packet, callID, ubiAccountIDs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseIsUserPlayingRequest decodes the parameters of Ubi Account Management::IsUserPlaying from the given stream
func parseIsUserPlayingRequest(stream *nex.ByteStreamIn) (types.List[types.String], error) {
	var ubiAccountIDs types.List[types.String]

	if err := ubiAccountIDs.ExtractFrom(stream); err != nil {
		return ubiAccountIDs, fmt.Errorf("failed to read ubiAccountIDs from parameters. %s", err.Error())
	}

	return ubiAccountIDs, nil
}
