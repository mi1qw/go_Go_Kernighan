package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func findInEnv[T any](name string, def T) T {
	val := os.Getenv(name)
	if val == "" {
		return def
	}

	result := reflect.New(reflect.TypeOf(def)).Interface()
	convErr := fmt.Errorf("cannot convert '%s' to type %T", val, def)

	switch v := result.(type) {
	case *string:
		*v = val
		return reflect.ValueOf(*v).Interface().(T)
	case *int:
		i := result.(*int)
		*i, _ = strconv.Atoi(val)
		// или
		//fmt.Sscanf(val, "%d", v)
		return reflect.ValueOf(*i).Interface().(T)
	default:
		panic(fmt.Errorf("unsupported type: %T", def))
	}
	/*
	   result := reflect.New(reflect.TypeOf(def)).Interface() - Здесь мы создаем переменную result, которая имеет тип interface{} и содержит указатель на новое значение, созданное с использованием типа def. Это позволяет нам динамически создать значение нужного типа.

	   switch v := result.(type) - Здесь мы используем switch с type assertion, чтобы определить конкретный тип v, к которому приводится result.

	   case *string: - Если тип def является указателем на string, то мы утверждаем тип result как указатель на string, и в этом случае записываем значение val в этот указатель, чтобы выполнить преобразование строки окружения в строку.

	   case *int: - Если тип def является указателем на int, то мы утверждаем тип result как указатель на int, и в этом случае используем strconv.Atoi для преобразования строки окружения val в целое число и записываем это значение в указатель на int.

	   return reflect.ValueOf(*v).Interface().(T) - Здесь мы используем reflect.ValueOf(*v) для создания reflect.Value из указателя v, затем Interface() для преобразования его обратно в интерфейсное значение interface{}, и, наконец, используем (T) для приведения этого интерфейсного значения к типу T и возвращаем его.
	*/
	// или
	if strVal, ok := result.(*string); ok {
		*strVal = val
		return reflect.ValueOf(*strVal).Interface().(T)
	}

	if intVal, ok := result.(*int); ok {
		_, err := fmt.Sscanf(val, "%d", intVal)
		if err != nil {
			panic(convErr)
		}
		return reflect.ValueOf(*intVal).Interface().(T)
	}
	/*
	   if strVal, ok := result.(*string); ok - В этой строке мы пытаемся утвердить тип result как указатель на string (*string). Если result является указателем на string, то ok будет равно true, и strVal будет указывать на это значение типа string.

	   *strVal = val - Если тип result был успешно утвержден как *string, то мы присваиваем значение val этому указателю, чтобы поместить в него значение из окружения.

	   return reflect.ValueOf(*strVal).Interface().(T) - Здесь мы используем reflect.ValueOf(*strVal) для создания reflect.Value из *strVal, затем Interface() для преобразования его обратно в интерфейсное значение interface{}. Наконец, мы используем (T) для приведения этого интерфейсного значения к типу T, который определен через аргумент функции.
	*/
	// Добавьте обработку других типов, если необходимо.

	panic(convErr)
}

func main() {
	name := findInEnv("HOME", "default")
	fmt.Printf("Name: %s\n", name)

	//count := findInEnv("SHLVL", 42)
	//count := findInEnv("QT_ACCESSIBILITY", 42)
	count := findInEnv("SSH_AGENT_PID", 42)
	fmt.Printf("Count: %d\n", count)
}
