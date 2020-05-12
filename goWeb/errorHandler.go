package goWeb

func ErrorHandle(err error, info string)  {
	if err != nil {
		panic("ERROR: " + info + " " + err.Error())
	}
}
