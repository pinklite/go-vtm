// Copyright (C) 2018, Pulse Secure, LLC. 
// Licensed under the terms of the MPL 2.0. See LICENSE file for details.

// Go library for Pulse Virtual Traffic Manager REST version 4.0.
package vtm

import (
	"encoding/json"
)

type CacheSslSessionCacheStatistics struct {
	Statistics struct {
		Entries    *int `json:"entries"`
		Misses     *int `json:"misses"`
		Hits       *int `json:"hits"`
		HitRate    *int `json:"hit_rate"`
		EntriesMax *int `json:"entries_max"`
		Lookups    *int `json:"lookups"`
		Oldest     *int `json:"oldest"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetCacheSslSessionCacheStatistics() (*CacheSslSessionCacheStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/4.0/status/local_tm/statistics/cache/ssl_session_cache")
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(CacheSslSessionCacheStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
