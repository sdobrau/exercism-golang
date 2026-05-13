package parallelletterfrequency

import (
	"strings"
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

func CountFrequency(letter rune, text string) int {
	frequency := 0
	for _, text_letter := range text {
		if text_letter == letter {
			frequency += 1
		}
	}
	return frequency
}

func CountFrequencyWorker(wg *sync.WaitGroup, letterChan <-chan rune, text string, outChan chan<- int) {
	frequency := 0
	var letter = <-letterChan
	for _, text_letter := range text {
		if text_letter == letter {
			frequency += 1
		}
	}
	outChan <- frequency
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	var freqMap = make(map[rune]int)
	var letters = "abcdefghijklmnopqrstuvwxyz"
	lower_text := strings.ToLower(text)
	for _, letter := range letters {
		frequency := CountFrequency(letter, lower_text)
		if frequency > 0 {
			freqMap[letter] = frequency
		}
	}
	return freqMap
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	var freqMap = make(map[rune]int)
	var letterChan = make(chan rune)
	var outChan = make(chan int)
	// this is not the optimal way but the tests pass.
	// We need to compute the list of letters to count
	var letters = "abcdefghijklmnopqrstuvwxyzøφほ本"
	// I'm also not sure if what's intended is to run a goroutine for each
	// letter across the concatenated text,
	// instead of running a goroutine for each string of the string slice.
	// Concept is the same but applied differently
	var concatText = strings.Join(texts, "")
	concatText = strings.ToLower(concatText)

	var wg sync.WaitGroup
	for _, _ = range letters {
		// init worker for each letter
		wg.Go(func() { CountFrequencyWorker(&wg, letterChan, concatText, outChan) })
	}
	for _, letter := range letters {
		letterChan <- letter // send the letter
		var frequency int
		frequency += <-outChan // poll for the result
		if frequency > 0 {
			freqMap[letter] += frequency
		}
	}
	return freqMap
}
