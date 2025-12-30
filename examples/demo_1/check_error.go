package main

func check_error(err error) {
	if err != nil {
		panic(err)
	}
}
