package RC4

func RC4(key []byte, data []byte) []byte {
	result := make([]byte, len(data))
	var T, S [256]byte
	var j byte

	//Actually this is a bug in C FakeRC4 becasue it using strlen bu not len!
	Klen := 0
	for _, ch := range key {
		if ch == 0 {
			break
		}
		Klen++
	}

	for i := range S {
		S[i] = byte(i)
		T[i] = key[i%Klen]
	}

	for i := range S {
		j = j + S[i] + T[i]
		S[i], S[j] = S[j], S[i]
	}

	j = 0
	i := 0
	for x := range data {
		i = i + 1 // using %256 to avoid exceed the array limit
		j = j + S[i]

		//Swap S[i] & S[j]
		S[i], S[j] = S[j], S[i]
		t := S[i] + S[j]

		result[x] = data[x] ^ S[t] // XOR generated S[t] with Byte from the plaintext / cipher and append each Encrypted/Decrypted byte to result array
	}

	return result
}
