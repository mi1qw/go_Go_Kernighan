package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

func main() {
	// Устанавливаем соединение с X11 сервером.
	X, err := xgb.NewConn()
	if err != nil {
		log.Fatalf("Не удалось установить соединение с X11 сервером: %v", err)
	}
	defer X.Close()

	root := xproto.Setup(X).DefaultScreen(X).Root

	// Получаем идентификаторы всех окон на экране.
	//reply, err := xproto.QueryTree(X, xproto.Window(xproto.Setup(X).Root)).Reply()
	reply, err := xproto.QueryTree(X, xproto.Window(root)).Reply()
	if err != nil {
		log.Fatalf("Не удалось получить список окон: %v", err)
	}

	// Перебираем все окна и ищем окно по имени.
	for _, window := range reply.Children {
		w := window
		println(w)
		// Получаем имя окна.
		attributesReply, err := xproto.GetWindowAttributes(X, window).Reply()

		println(attributesReply)
		attributesReply.
		//name, err := xproto.GetWindowName(X, window).Reply()
		if err != nil {
			continue
		}

		// Сравниваем имя окна с искомым именем.
		//if name.Name != "имя программы" {
		//	continue
		//}

		// Выводим идентификатор окна.
		fmt.Printf("Идентификатор окна: %d\n", window)
		return
	}

	// Окно не найдено.
	fmt.Println("Окно не найдено.")
}
