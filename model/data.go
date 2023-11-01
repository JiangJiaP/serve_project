package model

type Data struct {
	MultiId   string `db:"multi_id"`
	RouteId   string `db:"route_id"`
	DeviceId  string `db:"device_id"`
	AddressId string `db:"address_id"`
	ServiceId string `db:"service_id"`
	DataId    string `db:"data_id"`
	UserId    string `db:"user_id"`
	CId       string `db:"cid"`
	MacId     string `db:"mac_id"`
	Ifn       string `db:"ifn"`
}

type RouterData struct {
	Router string `db:"router" json:"Router"`
	IpAddr string `db:"ip_addr"  json:"IPv6Addr"`
}

type ConnectData struct {
	Scid string `db:"scid" json:"scid"`
	Dcid string `db:"dcid" json:"dcid"`
}

type ConnectRouterData struct {
	Cid      string `db:"cid" json:"cid"`
	RouterId string `db:"router_id" json:"router_id"`
	Mac      string `db:"mac" json:"mac"`
	Ifn      string `db:"ifn" json:"ifn"`
}

type MultiData struct {
	UserId    string `db:"user_id" json:"user_id"`
	DeviceId  string `db:"device_id" json:"device_id"`
	AddressId string `db:"address_id" json:"address_id"`
	ServiceId string `db:"service_id" json:"service_id"`
	DataId    string `db:"data_id" json:"data_id"`
}
