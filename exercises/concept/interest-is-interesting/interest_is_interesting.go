package interest

import "fmt"
import "math"

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
        case balance < 0.0:
    		return float32(3.213)
        case balance >=0.0 && balance <1000.0:
    		return float32(0.5)
        case balance >=1000.0 && balance < 5000.0:
    		return float32(1.621)
        default:
    		return float32(2.475)
    }
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	switch {
        case balance < 0.0:
    		return float64(3.213/100.0*balance)
        case balance >=0.0 && balance <1000.0:
    		return float64(0.5/100.0*balance)
        case balance >=1000.0 && balance < 5000.0:
    		return float64(1.621/100.0*balance)
        default:
    		return float64(2.475/100.0*balance)
    }
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	switch {
        case balance < 0.0:
    		return float64(1.03213*balance)
        case balance >=0.0 && balance <1000.0:
    		return float64(1.005*balance)
        case balance >=1000.0 && balance < 5000.0:
    		return float64(1.01621*balance)
        default:
    		return float64(1.02475*balance)
    }
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	updatedBalance := balance
    yearCount := 1
    
    if balance == targetBalance || balance > targetBalance {
        return 0
    }
    
    for {    
        updatedBalance = AnnualBalanceUpdate(updatedBalance)
        if math.Abs(updatedBalance-targetBalance)<0.001 {
            break
        } else if updatedBalance < targetBalance {
            yearCount++
        } else {
        	break
        }
        fmt.Printf("ub = %f yc = %d\n", math.Abs(updatedBalance-targetBalance), yearCount)
    }
	
    return yearCount
}
