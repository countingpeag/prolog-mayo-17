package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("�Cu�ndo es S�bado?")
    today := time.Now().Weekday()
    switch time.Saturday {
    case today + 0:
        fmt.Println("Hoy.")
    case today + 1:
        fmt.Println("Ma�ana.")
    case today + 2:
        fmt.Println("Pasado ma�ana.")
    default:
        fmt.Println("Demasiado tarde.")
    }
}