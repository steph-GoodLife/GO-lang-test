package main

import (
    "flag"
    "fmt"
    "rand"
    "time"
    "regexp"
    "strconv"
    "log"
    
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    dice := flag.String("d", "d6", "the type of dice to roll. Format: dX where X is an integer Default")
    flag.Parse()

    matched, _ := regexp.Match("d\\d+", []byte(*dice))

    if matched == true {
        diceSide := (*dice)[1:]
        d, err := strconv.Atoi(diceSide)
        if err != nil {
            log.Fatal(err)
        }
        roll := rand.Intn(d) + 1
        fmt.Printf("you choose a %d\n", roll)
    }else{
        log.Fatal("Improper format for dice. Format should be dX where X is an integer")
    }

}