# Name

`diceware`

# Description

`diceware` generates human readable passwords.

# Purpose

I wrote this simply to play around a bit with some specific go patterns, such
as test coverage, concurrency, interfaces, and error handling.  While this
tool _seems_ to work fine for me, use it for your own passwords at your own
risk.

# Usage

    package main

    import (
      "fmt"
      "diceware"
      "flag"
    )

    func main() {
      var num = flag.Int("d", 8, "number of dice")
      flag.Parse();
      dice := diceware.New()
      if result, err := dice.Generate(\*num); err == nil {
        fmt.Printf("%s\\n", result)
      } else {
        fmt.Printf("error: %s\\n", err)
      }
    }

# License

MIT License, see LICENSE.md for details
