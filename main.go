package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func getNumberInput() (int, error) {
		var input string
		fmt.Print("Введите число: ")
		fmt.Scanln(&input)

		num, err := strconv.Atoi(input)
		if err != nil {
			return 0, fmt.Errorf("Ошибка: требуется ввести целое число")
		}

		return num, nil
	}

	func getHint(number, guess int) string {
					diff := number - guess
			if diff < 0 {
			diff = -diff
			}
			switch {
			case diff <= 5:
				return "🔥 Горячо"
			case diff <= 15:
				return "🙂 Тепло"
			default:
				return "🥶 Холодно"
			}
		
	}
func guess() {
	fmt.Println("Играй угадай числа от 1 до 100")
	fmt.Println("Угадайте число за 10 попыток!")
	var number int
	var maxNumber int
	rand.Seed(time.Now().UnixNano())
	number = rand.Intn(100) + 1
	var tries int
	var difficulty int
	fmt.Println("Выберите уровень сложности:")
	fmt.Println("1 - Easy (1–50, 15 попыток)")
	fmt.Println("2 - Medium (1–100, 10 попыток)")
	fmt.Println("3 - Hard (1–200, 5 попыток)")
	fmt.Scanln(&difficulty)
			switch difficulty{
		case 1:
			maxNumber = 50
			number = rand.Intn(maxNumber) + 1
			tries = 15
		case 2:
			maxNumber = 100
			number = rand.Intn(maxNumber) + 1
			tries = 10
		case 3:
			maxNumber = 200
			number = rand.Intn(maxNumber) + 1
			tries = 5
		default:
			fmt.Errorf("Ошибка: требуется выбрать сложность")
			return
		}

	for tries > 0 {

		guess, err := getNumberInput()

		if guess < 1 || guess > maxNumber {
			fmt.Printf("Число должно быть от 1 до %d\n", maxNumber)
			continue
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		if guess > number {
			fmt.Println("Секретное число меньше👇")
			fmt.Println(getHint(number, guess))
			tries--
		} else if guess < number {
			fmt.Println("Секретное число больше👆")
			fmt.Println(getHint(number, guess))
			tries--
		} else {
			color.Green("Позравляем! Вы угадали число! Но, так как эта игра просто тест, награды не предусмотрено.")
			break
		}
		if tries == 0 {
			color.Red("К сожалению, у Вас закончились попытки. Вы проиграли. Очень жаль.")
			break
		} else if guess > 100 || guess < 1 {
			fmt.Println("Вы ввели некорректное число. Попробуйте еще раз.")
		} 
	}
}

func main() {
	for {
		guess()
	var again string
	fmt.Print("Хотите сыграть еще? (y/n): ")
		fmt.Scanln(&again)
		if again == "y" || again == "Y" || again == "д" || again == "Д" {
			main()
		} else if again == "n" || again == "н" || again == "Н" || again == "N" {
			fmt.Println("Спасибо за игру!")
			break
		} else if again != "y" && again != "Y" && again != "н" && again != "Н" {
			fmt.Println("Игра завершена.Спасибо за игру.")
		}
	}
}