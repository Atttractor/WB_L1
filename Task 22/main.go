package main

import (
	"fmt"
	"math/big"
	"strconv"
)

type BigNum struct {
	num *big.Int
}

func getBigNum(n int) BigNum {
	bigNumber, _ := big.NewInt(0).SetString(strconv.Itoa(n), 10)
	return BigNum {
		num: bigNumber,
	}
}

func (bn1 BigNum) mul(bn2 BigNum) BigNum {
	res := BigNum {
		num: big.NewInt(0).Mul(bn1.num, bn2.num),
	}

	return res
}

func (bn1 BigNum) div(bn2 BigNum) BigNum {
	res := BigNum {
		num: big.NewInt(0).Div(bn1.num, bn2.num),
	}

	return res
}

func (bn1 BigNum) add(bn2 BigNum) BigNum {
	res := BigNum {
		num: big.NewInt(0).Add(bn1.num, bn2.num),
	}

	return res
}

func (bn1 BigNum) Sub(bn2 BigNum) BigNum {
	res := BigNum {
		num: big.NewInt(0).Sub(bn1.num, bn2.num),
	}

	return res
}

func main() {
	a := getBigNum(20000000)
	b := getBigNum(3000000000)
	fmt.Println(a.add(b).num)
}