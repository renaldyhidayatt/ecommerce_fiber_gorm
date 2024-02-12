package mapper

type Mapper struct {
	CategoryMapper CategoryMapping
	ProductMapper  ProductMapping
	OrderMapper    OrderMapping
	UserMapper     UserMapping
	ReviewMapper   ReviewMapping
}

func NewMapper() *Mapper {
	return &Mapper{
		CategoryMapper: NewCategoryMapper(),
		ProductMapper:  NewProductMapper(),
		OrderMapper:    NewOrderMapper(),
		UserMapper:     NewUserMapper(),
		ReviewMapper:   NewReviewMapper(),
	}
}
