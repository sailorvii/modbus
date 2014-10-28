// Copyright 2014 Quoc-Viet Nguyen. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.
package test

import (
	"github.com/goburrow/modbus"
	"log"
	"os"
	"testing"
)

const (
	testRTUDevice = "/dev/pts/6"
)

func TestRTUClientReadCoils(t *testing.T) {
	client := modbus.RTUClient(testRTUDevice)
	ClientTestReadCoils(t, client)
}

func TestRTUClientAdvancedUsage(t *testing.T) {
	var handler modbus.RTUClientHandler
	handler.Address = testRTUDevice
	handler.BaudRate = 19200
	handler.CharSize = 8
	handler.Parity = "E"
	handler.StopBits = 1
	handler.SlaveId = 17
	handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	err := handler.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer handler.Close()

	client := modbus.RTUClientWithHandler(&handler)
	results, err := client.ReadDiscreteInputs(15, 2)
	if err != nil || results == nil {
		t.Fatal(err, results)
	}
}
