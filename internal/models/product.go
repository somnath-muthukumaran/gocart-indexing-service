package models

type Product struct {
	ProductId          string   `bson:"productId" json:"productId"`
	ProductName        string   `bson:"productName" json:"productName"`
	Categories         []string `bson:"categories" json:"categories"`
	ActualPrice        float32  `bson:"actualPrice" json:"actualPrice"`
	DiscountedPrice    *float32 `bson:"discountedPrice" json:"discountedPrice"`
	DiscountPercentage *float32 `bson:"discountPercentage" json:"discountPercentage"`
	Rating             *float32 `bson:"rating" json:"rating"`
	RatingCount        *int32   `bson:"ratingCount" json:"ratingCount"`
	BaseModel
}
