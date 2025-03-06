package contacts

import "fmt"

const Email string = "qwerfg@mail.com" // глобальная для экспорта

var support string // локальная не для экспорта

func SetSupport(str string) { //  для экспорта функция
	support = str
}

func GetContact() string { // экспортируемая функция
	return fmt.Sprintf("%s <%s>", support, Email)
}
