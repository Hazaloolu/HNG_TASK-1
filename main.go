package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// isPrime checks if a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// isPerfect checks if a number is a perfect number
func isPerfect(n int) bool {
	if n < 2 {
		return false
	}
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n
}

func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// check if number is an armstrong Number
func isArmstrong(n int) bool {
	temp, sum, digits := n, 0, int(math.Log10(float64(n)))+1
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == n
}

// getFunFact fetches a fun fact from the Numbers API
func getFunFact(number int) string {
	url := fmt.Sprintf("http://numbersapi.com/%d/math?json", number)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching fun fact:", err)
		return "Fun fact not available."
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Error decoding response:", err)
		return "Fun fact not available."
	}
	if fact, ok := result["text"].(string); ok {
		return fact
	}
	return "Fun fact not available."
}

func classifyNumber(c *gin.Context) {
	numberStr := c.Query("number")

	// Validate input
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"number": numberStr,
			"error":  true,
		})
		return
	}

	// Properties
	prime := isPrime(number)
	perfect := isPerfect(number)
	armstrong := isArmstrong(number)
	dSum := digitSum(number)
	parity := "even"
	if number%2 != 0 {
		parity = "odd"
	}

	properties := []string{}
	if armstrong {
		properties = append(properties, "armstrong")
	}
	properties = append(properties, parity)

	funFact := getFunFact(number)

	// Return JSON response
	c.JSON(http.StatusOK, gin.H{
		"number":     number,
		"is_prime":   prime,
		"is_perfect": perfect,
		"properties": properties,
		"digit_sum":  dSum,
		"fun_fact":   funFact,
	})
}

func main() {
	r := gin.Default()
	r.GET("/api/classify-number", classifyNumber)

	// Start the server
	port := "8080"
	fmt.Println("Server running on port", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
