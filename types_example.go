package main

import "fmt"

// Примеры типов данных в Go

func typesExamples() {
	// ========== БАЗОВЫЕ ТИПЫ ==========
	
	// Целые числа
	var i int = 42           // int (зависит от платформы: 32 или 64 бита)
	var i8 int8 = 127        // -128 до 127
	var i16 int16 = 32767    // -32768 до 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	
	// Беззнаковые целые числа
	var u uint = 42          // uint (зависит от платформы)
	var u8 uint8 = 255       // 0 до 255 (синоним: byte)
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	
	// byte - синоним uint8
	var b byte = 255
	
	// rune - синоним int32, используется для Unicode символов
	var r rune = 'A'
	var r2 rune = 'Я'
	
	// Числа с плавающей точкой
	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793
	
	// Комплексные числа
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 1 + 2i
	
	// Булевы значения
	var isTrue bool = true
	var isFalse bool = false
	
	// Строки
	var str string = "Привет, Go!"
	var str2 string = `Многострочная
	строка`
	
	fmt.Println("Целые числа:", i, i8, i16, i32, i64)
	fmt.Println("Беззнаковые:", u, u8, u16, u32, u64)
	fmt.Println("byte:", b, "rune:", r, r2)
	fmt.Println("float:", f32, f64)
	fmt.Println("complex:", c64, c128)
	fmt.Println("bool:", isTrue, isFalse)
	fmt.Println("string:", str)
	fmt.Println("многострочная строка:", str2)
	
	// ========== СОСТАВНЫЕ ТИПЫ ==========
	
	// Массивы - фиксированного размера
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	var arr2 = [...]int{1, 2, 3}  // размер определяется автоматически
	
	// Срезы - динамические массивы
	var slice []int = []int{1, 2, 3, 4, 5}
	var slice2 = make([]int, 5)      // создание среза длиной 5
	var slice3 = make([]int, 5, 10)  // длина 5, емкость 10
	
	// Структуры
	type Person struct {
		Name string
		Age  int
	}
	var person Person = Person{Name: "Иван", Age: 30}
	var person2 = Person{"Мария", 25}  // без именованных полей
	
	// Указатели
	var x int = 42
	var ptr *int = &x  // указатель на x
	var value int = *ptr  // разыменование указателя
	
	// Карты (map)
	var m map[string]int = make(map[string]int)
	m["один"] = 1
	m["два"] = 2
	
	// Или с инициализацией
	var m2 = map[string]int{
		"один": 1,
		"два":  2,
	}
	
	// Каналы
	var ch chan int = make(chan int)
	var chBuffered chan int = make(chan int, 10)  // буферизованный канал
	
	// Интерфейсы
	type Writer interface {
		Write([]byte) (int, error)
	}
	
	// Функции как тип
	type MathFunc func(int, int) int
	var add MathFunc = func(a int, b int) int {
		return a + b
	}
	
	fmt.Println("Массив:", arr, arr2)
	fmt.Println("Срез:", slice, slice2, slice3)
	fmt.Println("Структура:", person, person2)
	fmt.Println("Указатель:", ptr, "значение:", value)
	fmt.Println("Карта:", m, m2)
	fmt.Println("Функция add:", add(5, 3))
	fmt.Println("Канал:", ch, chBuffered)
	
	// ========== НУЛЕВЫЕ ЗНАЧЕНИЯ ==========
	var zeroInt int           // 0
	var zeroFloat float64     // 0.0
	var zeroBool bool         // false
	var zeroString string     // ""
	var zeroSlice []int       // nil
	var zeroMap map[string]int // nil
	var zeroPtr *int          // nil
	var zeroChan chan int     // nil
	var zeroInterface interface{} // nil
	
	fmt.Println("Нулевые значения:")
	fmt.Println("int:", zeroInt, "float:", zeroFloat, "bool:", zeroBool)
	fmt.Println("string:", zeroString, "slice:", zeroSlice)
	fmt.Println("map:", zeroMap, "ptr:", zeroPtr, "chan:", zeroChan, "interface:", zeroInterface)
	
	// ========== ПРЕОБРАЗОВАНИЕ ТИПОВ ==========
	var intVal int = 42
	var floatVal float64 = float64(intVal)  // явное преобразование
	var intVal2 int = int(floatVal)
	
	// Преобразование строк
	var numStr string = "42"
	// Для преобразования строк используются функции из пакетов strconv, fmt и т.д.
	
	fmt.Println("Преобразование:", intVal, floatVal, intVal2)
	fmt.Println("Строка числа:", numStr)
	
	// ========== ТИПЫ ВЫВОДА (type inference) ==========
	// Go может автоматически определить тип
	var autoInt = 42           // int
	var autoFloat = 3.14       // float64
	var autoString = "hello"   // string
	var autoBool = true        // bool
	
	// Короткое объявление переменной
	shortInt := 100
	shortString := "Go"
	
	fmt.Println("Автоопределение:", autoInt, autoFloat, autoString, autoBool)
	fmt.Println("Короткое объявление:", shortInt, shortString)
}
