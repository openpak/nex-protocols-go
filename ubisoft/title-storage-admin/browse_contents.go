package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	title_storage_admin_types "github.com/PretendoNetwork/nex-protocols-go/v2/ubisoft/title-storage-admin/types"
)

// BrowseContentsRequest holds the request parameters of Title Storage Admin::BrowseContents
type BrowseContentsRequest struct{}

// WriteTo writes the BrowseContentsRequest to the given writable
func (m BrowseContentsRequest) WriteTo(writable types.Writable) {
	_ = writable
}

// ExtractFrom extracts the BrowseContentsRequest from the given readable
func (m *BrowseContentsRequest) ExtractFrom(readable types.Readable) error {
	_ = readable
	return nil
}

// BrowseContentsResponse holds the response parameters of Title Storage Admin::BrowseContents
type BrowseContentsResponse struct {
	Contents types.List[title_storage_admin_types.AdminTitleContent]
}

// WriteTo writes the BrowseContentsResponse to the given writable
func (m BrowseContentsResponse) WriteTo(writable types.Writable) {
	m.Contents.WriteTo(writable)
}

// ExtractFrom extracts the BrowseContentsResponse from the given readable
func (m *BrowseContentsResponse) ExtractFrom(readable types.Readable) error {
	if err := m.Contents.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to read BrowseContents.contents. %s", err.Error())
	}

	return nil
}

func (protocol *Protocol) handleBrowseContents(packet nex.PacketInterface) {
	if protocol.BrowseContents == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Title Storage Admin::BrowseContents not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.BrowseContents(nil, packet, callID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
