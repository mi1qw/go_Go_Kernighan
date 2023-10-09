package main

import (
	"fmt"
	"reflect"
)

type Class struct {
	name string
}
type StructFun[T, V comparable] struct {
	c    Class
	t    *StructFun[T, V]
	name string
	f    func(T) V
}

func main() {
	// стр. 385
	t := reflect.TypeOf(3)
	fmt.Println(t.String(), t.Name())
	fmt.Println(t)
	println()

	v := reflect.ValueOf(3)     // reflect.Value
	fmt.Println(v)              // "3"
	fmt.Printf("%v\n", v)       // "З"
	fmt.Println(v.String())     // Примечание: "<int Value>"
	println(v.Kind(), "Kind()") // число - базовый тип для reflect.Int, reflect.String и остальные
	fmt.Printf("%v \n\n", v.Type())

	i := v.Interface() // возвращает any
	i2 := i.(int)      // приводим к типу int, надо явно знать тип, неудобно(
	fmt.Printf("%v  \n\n", i2+1)

	sf := StructFun[Class, string]{
		c:    Class{"nameClass"},
		name: "name",
		f: func(t Class) string {
			return t.name
		},
	}
	//println(sf.sf(1))
	fr := reflect.TypeOf(sf)
	println(fr.String())
	fmt.Printf("%d - Struct\n", fr.Kind())

	v = reflect.ValueOf(sf)
	println(v.String())
	//pointer := v.Pointer()
	//pointer

	println(sf.saveType(Class{"nameClass!!!"}))
}
func (sf *StructFun[T, V]) saveType(arg T) reflect.Type {
	typeOf := reflect.TypeOf(arg)
	println(typeOf.Name())
	println(typeOf.String())

	newT := reflect.New(typeOf)
	println(newT.Type().String())

	if newT.Kind() == reflect.Ptr {
		// Создаем значение, которое вы хотите присвоить (здесь пустая структура)
		//valueToAssign := reflect.Zero(typeOf.Elem())
		//fmt.Println("New Value:", newValue.Interface())

		// Присваиваем значение newT
		//newT.Elem().Set(valueToAssign)
		//fmt.Println("New Value:", newT.Elem())
	} else {
		fmt.Println("Type is not a pointer.")
	}

	return typeOf
}

//println(newT.CanSet())
//newT.Set(reflect.ValueOf(arg))

//func (sf *StructFun[T, V]) saveType(arg T) V{
//	reflect.TypeOf(arg)
//}
