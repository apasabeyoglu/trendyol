package trendyol

type Campaign struct {
	Code             string
	Percentage       float64
	ProductCategory  []string
	RequiredQuantity int64
}

func (c *Campaign) ValidForCategory(category Category) bool {
	for {
		for _, cat := range c.ProductCategory {
			if category.CategoryTitle == cat {
				return true
			}
		}
		if 	category.Parent != nil {
			category = *category.Parent
		} else {
			return false
		}
	}
}
