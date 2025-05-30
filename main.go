package main

/*
 Система рейтингу гравців

 Це программа для керування системої рейтингу а саме:
 - гравцями
 - перемогами чи поразками гравців
 - рахунками гравців(балами)

 Функціонал программи:
 - Управління гравцями:
	-Реєстрація нових гравців
	- Видалення гравців із системи
	- Пошук гравця за ніком
	- Відображення всіх гравців
 - Управління рейтингом:
	- Оновлення рахунку гравця після матчу
	- Перегляд топ-10 найкращих гравців
	- Розрахунок середнього рейтингу
	- Пошук гравців за діапазоном рейтингу
 - Статистика:
	- Підрахунок кількості матчів гравця
	- Визначення найкращого та найгіршого гравця
	- Статистика переможців і програвших
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// глобальні змінні
var players []string
var scores = make(map[string]int)
var matchesPlayed = make(map[string]int)
var wins = make(map[string]int)
var losses = make(map[string]int)

// читачь строки
var reader *bufio.Reader = bufio.NewReader(os.Stdin)

// функція сортировки
func sortPlayer(arr []string, start, end int) {
	if start < end {
		pivotIndex := partition(arr, start, end)
		sortPlayer(arr, start, pivotIndex-1)
		sortPlayer(arr, pivotIndex+1, end)
	}
}

func partition(arr []string, start, end int) int {
	pivot := scores[arr[end]]
	i := start - 1

	for j := start; j < end; j++ {
		if scores[arr[j]] > pivot {
			i++
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	temp := arr[i+1]
	arr[i+1] = arr[end]
	arr[end] = temp

	return i + 1
}

// функція для зручного застосування сортировки
func callSortPlayer() {
	sortPlayer(players, 0, len(players)-1)
}

// Реєструє нового гравця з початковим рейтингом 1000
func registerPlayer(nickname string) bool {
	//перевірка на наявність гравця
	if !playerExists(nickname) {
		//додавання нового гравця
		players = append(players, nickname)
		scores[nickname] = 1000
	} else if playerExists(nickname) {
		return false
	}
	return true
}

// Видаляє гравця з системи
func removePlayer(nickname string) bool {
	//перевірка на наявність гравця
	if playerExists(nickname) {
		//видалення гравця
		delete(scores, nickname)
		index := findPlayerIndex(nickname)
		players = append(players[:index], players[index+1:]...)
	} else if !playerExists(nickname) {
		return false
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
	// викликає функцію сортування
	callSortPlayer()
	for i, nick := range players {
		fmt.Printf("%d. %s Рейтинг - %d | Матчів: %v | Перемог: %v | Поразок: %v\n", i+1, nick, scores[nick], matchesPlayed[nick], wins[nick], losses[nick])
	}
}

// Оновлює рейтинг після матчу (true = перемога, false = поразка)
func updateRating(nickname string, won bool, pointsChange int) {
	// перевірка на наявність гравця
	if playerExists(nickname) {
		//оновлює рейтинг згідно результату гри (перемоги чи поразки)
		if won == true {
			wins[nickname] += 1
			scores[nickname] += pointsChange
			matchesPlayed[nickname] += 1
		} else if won == false {
			losses[nickname] += 1
			scores[nickname] -= pointsChange
			matchesPlayed[nickname] += 1
		}
		fmt.Printf("\nРейтинг гравця \"%s\" оновлено! Новий рейтинг: %d ✅\n", nickname, scores[nickname])
	} else if !playerExists(nickname) {
		fmt.Printf("\nНажаль невдалося оновити рейтинг гравця \"%s\".Такого гравця не існує! ❌", nickname)
	}
}

// Повертає топ-10 гравців за рейтингом
func getTopPlayers(count int) []string {
	//викликає функцию сортування
	callSortPlayer()
	//слайс для зберегання топ ігроків
	var topPl []string
	//додавання та виведення топу ігроків
	for i := 0; i < count; i++ {
		topPl = append(topPl, players[i])
		fmt.Printf("%d. %s - %d\n", i+1, players[i], scores[players[i]])
	}
	return topPl
}

// Знаходить гравців у діапазоні рейтингу
func findPlayersByRatingRange(minRating, maxRating int) []string {
	//викликає функцию сортування
	callSortPlayer()
	//слайс для зберегання ігроків за діапозонном
	var rangePl []string
	//додавання та вивід гравців у діапазоні
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
	//змінна для зберегання середнього рейтингу ігроків
	var avg float64
	//розрахунок середнього рейтинга
	for i, _ := range scores {
		avg += float64(scores[i])
	}
	avg = avg / float64(len(scores))
	return avg
}

// Знаходить гравця з найвищим рейтингом
func getBestPlayer() string {
	//змінна для зберегання ніку кращого ігрока
	var bestPl string
	//змінна для зберегання рейтинга кращого ігрока
	var bestScore int
	//визначення кращого гравця
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
	//змінна для зберегання ніку гіршого ігрока
	var worstPl string
	//змінна для зберегання рейтинга гіршого ігрока
	var worstScore int
	//визначення гіршого гравця
	for nick, score := range scores {
		if score < worstScore || worstPl == "" {
			worstPl = nick
			worstScore = score
		}
	}
	return worstPl
}

// Розраховує відсоток перемог гравця
func calculateWinRate(nickname string) float64 {
	//змінна для зберегання відсотку перемог гравця
	var winRate float64
	//змінна зберегає зіграні матчі гравцем
	totalmatches := matchesPlayed[nickname]

	//визначення вінрейту
	if playerExists(nickname) && totalmatches != 0 {
		winRate = float64(wins[nickname]) / float64(matchesPlayed[nickname]) * 100
	} else if !playerExists(nickname) {
		fmt.Println("Такого гравця не існує! ❌")
	}
	return winRate
}

// Відображає детальну статистику гравця
func displayPlayerStats(nickname string) {
	//перевірка на наявність гравця та виведення статистики
	if playerExists(nickname) {
		fmt.Printf("\n📋 Статистика гравця \"%s\":\n", nickname)
		fmt.Printf("🏅 Поточний рейтинг: %d\n", scores[nickname])
		fmt.Printf("🎮 Зіграно матчів: %d\n", matchesPlayed[nickname])
		fmt.Printf("✅ Перемоги: %d\n", wins[nickname])
		fmt.Printf("❌ Поразки: %d\n", losses[nickname])
		fmt.Printf("📊 Відсоток перемог: %.1f%%\n", calculateWinRate(nickname))
	} else if !playerExists(nickname) {
		fmt.Println("\nТакого гравця не існує! ❌")
	}
}

// Показує загальну статистику системи
func displaySystemStats() {
	//змінни для зберегання загальну кількість матчів, загальних перемог, загальних поразок
	totalmatches := 0
	totalWins := 0
	totalLosses := 0

	//визначення загальну кількість матчів, загальних перемог, загальних поразок
	for _, nick := range players {
		totalmatches += matchesPlayed[nick]
		totalWins += wins[nick]
		totalLosses += losses[nick]
	}

	//виведення загальної статистики
	fmt.Printf("👥 Кількість Гравців: %d\n", len(players))
	fmt.Printf("🎮 Загальна кількість матчів:  %d\n", totalmatches)
	fmt.Printf("✅ Загальна кількість перемог: %d\n", totalWins)
	fmt.Printf("❌ Загальна кількість поразок:  %d\n", totalLosses)
	fmt.Printf("📊 Середній рейтинг: %.2f\n", calculateAverageRating())
	fmt.Printf("🏆 Найкращий гравець:  %s\n", getBestPlayer())
	fmt.Printf("🐢 Найгірший гравець: %s\n", getWorstPlayer())
}

// Відображає меню
func displayMenu() {
	fmt.Println("1. Зареєструвати гравця ✅")
	fmt.Println("2. Видалити гравця ❌")
	fmt.Println("3. Оновити рейтинг після матчу 🔄")
	fmt.Println("4. Знайти гравця 🔍")
	fmt.Println("5. Показати всіх гравців 👥")
	fmt.Println("6. Топ-10 гравців 🏆")
	fmt.Println("7. Пошук за діапазоном рейтингу 🔍")
	fmt.Println("8. Статистика гравця 📊")
	fmt.Println("9. Загальна статистика 📊")
	fmt.Println("10. Вихід 👋")
}

// Отримує текстове введення
func getStringInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	input, _ = reader.ReadString('\n')
	return input
}

// Отримує числове введення
func getIntInput(prompt string) int {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	//перевірка на пробіли
	if strings.ContainsAny(input, " \t") {
		fmt.Println("Некоректне введення. Ведіть ціле число без пробілів. ❌")
		return -1
	}

	//перевірка на введеня числа
	var value int
	_, err := fmt.Sscanf(input, "%d", &value)
	if err != nil {
		fmt.Println("Некоректне введення. Ведіть ціле число. ❌")
		return -1
	}
	return value
}

// ОПЦІЇ

// опція реєстрація гравця
func choiseRegPlayer() {
	fmt.Println("\n--- Реєстрація гравця --- 🎉")

	// отримання нікнейма
	input := getStringInput("Введіть нікнейм гравця: ")
	input = strings.TrimSpace(input)

	//вивід повідомлення про реестрацію
	if registerPlayer(input) {
		fmt.Printf("\nГравець \"%s\" успішно зареєстрований з рейтингом 1000! ✅\n", input)
	} else {
		fmt.Printf("\nНажаль гравця \"%s\" не зареестровано.Тому що гравець с таким ніком вже існує! ❌\n", input)
	}
}

// опція видалення гравця
func choiseDeletePlayer() {
	fmt.Println("\n--- Видалення гравця --- ❌")

	//отримання ніку
	input := getStringInput("Введіть нікнейм гравця: ")
	input = strings.TrimSpace(input)

	//вивід повідомлення про видалення
	if removePlayer(input) {
		fmt.Printf("\nГравець \"%s\" успішно видален! ✅\n", input)
	} else {
		fmt.Printf("\nНажаль гравця \"%s\" не видалено.Тому що такого гравця не існує! ❌\n", input)
	}
}

// опція оновлення рейтингу гравця
func choiseUpdateRating() {
	fmt.Println("\n--- Оновлення рейтингу --- 🔄")

	//отримання ніку, перемоги чи поразки, зміни рейтингу
	nick := getStringInput("Введіть нікнейм гравця: ")
	winOrLose := getIntInput("Гравець переміг чи програв? (1 - перемога 🏆, 0 - поразка ❌): ")
	ratingChange := getIntInput("Введіть зміну рейтингу: ")

	nick = strings.TrimSpace(nick)

	//вивід повідомлення про оновлення рейтингу
	if winOrLose == 1 {
		updateRating(nick, true, ratingChange)
	} else if winOrLose == 0 {
		updateRating(nick, false, ratingChange)
	} else {
		fmt.Println("Невірний вибір результату гри.Оновлення скасовано.❌")
	}
}

// опція знаходження гравця
func choiseFindPlayer() {
	fmt.Println("\n--- Знаходження Гравця --- 🔍")

	//отримання ніку
	input := getStringInput("Введіть нікнейм гравця: ")
	input = strings.TrimSpace(input)

	//вивід повідомлення про знаходження
	if playerExists(input) {
		fmt.Printf("\nГравець \"%s\" знайден. ✅\n", input)
	} else if !playerExists(input) {
		fmt.Printf("\nГравець \"%s\" не знайден. ❌\n", input)
	}
}

// опція для виводу всіх гравців
func choiseShowAllPlayer() {
	fmt.Println("\n--- Всі гравці --- 👥")
	//вивід всіх гравців
	displayAllPlayers()
}

// опція для виводу топ 10 гравців
func choiseTopPlayers() {
	fmt.Println("\n--- Топ-10 гравців --- 🏆")

	//отримання числа для відображення топу гравців
	input := getIntInput("Введіть кількість гравців для відображення: ")

	//вивід повідомлення про топ
	if input > 0 && input <= 10 {
		fmt.Println("\nТОП ГРАВЦІ:")
		getTopPlayers(input)
	} else {
		fmt.Println("Не вдалося показати гравців: кількість має бути від 1 до 10. ❌")
	}
}

// опція пошук за діапозоном рейтингу
func choiseFindByRatingRange() {
	fmt.Println("\n--- Пошук за діапазоном рейтингу --- 📈")

	//отримання мінімального значення
	minRating := getIntInput("Введіть мінімальне значення: ")
	//перевірка на коректність введення
	if minRating == -1 {
		fmt.Println("Пошук за діапазоном припинено - мінімальне значення введено некоректно. ❌")
		return
	}

	//отримання максимального значення
	maxRating := getIntInput("Введіть максимальне значення: ")
	//перевірка на коректність введення
	if maxRating == -1 {
		fmt.Println("Пошук за діапазоном припинено - максимальне значення введено некоректно. ❌")
		return
	}

	//вивід гравців за діапозоном
	fmt.Println("\nГравці за діапазоном рейтингу:")
	findPlayersByRatingRange(minRating, maxRating)
}

// опція статистика гравця
func choisePlayerStats() {
	fmt.Println("\n--- Статистика гравця --- 📊")

	//отримання нікнейма
	input := getStringInput("Введіть нікнейм гравця: ")
	input = strings.TrimSpace(input)

	//вивід статистики гравця
	displayPlayerStats(input)
}

// опція загальна статистика
func choiseGlobalStats() {
	fmt.Println("\n---Загальна Статистика --- 🌍📊")
	//вивід загальної статистики
	displaySystemStats()
}

func main() {
	fmt.Println("=== Система рейтингу гравців === 💻")
	//вивід меню
	displayMenu()
	//цикл програми
	for {
		//отримання вибору
		choise := getIntInput("> ")

		//отримання опцію за вибором користувачя
		switch choise {
		case 1:
			choiseRegPlayer()
		case 2:
			choiseDeletePlayer()
		case 3:
			choiseUpdateRating()
		case 4:
			choiseFindPlayer()
		case 5:
			choiseShowAllPlayer()
		case 6:
			choiseTopPlayers()
		case 7:
			choiseFindByRatingRange()
		case 8:
			choisePlayerStats()
		case 9:
			choiseGlobalStats()
		case 10:
			fmt.Println("\nДо побачення! 👋")
			return
		default:
			fmt.Println("\nТакого вибору не існує. ❌")
		}
	}
}
