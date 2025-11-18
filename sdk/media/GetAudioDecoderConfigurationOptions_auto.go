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

// Call_GetAudioDecoderConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioDecoderConfigurationOptionsResponse.
func Call_GetAudioDecoderConfigurationOptions(ctx context.Context, dev *onvif.Device, request media.GetAudioDecoderConfigurationOptions) (media.GetAudioDecoderConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioDecoderConfigurationOptionsResponse media.GetAudioDecoderConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioDecoderConfigurationOptionsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetAudioDecoderConfigurationOptions")
		return reply.Body.GetAudioDecoderConfigurationOptionsResponse, errors.Annotate(err, "reply")
	}
}
