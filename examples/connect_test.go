package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/gowvp/onvif"
	m "github.com/gowvp/onvif/media"
	"github.com/gowvp/onvif/sdk/media"
)

func TestConnect(t *testing.T) {
	dev, err := onvif.NewDevice(onvif.DeviceParams{
		Xaddr:    "192.168.1.112",
		Username: "admin",
		Password: "a1234567",
	})
	if err != nil {
		panic(err)
	}

	resp, err := media.Call_GetProfiles(context.TODO(), dev, m.GetProfiles{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp.Profiles)
}
