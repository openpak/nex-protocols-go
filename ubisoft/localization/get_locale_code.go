package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// GetLocaleCodeRequest holds the request parameters of Localization::GetLocaleCode
type GetLocaleCodeRequest struct{}

// WriteTo writes the GetLocaleCodeRequest to the given writable
func (m GetLocaleCodeRequest) WriteTo(writable types.Writable) {
	_ = writable
}

// ExtractFrom extracts the GetLocaleCodeRequest from the given readable
func (m *GetLocaleCodeRequest) ExtractFrom(readable types.Readable) error {
	_ = readable
	return nil
}

// GetLocaleCodeResponse holds the response parameters of Localization::GetLocaleCode
type GetLocaleCodeResponse struct {
	LocalCode types.String
}

// WriteTo writes the GetLocaleCodeResponse to the given writable
func (m GetLocaleCodeResponse) WriteTo(writable types.Writable) {
	m.LocalCode.WriteTo(writable)
}

// ExtractFrom extracts the GetLocaleCodeResponse from the given readable
func (m *GetLocaleCodeResponse) ExtractFrom(readable types.Readable) error {
	if err := m.LocalCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetLocaleCode.localCode. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetLocaleCode(packet nex.PacketInterface) {
	if protocol.GetLocaleCode == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Localization::GetLocaleCode not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.GetLocaleCode(nil, packet, callID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
