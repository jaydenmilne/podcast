package podcast

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/gdexlab/go-render/render"
)

func helper(a any) {
	out, _ := os.Create("test.txt")
	out.Write([]byte(render.AsCode(a)))
	out.Close()

	out2, _ := os.Create("test_json.txt")
	b, _ := json.MarshalIndent(a, "", "\t")
	out2.Write(b)
	out2.Close()
}

func helper2(b []byte) {
	out, _ := os.Create("test2.txt")
	out.Write(b)
	out.Close()
}

func TestGosBrokenXmlParser(t *testing.T) {
	const input = `
<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd">
    <itunes:type>serial</itunes:type>
    <title>Hiking Treks</title>
</rss>
	`

	type Model struct {
		XMLName    xml.Name `xml:"rss"`
		Title      string   `xml:"title"`
		ItunesType string   `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd type,omitempty"`
	}

	var u Model

	if err := xml.Unmarshal([]byte(input), &u); err != nil {
		t.Fatalf("failure to unmarshal: %s", err)
	}

	helper(u)

	bytes, _ := xml.MarshalIndent(&u, "", "\t")

	helper2(bytes)

}

func TestAppleSample(t *testing.T) {
	var pod RSSPodcast

	if err := xml.Unmarshal(ApplePodcastSample, &pod); err != nil {
		t.Fatalf("failure to unmarshal: %s", err)
	}

	helper(pod)

	//if !cmp.Equal(RSS_SAMPLE_EXPECTED, doc) {
	//	t.Errorf("document didn't match! %s", cmp.Diff(RSS_SAMPLE_EXPECTED, doc))
	//}
}

func TestItWorksAtA(t *testing.T) {
	a := Podcast{
		ItunesTitle: "itunes title",
		Items: []PodcastEpisode{
			PodcastEpisode{},
			PodcastEpisode{},
			PodcastEpisode{},
		},
	}

	b, _ := xml.MarshalIndent(a, "", "\t")
	fmt.Printf("%s", string(b))
}

func TestAmbiguousDeserializing(t *testing.T) {
	type Model struct {
		XMLName  xml.Name `xml:"enclosure"`
		Field    string   `xml:"https://www.rssboard.org/rss-specification field"`
		FieldTwo string   `xml:"http://www.example.org field"`
	}

	input := `<enclosure xmlns:example="http://www.example.org">
		<field>not namespaced</field>
		<example:field>namespaced</example:field>
	</enclosure>`

	var u Model

	decoder := xml.NewDecoder(strings.NewReader(input))
	decoder.DefaultSpace = "https://www.rssboard.org/rss-specification"

	decoder.Decode(&u)
	helper(u)

	//xmlbytes, _ := xml.MarshalIndent(&u, "", "\t")

	if u.Field != "not namespaced" || u.FieldTwo != "namespaced" {
		t.Errorf("wtf")
	}

	buf := new(bytes.Buffer)

	encoder := xml.NewEncoder(buf)
	encoder.Encode(&u)

	helper2(buf.Bytes())

}
