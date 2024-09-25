package main

func main() {
	cards := deck{}

	cards.print()

	cards.generate()

	cards.print()

	cards.shuffle()

	cards.saveToFile("something")
	cards.openFile("something")

	cards.print()

}
