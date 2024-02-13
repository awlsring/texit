package appinit

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
