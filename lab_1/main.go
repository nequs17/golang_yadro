package main

import (
	"fmt"
	"os"
)

func main() {
	var num1, num2, result float64
	var operator string

	for {
		fmt.Print("Введите первое число: ")
		if _, err := fmt.Scanln(&num1); err != nil {
			fmt.Println("Ошибка: введите корректное число.")
			fmt.Println(err)
			os.Exit(1)
		}

		for {
			fmt.Print("Выберите операцию (+, -, *, /) или введите 'q' для завершения: ")
			fmt.Scanln(&operator)

			if operator == "q" {
				fmt.Println("Программа завершена.")
				return
			}

			fmt.Print("Введите второе число: ")
			if _, err := fmt.Scanln(&num2); err != nil {
				fmt.Println("Ошибка: введите корректное число.")
				fmt.Println(err)
				os.Exit(1)
			}

			switch operator {
			case "+":
				result = num1 + num2
			case "-":
				result = num1 - num2
			case "*":
				result = num1 * num2
			case "/":
				if num2 == 0 {
					fmt.Println("Ошибка: деление на ноль невозможно.")
					continue
				}
				result = num1 / num2
			default:
				fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
				continue
			}

			fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
			num1 = result // Присваиваем результат предыдущей операции первому числу для возможности продолжения вычислений
		}
	}
}
