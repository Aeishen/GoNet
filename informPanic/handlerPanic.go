package informPanic

func Deal(f func()) func(){
	return func() {
		defer func() {
			if x := recover(); x!= nil{
				// 通知xxx
				Inform()
			}
		}()
		f()
	}
}
