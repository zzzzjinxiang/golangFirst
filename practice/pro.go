package practice

import "fmt"

func f(collecion DataCollection) {
	s := collecion.Load()
	if s != nil {
		fmt.Printf("")
	}
	for iter := collecion.Iterator(); ; {
		item, err := iter.Next()
		if err != nil {
			break;
		}
		dataset := item.(Dataset)
		if dataset != nil {
			fmt.Printf("")
		}
	}
}
