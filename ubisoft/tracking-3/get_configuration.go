package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// GetConfigurationRequest holds the request parameters of Tracking 3::GetConfiguration
type GetConfigurationRequest struct{}

// WriteTo writes the GetConfigurationRequest to the given writable
func (m GetConfigurationRequest) WriteTo(writable types.Writable) {
	_ = writable
}

// ExtractFrom extracts the GetConfigurationRequest from the given readable
func (m *GetConfigurationRequest) ExtractFrom(readable types.Readable) error {
	_ = readable
	return nil
}

// GetConfigurationResponse holds the response parameters of Tracking 3::GetConfiguration
type GetConfigurationResponse struct {
	Tags types.List[types.String]
}

// WriteTo writes the GetConfigurationResponse to the given writable
func (m GetConfigurationResponse) WriteTo(writable types.Writable) {
	m.Tags.WriteTo(writable)
}

// ExtractFrom extracts the GetConfigurationResponse from the given readable
func (m *GetConfigurationResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Tags.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read GetConfiguration.tags. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleGetConfiguration(packet nex.PacketInterface) {
	if protocol.GetConfiguration == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Tracking 3::GetConfiguration not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.GetConfiguration(nil, packet, callID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
