package diceware

import (
  "testing"
)

func TestNew(t *testing.T) {
  var dice Diceware
  if dice = New(); dice == nil {
    t.Error("New() should not return nil")
  }
}

func TestGenerate(t *testing.T) {
  dice := New()
  if words, err := dice.Generate(5); err != nil {
    t.Error("Generate() should not return an error")
  } else {
    if words == "" {
      t.Error("Generate() produced an empty string")
    } else if len(words) < 15 {
      t.Error("Generate() created a result that is too small")
    }
  }
}

func TestGenerateBig(t *testing.T) {
  dice := New()
  if words, err := dice.Generate(5000); err != nil {
    t.Error("Generate() should not return an error on large set")
  } else {
    if words == "" {
      t.Error("Generate() produced an empty string on large set")
    } else if len(words) < 15000 {
      t.Error("Generate() created a result that is too small on large set")
    }
  }
}

func TestGenerateZero(t *testing.T) {
  dice := New()
  if _, err := dice.Generate(0); err == nil {
    t.Error("Generate() should error on 0 or fewer words")
  }
}

func TestLookupMissing(t *testing.T) {
  dice := &diceDB{}
  if _, err := dice.lookup("99999"); err == nil {
    t.Error("lookup(99999) should have errored")
  }
}
