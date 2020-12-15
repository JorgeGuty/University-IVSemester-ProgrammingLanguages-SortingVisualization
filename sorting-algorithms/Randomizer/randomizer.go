package Randomizer

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func GetSeed() int {

	currentTime := time.Now()
	year := currentTime.Year()
	month := currentTime.Month()
	day := currentTime.Day()
	hour := currentTime.Hour()
	minutes := currentTime.Minute()
	seconds := currentTime.Second()

	seed := ((((year*100+int(month))*100+day)*100+hour)*100+minutes)*100 + seconds

	// increments until the first prime number is found
	for isPrime(seed) == false {
		seed++
	}

	fmt.Println(seed)

	return seed
}

func isPrime(pNumber int) bool {
	return big.NewInt(int64(pNumber)).ProbablyPrime(0)
}

func RandomNumber(pLastRandom int) int {
	// Uses linear congruential generator method to generate pseudo random int values.

	return (RANDOM_GENERATOR_MULTIPLIER * pLastRandom + RANDOM_GENERATOR_INCREMENT) % RANDOM_GENERATOR_MODULUS
}

func RandomArray(pArrayLength int) []int {

	randomNumber := GetSeed() // Random number starts with a seed value.
	var normalizedRandom int  // To store the random number normalized to the MAX_RANDOM_VALUE stipulated.
	var randomizedArray []int // To store the normalized random values.

	for iterationsQuantity := 0; iterationsQuantity < pArrayLength; iterationsQuantity++ {
		randomNumber = RandomNumber(randomNumber)
		normalizedRandom = randomNumber % MAX_RANDOM_VALUE + 1
		randomizedArray = append(randomizedArray, int(math.Abs(float64(normalizedRandom))))
	}

	return randomizedArray
}
