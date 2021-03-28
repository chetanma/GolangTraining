package main

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type HotelStore interface {
	AddHotels([]Hotel) error
	GetHotels() []Hotel
}

type HotelStoreJson struct {
	mutex sync.RWMutex
}

func (hs *HotelStoreJson) AddHotels(hotels []Hotel) error {
	hs.mutex.Lock()
	defer hs.mutex.Unlock()
	data, _ := json.Marshal(hotels)
	ioutil.WriteFile("data.json", data, 0644)
	return nil
}

func (hs *HotelStoreJson) GetHotels() []Hotel {
	hs.mutex.RLock()
	defer hs.mutex.RUnlock()

	data, _ := ioutil.ReadFile("data.json")

	var hotels []Hotel

	_ = json.Unmarshal(data, &hotels)

	return hotels

}
