package gofasta

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

// FastaFileToAlignment reads a FASTA file into an Alignment struct.
func FastaFileToAlignment(path string) (sequences Alignment) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return FastaToAlignment(file)
}

// FastaToAlignment reads a FASTA-formatted io.Reader stream into an Alignment struct.
func FastaToAlignment(file io.Reader) (sequences Alignment) {
	// file, err := os.Open(path)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

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
				sequences = append(sequences, NewCharSequence(name, desc, seqBuffer.String()))
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
		sequences = append(sequences, NewCharSequence(name, desc, seqBuffer.String()))
		seqBuffer.Reset()
	}
	return
}
