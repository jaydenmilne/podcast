package rss

import "encoding/xml"

// SpecSampleXML taken from https://www.rssboard.org/files/sample-rss-2.xml
const SpecSampleXML = `
<?xml version="1.0"?>
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
   <channel>
      <title>NASA Space Station News</title>
      <link>http://www.nasa.gov/</link>
      <description>A RSS news feed containing the latest NASA press releases on the International Space Station.</description>
      <language>en-us</language>
      <pubDate>Tue, 10 Jun 2003 04:00:00 GMT</pubDate>
      <lastBuildDate>Fri, 21 Jul 2023 09:04 EDT</lastBuildDate>
      <docs>https://www.rssboard.org/rss-specification</docs>
      <generator>Blosxom 2.1.2</generator>
      <managingEditor>neil.armstrong@example.com (Neil Armstrong)</managingEditor>
      <webMaster>sally.ride@example.com (Sally Ride)</webMaster>
      <atom:link href="https://www.rssboard.org/files/sample-rss-2.xml" rel="self" type="application/rss+xml" />
      <item>
         <title>Louisiana Students to Hear from NASA Astronauts Aboard Space Station</title>
         <link>http://www.nasa.gov/press-release/louisiana-students-to-hear-from-nasa-astronauts-aboard-space-station</link>
         <description>As part of the state's first Earth-to-space call, students from Louisiana will have an opportunity soon to hear from NASA astronauts aboard the International Space Station.</description>
         <pubDate>Fri, 21 Jul 2023 09:04 EDT</pubDate>
         <guid>http://www.nasa.gov/press-release/louisiana-students-to-hear-from-nasa-astronauts-aboard-space-station</guid>
      </item>
      <item>
         <description>NASA has selected KBR Wyle Services, LLC, of Fulton, Maryland, to provide mission and flight crew operations support for the International Space Station and future human space exploration.</description>
         <link>http://www.nasa.gov/press-release/nasa-awards-integrated-mission-operations-contract-iii</link>
         <pubDate>Thu, 20 Jul 2023 15:05 EDT</pubDate>
         <guid>http://www.nasa.gov/press-release/nasa-awards-integrated-mission-operations-contract-iii</guid>
      </item>
      <item>
         <title>NASA Expands Options for Spacewalking, Moonwalking Suits</title>
         <link>http://www.nasa.gov/press-release/nasa-expands-options-for-spacewalking-moonwalking-suits-services</link>
         <description>NASA has awarded Axiom Space and Collins Aerospace task orders under existing contracts to advance spacewalking capabilities in low Earth orbit, as well as moonwalking services for Artemis missions.</description>
         <enclosure url="http://www.nasa.gov/sites/default/files/styles/1x1_cardfeed/public/thumbnails/image/iss068e027836orig.jpg?itok=ucNUaaGx" length="1032272" type="image/jpeg" />
         <pubDate>Mon, 10 Jul 2023 14:14 EDT</pubDate>
         <guid>http://www.nasa.gov/press-release/nasa-expands-options-for-spacewalking-moonwalking-suits-services</guid>
      </item>
      <item>
         <title>NASA to Provide Coverage as Dragon Departs Station</title>
         <link>http://www.nasa.gov/press-release/nasa-to-provide-coverage-as-dragon-departs-station-with-science</link>
         <description>NASA is set to receive scientific research samples and hardware as a SpaceX Dragon cargo resupply spacecraft departs the International Space Station on Thursday, June 29.</description>
         <pubDate>Tue, 20 May 2003 08:56:02 GMT</pubDate>
         <guid>http://www.nasa.gov/press-release/nasa-to-provide-coverage-as-dragon-departs-station-with-science</guid>
      </item>
      <item>
         <title>NASA Plans Coverage of Roscosmos Spacewalk Outside Space Station</title>
         <link>http://liftoff.msfc.nasa.gov/news/2003/news-laundry.asp</link>
         <description>Compared to earlier spacecraft, the International Space Station has many luxuries, but laundry facilities are not one of them.  Instead, astronauts have other options.</description>
         <enclosure url="http://www.nasa.gov/sites/default/files/styles/1x1_cardfeed/public/thumbnails/image/spacex_dragon_june_29.jpg?itok=nIYlBLme" length="269866" type="image/jpeg" />
         <pubDate>Mon, 26 Jun 2023 12:45 EDT</pubDate>
         <guid>http://liftoff.msfc.nasa.gov/2003/05/20.html#item570</guid>
      </item>
   </channel>
</rss>
`

var SpecSampleExpected = RSS{
	Channel: Channel{
		XMLName: xml.Name{
			Space: "",
			Local: "channel",
		},
		Title: "NASA Space Station News",
		Link:  "",
		Description: &Description{
			XMLName: xml.Name{
				Space: "",
				Local: "description",
			},
			Value: "A RSS news feed containing the latest NASA press releases on the International Space Station.",
		},
		Language:       "en-us",
		Copyright:      "",
		ManagingEditor: "neil.armstrong@example.com (Neil Armstrong)",
		WebMaster:      "sally.ride@example.com (Sally Ride)",
		PubDate:        RFC2822Date("Tue, 10 Jun 2003 04:00:00 GMT"),
		LastBuildDate:  RFC2822Date("Fri, 21 Jul 2023 09:04 EDT"),
		Categories:     []Category(nil),
		Generator:      "Blosxom 2.1.2",
		Docs:           "https://www.rssboard.org/rss-specification",
		Cloud:          nil,
		TTL:            0,
		Image:          nil,
		Rating:         "",
		TextInput:      nil,
		SkipHours:      nil,
		SkipDays:       nil,
		Items: []Item{
			Item{
				XMLName: xml.Name{
					Space: "",
					Local: "item",
				},
				Title: "Louisiana Students to Hear from NASA Astronauts Aboard Space Station",
				Link:  "http://www.nasa.gov/press-release/louisiana-students-to-hear-from-nasa-astronauts-aboard-space-station",
				Description: &Description{
					XMLName: xml.Name{
						Space: "",
						Local: "description",
					},
					Value: "As part of the state's first Earth-to-space call, students from Louisiana will have an opportunity soon to hear from NASA astronauts aboard the International Space Station.",
				},
				Author:     "",
				Categories: []Category(nil),
				Comments:   "",
				Enclosure:  nil,
				GUID: &GUID{
					XMLName: xml.Name{
						Space: "",
						Local: "guid",
					},
					Value:       "http://www.nasa.gov/press-release/louisiana-students-to-hear-from-nasa-astronauts-aboard-space-station",
					IsPermaLink: nil,
				},
				PubDate: RFC2822Date("Fri, 21 Jul 2023 09:04 EDT"),
				Source:  nil,
			}, Item{
				XMLName: xml.Name{
					Space: "",
					Local: "item",
				},
				Title: "",
				Link:  "http://www.nasa.gov/press-release/nasa-awards-integrated-mission-operations-contract-iii",
				Description: &Description{
					XMLName: xml.Name{
						Space: "",
						Local: "description",
					},
					Value: "NASA has selected KBR Wyle Services, LLC, of Fulton, Maryland, to provide mission and flight crew operations support for the International Space Station and future human space exploration.",
				},
				Author:     "",
				Categories: []Category(nil),
				Comments:   "",
				Enclosure:  nil,
				GUID: &GUID{
					XMLName: xml.Name{
						Space: "",
						Local: "guid",
					},
					Value:       "http://www.nasa.gov/press-release/nasa-awards-integrated-mission-operations-contract-iii",
					IsPermaLink: nil,
				},
				PubDate: RFC2822Date("Thu, 20 Jul 2023 15:05 EDT"),
				Source:  nil,
			}, Item{
				XMLName: xml.Name{
					Space: "",
					Local: "item",
				},
				Title: "NASA Expands Options for Spacewalking, Moonwalking Suits",
				Link:  "http://www.nasa.gov/press-release/nasa-expands-options-for-spacewalking-moonwalking-suits-services",
				Description: &Description{
					XMLName: xml.Name{
						Space: "",
						Local: "description",
					},
					Value: "NASA has awarded Axiom Space and Collins Aerospace task orders under existing contracts to advance spacewalking capabilities in low Earth orbit, as well as moonwalking services for Artemis missions.",
				},
				Author:     "",
				Categories: []Category(nil),
				Comments:   "",
				Enclosure: &Enclosure{
					URL:    "",
					Length: 0,
					Type:   "",
				},
				GUID: &GUID{
					XMLName: xml.Name{
						Space: "",
						Local: "guid",
					},
					Value:       "http://www.nasa.gov/press-release/nasa-expands-options-for-spacewalking-moonwalking-suits-services",
					IsPermaLink: nil,
				},
				PubDate: RFC2822Date("Mon, 10 Jul 2023 14:14 EDT"),
				Source:  nil,
			}, Item{
				XMLName: xml.Name{
					Space: "",
					Local: "item",
				},
				Title: "NASA to Provide Coverage as Dragon Departs Station",
				Link:  "http://www.nasa.gov/press-release/nasa-to-provide-coverage-as-dragon-departs-station-with-science",
				Description: &Description{
					XMLName: xml.Name{
						Space: "",
						Local: "description",
					},
					Value: "NASA is set to receive scientific research samples and hardware as a SpaceX Dragon cargo resupply spacecraft departs the International Space Station on Thursday, June 29.",
				},
				Author:     "",
				Categories: []Category(nil),
				Comments:   "",
				Enclosure:  nil,
				GUID: &GUID{
					XMLName: xml.Name{
						Space: "",
						Local: "guid",
					},
					Value:       "http://www.nasa.gov/press-release/nasa-to-provide-coverage-as-dragon-departs-station-with-science",
					IsPermaLink: nil,
				},
				PubDate: RFC2822Date("Tue, 20 May 2003 08:56:02 GMT"),
				Source:  nil,
			}, Item{
				XMLName: xml.Name{
					Space: "",
					Local: "item",
				},
				Title: "NASA Plans Coverage of Roscosmos Spacewalk Outside Space Station",
				Link:  "http://liftoff.msfc.nasa.gov/news/2003/news-laundry.asp",
				Description: &Description{
					XMLName: xml.Name{
						Space: "",
						Local: "description",
					},
					Value: "Compared to earlier spacecraft, the International Space Station has many luxuries, but laundry facilities are not one of them.  Instead, astronauts have other options.",
				},
				Author:     "",
				Categories: []Category(nil),
				Comments:   "",
				Enclosure: &Enclosure{
					URL:    "",
					Length: 0,
					Type:   "",
				},
				GUID: &GUID{
					XMLName: xml.Name{
						Space: "",
						Local: "guid",
					},
					Value:       "http://liftoff.msfc.nasa.gov/2003/05/20.html#item570",
					IsPermaLink: nil,
				},
				PubDate: RFC2822Date("Mon, 26 Jun 2023 12:45 EDT"),
				Source:  nil,
			},
		},
	},
	Version: "2.0",
}

const MoreComplexSample = `
<?xml version="1.0"?>
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
   <channel>
      <title>Epic Podcast</title>
      <link>http://www.yodaspin.com</link>
      <description><![CDATA[<a href="www.starwars.jayd.ml">test</a>]]></description>
      <copyright>(c) some guy</copyright>
      <managingEditor>bob@contoso.com</managingEditor>
      <webMaster>steve@contoso.com</webMaster>
      <language>en-us</language>
      <pubDate>Tue, 10 Jun 2003 04:00:00 GMT</pubDate>
      <lastBuildDate>Fri, 21 Jul 2023 09:04 EDT</lastBuildDate>

      <category>bad</category>
      <category>good</category>
      <category domain="https://constoso.com">this one has a domain</category>
      <ttl>0118999</ttl>
      <image>
        <url>https://contoso.com/asdf.gif</url>
        <title>My Epic Picture</title>
        <link>https://contoso.com</link>
        <width>123456</width>
        <width>1234567</width>
        <description>some epic logo idk</description>
      </image>
      <rating>what even is this pics stuff</rating>
      <textInput>
        <title>text input title</title>
        <description>description of the text input</description>
        <name>name of the text input</name>
        <link>link of the text input</link>
      </textInput>
      <skipHours>
        <hour>1</hour>
        <hour>4</hour>
        <hour>9</hour>
      </skipHours>
      <skipDays>
        <day>Tuesday</day>
        <day>Saturday</day>
      </skipDays>
      <generator>by hand, the way you're supposed to</generator>
      <cloud domain="consoto.com" port="12345" path="/some/location" registerProcedure="what even is this 2000s rpc crap" procotol="soap"/>
      <docs>https://www.rssboard.org/rss-specification</docs>
      <item>
         <title>episode 1</title>
         <link>https://contoso.com/episode1</link>
         <description><![CDATA[<a href="www.starwars.jayd.ml">test</a>]]></description>
         <author>bob@consoto.com</author>
         <category>bad</category>
         <category>good</category>
         <category domain="https://constoso.com">this one has a domain</category>
         <comments>https://google.com</comments>
         <enclosure>
            <url>https://contoso.com/url</url>
            <length>117</length>
            <type>audio/x-midi</type>
         </enclosure>
         <guid isPermaLink="true">guid-1</guid>
         <pubDate>Fri, 21 Jul 2023 09:04 EDT</pubDate>
         <source url="has_url">https://stuff.com</source>
      </item>
      <item>
         <description>this one has no title but is explicitly not a permalink</description>
         <guid isPermaLink="false">link.com</guid>
      </item>
   </channel>
</rss>
`

var MoreComplexSampleExpected = RSS{
	Channel: Channel{
		XMLName: xml.Name{
			Space: "",
			Local: "channel",
		},
		Title: "Epic Podcast",
		Link:  "http://www.yodaspin.com",
		Description: &Description{
			XMLName: xml.Name{
				Space: "",
				Local: "description",
			},
			Value: "<a href=\"www.starwars.jayd.ml\">test</a>",
		},
		Language:       "en-us",
		Copyright:      "(c) some guy",
		ManagingEditor: "bob@contoso.com",
		WebMaster:      "steve@contoso.com",
		PubDate:        RFC2822Date("Tue, 10 Jun 2003 04:00:00 GMT"),
		LastBuildDate:  RFC2822Date("Fri, 21 Jul 2023 09:04 EDT"),
		Categories: []Category{
			Category{
				XMLName: xml.Name{
					Space: "",
					Local: "category",
				},
				Value:  "bad",
				Domain: "",
			}, Category{
				XMLName: xml.Name{
					Space: "",
					Local: "category",
				},
				Value:  "good",
				Domain: "",
			}, Category{
				XMLName: xml.Name{
					Space: "",
					Local: "category",
				},
				Value:  "this one has a domain",
				Domain: "https://constoso.com",
			},
		},
		Generator: "by hand, the way you're supposed to",
		Docs:      "https://www.rssboard.org/rss-specification",
		Cloud: &Cloud{
			XMLName: xml.Name{
				Space: "",
				Local: "cloud",
			},
			Domain:            "consoto.com",
			Port:              "12345",
			Path:              "/some/location",
			RegisterProcedure: "what even is this 2000s rpc crap",
			Protocol:          CloudProtocol(""),
		},
		TTL: 118999,
		Image: &Image{
			XMLName: xml.Name{
				Space: "",
				Local: "image",
			},
			URL:         "https://contoso.com/asdf.gif",
			Title:       "My Epic Picture",
			Link:        "https://contoso.com",
			Width:       1234567,
			Height:      0,
			Description: "some epic logo idk",
		},
		Rating: "what even is this pics stuff",
		TextInput: &TextInput{
			XMLName: xml.Name{
				Space: "",
				Local: "textInput",
			},
			Title:       "text input title",
			Description: "description of the text input",
			Name:        "name of the text input",
			Link:        "link of the text input",
		},
		SkipHours: &SkipHours{
			XMLName: xml.Name{
				Space: "",
				Local: "skipHours",
			},
			Hours: []string{
				"1", "4", "9",
			},
		},
		SkipDays: &SkipDays{
			XMLName: xml.Name{
				Space: "",
				Local: "skipDays",
			},
			Days: []SkipDay{
				SkipDay("Tuesday"), SkipDay("Saturday"),
			},
		},
		Items: []Item{
			Item{
				XMLName: xml.Name{
					Space: "",
					Local: "item",
				},
				Title: "episode 1",
				Link:  "https://contoso.com/episode1",
				Description: &Description{
					XMLName: xml.Name{
						Space: "",
						Local: "description",
					},
					Value: "<a href=\"www.starwars.jayd.ml\">test</a>",
				},
				Author: "bob@consoto.com",
				Categories: []Category{
					Category{
						XMLName: xml.Name{
							Space: "",
							Local: "category",
						},
						Value:  "bad",
						Domain: "",
					}, Category{
						XMLName: xml.Name{
							Space: "",
							Local: "category",
						},
						Value:  "good",
						Domain: "",
					}, Category{
						XMLName: xml.Name{
							Space: "",
							Local: "category",
						},
						Value:  "this one has a domain",
						Domain: "https://constoso.com",
					},
				},
				Comments: "",
				Enclosure: &Enclosure{
					URL:    "https://contoso.com/url",
					Length: 117,
					Type:   "audio/x-midi",
				},
				GUID: &GUID{
					XMLName: xml.Name{
						Space: "",
						Local: "guid",
					},
					Value:       "guid-1",
					IsPermaLink: &TRUE,
				},
				PubDate: RFC2822Date("Fri, 21 Jul 2023 09:04 EDT"),
				Source: &Source{
					XMLName: xml.Name{
						Space: "",
						Local: "source",
					},
					Value: "https://stuff.com",
					URL:   "",
				},
			}, Item{
				XMLName: xml.Name{
					Space: "",
					Local: "item",
				},
				Title: "",
				Link:  "",
				Description: &Description{
					XMLName: xml.Name{
						Space: "",
						Local: "description",
					},
					Value: "this one has no title but is explicitly not a permalink",
				},
				Author:     "",
				Categories: []Category(nil),
				Comments:   "",
				Enclosure:  nil,
				GUID: &GUID{
					XMLName: xml.Name{
						Space: "",
						Local: "guid",
					},
					Value:       "link.com",
					IsPermaLink: &FALSE,
				},
				PubDate: RFC2822Date(""),
				Source:  nil,
			},
		},
	},
	Version: "2.0",
}

var TRUE = true
var FALSE = false
