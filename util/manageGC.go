package util

import "fmt"

func acquire(name string) (string, error) {
	return name, nil
}

func relese(name string) {
	fmt.Println("releasing", name)
}

func generalWorker() {
	res, err := acquire("A")
	if err != nil {
		fmt.Println("Error on acuire:", err)
		return
	}
	defer relese(res)
	res1, err1 := acquire("B")
	if err1 != nil {
		fmt.Println("Error on acuire:", err1)
		return
	}
	defer relese(res1)
	fmt.Println("worker started")
}

func TryGeneralWorker() {
	generalWorker()
	generalWorker2()
}

func generalWorker2() {
	for _, v := range []string{"C", "D"} {
		// typically here would start anonymous function
		name, err := acquire(v)
		if err != nil {
			fmt.Println("Error on acuire:", err)
			return
		}
		defer relese(name)
		// and will end here
		// as defer binded to the end of func
		// not block
		// but here on purpose
		// i use it deferently as i want first of all
		// worker started2 and only then release
		// in reversed order
	}
	fmt.Println("worker started2")
}
