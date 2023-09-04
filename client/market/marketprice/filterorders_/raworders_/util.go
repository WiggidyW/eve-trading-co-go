package raworders_

func appendOrder(
	s *[]MarketOrder,
	price float64,
	quantity int64,
) {
	*s = append(
		*s,
		MarketOrder{
			Price:    price,
			Quantity: quantity,
		},
	)
}