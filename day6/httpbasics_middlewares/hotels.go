package main

type HotelStore interface {
	GetHotel(id int) *Hotel
	GetHotels() []Hotel
}

type HotelsDataStore struct {
	Data []Hotel
}

func NewHotelsDataStore() HotelsDataStore {
	data := make([]Hotel, 0, 100)
	data = append(data, Hotel{
		Name:     "Hotel1",
		Id:       1001,
		Location: "Pune",
	})

	data = append(data, Hotel{
		Name:     "Hotel1",
		Id:       1001,
		Location: "Pune",
	})

	data = append(data, Hotel{
		Name:     "Hotel2",
		Id:       1002,
		Location: "Mumbai",
	})

	data = append(data, Hotel{
		Name:     "Hotel3",
		Id:       1003,
		Location: "Pune",
	})

	data = append(data, Hotel{
		Name:     "Hotel4",
		Id:       1004,
		Location: "Pune",
	})

	data = append(data, Hotel{
		Name:     "Hotel5",
		Id:       1005,
		Location: "Hyd",
	})

	return HotelsDataStore{
		Data: data,
	}
}

func (hd *HotelsDataStore) GetHotel(id int) *Hotel {
	for _, h := range hd.Data {
		if h.Id == id {
			return &h
		}
	}

	return nil
}

func (hd *HotelsDataStore) GetHotels() []Hotel {
	return hd.Data
}
