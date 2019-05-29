package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ana-maria-ricardo/exercise/errorhandling"
)

func main() {

	new_err := errorhandling.OrderNotFoundError(errors.New("something's wrong"))
	fmt.Println(new_err)
	g, err := json.Marshal(new_err)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(g))
}
