package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type KeyType interface {
	int | string
}

type Entry[K KeyType, V comparable] struct {
	key K
	value V
}

type Map[K KeyType, V comparable] struct {
	entries []Entry[K, V]
}

func addEntry[K KeyType, V comparable](m *Map[K, V], key K, value V) {
	m.entries = append(m.entries, Entry[K, V]{key, value})
}

func getEntry[K KeyType, V comparable](m Map[K, V], key K) {
	for i := 0; i < len(m.entries); i++ {
		if (m.entries[i].key == key) {
			fmt.Print(m.entries[i].value)
			return
		}
	}
	fmt.Println("Key does not exist")
}

func mapSize[K KeyType, V comparable](m Map[K, V]) int {
	return len(m.entries)
}

func printMap[K KeyType, V comparable](m Map[K, V]) {
	mapsize := mapSize(m)
	if mapsize != 0 {
		for _, e := range m.entries {
			fmt.Printf("[%v] %v", e.key, e.value)
		}
	} else {
		fmt.Println("[]")
	}
	fmt.Println(mapsize)
}

func main() {
	var m Map[int, string]	
	inputReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\n> ")
		input, _ := inputReader.ReadString('\n')
		operation := strings.Split(input, " ")

		switch strings.TrimSpace(operation[0]) {
		case "add":
			if len(operation[1:]) == 2 {
				key, _ := strconv.Atoi(strings.TrimSpace(operation[1]))
				value := operation[2]
				addEntry(&m, key, value)
				printMap(m)
			} else {
				fmt.Println("Invalid number of arguments for 'add'")
			}
		case "get":
			if len(operation[1:]) == 1 {
				key, _ := strconv.Atoi(strings.TrimSpace(operation[1]))
				getEntry(m, key)
			} else {
				fmt.Println("Invalid number of arguments for 'get'")
			}
		case "size":
			mapSize(m)
		case "print":
			printMap(m)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Operation not supported")
		}
	}
}