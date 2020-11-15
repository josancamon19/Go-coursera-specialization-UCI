package main
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)


func GenDisplaceFn(a, v0, s0 float64) func (float64) float64 {
	// solving for equation s =½ a t^2 + vo t + so
	fn := func(t  float64 ) float64 {
		return 0.5*a*math.Pow(t,2) + v0*t + s0
	}
	return fn
}

func xentered (val string) {
	if val == "X" {
		os.Exit(0)
	}
}

func getInput(req string, s *bufio.Scanner) float64{
	fmt.Print(req)
	s.Scan()
	es := s.Text()
	xentered(es)
	val, err := strconv.ParseFloat(es, 64)
	if err != nil {
		fmt.Println("Error processing float64 conversion")
		os.Exit(1)
	}
	return val

}

func main () {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter 'X' when you want to exit")
	a := getInput("Enter the value for acceleration: ", scanner)
	v0 := getInput("Enter the value for initial velocity: ", scanner)
	s0 := getInput("Enter the value for initial displacement: ", scanner)
	fmt.Println(a,v0,s0)
	fnt := GenDisplaceFn(a,v0,s0)
	for {
		t := getInput("Enter the value for time: ", scanner)
		fmt.Println(fnt(t))
	}

}