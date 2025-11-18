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

// Call_RemoveAudioEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveAudioEncoderConfigurationResponse.
func Call_RemoveAudioEncoderConfiguration(ctx context.Context, dev *onvif.Device, request media.RemoveAudioEncoderConfiguration) (media.RemoveAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveAudioEncoderConfigurationResponse media.RemoveAudioEncoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveAudioEncoderConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemoveAudioEncoderConfiguration")
		return reply.Body.RemoveAudioEncoderConfigurationResponse, errors.Annotate(err, "reply")
	}
}
