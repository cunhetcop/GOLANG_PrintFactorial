// TÌM GIAI THỪA CỦA 1 SỐ N NHẬP VÀO BẰNG BÀN PHÍM
package methods

import (
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func FactorialCalculator() {

	time.AfterFunc(5*time.Second, func() {
		color.Red("Chương trình đã tự động kết thúc sau 5 giây")
		panic("timeout")
	})

	var input string
	count := 0
	color.Yellow("CHƯƠNG TRÌNH TÌM GIAI THỪA CỦA 1 SỐ N")
	color.Green("Nhập một số nguyên dương: ")
	for count < 3 {
		//nhập khoảng trắng
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Printf("Nhập sai %d lần \n", count+1)
			count++
			continue
		}
		//nhập sai
		n, err := strconv.ParseInt(input, 10, 64)
		if err != nil || n <= 0 {
			fmt.Printf("Nhập sai %d lần \n", count+1)
			count++
			continue
		}
		if n > 100 {
			fmt.Println("Số quá lớn, không thể tính được giai thừa")
			return
		}
		result, factors := factorial(n)
		fmt.Printf("%d! = %s\n", n, result.String())
		fmt.Printf("Các thừa số của %d! là: %v\n", n, factors)
		return
	}
	color.Red("Bạn đã nhập sai quá 3 lần. Game over.")
}

func factorial(n int64) (*big.Int, []int64) {
	result := new(big.Int).SetInt64(1)
	factors := []int64{}
	for i := int64(1); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
		factors = append(factors, i)
	}
	return result, factors
}