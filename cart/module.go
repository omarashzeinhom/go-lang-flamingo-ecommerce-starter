package cart

func (*service) Schema() []byte {
	//language=graphql
	return []byte(`
	query cartexample {
		Commerce_Cart {
		  cart {
			id
			itemCount
			deliveries {
			  deliveryInfo {
				code
			  }
			  cartitems {
				qty
				productName
			  }
			}
		  }
		}
	  }
	  
	`)
}
