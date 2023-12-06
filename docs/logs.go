package docs

import (
	"bufio"
	"log"
	"os"
)

type Logs struct{
	name_file string
	file *os.File
}

func NewLogs(name_file string) *Logs{
	return &Logs{name_file:name_file}
}

func (l *Logs)GenerateLogs() {
	fileName := l.name_file
	var err error
	l.file, err = os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer l.file.Close()
}


func (l *Logs) WriteLog(message string){
	var err error
	l.file, err = os.OpenFile(l.file.Name(), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(l.file)
	_, err = writer.WriteString(message+"\n")
	if err != nil {
		log.Fatal(err)
	}
	writer.Flush()
}
