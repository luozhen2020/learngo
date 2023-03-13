package persist

import "log"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		// save items
		for {
			itemCount++
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
		}
	}()
	return out
}
