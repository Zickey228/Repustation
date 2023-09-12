package main

func (cli *CLI) recvReview() {
	go func() {
		recvReview()
	}()

	go func() {
		checkReview()
	}()

}
