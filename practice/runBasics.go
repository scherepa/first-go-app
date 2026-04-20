package practice

import (
	"fmt"
	"svetlana/first-app/util"
)

func wrap(name string, fn func()) {
	startMarker(name)
	defer endMarker()
	fn()
}
func startMarker(name string) {
	fmt.Println("***", name)
}

func endMarker() {
	fmt.Println("***")
}

func RunLearning() {
	// we do not have to import those functions, they are public
	// available everywhere
	wrap("SayHello", SayHello)
	wrap("CountWordsPrint", CountWordsPrint)
	wrap("FindMaxInArr", FindMaxInArr)
	wrap("PlayAndPrint", PlayAndPrint)
	playMathSqrt(4.0)
	playMathSqrt(-4.0)
	fmt.Println("********")
	fmt.Println("Pay attention!\n", "\t* release will start in the end\n", "\t* First released is the last defer!")
	wrap("Defer", util.TryGeneralWorker)
	header, errHeader := util.CallURLGetHeader("https://www.example.com")
	fmt.Printf("response content type is:\n\t%s\nerror is:\n\t%v\n", header, errHeader)
	fmt.Println("********")
	util.NewPost("simple", "la la boo text")
	fmt.Println("********")
	util.NewPostClassic("classic", "la la boo text")
	fmt.Println("********")
}

func playMathSqrt(n float64) {
	res, err := util.SqrtCalc(n)
	if err != nil {
		fmt.Printf("Calculation failed. err: %v\n", err)
		// if needed to exit with error use log fatal
	}
	fmt.Printf("sqrt of %f returned is: %f\n", n, res)
}
