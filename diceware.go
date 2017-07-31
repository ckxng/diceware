// Package diceware generates diceware passwords.
package diceware

// New returns a Diceware password generator
// The entropy provided is 12.9 bits per word.
func New() (Diceware) {
  return &diceDB{}
}

// interface Diceware is capable of generating human readable passwords
type Diceware interface {

  // Generate a diceware passphrase containing the specified number of words.
  Generate(int) (string, error)
}
