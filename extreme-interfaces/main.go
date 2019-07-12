package main

type user struct {
	Name string
}

func main() {

	useInterface()
	useStruct()

}

func useInterface() user {
	return returnInterface().(user)
}
func useStruct() user {
	return returnUser()
}

func returnInterface() interface{} {
	return user{Name: "Roshan"}
}

func returnUser() user {
	return user{Name: "Roshan"}
}
