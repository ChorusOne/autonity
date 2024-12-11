package main

import (
	"encoding/csv"
	"math/big"
	"os"
	"strings"
	"text/template"

	"github.com/autonity/autonity/params"
)

type Alloc struct {
	Account string
	Value   string
}

type TemplateParams struct {
	ATNGenesis     []Alloc
	NTNgenesis     []Alloc
	NTNLS          []Alloc
	NTNLNS         []Alloc
	NTNGVBonds     []Alloc
	SDPDelegations []Alloc
}

func parseAmount(csvAmount string) string {
	values := strings.Split(csvAmount, ".")
	atn, ok := new(big.Int).SetString(values[0], 10)
	if !ok {
		panic("can't parse amount")
	}
	atn = new(big.Int).Mul(atn, big.NewInt(params.Ether))
	wei, ok := new(big.Int).SetString(values[1], 10)
	if !ok {
		panic("can't parse amount")
	}
	return new(big.Int).Add(atn, wei).String()
}

func parseAtnAllocs() []Alloc {
	file, err := os.Open("./genesisdata/picc-tiber-genesis-atn.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	data := make([]Alloc, 0)
	for i, record := range records {
		if i == 0 { // Skip the header row
			continue
		}
		data = append(data, Alloc{
			Account: record[0],
			Value:   parseAmount(record[1]),
		})
	}
	return data
}

func parseNTNallocs() (genesis, ls, lns, gvBonds []Alloc) {
	file, err := os.Open("./genesisdata/picc-tiber-genesis-ntn.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	for i, record := range records {
		if i == 0 { // Skip the header row
			continue
		}
		alloc := Alloc{
			Account: record[0],
			Value:   parseAmount(record[1]),
		}
		switch record[3] {
		case "0":
			genesis = append(genesis, alloc)
		case "1":
			ls = append(ls, alloc)
		case "2":
			lns = append(lns, alloc)
		case "3":
			gvBonds = append(gvBonds, alloc)
		}
	}
	return genesis, ls, lns, gvBonds
}

func parseSDPDelegation() []Alloc {
	file, err := os.Open("./genesisdata/picc-tiber-genesis-SDP-delegations.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	data := make([]Alloc, 0)
	for i, record := range records {
		if i == 0 { // Skip the header row
			continue
		}
		data = append(data, Alloc{
			Account: record[1],
			Value:   parseAmount(record[2]),
		})
	}
	return data
}

func main() {
	// Import csv data
	atnAllocs := parseAtnAllocs()
	ntnAllocs, lsntn, lnsntn, gvNtnBonds := parseNTNallocs()
	sdpDelegation := parseSDPDelegation()
	templateConfig := TemplateParams{
		ATNGenesis:     atnAllocs,
		NTNgenesis:     ntnAllocs,
		NTNLS:          lsntn,
		NTNLNS:         lnsntn,
		NTNGVBonds:     gvNtnBonds,
		SDPDelegations: sdpDelegation,
	}
	// Define the Go source file template

	const templateText = `// Code generated by picgen. DO NOT EDIT.
package params

import (
	"math/big"

	"github.com/autonity/autonity/common"
);

var PiccadillyATNallocs = []struct {
	Address common.Address
	Value   *big.Int
}{
{{- range .ATNGenesis }}
	{Address: common.HexToAddress("{{ .Account }}"), Value: mustParseString("{{ .Value }}")},
{{- end }}
}

var PiccadillyNTNallocs = []struct {
	Address common.Address
	Value   *big.Int
}{
{{- range .NTNgenesis }}
	{Address: common.HexToAddress("{{ .Account }}"), Value: mustParseString("{{ .Value }}")},
{{- end }}
}

var PiccadillyLSNTNallocs = []struct {
	Address common.Address
	Value   *big.Int
}{
{{- range .NTNLS }}
	{Address: common.HexToAddress("{{ .Account }}"), Value: mustParseString("{{ .Value }}")},
{{- end }}
}

var PiccadillyLNSNTNallocs = []struct {
	Address common.Address
	Value   *big.Int
}{
{{- range .NTNLNS }}
	{Address: common.HexToAddress("{{ .Account }}"), Value: mustParseString("{{ .Value }}")},
{{- end }}
}

var PiccadillyGVNTNallocs = []struct {
	Address common.Address
	Value   *big.Int
}{
{{- range .NTNGVBonds }}
	{Address: common.HexToAddress("{{ .Account }}"), Value: mustParseString("{{ .Value }}")},
{{- end }}
}

var PiccadillySDPDelegations = []struct {
	Address common.Address
	Value   *big.Int
}{
{{- range .SDPDelegations }}
	{Address: common.HexToAddress("{{ .Account }}"), Value: mustParseString("{{ .Value }}")},
{{- end }}
}

`
	// Create the output file
	outputFile, err := os.Create("gen_piccadilly_config.go")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Parse and execute the template
	tmpl, err := template.New("goSource").Parse(templateText)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(outputFile, templateConfig)
	if err != nil {
		panic(err)
	}

	println("Generated Piccadilly genesis config successfully!")
}