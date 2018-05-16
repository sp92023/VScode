package main

//var card string 可執行

func main() {
	//card := "Ace of Spades"  等同var card string = "Ace of Spades"
	//card = "Five of Diamonds"
	/*card := newCard()
	fmt.Println(card)
	fmt.Println(getTitle())*/
	//cards := []string{"Ace of Diamonds", newCard()} //slice 可改變長度的array
	//一個slice只能存一種type 不能int和string同時使用
	/*cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")*/

	/*for i, card := range cards {
		i = i + 1
		fmt.Println(card)
	}*/
	/*cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()*/
	/*greeting := "Hi there!"
	fmt.Println([]byte(greeting))*/
	/*cards := newDeck()
	//fmt.Println(cards.toString())
	cards.saveToFile("my_cards")*/
	/*cards := newDeckFromFile("my_cards")
	cards.print()*/
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

/*func newCard() string {
	return "Five of Diamonds"
}

func newCard() int {
	return 12
}

func getTitle() string {
	return "Harry"
}*/
