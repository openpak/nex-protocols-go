package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// SetLocaleCodeRequest holds the request parameters of Localization::SetLocaleCode
type SetLocaleCodeRequest struct {
	LocalCode types.String
}

// WriteTo writes the SetLocaleCodeRequest to the given writable
func (m SetLocaleCodeRequest) WriteTo(writable types.Writable) {
	m.LocalCode.WriteTo(writable)
}

// ExtractFrom extracts the SetLocaleCodeRequest from the given readable
func (m *SetLocaleCodeRequest) ExtractFrom(readable types.Readable) error {
	if err := m.LocalCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read SetLocaleCode.localCode. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleSetLocaleCode(packet nex.PacketInterface) {
	if protocol.SetLocaleCode == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Localization::SetLocaleCode not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var params SetLocaleCodeRequest

	if err := params.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.SetLocaleCode(err, packet, callID, params.LocalCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SetLocaleCode(nil, packet, callID, params.LocalCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
