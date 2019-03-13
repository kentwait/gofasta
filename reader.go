package gofasta

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

// FastaFileToCharAlignment reads a FASTA file into a character-based Alignment struct.
func FastaFileToCharAlignment(path string) (sequences Alignment) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return FastaToAlignment(file, false)
}

// FastaFileToCodonAlignment reads a FASTA file into a codon-based Alignment struct.
func FastaFileToCodonAlignment(path string) (sequences Alignment) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return FastaToAlignment(file, true)
}

// FastaToAlignment reads a FASTA-formatted io.Reader stream into an Alignment struct.
func FastaToAlignment(file io.Reader, toCodon bool) (sequences Alignment) {
	reader := bufio.NewReader(file)

	var err error
	var name, desc string
	var seqBuffer bytes.Buffer
	var splitted []string

	var line string
	for {
		line, err = reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if strings.HasPrefix(line, ">") {
			if seqBuffer.Len() > 0 {
				var sequence Sequence
				if toCodon == true {
					sequence = NewCodonSequence(name, desc, seqBuffer.String())
				} else {
					sequence = NewCharSequence(name, desc, seqBuffer.String())
				}
				sequences = append(sequences, sequence)
				seqBuffer.Reset()
				name, desc = "", ""
			}
			splitted = strings.SplitN(line[1:], " ", 2)
			name = splitted[0]
			if len(splitted) == 2 {
				desc = splitted[1]
			}
		} else if strings.HasPrefix(line, "\n") {
			continue
		} else if strings.HasPrefix(line, "#") {
			continue
		} else if strings.HasPrefix(line, ";") {
			continue
		} else if len(name) > 0 {
			seqBuffer.WriteString(line)
		}
		if err == io.EOF {
			break
		} else if err != nil {
			panic("[Error!] alignment file may be malformed")
		}
	}
	if seqBuffer.Len() > 0 {
		var sequence Sequence
		if toCodon == true {
			sequence = NewCodonSequence(name, desc, seqBuffer.String())
		} else {
			sequence = NewCharSequence(name, desc, seqBuffer.String())
		}
		sequences = append(sequences, sequence)
		seqBuffer.Reset()
	}
	return
}
