package podcast

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jaydenmilne/podcast/rss"
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

func ExamplePodcast_encode() {
	pod := RSSPodcast{
		Version: rss.RSSVersion,
		Channel: Podcast{
			Channel: rss.Channel{
				Title: "My Awesome Feed",
				Link:  "https://example.com",
				Description: rss.Description{
					Value: "<b>AN AMAZING FEED</b>",
				},
			},
			ItunesAuthor: "Dan Jones",
			Items: []Episode{
				{
					Item: rss.Item{
						Title: "Episode 1: The Pod Awakens",
						Enclosure: &rss.Enclosure{
							URL:    "https://example.com/ep01.mp3",
							Length: 123,
							Type:   "audio/mpeg",
						},
					},
				},
			},
			PodcastPeople: []PodcastPerson{
				{
					PersonName: "Steve",
				},
			},
		},
	}

	output, _ := xml.MarshalIndent(pod, "", "\t")
	fmt.Print(string(output))

	// Output: <rss version="2.0">
	//	<channel xmlns="https://www.rssboard.org/rss-specification">
	//		<title xmlns="https://www.rssboard.org/rss-specification">My Awesome Feed</title>
	//		<link xmlns="https://www.rssboard.org/rss-specification">https://example.com</link>
	//		<description xmlns="https://www.rssboard.org/rss-specification"><![CDATA[<b>AN AMAZING FEED</b>]]></description>
	//		<image xmlns="http://www.itunes.com/dtds/podcast-1.0.dtd" href=""></image>
	//		<author xmlns="http://www.itunes.com/dtds/podcast-1.0.dtd">Dan Jones</author>
	//		<person xmlns="https://podcastindex.org/namespace/1.0" href="" img="">Steve</person>
	//		<item xmlns="https://www.rssboard.org/rss-specification">
	//			<title xmlns="https://www.rssboard.org/rss-specification">Episode 1: The Pod Awakens</title>
	//			<enclosure xmlns="https://www.rssboard.org/rss-specification" url="https://example.com/ep01.mp3" length="123" type="audio/mpeg"></enclosure>
	//			<pubDate xmlns="https://www.rssboard.org/rss-specification"></pubDate>
	//		</item>
	//	</channel>
	// </rss>
}

func ExamplePodcast_decode() {
	feedXML := `
	<rss version="2.0">
	<channel xmlns="https://www.rssboard.org/rss-specification">
		<title xmlns="https://www.rssboard.org/rss-specification">My Awesome Feed</title>
		<link xmlns="https://www.rssboard.org/rss-specification">https://example.com</link>
		<description xmlns="https://www.rssboard.org/rss-specification"><![CDATA[<b>AN AMAZING FEED</b>]]></description>
		<image xmlns="http://www.itunes.com/dtds/podcast-1.0.dtd" href=""></image>
		<author xmlns="http://www.itunes.com/dtds/podcast-1.0.dtd">Dan Jones</author>
		<person xmlns="https://podcastindex.org/namespace/1.0" href="" img="">Steve</person>
		<item xmlns="https://www.rssboard.org/rss-specification">
			<title xmlns="https://www.rssboard.org/rss-specification">Episode 1: The Pod Awakens</title>
			<enclosure xmlns="https://www.rssboard.org/rss-specification" url="https://example.com/ep01.mp3" length="123" type="audio/mpeg"></enclosure>
			<pubDate xmlns="https://www.rssboard.org/rss-specification"></pubDate>
		</item>
	</channel>
</rss>`

	decoder := xml.NewDecoder(strings.NewReader(feedXML))
	decoder.DefaultSpace = "https://www.rssboard.org/rss-specification"

	var decoded RSSPodcast
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
	//			"Space": "",
	//			"Local": ""
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
	//		"ItunesImage": {
	//			"XMLName": {
	//				"Space": "http://www.itunes.com/dtds/podcast-1.0.dtd",
	//				"Local": "image"
	//			},
	//			"Href": ""
	//		},
	//		"ItunesCategory": null,
	//		"ItunesExplicit": null,
	//		"ItunesAuthor": "Dan Jones",
	//		"ItunesTitle": "",
	//		"ItunesType": "",
	//		"ItunesNewFeedURL": "",
	//		"ItunesComplete": "",
	//		"ItunesBlock": "",
	//		"ItunesApplePodcastVerify": "",
	//		"PodcastTxt": null,
	//		"PodcastPodroll": null,
	//		"PodcastLocked": null,
	//		"PodcastFunding": null,
	//		"PodcastPeople": [
	//			{
	//				"XMLName": {
	//					"Space": "https://podcastindex.org/namespace/1.0",
	//					"Local": "person"
	//				},
	//				"PersonName": "Steve",
	//				"Group": "",
	//				"Role": "",
	//				"Href": "",
	//				"Img": ""
	//			}
	//		],
	//		"PodcastLocation": null,
	//		"PodcastTrailers": null,
	//		"PodcastValue": null,
	//		"PodcastMedium": "",
	//		"PodcastLiveItem": null,
	//		"PodcastGUID": "",
	//		"PodcastBlock": null,
	//		"PodcastLicense": null,
	//		"Items": [
	//			{
	//				"XMLName": {
	//					"Space": "",
	//					"Local": ""
	//				},
	//				"Title": "Episode 1: The Pod Awakens",
	//				"Link": "",
	//				"Description": null,
	//				"Author": "",
	//				"Categories": null,
	//				"Comments": "",
	//				"Enclosure": {
	//					"XMLName": {
	//						"Space": "https://www.rssboard.org/rss-specification",
	//						"Local": "enclosure"
	//					},
	//					"URL": "https://example.com/ep01.mp3",
	//					"Length": 123,
	//					"Type": "audio/mpeg"
	//				},
	//				"GUID": null,
	//				"PubDate": "",
	//				"Source": null,
	//				"ItunesDuration": "",
	//				"ItunesImage": null,
	//				"ItunesExplicit": null,
	//				"ItunesTitle": "",
	//				"ItunesEpisode": 0,
	//				"ItunesSeason": 0,
	//				"ItunesEpisodeType": "",
	//				"ItunesBlock": "",
	//				"PodcastTranscript": null,
	//				"PodcastChapters": null,
	//				"PodcastSoundbite": null,
	//				"PodcastPeople": null,
	//				"PodcastSeason": null,
	//				"PodcastEpisode": null,
	//				"PodcastLicense": null,
	//				"PodcastAlternateEnclosures": null,
	//				"PodcastValue": null,
	//				"PodcastImages": null,
	//				"PodcastSocialInteracts": null,
	//				"PodcastUpdateFrequency": null,
	//				"PodcastPodping": null
	//			}
	//		]
	//	},
	//	"Version": "2.0"
	// }
}
