package helper

func DefaultPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func CustomPanic(err error, detail string) {
	if err != nil {
		panic(detail)
	}
}
