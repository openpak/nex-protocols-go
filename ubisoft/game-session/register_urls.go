package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// RegisterURLsRequest holds the request parameters of Game Session::RegisterURLs
type RegisterURLsRequest struct {
	StationURLs types.List[types.StationURL]
}

// WriteTo writes the RegisterURLsRequest to the given writable
func (m RegisterURLsRequest) WriteTo(writable types.Writable) {
	m.StationURLs.WriteTo(writable)
}

// ExtractFrom extracts the RegisterURLsRequest from the given readable
func (m *RegisterURLsRequest) ExtractFrom(readable types.Readable) error {
	if err := m.StationURLs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read RegisterURLs.stationURLs. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleRegisterURLs(packet nex.PacketInterface) {
	if protocol.RegisterURLs == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Game Session::RegisterURLs not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params RegisterURLsRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.RegisterURLs(err, packet, callID, params.StationURLs)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RegisterURLs(nil, packet, callID, params.StationURLs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
