// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type BandwidthStatistics struct {
	Statistics struct {
		BytesOut    *int `json:"bytes_out"`
		BytesOutLo  *int `json:"bytes_out_lo"`
		BytesDropLo *int `json:"bytes_drop_lo"`
		Guarantee   *int `json:"guarantee"`
		PktsDropLo  *int `json:"pkts_drop_lo"`
		PktsDrop    *int `json:"pkts_drop"`
		PktsDropHi  *int `json:"pkts_drop_hi"`
		Maximum     *int `json:"maximum"`
		BytesDrop   *int `json:"bytes_drop"`
		BytesOutHi  *int `json:"bytes_out_hi"`
		BytesDropHi *int `json:"bytes_drop_hi"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetBandwidthStatistics(name string) (*BandwidthStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/bandwidth/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(BandwidthStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
