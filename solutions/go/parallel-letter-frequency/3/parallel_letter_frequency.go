package parallelletterfrequency

import (
	"regexp"
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

var mu sync.Mutex

func CountFrequencyWorker(wg *sync.WaitGroup, textChan <-chan string, freqMap *FreqMap) {
	reDigit := regexp.MustCompile("[0-9]")
	rePunct := regexp.MustCompile("[[:punct:]]")
	reSpace := regexp.MustCompile("[[:space:]]")
	for text := range textChan {
		for _, letter := range text {
			if reDigit.MatchString(string(letter)) ||
				rePunct.MatchString(string(letter)) ||
				reSpace.MatchString(string(letter)) {
					// skip
					continue
				}
			mu.Lock() // prevent concurrent writes
			(*freqMap)[letter]++
			mu.Unlock()
		}
	}
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	var freqMap = make(map[rune]int)
	lower_text := strings.ToLower(text)
	for _, letter := range lower_text {
		freqMap[letter] += 1
	}
	return freqMap
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	var freqMap = make(FreqMap)
	var textChan = make(chan string)

	var wg sync.WaitGroup
	for _, _ = range texts {
		// init worker for each text
		wg.Go(func() { CountFrequencyWorker(&wg, textChan, &freqMap) })
	}
	for _, text := range texts {
		textChan <- strings.ToLower(text) // send text to textChan
	}
	close(textChan) // otherwise waits forever to close
	wg.Wait() // wait until all finished
	return freqMap
}
