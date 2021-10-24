package protein

import "errors"

var CodonsMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine",
	"UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cystenine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP", "UGA": "STOP",
}

var ErrInvalidBase = errors.New("invalid base codon")
var ErrStop = errors.New("stop codon")

func FromCodon(input string) (string, error) {
	var protein string
	if _, ok := CodonsMap[input]; !ok {
		return "", ErrInvalidBase
	}
	protein = CodonsMap[input]
	if protein == "STOP" {
		return "", ErrStop
	}

	return protein, nil

}

func FromRNA(input string) ([]string, error) {
	var proteins = make([]string, 0)
	for i := 0; i+3 <= len(input); i = i + 3 {
		codon := input[i : i+3]
		protein, err := FromCodon(codon)
		if err == ErrInvalidBase {
			return proteins, err
		} else if err == ErrStop {
			break
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}
