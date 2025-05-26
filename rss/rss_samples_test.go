package rss

import (
	_ "embed"
	"encoding/xml"
)

// SpecSampleXML taken from https://www.rssboard.org/files/sample-rss-2.xml
//
//go:embed samples/sample-rss-2.xml
var SpecSampleXML []byte

var SpecSampleExpected = RSS{
	XMLName: xml.Name{
		Space: "https://www.rssboard.org/rss-specification",
		Local: "rss",
	},
	Channel: Channel{
		XMLName: xml.Name{
			Space: "https://www.rssboard.org/rss-specification",
			Local: "channel",
		},
		Title: "NASA Space Station News",
		Link:  "http://www.nasa.gov/",
		Description: &Description{
			XMLName: xml.Name{
				Space: "https://www.rssboard.org/rss-specification",
				Local: "description",
			},
			Value: "A RSS news feed containing the latest NASA press releases on the International Space Station.",
		},
		Language:       "en-us",
		Copyright:      "",
		ManagingEditor: "neil.armstrong@example.com (Neil Armstrong)",
		WebMaster:      "sally.ride@example.com (Sally Ride)",
		PubDate:        "Tue, 10 Jun 2003 04:00:00 GMT",
		LastBuildDate:  "Fri, 21 Jul 2023 09:04 EDT",
		Categories:     nil,
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
					Space: "https://www.rssboard.org/rss-specification",
					Local: "item",
				},
				Title: "Louisiana Students to Hear from NASA Astronauts Aboard Space Station",
				Link:  "http://www.nasa.gov/press-release/louisiana-students-to-hear-from-nasa-astronauts-aboard-space-station",
				Description: &Description{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "description",
					},
					Value: "As part of the state's first Earth-to-space call, students from Louisiana will have an opportunity soon to hear from NASA astronauts aboard the International Space Station.",
				},
				Author:     "",
				Categories: nil,
				Comments:   "",
				Enclosure:  nil,
				GUID: &GUID{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "guid",
					},
					Value:       "http://www.nasa.gov/press-release/louisiana-students-to-hear-from-nasa-astronauts-aboard-space-station",
					IsPermaLink: nil,
				},
				PubDate: "Fri, 21 Jul 2023 09:04 EDT",
				Source:  nil,
			},
			Item{
				XMLName: xml.Name{
					Space: "https://www.rssboard.org/rss-specification",
					Local: "item",
				},
				Title: "",
				Link:  "http://www.nasa.gov/press-release/nasa-awards-integrated-mission-operations-contract-iii",
				Description: &Description{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "description",
					},
					Value: "NASA has selected KBR Wyle Services, LLC, of Fulton, Maryland, to provide mission and flight crew operations support for the International Space Station and future human space exploration.",
				},
				Author:     "",
				Categories: nil,
				Comments:   "",
				Enclosure:  nil,
				GUID: &GUID{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "guid",
					},
					Value:       "http://www.nasa.gov/press-release/nasa-awards-integrated-mission-operations-contract-iii",
					IsPermaLink: nil,
				},
				PubDate: "Thu, 20 Jul 2023 15:05 EDT",
				Source:  nil,
			},
			Item{
				XMLName: xml.Name{
					Space: "https://www.rssboard.org/rss-specification",
					Local: "item",
				},
				Title: "NASA Expands Options for Spacewalking, Moonwalking Suits",
				Link:  "http://www.nasa.gov/press-release/nasa-expands-options-for-spacewalking-moonwalking-suits-services",
				Description: &Description{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "description",
					},
					Value: "NASA has awarded Axiom Space and Collins Aerospace task orders under existing contracts to advance spacewalking capabilities in low Earth orbit, as well as moonwalking services for Artemis missions.",
				},
				Author:     "",
				Categories: nil,
				Comments:   "",
				Enclosure: &Enclosure{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "enclosure",
					},
					URL:    "http://www.nasa.gov/sites/default/files/styles/1x1_cardfeed/public/thumbnails/image/iss068e027836orig.jpg?itok=ucNUaaGx",
					Length: 1032272,
					Type:   "image/jpeg",
				},
				GUID: &GUID{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "guid",
					},
					Value:       "http://www.nasa.gov/press-release/nasa-expands-options-for-spacewalking-moonwalking-suits-services",
					IsPermaLink: nil,
				},
				PubDate: "Mon, 10 Jul 2023 14:14 EDT",
				Source:  nil,
			},
			Item{
				XMLName: xml.Name{
					Space: "https://www.rssboard.org/rss-specification",
					Local: "item",
				},
				Title: "NASA to Provide Coverage as Dragon Departs Station",
				Link:  "http://www.nasa.gov/press-release/nasa-to-provide-coverage-as-dragon-departs-station-with-science",
				Description: &Description{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "description",
					},
					Value: "NASA is set to receive scientific research samples and hardware as a SpaceX Dragon cargo resupply spacecraft departs the International Space Station on Thursday, June 29.",
				},
				Author:     "",
				Categories: nil,
				Comments:   "",
				Enclosure:  nil,
				GUID: &GUID{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "guid",
					},
					Value:       "http://www.nasa.gov/press-release/nasa-to-provide-coverage-as-dragon-departs-station-with-science",
					IsPermaLink: nil,
				},
				PubDate: "Tue, 20 May 2003 08:56:02 GMT",
				Source:  nil,
			},
			Item{
				XMLName: xml.Name{
					Space: "https://www.rssboard.org/rss-specification",
					Local: "item",
				},
				Title: "NASA Plans Coverage of Roscosmos Spacewalk Outside Space Station",
				Link:  "http://liftoff.msfc.nasa.gov/news/2003/news-laundry.asp",
				Description: &Description{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "description",
					},
					Value: "Compared to earlier spacecraft, the International Space Station has many luxuries, but laundry facilities are not one of them.  Instead, astronauts have other options.",
				},
				Author:     "",
				Categories: nil,
				Comments:   "",
				Enclosure: &Enclosure{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "enclosure",
					},
					URL:    "http://www.nasa.gov/sites/default/files/styles/1x1_cardfeed/public/thumbnails/image/spacex_dragon_june_29.jpg?itok=nIYlBLme",
					Length: 269866,
					Type:   "image/jpeg",
				},
				GUID: &GUID{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "guid",
					},
					Value:       "http://liftoff.msfc.nasa.gov/2003/05/20.html#item570",
					IsPermaLink: nil,
				},
				PubDate: "Mon, 26 Jun 2023 12:45 EDT",
				Source:  nil,
			},
		},
	},
	Version: "2.0",
}

//go:embed samples/more-complex-sample.xml
var MoreComplexSample []byte

var MoreComplexSampleExpected = RSS{
	Channel: Channel{
		XMLName: xml.Name{
			Space: "https://www.rssboard.org/rss-specification",
			Local: "channel",
		},
		Title: "Epic Podcast",
		Link:  "http://www.yodaspin.com",
		Description: &Description{
			XMLName: xml.Name{
				Space: "https://www.rssboard.org/rss-specification",
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
					Space: "https://www.rssboard.org/rss-specification",
					Local: "category",
				},
				Value:  "bad",
				Domain: "",
			}, Category{
				XMLName: xml.Name{
					Space: "https://www.rssboard.org/rss-specification",
					Local: "category",
				},
				Value:  "good",
				Domain: "",
			}, Category{
				XMLName: xml.Name{
					Space: "https://www.rssboard.org/rss-specification",
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
				Space: "https://www.rssboard.org/rss-specification",
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
				Space: "https://www.rssboard.org/rss-specification",
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
				Space: "https://www.rssboard.org/rss-specification",
				Local: "textInput",
			},
			Title:       "text input title",
			Description: "description of the text input",
			Name:        "name of the text input",
			Link:        "link of the text input",
		},
		SkipHours: &SkipHours{
			XMLName: xml.Name{
				Space: "https://www.rssboard.org/rss-specification",
				Local: "skipHours",
			},
			Hours: []string{
				"1", "4", "9",
			},
		},
		SkipDays: &SkipDays{
			XMLName: xml.Name{
				Space: "https://www.rssboard.org/rss-specification",
				Local: "skipDays",
			},
			Days: []SkipDay{
				SkipDay("Tuesday"), SkipDay("Saturday"),
			},
		},
		Items: []Item{
			Item{
				XMLName: xml.Name{
					Space: "https://www.rssboard.org/rss-specification",
					Local: "item",
				},
				Title: "episode 1",
				Link:  "https://contoso.com/episode1",
				Description: &Description{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "description",
					},
					Value: "<a href=\"www.starwars.jayd.ml\">test</a>",
				},
				Author: "bob@consoto.com",
				Categories: []Category{
					Category{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "category",
						},
						Value:  "bad",
						Domain: "",
					}, Category{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "category",
						},
						Value:  "good",
						Domain: "",
					}, Category{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "category",
						},
						Value:  "this one has a domain",
						Domain: "https://constoso.com",
					},
				},
				Comments: "",
				Enclosure: &Enclosure{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "enclosure",
					},
					URL:    "https://contoso.com/url",
					Length: 117,
					Type:   "audio/x-midi",
				},
				GUID: &GUID{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "guid",
					},
					Value:       "guid-1",
					IsPermaLink: &TRUE,
				},
				PubDate: RFC2822Date("Fri, 21 Jul 2023 09:04 EDT"),
				Source: &Source{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "source",
					},
					Value: "https://stuff.com",
					URL:   "",
				},
			}, Item{
				XMLName: xml.Name{
					Space: "https://www.rssboard.org/rss-specification",
					Local: "item",
				},
				Title: "",
				Link:  "",
				Description: &Description{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
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
						Space: "https://www.rssboard.org/rss-specification",
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
	XMLName: xml.Name{
		Space: "https://www.rssboard.org/rss-specification",
		Local: "rss",
	},
	Version: "2.0",
}

var TRUE = true
var FALSE = false
