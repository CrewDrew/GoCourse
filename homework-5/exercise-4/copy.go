package main

import (
	"flag"
	"log"
	"os"
)

func main() {

	flag.Parse()

	arguments := flag.Args()

	if len(arguments) != 2 {
		log.Fatalln("Не верно заданы параметры. Синтаксис: copy filesourcename filerecievername")
	}

	sourceFileName := arguments[0]
	receiverFileName := arguments[1]

	if sourceFileName == receiverFileName {
		log.Fatalln("Файл источник равен файлу приёмнику")
	}

	file, err := os.Open(sourceFileName)
	if err != nil {
		log.Fatalln("Не удалось прочитать исходный файл", sourceFileName)
	}
	defer file.Close()

	receiver, err := os.Create(receiverFileName)

	if err != nil {
		log.Fatalln("Не удалось создать файл приемник")
	}
	defer receiver.Close()

	bs := make([]byte, 1024)

	for {
		n, _ := file.Read(bs)
		if n == 0 { // если конец файла
			break // выходим из цикла
		}
		receiver.Write(bs[:n])
	}
}
