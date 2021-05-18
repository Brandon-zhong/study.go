package some

func allocate() {
	_ = make([]byte, 1<<20)
}

func GcDemo() {
	for n := 0; n < 100000; n++ {
		allocate()
	}
}