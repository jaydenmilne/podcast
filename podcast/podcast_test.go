package podcast

import (
	"bytes"
	"encoding/xml"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jaydenm/podcast/rss"
)

var testCases = []struct {
	name     string
	testFile []byte
	expected RSSPodcast
}{
	{"apple_sample", ApplePodcastSample, ApplePodcastSampleExpected},
	{"more_complex_sample", MoreComplexSample, MoreComplexSampleExpected},
	{"podcasting_2.0_example", Podcasting20Example, Podcasting20ExampleExpected},
}

func TestUnmarshal(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var pod RSSPodcast

			decoder := rss.GetDecoder(bytes.NewReader(tc.testFile))
			if err := decoder.Decode(&pod); err != nil {
				t.Fatalf("failure to unmarshal: %s", err)
			}

			if !cmp.Equal(tc.expected, pod) {
				t.Errorf("document didn't match! %s", cmp.Diff(tc.expected, pod))
			}
		})
	}
}

func TestMarshalRoundTrip(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var doc RSSPodcast

			decoder := rss.GetDecoder(bytes.NewReader(tc.testFile))

			if err := decoder.Decode(&doc); err != nil {
				t.Fatalf("failure to unmarshal: %s", err)
			}
			var err error
			var marshalled []byte
			if marshalled, err = xml.Marshal(&doc); err != nil {
				t.Fatalf("failure to marshal: %s", err)

			}
			os.WriteFile("feed.xml", marshalled, 0644)

			var docRoundTwo RSSPodcast
			decoder2 := rss.GetDecoder(bytes.NewReader(marshalled))

			if err := decoder2.Decode(&docRoundTwo); err != nil {
				t.Fatalf("failure to unmarshal: %s", err)
			}

			if !cmp.Equal(doc, docRoundTwo) {
				t.Errorf("document didn't match! %s", cmp.Diff(doc, docRoundTwo))
			}
		})
	}
}
