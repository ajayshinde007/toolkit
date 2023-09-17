package tools

import (
	"crypto/rand"
)

// Tools is the type used to instantiate this module .
// Any varibale of this type will have access to all the methods with the receiver
type Tools struct{}

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// RandomString returns string of random characters of length  n  using randomStringSource
// as a source
func (t *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len((r)))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
		// fmt.Println(s)
	}
	return string(s)
}

func (t *Tools) RandomString1(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)
	//fmt.Println(s)
	//fmt.Println("\n----------------------------\n")
	//fmt.Println(r)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		//fmt.Println(i, p)
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}

// func generateRandomString(n int) string {
// 	s := make([]rune, n)

// 	r := []rune(randomStringSource) // converts string char to int array with ASCII value

// 	for _, i := range s {
// 		p, _ := rand.Prime(rand.Reader, len((r)))
// 		x, y := p.Uint64(), uint64(len(r))
// 		println("x :", p.String())
// 		println("y :", y)
// 		s[i] = r[x%y]
// 	}
// 	return string(s) // converts  int array with ASCII value to string
// }

// func main() {

// 	var t Tools
// 	println(t.RandomString(5))

// 	// println(generateRandomString(5))
// 	// s := make([]rune, 3)
// 	// println(s[2])

// 	// var i [3]int32
// 	// println(i[2])

// 	// r := []int32("12345")

// 	// // println(string(r[3]))
// 	// p, _ := rand.Prime(rand.Reader, len((r)))
// 	// println(p.Uint64())
// 	// p1, _ := rand.Prime(rand.Reader, 2)
// 	// println(p1.Uint64())
// 	// p2, _ := rand.Prime(rand.Reader, 3)
// 	// println(p2.Uint64())
// 	// p3, _ := rand.Prime(rand.Reader, 4)
// 	// println(p3.Uint64())
// 	// p4, _ := rand.Prime(rand.Reader, 5)
// 	// println(p4.Uint64())
// 	// p5, _ := rand.Prime(rand.Reader, 6)
// 	// println(p5.Uint64())

// }
