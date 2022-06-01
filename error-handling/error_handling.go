package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var resource Resource

	defer func() {
		if r := recover(); r != nil {
			if v, ok := r.(FrobError); ok {
				resource.Defrob(v.defrobTag)
				err = v.inner
			} else {
				err = r.(error)
			}
			defer resource.Close()
		}
	}()

	for {
		resource, err = opener()

		if err != nil {
			if _, ok := err.(TransientError); ok {
				continue
			} else {
				return
			}
		}

		break
	}

	resource.Frob(input)
	defer resource.Close()

	return
}
