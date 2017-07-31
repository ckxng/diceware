package diceware

import (
  "encoding/json"
  "crypto/rand"
  "math/big"
  "strings"
  "strconv"
  "bytes"
  "errors"
  "fmt"
  "sync"
  )

// type diceDB contains a db of diceware words
type diceDB struct {
  db      map[string]string
  config  struct {
    loadOnce  sync.Once
  }
}

// Generate create a diceware password in words length, separated by dashes.
func (dice *diceDB) Generate(words int) (string, error) {
  if words < 1 {
    return "", errors.New("cannot generate 0 or fewer words")
  }
  worda := make([]string, words);
  for i := 0; i < words; i++ {
    if key, err := dice.rand5(); err == nil {
      if newWord, err := dice.lookup(key); err == nil {
        worda[i] = newWord;
      } else {
        return "", err
      }
    } else {
      return "", err
    }
  }
  return strings.Join(worda, "-"), nil
}

// rand5 uses crypto.rand to roll 5 dice and returns a string of 5 numbers.
func (dice *diceDB) rand5() (string, error) {
  var roll5 bytes.Buffer
  for i := 0; i < 5; i++ {
    if roll, err := rand.Int(rand.Reader, big.NewInt(6)); err == nil {
      roll5.WriteString(strconv.Itoa((int(roll.Uint64())+1)))
    } else {
      return "", err
    }
  }
  return roll5.String(), nil
}

// lookup returns the corresponding word.
func (dice *diceDB) lookup(key string) (string, error) {
  if err := dice.loadDiceDB(); err != nil {
    return "", err
  }

  if val, ok := dice.db[key]; ok {
    return val, nil
  } else {
    return "", errors.New(fmt.Sprintf("Key '%s' not found in dice.db", key))
  }
}

// loadDiceDB unmarshals the JSON string from loadDiceString() only once.
// This function should be called at the beginning of any function that reads
// the database, for lazy initialization.
func (dice *diceDB) loadDiceDB() (error) {
  var err error
  dice.config.loadOnce.Do(func() {
    err = json.Unmarshal(loadDiceString(), &dice.db)
  })
  if err != nil {
    return err
  }
  return nil
}
