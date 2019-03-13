package gofasta

var Bases = [4]string{"T", "C", "A", "G"}
var Codons = [64]string{
	"TTT", "TTC", "TTA", "TTG",
	"TCT", "TCC", "TCA", "TCG",
	"TAT", "TAC", "TAA", "TAG",
	"TGT", "TGC", "TGA", "TGG",
	"CTT", "CTC", "CTA", "CTG",
	"CCT", "CCC", "CCA", "CCG",
	"CAT", "CAC", "CAA", "CAG",
	"CGT", "CGC", "CGA", "CGG",
	"ATT", "ATC", "ATA", "ATG",
	"ACT", "ACC", "ACA", "ACG",
	"AAT", "AAC", "AAA", "AAG",
	"AGT", "AGC", "AGA", "AGG",
	"GTT", "GTC", "GTA", "GTG",
	"GCT", "GCC", "GCA", "GCG",
	"GAT", "GAC", "GAA", "GAG",
	"GGT", "GGC", "GGA", "GGG",
}
var StopCodons = [3]string{"TGA", "TAG", "TAA"}
var AminoAcids = [20]string{
	"A",
	"R",
	"N",
	"D",
	"C",
	"Q",
	"E",
	"G",
	"H",
	"I",
	"L",
	"K",
	"M",
	"F",
	"P",
	"S",
	"T",
	"W",
	"Y",
	"V",
}
var GeneticCode = map[string]string{
	"TTT": "F",
	"TTC": "F",
	"TTA": "L",
	"TTG": "L",
	"TCT": "S",
	"TCC": "S",
	"TCA": "S",
	"TCG": "S",
	"TAT": "Y",
	"TAC": "Y",
	"TAA": "*",
	"TAG": "*",
	"TGT": "C",
	"TGC": "C",
	"TGA": "*",
	"TGG": "W",
	"CTT": "L",
	"CTC": "L",
	"CTA": "L",
	"CTG": "L",
	"CCT": "P",
	"CCC": "P",
	"CCA": "P",
	"CCG": "P",
	"CAT": "H",
	"CAC": "H",
	"CAA": "Q",
	"CAG": "Q",
	"CGT": "R",
	"CGC": "R",
	"CGA": "R",
	"CGG": "R",
	"ATT": "I",
	"ATC": "I",
	"ATA": "I",
	"ATG": "M",
	"ACT": "T",
	"ACC": "T",
	"ACA": "T",
	"ACG": "T",
	"AAT": "N",
	"AAC": "N",
	"AAA": "K",
	"AAG": "K",
	"AGT": "S",
	"AGC": "S",
	"AGA": "R",
	"AGG": "R",
	"GTT": "V",
	"GTC": "V",
	"GTA": "V",
	"GTG": "V",
	"GCT": "A",
	"GCC": "A",
	"GCA": "A",
	"GCG": "A",
	"GAT": "D",
	"GAC": "D",
	"GAA": "E",
	"GAG": "E",
	"GGT": "G",
	"GGC": "G",
	"GGA": "G",
	"GGG": "G",
	"---": "-",
}
