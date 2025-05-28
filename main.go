package main

import (
	"fmt"
)

var players []string
var scores = make(map[string]int)
var matchesPlayed = make(map[string]int)
var wins = make(map[string]int)
var losses = make(map[string]int)

// Реєструє нового гравця з початковим рейтингом 1000
func registerPlayer(nickname string) bool {
	players = append(players, nickname)
	scores[nickname] = 1000
	return true
}

// Видаляє гравця з системи
func removePlayer(nickname string) bool {
	if playerExists(nickname) {
		delete(scores, nickname)
		index := findPlayerIndex(nickname)
		players = append(players[:index], players[index+1:]...)
	} else if !playerExists(nickname) {
		fmt.Println("Такого гравця не існує!")
	}
	return true
}

// Знаходить індекс гравця в слайсі
func findPlayerIndex(nickname string) int {
	for i, name := range players {
		if name == nickname {
			return i
		}
	}
	return -1
}

// Перевіряє чи існує гравець
func playerExists(nickname string) bool {
	for _, name := range players {
		if name == nickname {
			return true
		}
	}
	return false
}

// Відображає всіх гравців з їх статистикою
func displayAllPlayers() {
	for i, _ := range players {
		fmt.Printf("%d. %v\n", i+1, players[i])
	}
}

// Відображає меню
func displayMenu() {
	fmt.Println("1. Зареєструвати гравця")
	fmt.Println("2. Видалити гравця")
	fmt.Println("3. Оновити рейтинг після матчу")
	fmt.Println("4. Знайти гравця")
	fmt.Println("5. Показати всіх гравців")
	fmt.Println("6. Топ-10 гравців")
	fmt.Println("7. Пошук за діапазоном рейтингу")
	fmt.Println("8. Статистика гравця")
	fmt.Println("9. Загальна статистика")
	fmt.Println("10. Вихід")
}

// Отримує текстове введення
func getStringInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}

// Отримує числове введення
func getIntInput(prompt string) int {
	var input int
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}

func main() {

}
