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

// Call_AddPTZConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddPTZConfigurationResponse.
func Call_AddPTZConfiguration(ctx context.Context, dev *onvif.Device, request media.AddPTZConfiguration) (media.AddPTZConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddPTZConfigurationResponse media.AddPTZConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddPTZConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "AddPTZConfiguration")
		return reply.Body.AddPTZConfigurationResponse, errors.Annotate(err, "reply")
	}
}
