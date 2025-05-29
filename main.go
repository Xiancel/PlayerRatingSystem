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
	//сделать функцию сортировки и добавть сюда
	for i, _ := range players {
		fmt.Printf("%d. %v\n", i+1, players[i])
	}
}

// Оновлює рейтинг після матчу (true = перемога, false = поразка)
func updateRating(nickname string, won bool, pointsChange int) {
	if playerExists(nickname) {
		if won == true {
			wins[nickname] = 1
			scores[nickname] += pointsChange
		} else if won == false {
			losses[nickname] = 1
			scores[nickname] -= pointsChange
		}
	} else if !playerExists(nickname) {
		fmt.Println("Такого гравця не існує!")
	}
	fmt.Printf("Рейтинг гравця \"%s\" оновлено! Новий рейтинг: %d\n", nickname, scores[nickname])
}

// Повертає топ-10 гравців за рейтингом
func getTopPlayers(count int) []string {
	//сделать функцию сортировки и добавть сюда
	var topPl []string
	for i := 0; i < count; i++ {
		topPl = append(topPl, players[i])
		fmt.Printf("%d. %s - %d\n", i+1, players[i], scores[players[i]])
	}
	return topPl
}

// Знаходить гравців у діапазоні рейтингу
func findPlayersByRatingRange(minRating, maxRating int) []string {
	var rangePl []string
	for i, _ := range players {
		rating := scores[players[i]]
		if rating >= minRating && rating <= maxRating {
			rangePl = append(rangePl, players[i])
			fmt.Printf("%d. %s - %d\n", i+1, players[i], rating)
		}
	}
	return rangePl
}

// Розраховує середній рейтинг всіх гравців
func calculateAverageRating() float64 {
	var avg float64
	for i, _ := range scores {
		avg += float64(scores[i])
	}
	avg = avg / float64(len(scores))
	return avg
}

// Знаходить гравця з найвищим рейтингом
func getBestPlayer() string {
	var bestPl string
	var bestScore int
	for nick, score := range scores {
		if score > bestScore || bestPl == "" {
			bestPl = nick
			bestScore = score
		}
	}
	return bestPl
}

// Знаходить гравця з найнижчим рейтингом
func getWorstPlayer() string {
	var worstPl string
	var worstScore int
	for nick, score := range scores {
		if score < worstScore || worstPl == "" {
			worstPl = nick
			worstScore = score
		}
	}
	return worstPl
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
