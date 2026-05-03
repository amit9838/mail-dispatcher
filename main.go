package main

import (
	"sync"

	"github.com/amit9838/email-dispatcher/internal"
)

func main() {
	recipientChannel := make(chan internal.Recipient)

	go func() {
		internal.LoadRecipient("./emails.csv", recipientChannel)
	}()

	var wg sync.WaitGroup
	workerCount := 5

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go internal.EmailWorker(i, recipientChannel, &wg)
	}

	wg.Wait()
}
