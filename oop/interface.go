package main

type Writer interface {
	Write([]byte) (int, error)
}

type FileWriter struct {
	Filename string
}

func (f FileWriter) Write(data []byte) (int, error) {
	return len(data), nil
}
