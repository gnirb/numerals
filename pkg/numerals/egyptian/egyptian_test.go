package egyptian

import (
	"testing"

	"github.com/gnirb/numerals/pkg/numerals"
)

type tests struct {
	name    string
	integer int
	number  numerals.Number
}

var (
	egyptianNumerals = []tests{
		{
			"numeral zero",
			0,
			"",
		},
		{
			"numeral π€",
			1,
			"π€",
		},
		{
			"numeral π",
			10,
			"π",
		},
		{
			"numeral π’",
			100,
			"π’",
		},
		{
			"numeral πΌ",
			1000,
			"πΌ",
		},
		{
			"numeral π­",
			10000,
			"π­",
		},
		{
			"numeral π",
			100000,
			"π",
		},
		{
			"numeral π¨",
			1000000,
			"π¨",
		},
	}
	samples = []tests{
		// Samples
		{
			"sample πΌπΌπΌπΌπ’π’π’π’π’π’πππ€π€",
			4622,
			"πΌπΌπΌπΌπ’π’π’π’π’π’πππ€π€",
		},
		{
			"sample π’π’π’π’π’π’π’ππππππ€",
			751,
			"π’π’π’π’π’π’π’ππππππ€",
		},
	}
	all = [][]tests{
		egyptianNumerals,
		samples,
	}
	converter = NewEgyptianConverter()
)

func TestParse(t *testing.T) {
	for _, tt := range all {
		for _, ttt := range tt {
			t.Run(ttt.name, func(t *testing.T) {
				if got := converter.Parse(ttt.number); got != ttt.integer {
					t.Errorf("Parse() = %v, want %v", got, ttt.number)
				}
			})
		}
	}
}

func TestFormat(t *testing.T) {
	for _, tt := range all {
		for _, ttt := range tt {
			t.Run(ttt.name, func(t *testing.T) {
				if got := converter.Format(ttt.integer); got != ttt.number {
					t.Errorf("Format() = %v, want %v", got, ttt.number)
				}
			})
		}
	}
}

func BenchmarkParse_numerals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range egyptianNumerals {
			converter.Parse(tt.number)
		}
	}
}

func BenchmarkParse_samples(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range samples {
			converter.Parse(tt.number)
		}
	}
}

func BenchmarkParse_all(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range all {
			for _, ttt := range tt {
				converter.Parse(ttt.number)
			}
		}
	}
}

func BenchmarkFormat_numerals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range egyptianNumerals {
			converter.Format(tt.integer)
		}
	}
}

func BenchmarkFormat_samples(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range samples {
			converter.Format(tt.integer)
		}
	}
}

func BenchmarkFormat_all(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range all {
			for _, ttt := range tt {
				converter.Format(ttt.integer)
			}
		}
	}
}
