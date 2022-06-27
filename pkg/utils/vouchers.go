package utils

type Voucher struct {
	Points        int    `json:"points"`
	ValueInDollar string `json:"value_in_dollar"`
	Qty           int    `json:"qty"`
}

// BreakdwonVouchersToQtyOfOneUnit
func BreakdwonVouchersToQtyOfOneUnit(vs []Voucher) []Voucher {

	vouchers := []Voucher{}
	for _, v := range vs {

		count := v.Qty

		if count == 1 {
			vouchers = append(vouchers, v)
		}

		if count > 1 {
			tpoints := v.Points / v.Qty
			for count > 0 {
				tv := Voucher{
					Points:        tpoints,
					ValueInDollar: v.ValueInDollar,
					Qty:           1,
				}
				vouchers = append(vouchers, tv)
				count--
			}
			//
		}
		//
	}
	return vouchers
	//
}
