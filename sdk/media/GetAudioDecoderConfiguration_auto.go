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

// Call_GetAudioDecoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioDecoderConfigurationResponse.
func Call_GetAudioDecoderConfiguration(ctx context.Context, dev *onvif.Device, request media.GetAudioDecoderConfiguration) (media.GetAudioDecoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioDecoderConfigurationResponse media.GetAudioDecoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioDecoderConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetAudioDecoderConfiguration")
		return reply.Body.GetAudioDecoderConfigurationResponse, errors.Annotate(err, "reply")
	}
}
