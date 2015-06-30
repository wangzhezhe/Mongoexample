package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Setenv("MONGODB_PORT_27017_TCP_ADDR", "10.10.105.204")
	//os.Setenv("TESTENV", "123")

	fmt.Println(os.Getenv("abc"))
}
