package kata

type VigenèreCipher struct {
	Key   string
	Alpha string
}

func (c VigenèreCipher) Encode(in string) string {
	inRunes := []rune(in)
	alphaRunes := []rune(c.Alpha)
	keyRunes := []rune(c.Key)

	alphaRunesMap := make(map[rune]int)
	for pos, val := range []rune(c.Alpha) {
		alphaRunesMap[val] = pos
	}

	outRunes := make([]rune, 0, len(inRunes))

	for pos, val := range inRunes {
		x, isAlpha := alphaRunesMap[val]
		if !isAlpha {
			outRunes = append(outRunes, val)
			continue
		}

		y := alphaRunesMap[keyRunes[pos%len(keyRunes)]]

		outRunes = append(outRunes, alphaRunes[(x+y)%len(alphaRunes)])
	}

	return string(outRunes)
}

func (c VigenèreCipher) Decode(in string) string {
	inRunes := []rune(in)
	alphaRunes := []rune(c.Alpha)
	keyRunes := []rune(c.Key)

	alphaRunesMap := make(map[rune]int)
	for pos, val := range []rune(c.Alpha) {
		alphaRunesMap[val] = pos
	}
	outRunes := make([]rune, 0, len(inRunes))
	for pos, val := range inRunes {
		f, isAlpha := alphaRunesMap[val]
		if !isAlpha {
			outRunes = append(outRunes, val)
			continue
		}

		y := alphaRunesMap[keyRunes[pos%len(keyRunes)]]

		outRunes = append(outRunes, alphaRunes[(len(alphaRunes)+f-y)%len(alphaRunes)])

	}
	return string(outRunes)
}
