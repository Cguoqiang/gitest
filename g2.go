package main

func G2(wg *sync.WaitGroup){
	wg.Done()
	fmt.Println("g2....")
	fmt.Println("g2....")
}
