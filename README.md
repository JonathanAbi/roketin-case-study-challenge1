# Coding Case Study - Challenge 1: Earth to Roketin Time Conversion

This section contains the solution for Challenge 1 of the coding case study.

## Problem Description

On Earth:
* 1 day = 24 hours
* 1 hour = 60 minutes
* 1 minute = 60 seconds

On Planet Roketin:
* 1 day = 10 hours
* 1 hour = 100 minutes
* 1 minute = 100 seconds

The task is to create a program that converts time from Earth to Planet Roketin.
The input consists of 3 integers (Earth hours, minutes, seconds).

**Expected Output Format Example:**
`On Earth HH:MM:SS, on planet Roketin Planet RHH:RMM:RSS`

## Solution Approach

This solution assumes that the total duration of one day on Earth is equivalent to the total duration of one day on Planet Roketin. The conversion is performed using the following steps:

1.  Calculate the total seconds elapsed on Earth based on the input hours, minutes, and seconds.
    * Total seconds in an Earth day = $24 \times 60 \times 60 = 86400$ seconds.
2.  Calculate the conversion ratio between the total seconds per day on Planet Roketin and the total seconds per day on Earth.
    * Total seconds in a Roketin day = $10 \times 100 \times 100 = 100000$ Roketin seconds.
    * `ratio = total_roketin_seconds_per_day / total_earth_seconds_per_day`.
3.  Multiply the total Earth seconds by this ratio to get the equivalent total seconds on Planet Roketin.
    * `roketin_total_seconds = total_earth_seconds * ratio`.
4.  Convert the total Roketin seconds back into Planet Roketin's hours, minutes, and seconds format.
    * The ratio calculation uses `float64` for precision, which is then converted back to `int` for the final results.

## Programming Language

* **Go (Golang)**

## How to Run

1.  **Prerequisites:**
    * Ensure you have the Go compiler installed.

2.  **Running the Program:**
    * Save the solution code (e.g., in a file named `main.go`).
    * Open a terminal or command prompt.
    * Navigate to the directory where you saved the file.
    * Execute the command:
      ```bash
      go run main.go
      ```

3.  **Input and Output:**
    * The program will prompt you to enter Earth's hours, minutes, and seconds interactively via the console.
    * Basic input validation is included to ensure hours, minutes, and seconds are within their valid ranges.
    * The conversion result will be printed directly to the console.
