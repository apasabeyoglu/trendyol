package trendyol

type Category struct {
	ID            int64
	Parent        *Category
	CategoryTitle string
}
