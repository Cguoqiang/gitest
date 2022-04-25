package main

func G1(wg *sync.WaitGroup){
	wg.Done()
	fmt.Println("g1....")
	fmt.Println("g1....")
	fmt.Println("g1....")
	fmt.Println("g1....")
	fmt.Println("g1..22..")
}
