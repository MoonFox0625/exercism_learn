package strand

var transcribe = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	if dna == "" {
		return ""
	}
	rna := []rune{}
	for _, val := range dna {
		rna = append(rna, transcribe[val])
	}
	return string(rna)

}
