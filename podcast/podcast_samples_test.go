package podcast

import (
	_ "embed"
	"encoding/xml"

	"github.com/jaydenm/rssfeed/rss"
)

//go:embed samples/apple.rss
var ApplePodcastSample []byte

var f bool = false
var t bool = true

var ApplePodcastSampleExpected RSSPodcast = RSSPodcast{
	XMLName: xml.Name{
		Space: "https://www.rssboard.org/rss-specification",
		Local: "rss",
	},
	Channel: Podcast{
		Channel: rss.Channel{
			XMLName: xml.Name{
				Space: "",
				Local: "",
			},
			Title: "Hiking Treks",
			Link:  "https://www.apple.com/itunes/podcasts/",
			Description: &rss.Description{
				XMLName: xml.Name{
					Space: "https://www.rssboard.org/rss-specification",
					Local: "description",
				},
				Value: "\n      Love to get outdoors and discover nature's treasures? Hiking Treks is the\n      show for you. We review hikes and excursions, review outdoor gear and interview\n      a variety of naturalists and adventurers. Look for new episodes each week.\n    ",
			},
			Language:       "en-us",
			Copyright:      "© 2020 John Appleseed",
			ManagingEditor: "",
			WebMaster:      "",
			PubDate:        "",
			LastBuildDate:  "",
			Categories:     nil,
			Generator:      "",
			Docs:           "",
			Cloud:          nil,
			TTL:            0,
			Image:          nil,
			Rating:         "",
			TextInput:      nil,
			SkipHours:      nil,
			SkipDays:       nil,
			Items:          nil,
		},
		ItunesImage: ItunesImageTag{
			XMLName: xml.Name{
				Space: "http://www.itunes.com/dtds/podcast-1.0.dtd",
				Local: "image",
			},
			Href: "https://applehosted.podcasts.apple.com/hiking_treks/artwork.png",
		},
		ItunesCategory: []ItunesCategory{
			ItunesCategory{
				XMLName: xml.Name{
					Space: "http://www.itunes.com/dtds/podcast-1.0.dtd",
					Local: "category",
				},
				SubCategory: &struct {
					XMLName xml.Name "xml:\"http://www.itunes.com/dtds/podcast-1.0.dtd category\""
					Text    string   "xml:\"text,attr\""
				}{
					XMLName: xml.Name{
						Space: "http://www.itunes.com/dtds/podcast-1.0.dtd",
						Local: "category",
					},
					Text: "Wilderness",
				},
				Text: "Sports",
			},
		},
		ItunesExplicit:           &f,
		ItunesAuthor:             "The Sunset Explorers",
		ItunesTitle:              "",
		ItunesType:               "serial",
		ItunesNewFeedURL:         "",
		ItunesComplete:           "",
		ItunesBlock:              "",
		ItunesApplePodcastVerify: "",
		PodcastTxt:               nil,
		Items: []Episode{
			Episode{
				Item: rss.Item{
					XMLName: xml.Name{
						Space: "",
						Local: "",
					},
					Title: "",
					Link:  "",
					Description: &rss.Description{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "description",
						},
						Value: "\n          The Sunset Explorers share tips, techniques and recommendations for\n          great hikes and adventures around the United States. Listen on \n          <a href=\"https://www.apple.com/itunes/podcasts/\">Apple Podcasts</a>.\n      ",
					},
					Author:     "",
					Categories: nil,
					Comments:   "",
					Enclosure: &rss.Enclosure{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "enclosure",
						},
						URL:    "http://example.com/podcasts/everything/AllAboutEverythingEpisode4.mp3",
						Length: 498537,
						Type:   "audio/mpeg",
					},
					GUID: &rss.GUID{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "guid",
						},
						Value:       "D03EEC9B-B1B4-475B-92C8-54F853FA2A22",
						IsPermaLink: nil,
					},
					PubDate: "Tue, 8 Jan 2019 01:15:00 GMT",
					Source:  nil,
				},
				ItunesDuration:    "1079",
				ItunesImage:       nil,
				ItunesExplicit:    &f,
				ItunesTitle:       "Hiking Treks Trailer",
				ItunesEpisode:     0,
				ItunesSeason:      0,
				ItunesEpisodeType: "trailer",
				PodcastTranscript: nil,
				ItunesBlock:       "",
			},
			Episode{
				Item: rss.Item{
					XMLName: xml.Name{
						Space: "",
						Local: "",
					},
					Title: "S02 EP04 Mt. Hood, Oregon",
					Link:  "",
					Description: &rss.Description{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "description",
						},
						Value: "\n        Tips for trekking around the tallest mountain in Oregon\n      ",
					},
					Author:     "",
					Categories: nil,
					Comments:   "",
					Enclosure: &rss.Enclosure{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "enclosure",
						},
						URL:    "http://example.com/podcasts/everything/mthood.m4a",
						Length: 8727310,
						Type:   "audio/x-m4a",
					},
					GUID: &rss.GUID{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "guid",
						},
						Value:       "22BCFEBF-44FB-4A19-8229-7AC678629F57",
						IsPermaLink: nil,
					},
					PubDate: "Tue, 07 May 2019 12:00:00 GMT",
					Source:  nil,
				},
				ItunesDuration:    "1024",
				ItunesImage:       nil,
				ItunesExplicit:    &f,
				ItunesTitle:       "",
				ItunesEpisode:     4,
				ItunesSeason:      2,
				ItunesEpisodeType: "full",
				PodcastTranscript: nil,
				ItunesBlock:       "",
			},
			Episode{
				Item: rss.Item{
					XMLName: xml.Name{
						Space: "",
						Local: "",
					},
					Title: "S02 EP03 Bouldering Around Boulder",
					Link:  "href=\"http://example.com/podcasts/everything/",
					Description: &rss.Description{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "description",
						},
						Value: "\n        We explore fun walks to climbing areas about the beautiful Colorado city of Boulder.\n      ",
					},
					Author:     "",
					Categories: nil,
					Comments:   "",
					Enclosure: &rss.Enclosure{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "enclosure",
						},
						URL:    "http://example.com/podcasts/boulder.mp4",
						Length: 5650889,
						Type:   "video/mp4",
					},
					GUID: &rss.GUID{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "guid",
						},
						Value:       "BE486CAA-B3D5-4FB0-8298-EFEBE71C5982",
						IsPermaLink: nil,
					},
					PubDate: "Tue, 30 Apr 2019 13:00:00 EST",
					Source:  nil,
				},
				ItunesDuration: "3627",
				ItunesImage: &ItunesImageTag{
					XMLName: xml.Name{
						Space: "http://www.itunes.com/dtds/podcast-1.0.dtd",
						Local: "image",
					},
					Href: "http://example.com/podcasts/everything/AllAboutEverything/Episode2.jpg",
				},
				ItunesExplicit:    &f,
				ItunesTitle:       "",
				ItunesEpisode:     3,
				ItunesSeason:      2,
				ItunesEpisodeType: "full",
				PodcastTranscript: nil,
				ItunesBlock:       "",
			},
			Episode{
				Item: rss.Item{
					XMLName: xml.Name{
						Space: "",
						Local: "",
					},
					Title: "S02 EP02 Caribou Mountain, Maine",
					Link:  "",
					Description: &rss.Description{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "description",
						},
						Value: "\n        Put your fitness to the test with this invigorating hill climb.\n      ",
					},
					Author:     "",
					Categories: nil,
					Comments:   "",
					Enclosure: &rss.Enclosure{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "enclosure",
						},
						URL:    "http://example.com/podcasts/everything/caribou.m4v",
						Length: 5650889,
						Type:   "audio/x-m4v",
					},
					GUID: &rss.GUID{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "guid",
						},
						Value:       "142FAFE9-B1DF-4F6D-BAA8-79BDBAF653A9",
						IsPermaLink: nil,
					},
					PubDate: "Tue, 23 May 2019 02:00:00 -0700",
					Source:  nil,
				},
				ItunesDuration: "2434",
				ItunesImage: &ItunesImageTag{
					XMLName: xml.Name{
						Space: "http://www.itunes.com/dtds/podcast-1.0.dtd",
						Local: "image",
					},
					Href: "http://example.com/podcasts/everything/AllAboutEverything/Episode3.jpg",
				},
				ItunesExplicit:    &f,
				ItunesTitle:       "",
				ItunesEpisode:     2,
				ItunesSeason:      2,
				ItunesEpisodeType: "full",
				PodcastTranscript: nil,
				ItunesBlock:       "",
			},
			Episode{
				Item: rss.Item{
					XMLName: xml.Name{
						Space: "",
						Local: "",
					},
					Title: "S02 EP01 Stawamus Chief",
					Link:  "",
					Description: &rss.Description{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "description",
						},
						Value: "\n        We tackle Stawamus Chief outside of Vancouver, BC and you should too!\n      ",
					},
					Author:     "",
					Categories: nil,
					Comments:   "",
					Enclosure: &rss.Enclosure{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "enclosure",
						},
						URL:    "http://example.com/podcasts/everything/AllAboutEverythingEpisode4.mp3",
						Length: 498537,
						Type:   "audio/mpeg",
					},
					GUID: &rss.GUID{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "guid",
						},
						Value:       "5F1DBAEB-3327-49FB-ACB3-DB0158A1D0A3",
						IsPermaLink: nil,
					},
					PubDate: "2019-02-16T07:00:00.000Z",
					Source:  nil,
				},
				ItunesDuration:    "13:24",
				ItunesImage:       nil,
				ItunesExplicit:    &f,
				ItunesTitle:       "",
				ItunesEpisode:     1,
				ItunesSeason:      2,
				ItunesEpisodeType: "full",
				PodcastTranscript: nil,
				ItunesBlock:       "",
			},
			Episode{
				Item: rss.Item{
					XMLName: xml.Name{
						Space: "",
						Local: "",
					},
					Title: "S01 EP04 Kuliouou Ridge Trail",
					Link:  "",
					Description: &rss.Description{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "description",
						},
						Value: "\n        Oahu, Hawaii, has some picturesque hikes and this is one of the best!\n      ",
					},
					Author:     "",
					Categories: nil,
					Comments:   "",
					Enclosure: &rss.Enclosure{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "enclosure",
						},
						URL:    "http://example.com/podcasts/everything/AllAboutEverythingEpisode4.mp3",
						Length: 498537,
						Type:   "audio/mpeg",
					},
					GUID: &rss.GUID{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "guid",
						},
						Value:       "B5FCEB80-317C-4CD0-A84B-807065B43FB9",
						IsPermaLink: nil,
					},
					PubDate: "Tue, 27 Nov 2018 01:15:00 +0000",
					Source:  nil,
				},
				ItunesDuration:    "929",
				ItunesImage:       nil,
				ItunesExplicit:    &f,
				ItunesTitle:       "",
				ItunesEpisode:     4,
				ItunesSeason:      1,
				ItunesEpisodeType: "full",
				PodcastTranscript: nil,
				ItunesBlock:       "",
			},
			Episode{
				Item: rss.Item{
					XMLName: xml.Name{
						Space: "",
						Local: "",
					},
					Title: "S01 EP03 Blood Mountain Loop",
					Link:  "",
					Description: &rss.Description{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "description",
						},
						Value: "\n        Hiking the Appalachian Trail and Freeman Trail in Georgia\n      ",
					},
					Author:     "",
					Categories: nil,
					Comments:   "",
					Enclosure: &rss.Enclosure{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "enclosure",
						},
						URL:    "http://example.com/podcasts/everything/AllAboutEverythingEpisode4.mp3",
						Length: 498537,
						Type:   "audio/mpeg",
					},
					GUID: &rss.GUID{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "guid",
						},
						Value:       "F0C5D763-ED85-4449-9C09-81FEBDF6F126",
						IsPermaLink: nil,
					},
					PubDate: "Tue, 23 Oct 2018 01:15:00 +0000",
					Source:  nil,
				},
				ItunesDuration:    "1440",
				ItunesImage:       nil,
				ItunesExplicit:    &f,
				ItunesTitle:       "",
				ItunesEpisode:     3,
				ItunesSeason:      1,
				ItunesEpisodeType: "full",
				PodcastTranscript: nil,
				ItunesBlock:       "",
			},
			Episode{
				Item: rss.Item{
					XMLName: xml.Name{
						Space: "",
						Local: "",
					},
					Title: "S01 EP02 Garden of the Gods Wilderness",
					Link:  "",
					Description: &rss.Description{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "description",
						},
						Value: "\n        Wilderness Area Garden of the Gods in Illinois is a delightful spot for \n        an extended hike.\n      ",
					},
					Author:     "",
					Categories: nil,
					Comments:   "",
					Enclosure: &rss.Enclosure{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "enclosure",
						},
						URL:    "http://example.com/podcasts/everything/AllAboutEverythingEpisode4.mp3",
						Length: 498537,
						Type:   "audio/mpeg",
					},
					GUID: &rss.GUID{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "guid",
						},
						Value:       "821DD0B2-571D-4DFD-8E11-556E8C1EFE6A",
						IsPermaLink: nil,
					},
					PubDate: "Tue, 18 Sep 2018 01:15:00 +0000",
					Source:  nil,
				},
				ItunesDuration:    "839",
				ItunesImage:       nil,
				ItunesExplicit:    &f,
				ItunesTitle:       "",
				ItunesEpisode:     2,
				ItunesSeason:      1,
				ItunesEpisodeType: "full",
				PodcastTranscript: nil,
				ItunesBlock:       "",
			},
			Episode{
				Item: rss.Item{
					XMLName: xml.Name{
						Space: "",
						Local: "",
					},
					Title: "S01 EP01 Upper Priest Lake Trail to Continental Creek Trail",
					Link:  "",
					Description: &rss.Description{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "description",
						},
						Value: "\n        We check out this powerfully scenic hike following the river in the Idaho\n        Panhandle National Forests.\n      ",
					},
					Author:     "",
					Categories: nil,
					Comments:   "",
					Enclosure: &rss.Enclosure{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "enclosure",
						},
						URL:    "http://example.com/podcasts/everything/AllAboutEverythingEpisode4.mp3",
						Length: 498537,
						Type:   "audio/mpeg",
					},
					GUID: &rss.GUID{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "guid",
						},
						Value:       "EABDA7EE-1AC6-4B60-9E11-6B3F30B72F87",
						IsPermaLink: nil,
					},
					PubDate: "Tue, 14 Aug 2018 01:15:00 +0000",
					Source:  nil,
				},
				ItunesDuration:    "1399",
				ItunesImage:       nil,
				ItunesExplicit:    &f,
				ItunesTitle:       "",
				ItunesEpisode:     1,
				ItunesSeason:      1,
				ItunesEpisodeType: "full",
				PodcastTranscript: nil,
				ItunesBlock:       "",
			},
		},
	},
	Version: "2.0",
}

//go:embed samples/more-complex.rss
var MoreComplexSample []byte

var MoreComplexSampleExpected = RSSPodcast{
	XMLName: xml.Name{
		Space: "https://www.rssboard.org/rss-specification",
		Local: "rss",
	},
	Channel: Podcast{
		Channel: rss.Channel{
			XMLName: xml.Name{
				Space: "",
				Local: "",
			},
			Title: "Epic Podcast",
			Link:  "http://www.yodaspin.com",
			Description: &rss.Description{
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
			PubDate:        "Tue, 10 Jun 2003 04:00:00 GMT",
			LastBuildDate:  "Fri, 21 Jul 2023 09:04 EDT",
			Categories: []rss.Category{
				rss.Category{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "category",
					},
					Value:  "bad",
					Domain: "",
				},
				rss.Category{
					XMLName: xml.Name{
						Space: "https://www.rssboard.org/rss-specification",
						Local: "category",
					},
					Value:  "good",
					Domain: "",
				},
				rss.Category{
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
			Cloud: &rss.Cloud{
				XMLName: xml.Name{
					Space: "https://www.rssboard.org/rss-specification",
					Local: "cloud",
				},
				Domain:            "consoto.com",
				Port:              "12345",
				Path:              "/some/location",
				RegisterProcedure: "what even is this 2000s rpc crap",
				Protocol:          "",
			},
			TTL: 118999,
			Image: &rss.Image{
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
			TextInput: &rss.TextInput{
				XMLName: xml.Name{
					Space: "https://www.rssboard.org/rss-specification",
					Local: "textInput",
				},
				Title:       "text input title",
				Description: "description of the text input",
				Name:        "name of the text input",
				Link:        "link of the text input",
			},
			SkipHours: &rss.SkipHours{
				XMLName: xml.Name{
					Space: "https://www.rssboard.org/rss-specification",
					Local: "skipHours",
				},
				Hours: []string{
					"1",
					"4",
					"9",
				},
			},
			SkipDays: &rss.SkipDays{
				XMLName: xml.Name{
					Space: "https://www.rssboard.org/rss-specification",
					Local: "skipDays",
				},
				Days: []rss.SkipDay{
					"Tuesday",
					"Saturday",
				},
			},
			Items: nil,
		},
		ItunesImage: ItunesImageTag{
			XMLName: xml.Name{
				Space: "http://www.itunes.com/dtds/podcast-1.0.dtd",
				Local: "image",
			},
			Href: "https://example.com/image.heic",
		},
		ItunesCategory: []ItunesCategory{
			ItunesCategory{
				XMLName: xml.Name{
					Space: "http://www.itunes.com/dtds/podcast-1.0.dtd",
					Local: "category",
				},
				SubCategory: &struct {
					XMLName xml.Name "xml:\"http://www.itunes.com/dtds/podcast-1.0.dtd category\""
					Text    string   "xml:\"text,attr\""
				}{
					XMLName: xml.Name{
						Space: "http://www.itunes.com/dtds/podcast-1.0.dtd",
						Local: "category",
					},
					Text: "InnerCategory",
				},
				Text: "OuterCategory",
			},
		},
		ItunesExplicit:           &t,
		ItunesAuthor:             "the guy that made this idk",
		ItunesTitle:              "different title in itunes to be annoying",
		ItunesType:               "serial",
		ItunesNewFeedURL:         "https://new-url.com/feed.xml",
		ItunesComplete:           "Yes",
		ItunesBlock:              "Yes",
		ItunesApplePodcastVerify: "asdfjkl;",
		PodcastTxt: &PodcastTxt{
			XMLName: xml.Name{
				Space: "https://podcastindex.org/namespace/1.0",
				Local: "txt",
			},
			Value:   "123456",
			Purpose: "applepodcastverify",
		},
		Items: []Episode{
			Episode{
				Item: rss.Item{
					XMLName: xml.Name{
						Space: "",
						Local: "",
					},
					Title: "episode 1",
					Link:  "https://contoso.com/episode1",
					Description: &rss.Description{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "description",
						},
						Value: "<a href=\"www.starwars.jayd.ml\">test</a>",
					},
					Author: "bob@consoto.com",
					Categories: []rss.Category{
						rss.Category{
							XMLName: xml.Name{
								Space: "https://www.rssboard.org/rss-specification",
								Local: "category",
							},
							Value:  "bad",
							Domain: "",
						},
						rss.Category{
							XMLName: xml.Name{
								Space: "https://www.rssboard.org/rss-specification",
								Local: "category",
							},
							Value:  "good",
							Domain: "",
						},
						rss.Category{
							XMLName: xml.Name{
								Space: "https://www.rssboard.org/rss-specification",
								Local: "category",
							},
							Value:  "this one has a domain",
							Domain: "https://constoso.com",
						},
					},
					Comments: "",
					Enclosure: &rss.Enclosure{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "enclosure",
						},
						URL:    "https://contoso.com/url",
						Length: 117,
						Type:   "audio/x-midi",
					},
					GUID: &rss.GUID{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "guid",
						},
						Value:       "guid-1",
						IsPermaLink: &t,
					},
					PubDate: "Fri, 21 Jul 2023 09:04 EDT",
					Source: &rss.Source{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "source",
						},
						Value: "https://stuff.com",
						URL:   "",
					},
				},
				ItunesDuration: "234567",
				ItunesImage: &ItunesImageTag{
					XMLName: xml.Name{
						Space: "http://www.itunes.com/dtds/podcast-1.0.dtd",
						Local: "image",
					},
					Href: "https://episodeimage.com/image.bmp",
				},
				ItunesExplicit:    &f,
				ItunesTitle:       "different podcast title in itunes to be annoying",
				ItunesEpisode:     117,
				ItunesSeason:      343,
				ItunesEpisodeType: "bonus",
				PodcastTranscript: []PodcastTranscript{{
					XMLName: xml.Name{
						Space: "https://podcastindex.org/namespace/1.0",
						Local: "transcript",
					},
					URL:      "https://transcript.org/stuff.xml",
					Type:     "text/plain",
					Language: "en-us",
					Rel:      "captions",
				}},
				ItunesBlock: "Yes",
			},
			Episode{
				Item: rss.Item{
					XMLName: xml.Name{
						Space: "",
						Local: "",
					},
					Title: "",
					Link:  "",
					Description: &rss.Description{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "description",
						},
						Value: "this one has no title but is explicitly not a permalink",
					},
					Author:     "",
					Categories: nil,
					Comments:   "",
					Enclosure:  nil,
					GUID: &rss.GUID{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "guid",
						},
						Value:       "link.com",
						IsPermaLink: &f,
					},
					PubDate: "",
					Source:  nil,
				},
				ItunesDuration:    "",
				ItunesImage:       nil,
				ItunesExplicit:    nil,
				ItunesTitle:       "",
				ItunesEpisode:     0,
				ItunesSeason:      0,
				ItunesEpisodeType: "",
				PodcastTranscript: nil,
				ItunesBlock:       "",
			},
		},
	},
	Version: "2.0",
}

//go:embed samples/example.xml
var Podcasting20Example []byte

var Podcasting20ExampleExpected RSSPodcast = RSSPodcast{
	XMLName: xml.Name{
		Space: "https://www.rssboard.org/rss-specification",
		Local: "rss",
	},
	Channel: Podcast{
		Channel: rss.Channel{
			XMLName: xml.Name{
				Space: "",
				Local: "",
			},
			Title: "Podcasting 2.0 Namespace Example",
			Link:  "http://example.com/podcast",
			Description: &rss.Description{
				XMLName: xml.Name{
					Space: "https://www.rssboard.org/rss-specification",
					Local: "description",
				},
				Value: "This is a fake show that exists only as an example of the \"podcast\" namespace tag usage.",
			},
			Language:       "en-US",
			Copyright:      "",
			ManagingEditor: "john@example.com (John Doe)",
			WebMaster:      "support@example.com (Tech Support)",
			PubDate:        "Fri, 09 Oct 2020 04:30:38 GMT",
			LastBuildDate:  "Fri, 09 Oct 2020 04:30:38 GMT",
			Categories:     nil,
			Generator:      "Freedom Controller",
			Docs:           "http://blogs.law.harvard.edu/tech/rss",
			Cloud:          nil,
			TTL:            0,
			Image: &rss.Image{
				XMLName: xml.Name{
					Space: "https://www.rssboard.org/rss-specification",
					Local: "image",
				},
				URL:         "https://example.com/images/pci_avatar-massive.jpg",
				Title:       "Podcast Feed Template",
				Link:        "https://example.com/show",
				Width:       0,
				Height:      0,
				Description: "",
			},
			Rating:    "",
			TextInput: nil,
			SkipHours: nil,
			SkipDays:  nil,
			Items:     nil,
		},
		ItunesImage: ItunesImageTag{
			XMLName: xml.Name{
				Space: "http://www.itunes.com/dtds/podcast-1.0.dtd",
				Local: "image",
			},
			Href: "https://example.com/images/pci_avatar-massive.jpg",
		},
		ItunesCategory: []ItunesCategory{
			ItunesCategory{
				XMLName: xml.Name{
					Space: "http://www.itunes.com/dtds/podcast-1.0.dtd",
					Local: "category",
				},
				SubCategory: nil,
				Text:        "News",
			},
			ItunesCategory{
				XMLName: xml.Name{
					Space: "http://www.itunes.com/dtds/podcast-1.0.dtd",
					Local: "category",
				},
				SubCategory: nil,
				Text:        "Technology",
			},
		},
		ItunesExplicit:           &f,
		ItunesAuthor:             "John Doe",
		ItunesTitle:              "",
		ItunesType:               "episodic",
		ItunesNewFeedURL:         "",
		ItunesComplete:           "",
		ItunesBlock:              "",
		ItunesApplePodcastVerify: "",
		PodcastTxt:               nil,
		PodcastPodroll:           nil,
		PodcastLocked: &PodcastLocked{
			XMLName: xml.Name{
				Space: "https://podcastindex.org/namespace/1.0",
				Local: "locked",
			},
			Value: "yes",
			Owner: "podcastowner@example.com",
		},
		PodcastFunding: []PodcastFunding{
			PodcastFunding{
				XMLName: xml.Name{
					Space: "https://podcastindex.org/namespace/1.0",
					Local: "funding",
				},
				Value: "Support the show!",
				URL:   "https://example.com/donate",
			},
		},
		PodcastPeople: nil,
		PodcastLocation: &PodcastLocation{
			XMLName: xml.Name{
				Space: "https://podcastindex.org/namespace/1.0",
				Local: "location",
			},
			Text: "Austin, TX",
			Geo:  "geo:30.2672,97.7431",
			Osm:  "R113314",
		},
		PodcastTrailers: []PodcastTrailer{
			PodcastTrailer{
				XMLName: xml.Name{
					Space: "https://podcastindex.org/namespace/1.0",
					Local: "trailer",
				},
				Text:    "Coming April 1st, 2021",
				Pubdate: "Thu, 01 Apr 2021 08:00:00 EST",
				URL:     "https://example.org/trailers/teaser",
				Length:  12345678,
				Type:    "audio/mp3",
				Season:  "",
			},
		},
		PodcastValue: []PodcastValue{
			PodcastValue{
				XMLName: xml.Name{
					Space: "https://podcastindex.org/namespace/1.0",
					Local: "value",
				},
				Type:      "lightning",
				Method:    "keysend",
				Suggested: "0.00000005000",
				Recipients: []PodcastValueRecipient{
					PodcastValueRecipient{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "valueRecipient",
						},
						Name:        "podcaster",
						CustomKey:   "",
						CustomValue: "",
						Type:        "node",
						Address:     "036557ea56b3b86f08be31bcd2557cae8021b0e3a9413f0c0e52625c6696972e57",
						Split:       "99",
						Fee:         nil,
					},
					PodcastValueRecipient{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "valueRecipient",
						},
						Name:        "hosting company",
						CustomKey:   "",
						CustomValue: "",
						Type:        "node",
						Address:     "036557ea56b3b86f08be31bcd2557cae8021b0e3a9413f0c0e52625c6696972e57",
						Split:       "1",
						Fee:         nil,
					},
				},
				ValueTimeSplits: nil,
			},
		},
		PodcastMedium: "podcast",
		PodcastLiveItem: []PodcastLiveItem{
			PodcastLiveItem{
				Episode: Episode{
					Item: rss.Item{
						XMLName: xml.Name{
							Space: "",
							Local: "",
						},
						Title: "Podcasting 2.0 Live Show",
						Link:  "https://example.com/podcast/live",
						Description: &rss.Description{
							XMLName: xml.Name{
								Space: "https://www.rssboard.org/rss-specification",
								Local: "description",
							},
							Value: "A look into the future of podcasting and how we get to Podcasting 2.0!",
						},
						Author:     "",
						Categories: nil,
						Comments:   "",
						Enclosure: &rss.Enclosure{
							XMLName: xml.Name{
								Space: "https://www.rssboard.org/rss-specification",
								Local: "enclosure",
							},
							URL:    "https://example.com/pc20/livestream?format=.mp3",
							Length: 312,
							Type:   "audio/mpeg",
						},
						GUID: &rss.GUID{
							XMLName: xml.Name{
								Space: "https://www.rssboard.org/rss-specification",
								Local: "guid",
							},
							Value:       "https://example.com/live",
							IsPermaLink: &t,
						},
						PubDate: "",
						Source:  nil,
					},
					ItunesDuration:    "",
					ItunesImage:       nil,
					ItunesExplicit:    nil,
					ItunesTitle:       "",
					ItunesEpisode:     0,
					ItunesSeason:      0,
					ItunesEpisodeType: "",
					ItunesBlock:       "",
					PodcastTranscript: nil,
					PodcastChapters:   nil,
					PodcastSoundbite:  nil,
					PodcastPeople:     nil,
					PodcastSeason:     nil,
					PodcastEpisode:    nil,
					PodcastLicense:    nil,
					PodcastAlternateEnclosures: []PodcastAlternateEnclosure{
						PodcastAlternateEnclosure{
							XMLName: xml.Name{
								Space: "https://podcastindex.org/namespace/1.0",
								Local: "alternateEnclosure",
							},
							Type:             "audio/mpeg",
							Length:           "312",
							Bitrate:          0.0,
							Title:            "",
							Height:           0,
							Lang:             0,
							Rel:              0,
							Codecs:           0,
							Default:          &t,
							PodcastIntegrity: nil,
							Source: []PodcastSource{
								PodcastSource{
									XMLName: xml.Name{
										Space: "https://podcastindex.org/namespace/1.0",
										Local: "source",
									},
									URI:         "https://example.com/pc20/livestream",
									ContentType: "",
								},
							},
						},
					},
					PodcastValue:           nil,
					PodcastImages:          nil,
					PodcastSocialInteracts: nil,
					PodcastUpdateFrequency: nil,
					PodcastPodping:         nil,
				},
				XMLName: xml.Name{
					Space: "https://podcastindex.org/namespace/1.0",
					Local: "liveItem",
				},
				Status: "live",
				Start:  "2021-09-26T07:30:00.000-0600",
				End:    "2021-09-26T09:30:00.000-0600",
				PodcastContentLinks: []PodcastContentLink{
					PodcastContentLink{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "contentLink",
						},
						Value: "YouTube!",
						Href:  "https://youtube.com/pc20/livestream",
					},
					PodcastContentLink{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "contentLink",
						},
						Value: "Twitch!",
						Href:  "https://twitch.com/pc20/livestream",
					},
				},
			},
		},
		PodcastGUID: "y0ur-gu1d-g035-h3r3",
		PodcastBlock: []PodcastBlock{
			PodcastBlock{
				XMLName: xml.Name{
					Space: "https://podcastindex.org/namespace/1.0",
					Local: "block",
				},
				Value: "yes",
				ID:    "",
			},
			PodcastBlock{
				XMLName: xml.Name{
					Space: "https://podcastindex.org/namespace/1.0",
					Local: "block",
				},
				Value: "no",
				ID:    "google",
			},
			PodcastBlock{
				XMLName: xml.Name{
					Space: "https://podcastindex.org/namespace/1.0",
					Local: "block",
				},
				Value: "no",
				ID:    "amazon",
			},
		},
		PodcastLicense: &PodcastLicense{
			XMLName: xml.Name{
				Space: "https://podcastindex.org/namespace/1.0",
				Local: "license",
			},
			Value: "my-podcast-license-v1",
			URL:   "https://example.org/mypodcastlicense/full.pdf",
		},
		Items: []Episode{
			Episode{
				Item: rss.Item{
					XMLName: xml.Name{
						Space: "",
						Local: "",
					},
					Title: "Episode 3 - The Future",
					Link:  "https://example.com/podcast/ep0003",
					Description: &rss.Description{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "description",
						},
						Value: "<p>A look into the future of podcasting and how we get to Podcasting 2.0!</p>",
					},
					Author:     "John Doe (john@example.com)",
					Categories: nil,
					Comments:   "",
					Enclosure: &rss.Enclosure{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "enclosure",
						},
						URL:    "https://example.com/file-03.mp3",
						Length: 43200000,
						Type:   "audio/mpeg",
					},
					GUID: &rss.GUID{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "guid",
						},
						Value:       "https://example.com/ep0003",
						IsPermaLink: &t,
					},
					PubDate: "Fri, 09 Oct 2020 04:30:38 GMT",
					Source:  nil,
				},
				ItunesDuration: "",
				ItunesImage: &ItunesImageTag{
					XMLName: xml.Name{
						Space: "http://www.itunes.com/dtds/podcast-1.0.dtd",
						Local: "image",
					},
					Href: "https://example.com/ep0003/artMd.jpg",
				},
				ItunesExplicit:    &f,
				ItunesTitle:       "",
				ItunesEpisode:     0,
				ItunesSeason:      0,
				ItunesEpisodeType: "",
				ItunesBlock:       "",
				PodcastTranscript: []PodcastTranscript{
					PodcastTranscript{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "transcript",
						},
						URL:      "https://example.com/ep3/transcript.txt",
						Type:     "text/plain",
						Language: "",
						Rel:      "",
					},
				},
				PodcastChapters: &PodcastChapters{
					XMLName: xml.Name{
						Space: "https://podcastindex.org/namespace/1.0",
						Local: "chapters",
					},
					URL:  "https://example.com/ep3_chapters.json",
					Type: "application/json",
				},
				PodcastSoundbite: []PodcastSoundbite{
					PodcastSoundbite{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "soundbite",
						},
						StartTime: 33.833,
						Duration:  60.0,
						Value:     "",
					},
				},
				PodcastPeople: []PodcastPerson{
					PodcastPerson{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "person",
						},
						Text:  "Adam Curry",
						Group: "",
						Role:  "",
						Href:  "https://www.podchaser.com/creators/adam-curry-107ZzmWE5f",
						Img:   "https://example.com/images/adamcurry.jpg",
					},
					PodcastPerson{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "person",
						},
						Text:  "Dave Jones",
						Group: "",
						Role:  "guest",
						Href:  "https://github.com/daveajones/",
						Img:   "https://example.com/images/davejones.jpg",
					},
					PodcastPerson{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "person",
						},
						Text:  "Becky Smith",
						Group: "visuals",
						Role:  "cover art designer",
						Href:  "https://example.com/artist/beckysmith",
						Img:   "",
					},
				},
				PodcastSeason: &PodcastSeason{
					XMLName: xml.Name{
						Space: "https://podcastindex.org/namespace/1.0",
						Local: "season",
					},
					Value: 1,
					Name:  "Podcasting 2.0",
				},
				PodcastEpisode: &PodcastEpisode{
					XMLName: xml.Name{
						Space: "https://podcastindex.org/namespace/1.0",
						Local: "episode",
					},
					Value:   3.0,
					Display: "",
				},
				PodcastLicense: nil,
				PodcastAlternateEnclosures: []PodcastAlternateEnclosure{
					PodcastAlternateEnclosure{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "alternateEnclosure",
						},
						Type:             "audio/mpeg",
						Length:           "43200000",
						Bitrate:          128000.0,
						Title:            "Standard",
						Height:           0,
						Lang:             0,
						Rel:              0,
						Codecs:           0,
						Default:          &t,
						PodcastIntegrity: nil,
						Source: []PodcastSource{
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "https://example.com/file-03.mp3",
								ContentType: "",
							},
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "ipfs://someRandomMpegFile03",
								ContentType: "",
							},
						},
					},
					PodcastAlternateEnclosure{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "alternateEnclosure",
						},
						Type:             "audio/opus",
						Length:           "32400000",
						Bitrate:          96000.0,
						Title:            "High quality",
						Height:           0,
						Lang:             0,
						Rel:              0,
						Codecs:           0,
						Default:          nil,
						PodcastIntegrity: nil,
						Source: []PodcastSource{
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "https://example.com/file-high-03.opus",
								ContentType: "",
							},
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "ipfs://someRandomHighBitrateOpusFile03",
								ContentType: "",
							},
						},
					},
					PodcastAlternateEnclosure{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "alternateEnclosure",
						},
						Type:             "audio/aac",
						Length:           "54000000",
						Bitrate:          160000.0,
						Title:            "High quality AAC",
						Height:           0,
						Lang:             0,
						Rel:              0,
						Codecs:           0,
						Default:          nil,
						PodcastIntegrity: nil,
						Source: []PodcastSource{
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "https://example.com/file-proprietary-03.aac",
								ContentType: "",
							},
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "ipfs://someRandomProprietaryAACFile03",
								ContentType: "",
							},
						},
					},
					PodcastAlternateEnclosure{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "alternateEnclosure",
						},
						Type:             "audio/opus",
						Length:           "5400000",
						Bitrate:          16000.0,
						Title:            "Low bandwidth",
						Height:           0,
						Lang:             0,
						Rel:              0,
						Codecs:           0,
						Default:          nil,
						PodcastIntegrity: nil,
						Source: []PodcastSource{
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "https://example.com/file-low-03.opus",
								ContentType: "",
							},
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "ipfs://someRandomLowBitrateOpusFile03",
								ContentType: "",
							},
						},
					},
					PodcastAlternateEnclosure{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "alternateEnclosure",
						},
						Type:    "video/mp4",
						Length:  "7924786",
						Bitrate: 511276.53,
						Title:   "Video version",
						Height:  720,
						Lang:    0,
						Rel:     0,
						Codecs:  0,
						Default: nil,
						PodcastIntegrity: &PodcastIntegrity{
							XMLName: xml.Name{
								Space: "https://podcastindex.org/namespace/1.0",
								Local: "integrity",
							},
							Type:  "sri",
							Value: "sha384-ExVqijgYHm15PqQqdXfW95x+Rs6C+d6E/ICxyQOeFevnxNLR/wtJNrNYTjIysUBo",
						},
						Source: []PodcastSource{
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "https://example.com/file-720.mp4",
								ContentType: "",
							},
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "ipfs://QmX33FYehk6ckGQ6g1D9D3FqZPix5JpKstKQKbaS8quUFb",
								ContentType: "",
							},
						},
					},
				},
				PodcastValue: []PodcastValue{
					PodcastValue{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "value",
						},
						Type:      "lightning",
						Method:    "keysend",
						Suggested: "0.00000005000",
						Recipients: []PodcastValueRecipient{
							PodcastValueRecipient{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "valueRecipient",
								},
								Name:        "podcaster",
								CustomKey:   "",
								CustomValue: "",
								Type:        "node",
								Address:     "036557ea56b3b86f08be31bcd2557cae8021b0e3a9413f0c0e52625c6696972e57",
								Split:       "49",
								Fee:         nil,
							},
							PodcastValueRecipient{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "valueRecipient",
								},
								Name:        "hosting company",
								CustomKey:   "",
								CustomValue: "",
								Type:        "node",
								Address:     "036557ea56b3b86f08be31bcd2557cae8021b0e3a9413f0c0e52625c6696972e57",
								Split:       "1",
								Fee:         nil,
							},
							PodcastValueRecipient{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "valueRecipient",
								},
								Name:        "Gigi (Guest)",
								CustomKey:   "",
								CustomValue: "",
								Type:        "node",
								Address:     "02e12fea95f576a680ec1938b7ed98ef0855eadeced493566877d404e404bfbf52",
								Split:       "50",
								Fee:         nil,
							},
						},
						ValueTimeSplits: nil,
					},
				},
				PodcastImages: &PodcastImages{
					XMLName: xml.Name{
						Space: "https://podcastindex.org/namespace/1.0",
						Local: "images",
					},
					Srcset: "https://example.com/images/ep3/pci_avatar-massive.jpg 1500w,\n                                        https://example.com/images/ep3/pci_avatar-middle.jpg 600w,\n                                        https://example.com/images/ep3/pci_avatar-small.jpg 300w,\n                                        https://example.com/images/ep3/pci_avatar-tiny.jpg 150w",
				},
				PodcastSocialInteracts: []PodcastSocialInteract{
					PodcastSocialInteract{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "socialInteract",
						},
						Priority:   1,
						URI:        "https://podcastindex.social/web/@dave/108013847520053258",
						Protocol:   "activitypub",
						AccountId:  "@dave",
						AccountUrl: "https://podcastindex.social/web/@dave",
					},
					PodcastSocialInteract{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "socialInteract",
						},
						Priority:   2,
						URI:        "https://twitter.com/PodcastindexOrg/status/1507120226361647115",
						Protocol:   "twitter",
						AccountId:  "@podcastindexorg",
						AccountUrl: "https://twitter.com/PodcastindexOrg",
					},
				},
				PodcastUpdateFrequency: nil,
				PodcastPodping:         nil,
			},
			Episode{
				Item: rss.Item{
					XMLName: xml.Name{
						Space: "",
						Local: "",
					},
					Title: "Episode 2 - The Present",
					Link:  "https://example.com/podcast/ep0002",
					Description: &rss.Description{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "description",
						},
						Value: "<p>Where are we at now in the podcasting era. What are the current challenges?</p>",
					},
					Author:     "John Doe (john@example.com)",
					Categories: nil,
					Comments:   "",
					Enclosure: &rss.Enclosure{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "enclosure",
						},
						URL:    "https://example.com/file-02.mp3",
						Length: 43113000,
						Type:   "audio/mpeg",
					},
					GUID: &rss.GUID{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "guid",
						},
						Value:       "https://example.com/ep0002",
						IsPermaLink: &t,
					},
					PubDate: "Thu, 08 Oct 2020 04:30:38 GMT",
					Source:  nil,
				},
				ItunesDuration: "",
				ItunesImage: &ItunesImageTag{
					XMLName: xml.Name{
						Space: "http://www.itunes.com/dtds/podcast-1.0.dtd",
						Local: "image",
					},
					Href: "https://example.com/ep0002/artMd.jpg",
				},
				ItunesExplicit:    &f,
				ItunesTitle:       "",
				ItunesEpisode:     0,
				ItunesSeason:      0,
				ItunesEpisodeType: "",
				ItunesBlock:       "",
				PodcastTranscript: []PodcastTranscript{
					PodcastTranscript{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "transcript",
						},
						URL:      "https://example.com/ep2/transcript.txt",
						Type:     "text/plain",
						Language: "",
						Rel:      "",
					},
				},
				PodcastChapters: &PodcastChapters{
					XMLName: xml.Name{
						Space: "https://podcastindex.org/namespace/1.0",
						Local: "chapters",
					},
					URL:  "https://example.com/ep2_chapters.json",
					Type: "application/json",
				},
				PodcastSoundbite: []PodcastSoundbite{
					PodcastSoundbite{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "soundbite",
						},
						StartTime: 45.4,
						Duration:  56.0,
						Value:     "",
					},
				},
				PodcastPeople: []PodcastPerson{
					PodcastPerson{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "person",
						},
						Text:  "Adam Curry",
						Group: "",
						Role:  "",
						Href:  "https://en.wikipedia.org/wiki/Adam_Curry",
						Img:   "http://example.com/images/adamcurry.jpg",
					},
					PodcastPerson{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "person",
						},
						Text:  "Dave Jones",
						Group: "",
						Role:  "guest",
						Href:  "https://example.com/blog/daveajones/",
						Img:   "http://example.com/images/davejones.jpg",
					},
					PodcastPerson{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "person",
						},
						Text:  "Marcus Brown",
						Group: "visuals",
						Role:  "cover art designer",
						Href:  "https://example.com/artist/marcusbrown",
						Img:   "",
					},
				},
				PodcastSeason: &PodcastSeason{
					XMLName: xml.Name{
						Space: "https://podcastindex.org/namespace/1.0",
						Local: "season",
					},
					Value: 1,
					Name:  "Podcasting 2.0",
				},
				PodcastEpisode: &PodcastEpisode{
					XMLName: xml.Name{
						Space: "https://podcastindex.org/namespace/1.0",
						Local: "episode",
					},
					Value:   2.0,
					Display: "",
				},
				PodcastLicense: nil,
				PodcastAlternateEnclosures: []PodcastAlternateEnclosure{
					PodcastAlternateEnclosure{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "alternateEnclosure",
						},
						Type:             "audio/mpeg",
						Length:           "43200000",
						Bitrate:          128000.0,
						Title:            "Standard",
						Height:           0,
						Lang:             0,
						Rel:              0,
						Codecs:           0,
						Default:          &t,
						PodcastIntegrity: nil,
						Source: []PodcastSource{
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "https://example.com/file-02.mp3",
								ContentType: "",
							},
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "ipfs://someRandomMpegFile02",
								ContentType: "",
							},
						},
					},
					PodcastAlternateEnclosure{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "alternateEnclosure",
						},
						Type:             "audio/opus",
						Length:           "32400000",
						Bitrate:          96000.0,
						Title:            "High quality",
						Height:           0,
						Lang:             0,
						Rel:              0,
						Codecs:           0,
						Default:          nil,
						PodcastIntegrity: nil,
						Source: []PodcastSource{
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "https://example.com/file-high-02.opus",
								ContentType: "",
							},
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "ipfs://someRandomHighBitrateOpusFile02",
								ContentType: "",
							},
						},
					},
					PodcastAlternateEnclosure{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "alternateEnclosure",
						},
						Type:             "audio/aac",
						Length:           "54000000",
						Bitrate:          160000.0,
						Title:            "High quality AAC",
						Height:           0,
						Lang:             0,
						Rel:              0,
						Codecs:           0,
						Default:          nil,
						PodcastIntegrity: nil,
						Source: []PodcastSource{
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "https://example.com/file-proprietary-02.aac",
								ContentType: "",
							},
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "ipfs://someRandomProprietaryAACFile02",
								ContentType: "",
							},
						},
					},
					PodcastAlternateEnclosure{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "alternateEnclosure",
						},
						Type:             "audio/opus",
						Length:           "5400000",
						Bitrate:          16000.0,
						Title:            "Low bandwidth",
						Height:           0,
						Lang:             0,
						Rel:              0,
						Codecs:           0,
						Default:          nil,
						PodcastIntegrity: nil,
						Source: []PodcastSource{
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "https://example.com/file-low-02.opus",
								ContentType: "",
							},
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "ipfs://someRandomLowBitrateOpusFile02",
								ContentType: "",
							},
						},
					},
				},
				PodcastValue: []PodcastValue{
					PodcastValue{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "value",
						},
						Type:      "lightning",
						Method:    "keysend",
						Suggested: "0.00000005000",
						Recipients: []PodcastValueRecipient{
							PodcastValueRecipient{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "valueRecipient",
								},
								Name:        "podcaster",
								CustomKey:   "",
								CustomValue: "",
								Type:        "node",
								Address:     "036557ea56b3b86f08be31bcd2557cae8021b0e3a9413f0c0e52625c6696972e57",
								Split:       "49",
								Fee:         nil,
							},
							PodcastValueRecipient{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "valueRecipient",
								},
								Name:        "hosting company",
								CustomKey:   "",
								CustomValue: "",
								Type:        "node",
								Address:     "036557ea56b3b86f08be31bcd2557cae8021b0e3a9413f0c0e52625c6696972e57",
								Split:       "1",
								Fee:         nil,
							},
							PodcastValueRecipient{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "valueRecipient",
								},
								Name:        "Paul Itoi (Guest)",
								CustomKey:   "",
								CustomValue: "",
								Type:        "node",
								Address:     "03a9a8d953fe747d0dd94dd3c567ddc58451101e987e2d2bf7a4d1e10a2c89ff38",
								Split:       "50",
								Fee:         nil,
							},
						},
						ValueTimeSplits: nil,
					},
				},
				PodcastImages: &PodcastImages{
					XMLName: xml.Name{
						Space: "https://podcastindex.org/namespace/1.0",
						Local: "images",
					},
					Srcset: "https://example.com/images/ep2/pci_avatar-massive.jpg 1500w,\n                                        https://example.com/images/ep2/pci_avatar-middle.jpg 600w,\n                                        https://example.com/images/ep2/pci_avatar-small.jpg 300w,\n                                        https://example.com/images/ep2/pci_avatar-tiny.jpg 150w",
				},
				PodcastSocialInteracts: nil,
				PodcastUpdateFrequency: nil,
				PodcastPodping:         nil,
			},
			Episode{
				Item: rss.Item{
					XMLName: xml.Name{
						Space: "",
						Local: "",
					},
					Title: "Episode 1 - The Past",
					Link:  "https://example.com/podcast/ep0001",
					Description: &rss.Description{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "description",
						},
						Value: "<p>How did podcasting get started? What was it like in the beginning?</p>",
					},
					Author:     "John Doe (john@example.com)",
					Categories: nil,
					Comments:   "",
					Enclosure: &rss.Enclosure{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "enclosure",
						},
						URL:    "https://example.com/file-01.mp3",
						Length: 43111403,
						Type:   "audio/mpeg",
					},
					GUID: &rss.GUID{
						XMLName: xml.Name{
							Space: "https://www.rssboard.org/rss-specification",
							Local: "guid",
						},
						Value:       "https://example.com/ep0001",
						IsPermaLink: &t,
					},
					PubDate: "Wed, 07 Oct 2020 04:30:38 GMT",
					Source:  nil,
				},
				ItunesDuration: "",
				ItunesImage: &ItunesImageTag{
					XMLName: xml.Name{
						Space: "http://www.itunes.com/dtds/podcast-1.0.dtd",
						Local: "image",
					},
					Href: "https://example.com/ep0001/artMd.jpg",
				},
				ItunesExplicit:    &f,
				ItunesTitle:       "",
				ItunesEpisode:     0,
				ItunesSeason:      0,
				ItunesEpisodeType: "",
				ItunesBlock:       "",
				PodcastTranscript: []PodcastTranscript{
					PodcastTranscript{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "transcript",
						},
						URL:      "https://example.com/ep1/transcript.txt",
						Type:     "text/plain",
						Language: "",
						Rel:      "",
					},
				},
				PodcastChapters: &PodcastChapters{
					XMLName: xml.Name{
						Space: "https://podcastindex.org/namespace/1.0",
						Local: "chapters",
					},
					URL:  "https://example.com/ep1_chapters.json",
					Type: "application/json",
				},
				PodcastSoundbite: []PodcastSoundbite{
					PodcastSoundbite{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "soundbite",
						},
						StartTime: 29.32,
						Duration:  34.0,
						Value:     "",
					},
				},
				PodcastPeople: []PodcastPerson{
					PodcastPerson{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "person",
						},
						Text:  "Adam Curry",
						Group: "",
						Role:  "",
						Href:  "https://curry.com",
						Img:   "http://example.com/images/adamcurry.jpg",
					},
					PodcastPerson{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "person",
						},
						Text:  "Dave Jones",
						Group: "",
						Role:  "guest",
						Href:  "https://www.imdb.com/name/nm0427852888/",
						Img:   "http://example.com/images/davejones.jpg",
					},
					PodcastPerson{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "person",
						},
						Text:  "Jebick Morton",
						Group: "visuals",
						Role:  "cover art designer",
						Href:  "https://example.com/artist/jebickmorton",
						Img:   "",
					},
				},
				PodcastSeason: &PodcastSeason{
					XMLName: xml.Name{
						Space: "https://podcastindex.org/namespace/1.0",
						Local: "season",
					},
					Value: 1,
					Name:  "Podcasting 2.0",
				},
				PodcastEpisode: &PodcastEpisode{
					XMLName: xml.Name{
						Space: "https://podcastindex.org/namespace/1.0",
						Local: "episode",
					},
					Value:   1.0,
					Display: "",
				},
				PodcastLicense: nil,
				PodcastAlternateEnclosures: []PodcastAlternateEnclosure{
					PodcastAlternateEnclosure{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "alternateEnclosure",
						},
						Type:             "audio/mpeg",
						Length:           "43203200",
						Bitrate:          128000.0,
						Title:            "Standard",
						Height:           0,
						Lang:             0,
						Rel:              0,
						Codecs:           0,
						Default:          &t,
						PodcastIntegrity: nil,
						Source: []PodcastSource{
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "https://example.com/file-01.mp3",
								ContentType: "",
							},
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "ipfs://someRandomMpegFile01",
								ContentType: "",
							},
						},
					},
					PodcastAlternateEnclosure{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "alternateEnclosure",
						},
						Type:             "audio/opus",
						Length:           "32406000",
						Bitrate:          96000.0,
						Title:            "High quality",
						Height:           0,
						Lang:             0,
						Rel:              0,
						Codecs:           0,
						Default:          nil,
						PodcastIntegrity: nil,
						Source: []PodcastSource{
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "https://example.com/file-high-01.opus",
								ContentType: "",
							},
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "ipfs://someRandomHighBitrateOpusFile01",
								ContentType: "",
							},
						},
					},
					PodcastAlternateEnclosure{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "alternateEnclosure",
						},
						Type:             "audio/aac",
						Length:           "5400300",
						Bitrate:          160000.0,
						Title:            "High quality AAC",
						Height:           0,
						Lang:             0,
						Rel:              0,
						Codecs:           0,
						Default:          nil,
						PodcastIntegrity: nil,
						Source: []PodcastSource{
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "https://example.com/file-proprietary-01.aac",
								ContentType: "",
							},
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "ipfs://someRandomProprietaryAACFile01",
								ContentType: "",
							},
						},
					},
					PodcastAlternateEnclosure{
						XMLName: xml.Name{
							Space: "https://podcastindex.org/namespace/1.0",
							Local: "alternateEnclosure",
						},
						Type:             "audio/opus",
						Length:           "5042000",
						Bitrate:          16000.0,
						Title:            "Low bandwidth",
						Height:           0,
						Lang:             0,
						Rel:              0,
						Codecs:           0,
						Default:          nil,
						PodcastIntegrity: nil,
						Source: []PodcastSource{
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "https://example.com/file-low-01.opus",
								ContentType: "",
							},
							PodcastSource{
								XMLName: xml.Name{
									Space: "https://podcastindex.org/namespace/1.0",
									Local: "source",
								},
								URI:         "ipfs://someRandomLowBitrateOpusFile01",
								ContentType: "",
							},
						},
					},
				},
				PodcastValue: nil,
				PodcastImages: &PodcastImages{
					XMLName: xml.Name{
						Space: "https://podcastindex.org/namespace/1.0",
						Local: "images",
					},
					Srcset: "https://example.com/images/ep1/pci_avatar-massive.jpg 1500w,\n                                        https://example.com/images/ep1/pci_avatar-middle.jpg 600w,\n                                        https://example.com/images/ep1/pci_avatar-small.jpg 300w,\n                                        https://example.com/images/ep1/pci_avatar-tiny.jpg 150w",
				},
				PodcastSocialInteracts: nil,
				PodcastUpdateFrequency: nil,
				PodcastPodping:         nil,
			},
		},
	},
	Version: "2.0",
}
