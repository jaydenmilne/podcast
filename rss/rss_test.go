package rss

import (
	"bytes"
	"encoding/xml"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var testCases = []struct {
	name     string
	testFile []byte
	expected RSS
}{
	{"rss_spec_sample", SpecSampleXML, SpecSampleExpected},
	{"more_complex_sample", MoreComplexSample, MoreComplexSampleExpected},
}

func TestUnmarshal(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var doc RSS

			decoder := GetDecoder(bytes.NewReader(tc.testFile))

			if err := decoder.Decode(&doc); err != nil {
				t.Fatalf("failure to unmarshal: %s", err)
			}

			if !cmp.Equal(tc.expected, doc) {
				t.Errorf("document didn't match! %s", cmp.Diff(tc.expected, doc))
			}
		})
	}
}

func TestMarshalRoundTrip(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var doc RSS

			decoder := GetDecoder(bytes.NewReader(tc.testFile))

			if err := decoder.Decode(&doc); err != nil {
				t.Fatalf("failure to unmarshal: %s", err)
			}
			var err error
			var marshalled []byte
			if marshalled, err = xml.Marshal(&doc); err != nil {
				t.Fatalf("failure to marshal: %s", err)

			}
			var docRoundTwo RSS
			decoder2 := GetDecoder(bytes.NewReader(marshalled))

			if err := decoder2.Decode(&docRoundTwo); err != nil {
				t.Fatalf("failure to unmarshal: %s", err)
			}

			if !cmp.Equal(doc, docRoundTwo) {
				t.Errorf("document didn't match! %s", cmp.Diff(doc, docRoundTwo))
			}

			os.WriteFile("feed.xml", marshalled, 0644)
		})
	}
}
