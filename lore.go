package main

import (
	"bufio"
)

func Lore(writerPtr *bufio.Writer) {
	writer := *writerPtr
	writer.WriteString("Hello, welcome to the Cafe of Rest, the best local coffee shop in all of Sarn!\n")
	writer.WriteString("I am Clarissa, I operate this gem of hope, which I named in inspiration after my lost love.\n")
	writer.WriteString("I cling onto the hope of what our Eternal Empire once represented, and what it may still return to;\nSomething beautiful and majestic, in the face of dark magic and Thaulmaturgy.\n")
	writer.WriteString("Please enjoy from our selection of deliciously brewed espresso beans,\ncaptured by our resident Exile's from the corrupted monsters in the Riverways and near the Forrested areas.\n")
	writer.Flush()
}
