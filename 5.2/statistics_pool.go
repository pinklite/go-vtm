// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 5.2.
package vtm

import (
	"encoding/json"
)

type PoolStatistics struct {
	Statistics struct {
		Algorithm          *string `json:"algorithm"`
		Persistence        *string `json:"persistence"`
		BwLimitBytesDropLo *int    `json:"bw_limit_bytes_drop_lo"`
		MaxQueueTime       *int    `json:"max_queue_time"`
		Disabled           *int    `json:"disabled"`
		ConnsQueued        *int    `json:"conns_queued"`
		SessionMigrated    *int    `json:"session_migrated"`
		State              *string `json:"state"`
		BytesInLo          *int    `json:"bytes_in_lo"`
		BytesOut           *int    `json:"bytes_out"`
		BwLimitBytesDropHi *int    `json:"bw_limit_bytes_drop_hi"`
		Draining           *int    `json:"draining"`
		MinQueueTime       *int    `json:"min_queue_time"`
		BytesOutHi         *int    `json:"bytes_out_hi"`
		TotalConn          *int    `json:"total_conn"`
		BwLimitBytesDrop   *int    `json:"bw_limit_bytes_drop"`
		BytesInHi          *int    `json:"bytes_in_hi"`
		BytesOutLo         *int    `json:"bytes_out_lo"`
		QueueTimeouts      *int    `json:"queue_timeouts"`
		BwLimitPktsDropHi  *int    `json:"bw_limit_pkts_drop_hi"`
		Nodes              *int    `json:"nodes"`
		BwLimitPktsDropLo  *int    `json:"bw_limit_pkts_drop_lo"`
		BwLimitPktsDrop    *int    `json:"bw_limit_pkts_drop"`
		MeanQueueTime      *int    `json:"mean_queue_time"`
		BytesIn            *int    `json:"bytes_in"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetPoolStatistics(name string) (*PoolStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/5.2/status/local_tm/statistics/pools/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(PoolStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
