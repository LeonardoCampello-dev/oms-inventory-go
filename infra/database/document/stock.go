package stockdocument

type Document struct {
	Sku       string `bson:"sku"`
	AccountId string `bson:"accountId"`
	Quantity  int    `bson:"quantity"`
	Id        string `bson:"id"`
}
