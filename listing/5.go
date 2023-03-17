package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test2() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test2()
	if err != nil {
		println("error") // напишет это, т.к. тип хранящего внутри обхекта не nil, a указатель (*customError)
		return
	}
	println("ok")
}

// Для того, чтобы получить ok,  test должен возвращать объекты типа интерфейс error
