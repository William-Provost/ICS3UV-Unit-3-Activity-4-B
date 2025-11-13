/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-13
 * @fileoverview This program tells you what to do in different weather.
 */

// variables
let chanceOfRainAsString = "";
let chanceOfRainAsNumber = 0;

// input % chance of rain
chanceOfRainAsString = prompt("Enter the chance of rain (0-100): ") || "0";
chanceOfRainAsNumber = parseInt(chanceOfRainAsString);

// check the chance of rain
if (chanceOfRainAsNumber > 80) {
  // what to bring if raining
  console.log("Bring an umbrella.");
  console.log("Also bring rain boots!");
} else {
  // what to bring if it is sunny
  console.log("Bring a sun hat.");
  console.log("Also bring a swimsuit!");
}

console.log("\nDone.");