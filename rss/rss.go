// Package rss implements basic types for RSS 2.0 feeds. You should be able to
// use this to generate and parse simple RSS feeds if you wish.
//
// https://www.rssboard.org/rss-specification
//
// # Any Apple documentation is also included
//
// # Copyright
//
// Many of these docstrings are taken from the RSS Spec, which is
//   - License Link: https://creativecommons.org/licenses/by-sa/1.0/
//   - Changes were made
//   - Creator: RSS Advisory Board
//   - Title: RSS 2.0 Specification
//
// Original Notice:
//
// This document is authored by the RSS Advisory Board and is offered under the
// terms of the Creative Commons Attribution/Share Alike license. It is a
// derivative work of an original document titled RSS 2.0 published by the
// Berkman Klein Center for Internet & Society authored by Dave Winer.
package rss

import (
	"encoding/xml"
	"io"
)

// # RSS 2.0 (required)
//
//
//
// Example: `http://www.goupstate.com/`
//
// # Apple Podcasts (recommended)
//
//
//

// https://datatracker.ietf.org/doc/html/rfc2822
// TL;DR use [RFC2822DateFormatSpecifier]
type RFC2822Date string

// Provided for convenience
const RFC2822DateFormatSpecifier = "Mon Jan 02 15:04:05 -0700 2006"

const RSSVersion = "2.0"
const RSSNamespace = "https://www.rssboard.org/rss-specification"

type RSS struct {
	XMLName xml.Name `xml:"https://www.rssboard.org/rss-specification rss"`
	Channel Channel  `xml:"channel"`
	Version string   `xml:"version,attr,omitempty"`
}

type Channel struct {
	XMLName xml.Name `xml:"https://www.rssboard.org/rss-specification channel"`

	// # RSS 2.0 (required)
	//
	// The name of the channel. It's how people refer to your service. If you
	// have an HTML website that contains the same information as your RSS file,
	// the title of your channel should be the same as the title of your website.
	//
	// Example: `GoUpstate.com News Headlines`
	//
	// # Apple Podcasts (required)
	//
	// The show title.
	//
	// It’s important to have a clear, concise name for your podcast. Make your
	// title specific. A show titled Our Community Bulletin is too vague to
	// attract many subscribers, no matter how compelling the content.
	//
	// Pay close attention to the title as Apple Podcasts uses this field for
	// search.
	//
	// If you include a long list of keywords in an attempt to game podcast
	// search, your show may be removed from the Apple directory.
	Title string `xml:"https://www.rssboard.org/rss-specification title"`

	// # RSS 2.0 (required)
	//
	// The URL to the HTML website corresponding to the channel.
	//
	// Example:
	//  http://www.goupstate.com/
	//
	// # Apple Podcasts (recommended)
	//
	// The website associated with a podcast. Use the full URL.
	//
	// Typically a home page for a podcast or a dedicated portion of a larger
	// website. For example:
	//
	//  <link>
	//   http://www.example.com
	//  </link>
	//
	// or
	//
	//  <link>
	//   http://www.example.com/podcast
	//  </link>
	Link string `xml:"https://www.rssboard.org/rss-specification link"`

	// # RSS 2.0 (required)
	//
	// Phrase or sentence describing the channel.
	//
	// Example:
	//  The latest news from GoUpstate.com, a Spartanburg Herald-Journal Web site.
	//
	// # Apple Podcasts (required)
	//
	// The show description.
	//
	// Where description is text containing one or more sentences describing
	// your podcast to potential listeners. The maximum amount of text allowed
	// for this tag is 4000 bytes.
	//
	// To include links in your description or rich HTML, adhere to the
	// following technical guidelines: enclose all portions of your XML that
	// contain embedded HTML in a CDATA section to prevent formatting issues,
	// and to ensure proper link functionality. For example:
	//   <![CDATA[
	//     <a href="http://www.apple.com">Apple</a>
	//   ]]>
	//
	Description *Description `xml:"https://www.rssboard.org/rss-specification description"`

	// # RSS 2.0 (optional)
	//
	// The language the channel is written in. This allows aggregators to group
	// all Italian language sites, for example, on a single page. A list of
	// allowable values for this element, as provided by Netscape, is [here].
	//
	// You may also use [values defined] by the W3C.
	//
	// [here]: https://cyber.harvard.edu/rss/languages.html
	// [values defined]: https://www.w3.org/TR/REC-html40/struct/dirlang.html#langcodes
	//
	// Example:
	//  en-us
	//
	// # Apple Podcasts (required)
	//
	// The language spoken on the show.
	//
	// Because Apple Podcasts is available in territories around the world, it
	// is critical to specify the language of a podcast. Apple Podcasts only
	// supports values from the [ISO 639] list (two-letter language codes, with
	// some possible modifiers, such as "fr-ca").
	//
	// Invalid language codes will cause your feed to fail Apple validation.
	//
	// [ISO 639]: https://www.loc.gov/standards/iso639-2/php/code_list.php
	Language string `xml:"https://www.rssboard.org/rss-specification language,omitempty"`

	// # RSS 2.0 (optional)
	//
	// Copyright notice for content in the channel.
	//
	// Example:
	//  Copyright 2002, Spartanburg Herald-Journal
	//
	// # Apple Podcasts (situational)
	//
	// The show copyright details.
	//
	// If your show is copyrighted you should use this tag. For example:
	//  <copyright>Copyright 1995-2024 John John Appleseed</copyright>
	//
	Copyright string `xml:"https://www.rssboard.org/rss-specification copyright,omitempty"`

	// # RSS 2.0 (optional)
	//
	// Email address for person responsible for editorial content.
	//
	// Example:
	//  geo@herald.com (George Matesky)
	ManagingEditor string `xml:"https://www.rssboard.org/rss-specification managingEditor,omitempty"`

	// # RSS 2.0 (optional)
	//
	// Email address for person responsible for technical issues relating to channel.
	//
	// Example: `betty@herald.com (Betty Guernsey)`
	WebMaster string `xml:"https://www.rssboard.org/rss-specification webMaster,omitempty"`

	// # RSS 2.0 (optional)
	//
	// The publication date for the content in the channel. For example, the New
	// York Times publishes on a daily basis, the publication date flips once
	// every 24 hours. That's when the pubDate of the channel changes. All
	// date-times in RSS conform to the Date and Time Specification of [RFC 822],
	// with the exception that the year may be expressed with two characters or
	// four characters (four preferred).
	//
	// Example:
	//  Sat, 07 Sep 2002 09:42:31 GMT
	// [RFC 822]: https://www.ietf.org/rfc/rfc822.txt

	PubDate RFC2822Date `xml:"https://www.rssboard.org/rss-specification pubDate,omitempty"`

	// # RSS 2.0 (optional)
	//
	// The last time the content of the channel changed.
	//
	// Example:
	//  Sat, 07 Sep 2002 09:42:31 GMT
	// [RFC 822]: https://www.ietf.org/rfc/rfc822.txt

	LastBuildDate RFC2822Date `xml:"https://www.rssboard.org/rss-specification lastBuildDate,omitempty"`

	// # RSS 2.0 (optional)
	//
	// Specify one or more categories that the channel belongs to. Follows the
	// same rules as the <item>-level category element.
	//
	// Example:
	//  <category>Newspapers</category>
	//
	Categories []Category `xml:"https://www.rssboard.org/rss-specification category,omitempty"`

	// # RSS 2.0 (optional)
	//
	// A string indicating the program used to generate the channel.
	//
	// Example:
	//  MightyInHouse Content System v2.3
	//
	// # Apple Podcasts (situational):
	//
	// Specifies the program or hosting provider used to create the RSS feed.
	//
	// Hosting providers use this tag to identify themselves as the creator of an RSS feed.
	Generator string `xml:"https://www.rssboard.org/rss-specification generator,omitempty"`

	// # RSS 2.0 (optional)
	//
	// A URL that points to the documentation for the format used in the RSS
	// file. It's probably a pointer to this page. It's for people who might
	// stumble across an RSS file on a Web server 25 years from now and wonder
	// what it is.
	//
	// Example:
	//  https://www.rssboard.org/rss-specification
	//
	Docs string `xml:"https://www.rssboard.org/rss-specification docs,omitempty"`

	// # RSS 2.0 (optional)
	//
	// Allows processes to register with a cloud to be notified of updates to
	// the channel, implementing a lightweight publish-subscribe protocol for
	// RSS feeds. More info [here].
	//
	// Example:
	//  http://blogs.law.harvard.edu/tech/rss
	//
	// [here]: https://www.rssboard.org/rss-specification#ltcloudgtSubelementOfLtchannelgt
	Cloud *Cloud `xml:"https://www.rssboard.org/rss-specification cloud,omitempty"`

	// # RSS 2.0 (optional)
	//
	// ttl stands for time to live. It's a number of minutes that indicates how
	// long a channel can be cached before refreshing from the source. More info
	// [here].
	//
	// [here]: https://www.rssboard.org/rss-specification#ltttlgtSubelementOfLtchannelgt
	//
	// Example:
	//  <ttl>60</ttl>
	//
	TTL int `xml:"https://www.rssboard.org/rss-specification ttl,omitempty"`

	// # RSS 2.0 (optional)
	//
	// Specifies a GIF, JPEG or PNG image that can be displayed with the
	// channel. More info [here].
	//
	// [here]: https://cyber.harvard.edu/rss/rss.html#ltttlgtSubelementOfLtchannelgt
	//
	Image *Image `xml:"https://www.rssboard.org/rss-specification image,omitempty"`

	// # RSS 2.0 (optional)
	//
	// The [PICS] rating for the channel.
	//
	// [PICS]: http://www.w3.org/PICS/
	Rating string `xml:"https://www.rssboard.org/rss-specification rating,omitempty"`

	// TextInput specifies a text input box that can be displayed with the
	// channel. More info [here].
	//
	// [here]: https://cyber.harvard.edu/rss/rss.html#ltimagegtSubelementOfLtchannelgt
	TextInput *TextInput `xml:"https://www.rssboard.org/rss-specification textInput,omitempty"`

	// SkipHours is a hint for aggregators telling them which hours they can
	// skip. More info [here].
	//
	// [here]: https://cyber.harvard.edu/rss/skipHoursDays.html#skiphours
	//
	SkipHours *SkipHours `xml:"https://www.rssboard.org/rss-specification skipHours,omitempty"`

	// A hint for aggregators telling them which days they can skip. More info
	// [here].
	//
	// [here]: https://cyber.harvard.edu/rss/skipHoursDays.html#skipdays
	//
	SkipDays *SkipDays `xml:"https://www.rssboard.org/rss-specification skipDays,omitempty"`

	Items []Item `xml:"https://www.rssboard.org/rss-specification item"`
}

type Description struct {
	XMLName xml.Name `xml:"https://www.rssboard.org/rss-specification description"`
	Value   string   `xml:",cdata"`
}

type CloudProtocol string

const (
	CloudProtocolXMLRPC   CloudProtocol = "xml-rpc"
	CloudProtocolSOAP     CloudProtocol = "soap"
	CloudProtocolHTTPPost CloudProtocol = "http-post"
)

// Cloud specifies a Web application that supports the rssCloud interface,
// described below, which can be implemented in either XML-RPC or SOAP.
//
// See [original docs]
//
// [original docs]: https://cyber.harvard.edu/rss/soapMeetsRss.html#rsscloudInterface
type Cloud struct {
	XMLName xml.Name `xml:"https://www.rssboard.org/rss-specification cloud"`

	// Domain is the domain name or IP address of the cloud
	//
	// Example:
	//  radio.xmlstoragesystem.com
	Domain string `xml:"domain,attr"`

	// Port is the TCP port that the cloud is running on
	//
	// Example:
	//  80
	Port string `xml:"port,attr"`

	// Path is the location of its responder
	//
	// Example:
	//  /RPC2
	Path string `xml:"path,attr"`

	// RegisterProcedure is the name of the procedure to call to request notification
	//
	// Example:
	//  xmlStorageSystem.rssPleaseNotify
	RegisterProcedure string `xml:"registerProcedure,attr"`

	// Protocol is xml-rpc, soap or http-post (case-sensitive),
	// indicating which protocol is to be used
	//
	// Example:
	//  xmlStorageSystem.rssPleaseNotify
	Protocol CloudProtocol `xml:"protocol,attr"`
}

// Image is an optional sub-element of [Channel], which contains three
// required and three optional sub-elements.
//
// https://cyber.harvard.edu/rss/rss.html#ltimagegtSubelementOfLtchannelgt
type Image struct {
	XMLName xml.Name `xml:"https://www.rssboard.org/rss-specification image"`

	// URL (required) is the URL of a GIF, JPEG or PNG image that represents the channel
	URL string `xml:"https://www.rssboard.org/rss-specification url"`

	// Title (required) describes the image, it's used in the ALT attribute of the HTML
	// <img> tag when the channel is rendered in HTML.
	Title string `xml:"https://www.rssboard.org/rss-specification title"`

	// Link (required) is the URL of the site, when the channel is rendered, the image is
	// a link to the site. (Note, in practice the image <title> and <link> should have the same value as the channel's <title> and <link>.
	Link string `xml:"https://www.rssboard.org/rss-specification link"`

	// Width of image in pixels. Default 88, max 144. Optional
	Width int `xml:"https://www.rssboard.org/rss-specification width,omitempty"`

	// Height of image in pixels. Default 31, max 400. Optional
	Height int `xml:"https://www.rssboard.org/rss-specification height,omitempty"`

	// Description contains text that is included in the TITLE attribute of the
	// link formed around the image in the HTML rendering. Optional.
	Description string `xml:"https://www.rssboard.org/rss-specification description,omitempty"`
}

// TextInput: The purpose of the <textInput> element is something of a
// mystery. You can use it to specify a search engine box. Or to allow a reader
// to provide feedback. Most aggregators ignore it.
//
// https://cyber.harvard.edu/rss/rss.html#lttextinputgtSubelementOfLtchannelgt
type TextInput struct {
	XMLName xml.Name `xml:"https://www.rssboard.org/rss-specification textInput"`

	// Title (required) is the label of the Submit button in the text input area.
	Title string `xml:"https://www.rssboard.org/rss-specification title"`

	// Description (required) explains the text input area.
	Description string `xml:"https://www.rssboard.org/rss-specification description"`

	// Name (required) of the text object in the text input area.
	// <img> tag when the channel is rendered in HTML.
	Name string `xml:"https://www.rssboard.org/rss-specification name"`

	// Link (required) is the URL of the CGI script that processes text input requests.
	Link string `xml:"https://www.rssboard.org/rss-specification link"`
}

// SkipHours is an XML element that contains up to 24 <hour>
// sub-elements whose value is a number between 0 and 23, representing a time in
// GMT, when aggregators, if they support the feature, may not read the channel
// on hours listed in the skipHours element.
//
// The hour beginning at midnight is hour zero.
//
// https://cyber.harvard.edu/rss/skipHoursDays.html#skiphours
type SkipHours struct {
	XMLName xml.Name `xml:"https://www.rssboard.org/rss-specification skipHours"`

	Hours []string `xml:"https://www.rssboard.org/rss-specification hour"`
}

// Values a <day> element can have in RSSChannelSkipDays
type SkipDay string

const (
	SkipDayMonday    SkipDay = "Monday"
	SkipDayTuesday   SkipDay = "Tuesday"
	SkipDayWednesday SkipDay = "Wednesday"
	SkipDayThursday  SkipDay = "Thursday"
	SkipDayFriday    SkipDay = "Friday"
	SkipDaySaturday  SkipDay = "Saturday"
	SkipDaySunday    SkipDay = "Sunday"
)

// SkipDays is an XML element that contains up to seven <day>
// sub-elements whose value is Monday, Tuesday, Wednesday, Thursday, Friday,
// Saturday or Sunday. Aggregators may not read the channel during days listed
// in the skipDays element.
//
// https://cyber.harvard.edu/rss/skipHoursDays.html#skipdays
type SkipDays struct {
	XMLName xml.Name `xml:"https://www.rssboard.org/rss-specification skipDays"`

	Days []SkipDay `xml:"https://www.rssboard.org/rss-specification day"`
}

// # RSS 2.0
//
// Item may represent a "story" -- much like a story in a newspaper or magazine;
// if so its description is a synopsis of the story, and the link points to the
// full story. An item may also be complete in itself, if so, the description
// contains the text (entity-encoded HTML is allowed; see [examples]), and the
// link and title may be omitted. All elements of an item are optional, however
// at least one of title or description must be present.
//
// [examples]: https://cyber.harvard.edu/rss/encodingDescriptions.html
type Item struct {
	XMLName xml.Name `xml:"https://www.rssboard.org/rss-specification item"`

	// # RSS 2.0 (optional, one of title or description must be present)
	//
	// The title of the item.
	//
	// Example:
	//  Venice Film Festival Tries to Quit Sinking
	//
	// # Apple Podcasts (required)
	//
	// An episode title.
	//
	// title is a string containing a clear, concise name for your episode.
	// Don’t specify the episode number or season number in this tag. Instead,
	// specify those details in the appropriate tags ( <itunes:episode>,
	// <itunes:season>). Also, don’t repeat the title of your show within your
	// episode title.
	//
	// Separating episode and season number from the title makes it possible for
	// Apple to easily index and order content from all shows.
	//
	Title string `xml:"https://www.rssboard.org/rss-specification title,omitempty"`

	// # RSS 2.0 (optional)
	//
	// The URL of the item.
	//
	// Example:
	//  http://nytimes.com/2004/12/07FEST.html
	//
	// # Apple Podcasts (recommended)
	//
	// An episode link URL.
	// This is used when an episode has a corresponding webpage. Use the full URL. For example:
	//  <link>
	//   https://www.example.com
	//  </link>
	//
	// or
	//  <link>
	//   https://www.example.com/podcast
	//  </link>
	//
	Link string `xml:"https://www.rssboard.org/rss-specification link,omitempty"`

	// # RSS 2.0 (optional, one of title or description must be present)
	//
	// The item synopsis.
	//
	// Example:
	//  Some of the most heated chatter at the Venice Film Festival this week was about the way that the arrival of the stars at the Palazzo del Cinema was being staged.
	//
	// # Apple Podcasts (recommended)
	//
	// An episode description.
	//
	// description is text containing one or more sentences describing your
	// episode to potential listeners. You can specify up to 10,000 characters.
	// You can use rich text formatting and some HTML (<p>, <ol>, <ul>, <li>,
	// <a>) if wrapped in the <CDATA> tag.
	//
	// To include links in your description or rich HTML, adhere to the
	//  following technical guidelines: enclose all portions of your XML that
	// contain embedded HTML in a CDATA section to prevent formatting issues,
	// and to ensure proper link functionality. For example:
	//  <![CDATA[
	// 	 <a href="http://www.apple.com">Apple</a>
	//  ]]>
	//
	Description *Description `xml:"https://www.rssboard.org/rss-specification description"`

	// # RSS 2.0 (optional)
	//
	// Email address of the author of the item. [More].
	//
	// [More]: https://cyber.harvard.edu/rss/rss.html#ltauthorgtSubelementOfLtitemgt
	//
	// Example:
	//  oprah\@oxygen.net
	//
	Author string `xml:"https://www.rssboard.org/rss-specification author,omitempty"`

	// # RSS 2.0 (optional)
	//
	// Includes the item in one or more categories. [More].
	//
	// [More]: https://www.rssboard.org/rss-specification#ltcategorygtSubelementOfLtitemgt
	//
	Categories []Category `xml:"https://www.rssboard.org/rss-specification category,omitempty"`

	// # RSS 2.0 (optional)
	//
	// URL of a page for comments relating to the item. [More].
	//
	// [More]: https://www.rssboard.org/rss-specification#ltcommentsgtSubelementOfLtitemgt
	//
	Comments string `xml:"https://www.rssboard.org/rss-specification url,omitempty"`

	// # RSS 2.0 (optional)
	//
	// Describes a media object that is attached to the item. [More].
	//
	// [More]: https://www.rssboard.org/rss-specification#ltenclosuregtSubelementOfLtitemgt
	//
	// # Apple Podcasts (required)
	// 	The episode content, file size, and file type information.
	//
	Enclosure *Enclosure `xml:"https://www.rssboard.org/rss-specification enclosure,omitempty"`

	// GUID is a string that uniquely identifies the item.
	GUID *GUID `xml:"https://www.rssboard.org/rss-specification guid,omitempty"`

	// # RSS 2.0 (optional)
	//
	// Its value is a [date], indicating when the item was published. If it's a
	// date in the future, aggregators may choose to not display the item until
	// that date.
	//
	// Example:
	//  <pubDate>Sun, 19 May 2002 15:21:36 GMT</pubDate>
	//
	// [date]: http://asg.web.cmu.edu/rfc/rfc822.html
	//
	// # Apple Podcasts (recommended)
	//
	// The date and time when an episode was released.
	//
	// Format the date using the [RFC 2822] specifications.
	//
	// [RFC 2822]: http://www.faqs.org/rfcs/rfc2822.html
	//
	// For example:
	//  Sat, 01 Apr 2023 19:00:00 GMT.
	//
	//
	PubDate RFC2822Date `xml:"https://www.rssboard.org/rss-specification pubDate,omitEmpty"`

	// # RSS 2.0 (optional)
	//
	// The RSS channel that the item came from.
	Source *Source `xml:"https://www.rssboard.org/rss-specification source,omitempty"`
}

// Category is an optional sub-element of <item>.
//
// It has one optional attribute, domain, a string that identifies a
// categorization taxonomy.
//
// The value of the element is a forward-slash-separated string that identifies
// a hierarchic location in the indicated taxonomy. Processors may establish
// conventions for the interpretation of categories. Two examples are provided
// below:
//
//	<category>Grateful Dead</category>
//
//	<category domain="http://www.fool.com/cusips">MSFT</category>
//
// You may include as many category elements as you need to, for different
// domains, and to have an item cross-referenced in different parts of the same
// domain.
//
// https://www.rssboard.org/rss-specification#ltcategorygtSubelementOfLtitemgt
type Category struct {
	XMLName xml.Name `xml:"https://www.rssboard.org/rss-specification category"`
	Value   string   `xml:",chardata"`

	// Domain is optional
	Domain string `xml:"domain,attr,omitempty"`
}

// # RSS 2.0 (optional)
//
// Source  is an optional sub-element of <item>.
//
// Its value is the name of the RSS channel that the item came from, derived
// from its <title>. It has one required attribute, url, which links to the
// XMLization of the source.
//
// Example:
//
//	<source url="http://www.tomalak.org/links2.xml">Tomalak's Realm</source>
//
// The purpose of this element is to propagate credit for links, to
// publicize the sources of news items. It can be used in the Post command
// of an aggregator. It should be generated automatically when forwarding an
// item from an aggregator to a weblog authoring tool.
type Source struct {
	XMLName xml.Name `xml:"https://www.rssboard.org/rss-specification source"`
	Value   string   `xml:",chardata"`

	// URL is required
	URL string `xml:"https://www.rssboard.org/rss-specification url,omitempty"`
}

// # RSS 2.0
//
// Enclosure is an optional sub-element of <item>.
//
// It has three required attributes.
//
// # Apple Podcasts (required)
//
// The <enclosure> tag has three attributes: URL, length, and type:
//
//   - URL. The URL attribute points to your podcast media file. Specify the
//     full file extension within the URL attribute. This determines whether
//     or not content appears in the podcast directory. Supported file
//     formats include M4A, MP3, MOV, MP4, M4V, and PDF.
//
//   - Length. The length attribute is the file size in bytes. You can find
//     this information in the properties of your podcast file (on a Mac,
//     choose File > Get Info and refer to the size field).
//
//   - Type. The type attribute provides the correct category for the type of
//     file you are using. The type values for the supported file formats
//     are: audio/x-m4a, audio/mpeg, video/quicktime, video/mp4, video/x-m4v,
//     and application/pdf.
//
// For example:
//
//	<enclosure
//	 url="http://mypodcast.com/episode001.mp3"
//	 length="5650889"
//	 type="audio/mpeg
//	/>
type Enclosure struct {
	XMLName xml.Name `xml:"https://www.rssboard.org/rss-specification enclosure"`
	// # RSS 2.0 (required)
	//
	// URL says where the enclosure is located. The url must be an http url.
	//
	// # Apple Podcasts (required)
	//
	// The URL attribute points to your podcast media file. Specify the full
	// file extension within the URL attribute. This determines whether or not
	// content appears in the podcast directory. Supported file formats include
	// M4A, MP3, MOV, MP4, M4V, and PDF.
	//
	//   - Length. The length attribute is the file size in bytes. You can find
	//     this information in the properties of your podcast file (on a Mac,
	//     choose File > Get Info and refer to the size field).
	//
	//   - Type.
	URL string `xml:"url,attr"`

	// # RSS 2.0 (required)
	//
	// Length says how big it is in bytes
	//
	// # Apple Podcasts (required)
	//
	// Length. The length attribute is the file size in bytes. You can find this
	//  information in the properties of your podcast file (on a Mac, choose
	// File > Get Info and refer to the size field).
	//
	Length int `xml:"length,attr"`

	// # RSS 2.0 (required)
	//
	// Type says what its type is, a standard MIME type.
	//
	// # Apple Podcasts (required)
	//
	// The type attribute provides the correct category for the type of file you
	// are using. The type values for the supported file formats are:
	//
	//  - audio/x-m4a
	//  - audio/mpeg
	//  - video/quicktime
	//  - video/mp4
	//  - video/x-m4v
	//  - application/pdf
	Type string `xml:"type,attr"`
}

// GUID  is an optional sub-element of <item>.
//
// guid stands for globally unique identifier. It's a string that uniquely
// identifies the item. When present, an aggregator may choose to use this
// string to determine if an item is new.
//
//	<guid>http://some.server.com/weblogItem3207</guid>
//
// There are no rules for the syntax of a guid. Aggregators must view them
// as a string. It's up to the source of the feed to establish the
// uniqueness of the string.
type GUID struct {
	XMLName xml.Name `xml:"https://www.rssboard.org/rss-specification guid"`
	Value   string   `xml:",chardata"`

	// If the guid element has an attribute named isPermaLink with a value of
	// true, the reader may assume that it is a permalink to the item, that is,
	// a url that can be opened in a Web browser, that points to the full item
	// described by the <item> element. An example:
	//
	//	<guid isPermaLink="true">http://inessential.com/2002/09/01.php#a2</guid>
	//
	// isPermaLink is optional, its default value is true. If its value is
	// false, the guid may not be assumed to be a url, or a url to anything in
	// particular.
	IsPermaLink *bool `xml:"isPermaLink,attr,omitempty"`
}

// GetDecoder is a helper method to get an encoding/xml encoder with the default
// namespace set to https://www.rssboard.org/rss-specification.
//
// You probably want to decode XML files that don't specify a namespace, so you
// need to specify a default for go's XML decoder.
func GetDecoder(r io.Reader) *xml.Decoder {
	decoder := xml.NewDecoder(r)
	decoder.DefaultSpace = "https://www.rssboard.org/rss-specification"
	return decoder
}
