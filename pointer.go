package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(*address2)

	var address4 *Address = new(Address)
	address4.City = "Kediri"
	fmt.Println(address4)

	address := Address{"Subang", "Jawa Barat", ""}
	ChangeAddressToIndonesia(&address)

	fmt.Println(address)
}

func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}
