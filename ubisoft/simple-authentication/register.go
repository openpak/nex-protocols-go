package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleRegister(packet nex.PacketInterface) {
	if protocol.Register == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SimpleAuthentication::Register not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	vecMyURLs, err := parseRegisterRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.Register(err, packet, callID, vecMyURLs)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.Register(nil, packet, callID, vecMyURLs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseRegisterRequest decodes the parameters of Simple Authentication::Register from the given stream
func parseRegisterRequest(stream *nex.ByteStreamIn) (types.List[types.StationURL], error) {
	var vecMyURLs types.List[types.StationURL]

	if err := vecMyURLs.ExtractFrom(stream); err != nil {
		return vecMyURLs, fmt.Errorf("failed to read vecMyURLs from parameters. %s", err.Error())
	}

	return vecMyURLs, nil
}
