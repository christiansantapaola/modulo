package modulo

import "fmt"

// PowerMod calculate g^x mod n
func PowerMod(base, exp, modulus int64) (int64, error) {
	var res int64 = 1
	if exp < 0 {
		inv, err := InverseMod(base, modulus)
		if err != nil {
			return 0, fmt.Errorf("can't compute inverse of %d  modulo %d",
				base, modulus)
		}
		base = inv
		exp = -exp
	}

	if modulus == 1 {
		return 0, nil
	}
	base = base % modulus
	for exp > 0 {
		if exp%2 != 0 {
			res = (res * base) % modulus
		}
		exp = exp >> 1
		base = (base * base) % modulus
	}
	return res, nil
}

// Gcd compute common demoninator
func Gcd(n1, n2 int64) int64 {
	for n2 != 0 {
		n1, n2 = n2, n1%n2
	}
	return n1
}

// InverseMod calculate the inverse of base mod modulus
func InverseMod(base, modulus int64) (int64, error) {
	euclid := ExtendedEuclideanAlgorithm(base, modulus)
	inverse, _, err := euclid.Exec()
	if err != nil {
		return 0, err
	}
	for inverse < 0 {
		inverse += modulus
	}
	return inverse, nil
}

func RSAEncrypt(message, e, pubkey int64) int64 {
	cryptMessage, _ := PowerMod(message, e, pubkey)
	return cryptMessage
}

func RSADecrypt(cryptedMessage, privkey, pubkey int64) int64 {
	msg, _ := PowerMod(cryptedMessage, privkey, pubkey)
	return msg
}
