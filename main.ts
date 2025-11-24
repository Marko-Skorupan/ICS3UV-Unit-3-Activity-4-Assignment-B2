/**
 * @author Marko Skorupan
 * @version 1.0.0
 * @date 2025-11-21
 * @fileoverview "The program asks users what automobile addons they want."
 */

// CONSTANTS 
const basePrice = 25000;
const matsCost = 500;
const navCost = 1000;
const seatsCost = 500;
const warrantyCost = 2500;

let totalPrice: number = basePrice;

console.log("Base Price   25000.00");

// Floor mats
let mats = prompt("Do you want floor mats? (yes/no)");
if (mats == "yes") {
  console.log("Floor mats   500.00");
  totalPrice += matsCost;
}

// Navigation
let nav = prompt("Do you want a navigation system? (yes/no)");
if (nav == "yes") {
  console.log("Navigation system   1000.00");
  totalPrice += navCost;
}

// Heated seats
let seats = prompt("Do you want heated leather seats? (yes/no)");
if (seats == "yes") {
  console.log("Heated leather seats   500.00");
  totalPrice += seatsCost;
}

// Warranty
let warranty = prompt("Do you want the 5-year extended warranty? (yes/no)");
if (warranty == "yes") {
  console.log("5-year extended warranty   2500.00");
  totalPrice += warrantyCost;
}

// TAX + FINAL COST
let tax = totalPrice * 0.13;
let finalCost = totalPrice + tax;

console.log("13% Taxes   " + tax.toFixed(2));
console.log("Final cost of car   " + finalCost.toFixed(2));

console.log("\nDone.");
