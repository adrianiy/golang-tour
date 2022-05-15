// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// twofer is a package for exercism two-fer exercise
package twofer

// ShareWith should return a string with format "One for XX, One for me."
func ShareWith(name string) string {
    if name == "" {
        name = "you"
    }
    return "One for " + name + ", one for me."
}
