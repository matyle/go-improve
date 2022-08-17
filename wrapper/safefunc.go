package wrapper

func SafeFunc(f func()) {
	defer func() {
		if r := recover(); r != nil {
			f()
		}
	}()
	f()
}
