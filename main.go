/*
 * @author Marko Skorupan
 * @version 1.0.0
 *@date 2025-11-21
 * @fileoverview "The program asks users what automobile addons they want."
 */

package main

import "fmt"

func main () {
const BASE float64 = 25000
const FLOORMATS float64 = 500
const NAV float64 = 1000
const HEATED float64 = 500
const WARRANTY float64 = 25000

var answer string
total := BASE

fmt.Println("Base Price:   ",BASE )

fmt.Println("Do you want floor mats? (yes/no)")
fmt.Scan(&answer)
if answer == "yes" {
	total = total + FLOORMATS
	fmt.Println("Floor Mats   ", FLOORMATS)
}

fmt.Println("Do you want a navigation system? (yes/no)")
fmt.Scan(&answer)
if answer == "yes" {
	total = total + NAV
	fmt.Println("Navigation system   ", NAV)
}

fmt.Println("Do you want heated leather seats? (yes/no)")
fmt.Scan(&answer)
if answer == "yes" {
	total = total + HEATED
	fmt.Println("Heated leather seats   ", HEATED)
}

fmt.Println("Do you want 5-year extended warranty? (yes/no)")
fmt.Scan(&answer)
if answer == "yes" {
	total = total + WARRANTY
	fmt.Println("5-year extended warranty   ", WARRANTY)
}

tax := total * 0.13

fmt.Println("13% Taxes   ", tax)
fmt.Println("Final cost of car   ", total+tax)

fmt.Println("\nDone.")
}
