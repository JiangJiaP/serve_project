package model

type Data struct {
	MultiId string `db:"multi_id"`
	RouteId string `db:"route_id"`
	CId     string `db:"cid"`
	MacId   string `db:"mac_id"`
	Ifi     string `db:"ifi"`
}
