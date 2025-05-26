package rss

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/gdexlab/go-render/render"
	"github.com/google/go-cmp/cmp"
)

func helper(a any) {
	out, _ := os.Create("test.txt")
	out.Write([]byte(render.AsCode(a)))
	out.Close()
}

func TestUnmarshalSpecSample(t *testing.T) {
	var doc RSS

	if err := xml.Unmarshal([]byte(SpecSampleXML), &doc); err != nil {
		t.Fatalf("failure to unmarshal: %s", err)
	}

	if !cmp.Equal(SpecSampleExpected, doc) {
		t.Errorf("document didn't match! %s", cmp.Diff(SpecSampleExpected, doc))
	}
}

func TestMarshalSpecRoundTrip(t *testing.T) {
	var doc RSS

	if err := xml.Unmarshal([]byte(SpecSampleXML), &doc); err != nil {
		t.Fatalf("failure to unmarshal: %s", err)
	}
	var err error
	var marshalled []byte
	if marshalled, err = xml.Marshal(&doc); err != nil {
		t.Fatalf("failure to marshal: %s", err)

	}
	var docRoundTwo RSS
	if err := xml.Unmarshal(marshalled, &docRoundTwo); err != nil {
		t.Fatalf("failure to unmarshal: %s", err)
	}

	if !cmp.Equal(doc, docRoundTwo) {
		t.Errorf("document didn't match! %s", cmp.Diff(doc, docRoundTwo))
	}
}

func TestUnmarshalComplexSample(t *testing.T) {
	var doc RSS

	if err := xml.Unmarshal([]byte(MoreComplexSample), &doc); err != nil {
		t.Fatalf("failure to unmarshal: %s", err)
	}

	if !cmp.Equal(MoreComplexSampleExpected, doc) {
		t.Errorf("document didn't match! %s", cmp.Diff(MoreComplexSampleExpected, doc))
	}
}

func TestMarshalComplexRoundTrip(t *testing.T) {
	var doc RSS

	if err := xml.Unmarshal([]byte(MoreComplexSample), &doc); err != nil {
		t.Fatalf("failure to unmarshal: %s", err)
	}
	var err error
	var marshalled []byte
	if marshalled, err = xml.Marshal(&doc); err != nil {
		t.Fatalf("failure to marshal: %s", err)

	}
	var docRoundTwo RSS
	if err := xml.Unmarshal(marshalled, &docRoundTwo); err != nil {
		t.Fatalf("failure to unmarshal: %s", err)
	}

	if !cmp.Equal(doc, docRoundTwo) {
		t.Errorf("document didn't match! %s", cmp.Diff(doc, docRoundTwo))
	}
}
