// raindrops is a exerism example
package raindrops

import "strconv"

// Convert should convert an integer in a string emulating rain sounds
func Convert(number int) (rain string) {
    if number % 3 == 0 {
        rain += "Pling"
    }
    if number % 5 == 0 {
        rain += "Plang"
    }
    if number % 7 == 0 {
        rain += "Plong"
    }
    if rain == "" {
        rain = strconv.Itoa(number)
    }
    return
}
