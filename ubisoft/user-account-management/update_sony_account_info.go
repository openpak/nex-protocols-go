package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleUpdateSonyAccountInfo(packet nex.PacketInterface) {
	if protocol.UpdateSonyAccountInfo == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UserAccountManagement::UpdateSonyAccountInfo not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	ticketData, ticketSize, err := parseUpdateSonyAccountInfoRequest(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateSonyAccountInfo(err, packet, callID, ticketData, ticketSize)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateSonyAccountInfo(nil, packet, callID, ticketData, ticketSize)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}

// parseUpdateSonyAccountInfoRequest decodes the parameters of User Account Management::UpdateSonyAccountInfo from the given stream
func parseUpdateSonyAccountInfoRequest(stream *nex.ByteStreamIn) (types.QBuffer, types.UInt32, error) {
	var ticketData types.QBuffer
	var ticketSize types.UInt32

	if err := ticketData.ExtractFrom(stream); err != nil {
		return ticketData, ticketSize, fmt.Errorf("failed to read ticketData from parameters. %s", err.Error())
	}

	if err := ticketSize.ExtractFrom(stream); err != nil {
		return ticketData, ticketSize, fmt.Errorf("failed to read ticketSize from parameters. %s", err.Error())
	}

	return ticketData, ticketSize, nil
}
