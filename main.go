// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"github.com/itsubaki/q"
)

func main() {

	qmac := q.New()
	bob := qmac.Zero()

	fmt.Print("bob inital state: ", bob, qmac.State()[0])

	qmac.H(bob).CNOT(bob, qmac.One())
	qmac.Measure(bob, qmac.Zero())

	aqmac := q.New()
	alice := aqmac.One()

	fmt.Println("\nAlice inital state: ", alice, aqmac.State()[0])

	aqmac.H(alice).CNOT(aqmac.Zero(), bob)
	aqmac.Measure(bob, aqmac.One())

	for s := range qmac.State() {
		fmt.Println("bob states :", s, qmac.State())
	}
	for s := range aqmac.State() {
		fmt.Println("alice state :", s, aqmac.State())
	}

}
