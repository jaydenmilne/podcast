package rss

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
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

func ExampleRSS_decode() {
	var feedXML = `
	<rss xmlns="https://www.rssboard.org/rss-specification" version="2.0">
        <channel xmlns="https://www.rssboard.org/rss-specification">
                <title xmlns="https://www.rssboard.org/rss-specification">My Awesome Feed</title>
                <link xmlns="https://www.rssboard.org/rss-specification">https://example.com</link>
                <description xmlns="https://www.rssboard.org/rss-specification"><![CDATA[<b>AN AMAZING FEED</b>]]></description>
                <item xmlns="https://www.rssboard.org/rss-specification">
                        <title xmlns="https://www.rssboard.org/rss-specification">Post 1</title>
                        <pubDate xmlns="https://www.rssboard.org/rss-specification"></pubDate>
                </item>
        </channel>
</rss>`

	// There is a helper for this boilerplate rss.GetDecoder
	decoder := xml.NewDecoder(strings.NewReader(feedXML))
	decoder.DefaultSpace = "https://www.rssboard.org/rss-specification"

	var decoded RSS
	decoder.Decode(&decoded)

	var jsonBytes, _ = json.MarshalIndent(decoded, "", "\t")
	fmt.Print(string(jsonBytes))

	// Output: {
	//	"XMLName": {
	//		"Space": "https://www.rssboard.org/rss-specification",
	//		"Local": "rss"
	//	},
	//	"Channel": {
	//		"XMLName": {
	//			"Space": "https://www.rssboard.org/rss-specification",
	//			"Local": "channel"
	//		},
	//		"Title": "My Awesome Feed",
	//		"Link": "https://example.com",
	//		"Description": {
	//			"XMLName": {
	//				"Space": "https://www.rssboard.org/rss-specification",
	//				"Local": "description"
	//			},
	//			"Value": "\u003cb\u003eAN AMAZING FEED\u003c/b\u003e"
	//		},
	//		"Language": "",
	//		"Copyright": "",
	//		"ManagingEditor": "",
	//		"WebMaster": "",
	//		"PubDate": "",
	//		"LastBuildDate": "",
	//		"Categories": null,
	//		"Generator": "",
	//		"Docs": "",
	//		"Cloud": null,
	//		"TTL": 0,
	//		"Image": null,
	//		"Rating": "",
	//		"TextInput": null,
	//		"SkipHours": null,
	//		"SkipDays": null,
	//		"Items": [
	//			{
	//				"XMLName": {
	//					"Space": "https://www.rssboard.org/rss-specification",
	//					"Local": "item"
	//				},
	//				"Title": "Post 1",
	//				"Link": "",
	//				"Description": null,
	//				"Author": "",
	//				"Categories": null,
	//				"Comments": "",
	//				"Enclosure": null,
	//				"GUID": null,
	//				"PubDate": "",
	//				"Source": null
	//			}
	//		]
	//	},
	//	"Version": "2.0"
	// }
}

func ExampleRSS_encode() {
	var feed = RSS{
		Version: RSSVersion,
		Channel: Channel{
			Title: "My Awesome Feed",
			Link:  "https://example.com",
			Description: Description{
				Value: "<b>AN AMAZING FEED</b>",
			},
			Items: []Item{
				{
					Title: "Post 1",
				},
			},
		},
	}

	output, _ := xml.MarshalIndent(&feed, "", "\t")

	fmt.Print(string(output))

	// Output: <rss xmlns="https://www.rssboard.org/rss-specification" version="2.0">
	//	<channel xmlns="https://www.rssboard.org/rss-specification">
	//		<title xmlns="https://www.rssboard.org/rss-specification">My Awesome Feed</title>
	//		<link xmlns="https://www.rssboard.org/rss-specification">https://example.com</link>
	//		<description xmlns="https://www.rssboard.org/rss-specification"><![CDATA[<b>AN AMAZING FEED</b>]]></description>
	//		<item xmlns="https://www.rssboard.org/rss-specification">
	//			<title xmlns="https://www.rssboard.org/rss-specification">Post 1</title>
	//			<pubDate xmlns="https://www.rssboard.org/rss-specification"></pubDate>
	//		</item>
	//	</channel>
	// </rss>
}
