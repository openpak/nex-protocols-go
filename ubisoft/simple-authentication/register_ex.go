package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleRegisterEx(packet nex.PacketInterface) {
	if protocol.RegisterEx == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SimpleAuthentication::RegisterEx not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	vecMyURLs, hCustomData, err := parseRegisterExRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.RegisterEx(err, packet, callID, vecMyURLs, hCustomData)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RegisterEx(nil, packet, callID, vecMyURLs, hCustomData)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseRegisterExRequest decodes the parameters of Simple Authentication::RegisterEx from the given stream
func parseRegisterExRequest(stream *nex.ByteStreamIn) (types.List[types.StationURL], types.DataHolder, error) {
	var vecMyURLs types.List[types.StationURL]
	var hCustomData types.DataHolder

	if err := vecMyURLs.ExtractFrom(stream); err != nil {
		return vecMyURLs, hCustomData, fmt.Errorf("failed to read vecMyURLs from parameters. %s", err.Error())
	}

	if err := hCustomData.ExtractFrom(stream); err != nil {
		return vecMyURLs, hCustomData, fmt.Errorf("failed to read hCustomData from parameters. %s", err.Error())
	}

	return vecMyURLs, hCustomData, nil
}
