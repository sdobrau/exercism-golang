package interest

import "golang.org/x/crypto/ssh/test"

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
 	if balance < 0 {
        return float32(3.213)
    } else if balance >= 0.0 && balance < 1000.0 {
        return float32(0.5)
    } else if balance >= 1000.0 && balance < 5000.0 {
        return float32(1.621)
    } else if balance >= 5000 {
        return float32(2.475)
    }
    return float32(0.5)
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	currentBalance := balance
	for years := 0; years < 9999; years++ {
		currentBalance += AnnualBalanceUpdate(currentBalance)
		if currentBalance >= targetBalance {
			return years + 1
		}
	}
}
