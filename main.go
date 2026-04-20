package main

import (
	//"database/sql"
	"fmt"
	"log"
	"svetlana/first-app/config"
	"svetlana/first-app/util"

	// lets add .env
	//"os"

	"github.com/joho/godotenv"

	// This is the driver we just "got"
	//In Go, _ means “import for side effects only”
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sqlConnsMap := map[string]string{
		"mysql": "mysql",
	}
	for name, driver := range sqlConnsMap {
		config.InitConnection(name, driver)
		fmt.Printf("Connection '%s' is ready!\n", name)
		config.PingConn(name)
		fmt.Printf("\u2705 Connection '%s' was 📌 added to map and is ready!\n", name)
		defer config.DBRegistry[name].Close()
	}
	fmt.Println("********")
	// we do not have to import those functions, they are public
	// available everywhere
	SayHello()
	fmt.Println("********")
	CountWordsPrint()
	fmt.Println("********")
	FindMaxInArr()
	fmt.Println("********")
	PlayAndPrint()
	fmt.Println("********")
	playMathSqrt(4.0)
	playMathSqrt(-4.0)
	fmt.Println("********")
	fmt.Println("Pay attention!\n", "\t* release will start in the end\n", "\t* First released is the last defer!")
	util.TryGeneralWorker()
	fmt.Println("********")
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
