package util

import "log"

func a(f func()) {
	defer func() {
		if p := recover(); p != nil {
			log.Fatal(p)
		}
	}()
	f()
}
