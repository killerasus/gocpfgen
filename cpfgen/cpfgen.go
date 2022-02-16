package cpfgen

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	maxTries = 5
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CPFFormat(cpf []int) string {
	cpfStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(cpf)), ""), "[]")
	return fmt.Sprintf("%s.%s.%s-%s", cpfStr[0:3], cpfStr[3:6], cpfStr[6:9], cpfStr[9:])
}

func GenCPF() (cpf []int, ok bool) {
	cpf = make([]int, 11)

	ok = false
	tries := 0

	for !ok && tries < maxTries {
		for i := 0; i < len(cpf)-2; i++ {
			cpf[i] = rand.Intn(10)
		}

		cpf[9] = GenerateFirstDigit(cpf)
		cpf[10] = GenerateSecondDigit(cpf)

		ok = ValidateCPF(cpf)
		tries++
	}

	return
}

func GenerateFirstDigit(cpf []int) int {
	d := cpf[0]*10 + cpf[1]*9 + cpf[2]*8 + cpf[3]*7 + cpf[4]*6 + cpf[5]*5 + cpf[6]*4 + cpf[7]*3 + cpf[8]*2
	if d = (10 * d) % 11; d == 10 {
		d = 0
	}
	return d
}

func GenerateSecondDigit(cpf []int) int {
	d := cpf[0]*11 + cpf[1]*10 + cpf[2]*9 + cpf[3]*8 + cpf[4]*7 + cpf[5]*6 + cpf[6]*5 + cpf[7]*4 + cpf[8]*3 + cpf[9]*2
	if d = (10 * d) % 11; d == 10 {
		d = 0
	}
	return d
}

func ValidateCPF(cpf []int) bool {
	if IsRepeated(cpf) {
		return false
	}

	firstDigit := cpf[0]*10 + cpf[1]*9 + cpf[2]*8 + cpf[3]*7 + cpf[4]*6 + cpf[5]*5 + cpf[6]*4 + cpf[7]*3 + cpf[8]*2
	if firstDigit = (10 * firstDigit) % 11; firstDigit == 10 {
		firstDigit = 0
	}

	if firstDigit != cpf[9] {
		return false
	}

	secondDigit := cpf[0]*11 + cpf[1]*10 + cpf[2]*9 + cpf[3]*8 + cpf[4]*7 + cpf[5]*6 + cpf[6]*5 + cpf[7]*4 + cpf[8]*3 + cpf[9]*2
	if secondDigit = (10 * secondDigit) % 11; secondDigit == 10 {
		secondDigit = 0
	}

	if secondDigit != cpf[10] {
		return false
	}

	return true
}

func IsRepeated(cpf []int) bool {
	x := cpf[0]
	for i := 1; i < len(cpf); i++ {
		if x != cpf[i] {
			return false
		}
	}
	return true
}
