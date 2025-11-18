// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"

	"github.com/gowvp/onvif"
	"github.com/gowvp/onvif/media"
	"github.com/gowvp/onvif/sdk"
	"github.com/juju/errors"
)

// Call_RemovePTZConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemovePTZConfigurationResponse.
func Call_RemovePTZConfiguration(ctx context.Context, dev *onvif.Device, request media.RemovePTZConfiguration) (media.RemovePTZConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemovePTZConfigurationResponse media.RemovePTZConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemovePTZConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemovePTZConfiguration")
		return reply.Body.RemovePTZConfigurationResponse, errors.Annotate(err, "reply")
	}
}
