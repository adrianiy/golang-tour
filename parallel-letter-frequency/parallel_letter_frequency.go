// Package letter includes concurrency exercices
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ChannelFrequency returns word frequency to the channel.
func ChannelFrequency(s string, c chan<- FreqMap) {
	c <- Frequency(s)
}

// ParseList gets rune frequency of each list elemnt and send it to channel
func ParseList(l []string, c chan<- FreqMap) {
	for _, s := range l {
		go ChannelFrequency(s, c)
	}
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	m := FreqMap{}
	channel := make(chan FreqMap)

	go ParseList(l, channel)

	for range l {
		for k, v := range <-channel {
			m[k] += v
		}
	}

	return m
}
