package lesson11

import "fmt"

func Test() {
	main()

	ucr := UserCreateRequest{
		FirstName: "John",
	}

	fmt.Println(Validate(ucr))
}
