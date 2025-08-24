package models


//InputMessage representa el mensaje de entrada del tópico Pub/Sub
type InputMessage struct {
	Name					string `json:"name"`
	LogisticCondition		string `json:"logisticCondition"`
	Presentation			Presentation `json:"presentation"`
	WarehouseConfiguration	WarehouseConfiguration `json:"warehouseConfiguration"`
}

type Presentation struct {
	SellerCategoryID string `json:"sellerCategoryId"`
	ItemSize string `json:"ItemSize"`
	TransportInfo TransportInfo `json:"transportInfo"`
}

type TransportInfo struct {
	Name string `json:"name"`
}

type WarehouseConfiguration struct {
	Location string `json:"location"`
}