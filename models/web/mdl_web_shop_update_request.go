package web

type ShopUpdateRequest struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
