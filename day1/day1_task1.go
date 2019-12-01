package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"
import "math"
func main() {

  file, err := os.Open("day1_task_input.txt")
  if err != nil {
    fmt.Println("Error reading the file!")
    os.Exit(1)
  }
  defer file.Close()

  var lines []string
  fileScanner := bufio.NewScanner(file)
  for fileScanner.Scan() {
    lines = append(lines, fileScanner.Text())
  }
  var fuelNeed float64
  for _, massString := range lines {
    mass,_ := strconv.ParseInt(strings.TrimSuffix(massString, "\n"), 10, 0)
    fuelNeed = fuelNeed + calculateFuel(mass)
  }
  fmt.Printf("Final Requirement: %f\n", fuelNeed)
}

func calculateFuel(mass int64) float64 {
  fuelRequirement := math.Floor(float64(mass /3)) - 2
  fmt.Println("Fuel requirement:", int(fuelRequirement))
  return fuelRequirement
}
