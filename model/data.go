package model

type Data struct {
	MultiId string `db:"multi_id"`
	RouteId string `db:"route_id"`
	DeviceId string `db:"device_id"`
	AddressId string `db:"address_id"`
	ServiceId string  `db:"service_id"`
	DataId string `db:"data_id"`
	UserId string `db:"user_id"`
	CId     string `db:"cid"`
	MacId   string `db:"mac_id"`
	Ifn     string `db:"ifn"`
}
