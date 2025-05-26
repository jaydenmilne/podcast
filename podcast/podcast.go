package podcast

import (
	"encoding/xml"

	"github.com/jaydenm/rssfeed/rss"
)

const ItunesNamespaceURL = "http://www.itunes.com/dtds/podcast-1.0.dtd"
const PodcastNamepaceURL = "https://podcastindex.org/namespace/1.0"

type RSSPodcast struct {
	XMLName xml.Name `xml:"rss"`
	Channel Podcast  `xml:"channel"`
	Version string   `xml:"version,attr,omitempty"`
}

type Podcast struct {
	rss.Channel

	// # Apple Podcasts:
	//
	// ItunesImage (required) is the artwork for the show.
	//
	// Specify your show artwork by providing a URL linking to it.
	//
	// Depending on their device, subscribers see your podcast artwork in
	// varying sizes. Therefore, make sure your design is effective at both its
	// original size and at thumbnail size. You should include a show title,
	// brand, or source name as part of your podcast artwork. Here are
	// additional marketing best practices. For examples of podcast artwork,
	// see the Top Podcasts chart. To avoid technical issues when you update
	// your podcast artwork, be sure to:
	//
	//  - Change the artwork file name and URL at the same time
	//  - Make sure the file type in the URL matches the actual file type of the
	//    image file.
	//  - Verify the web server hosting your artwork allows HTTP head requests
	//    including Last Modified.
	//
	// Artwork must be a minimum size of 1400 x 1400 pixels and a maximum size
	// of 3000 x 3000 pixels, in JPEG or PNG format, 72 dpi, with appropriate
	// file extensions (.jpg, .png), and in the RGB colorspace. Confirm your
	// art does not contain an Alpha Channel. These requirements are different
	// from the standard RSS image tag specifications.
	ItunesImage ItunesImageTag `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd image"`

	// # Apple Podcasts:
	//
	// ItunesCategory (required) is the show category information. For a
	// complete list of categories and subcategories, see [Apple Podcast categories].
	//
	// [Apple Podcast Categories]: https://podcasters.apple.com/support/1691-apple-podcasts-categories
	//
	// Select the category that best reflects the content of your show. If available, you can also define a subcategory.
	//
	// Although you can specify more than one category and subcategory in your RSS feed, Apple Podcasts only recognizes the first category and subcategory.
	//
	// When specifying categories and subcategories, be sure to properly escape ampersands. For example:
	//
	// Single category:
	//  <itunes:category text="History" />
	//
	// Category with ampersand:
	//  <itunes:category text="Kids &amp; Family" />
	//
	// Category with subcategory:
	//  <itunes:category text="Society &amp; Culture">
	// 	 <itunes:category text="Documentary" />
	//  </itunes:category>
	//
	// Multiple categories:
	//  <itunes:category text="Society &amp; Culture">
	// 	 <itunes:category text="Documentary" />
	//  </itunes:category>
	//  <itunes:category text="Health">
	// 	 <itunes:category text="Mental Health" />
	//  </itunes:category>
	ItunesCategory string `xml:"itunes:category"`

	// # Apple Podcasts:
	//
	// ItunesExplicit (required) is the podcast parental advisory information.
	//
	// The explicit value can be one of the following:
	// - True. If you specify true, indicating the presence of explicit content,
	//   Apple Podcasts displays an [Explicit] parental advisory graphic for your
	//   podcast.
	//
	//   Podcasts containing explicit material aren’t available in some Apple Podcasts territories.
	//
	// - False. If you specify false, indicating that your podcast doesn’t
	//   contain explicit language or adult content, Apple Podcasts displays a
	//   [Clean] parental advisory graphic for your podcast.
	//
	// [Explicit]: https://help.apple.com/itc/podcasts_connect/#/itcfafb6d665
	// [Clean]: https://help.apple.com/itc/podcasts_connect/#/itcb343e8058
	ItunesExplicit *bool `xml:"itunes:explicit"`

	// # Apple Podcasts:
	//
	// ItunesAuthor (recommended) is the group responsible for creating the show
	//
	// Show author most often refers to the parent company or network of a
	// podcast, but it can also be used to identify the host(s) if none exists.
	//
	// Author information is especially useful if a company or organization
	// publishes multiple podcasts.
	ItunesAuthor string `xml:"itunes:author,omitempty"`

	// # Apple Podcasts:
	//
	// ItunesTitle (situational) is he show title specific for Apple Podcasts.
	//
	// <itunes:title> is a string containing a clear concise name of your show
	// on Apple Podcasts. Do not include episode or season number in the title.
	// There are dedicated tags for that information. See <itunes:episode> and
	// <itunes:season>.
	ItunesTitle string `xml:"itunes:title,omitempty"`

	// # Apple Podcasts:
	//
	// ItunesType (situational) is the type of show.
	//
	// If your show is Serial you must use this tag.
	//
	// Its values can be one of the following:
	//
	//  - Episodic (default). Specify episodic when episodes are intended to be
	//    consumed without any specific order. Apple Podcasts will present newest
	//    episodes first and display the publish date (required) of each episode. If
	//    organized into [seasons], the newest season will be presented first -
	//    otherwise, episodes will be grouped by year published, newest first.
	//
	//    For new subscribers, Apple Podcasts adds the newest, most recent episode
	//    in their Library.
	//
	//  - Serial. Specify serial when episodes are intended to be consumed in
	//    sequential order. Apple Podcasts will present the oldest episodes first
	//    and display the episode numbers (required) of each episode. If organized
	//    into seasons, the newest season will be presented first and
	//    <itunes:episode> numbers must be given for each episode.
	//
	// Each show type has different behavior for automatic downloads. [Learn more].
	//
	// [seasons]: https://help.apple.com/itc/podcasts_connect/#/itc77382b700
	// [Learn more]: https://podcasters.apple.com/support/1662-automatic-downloads-on-apple-podcasts
	//
	ItunesType ItunesShowType `xml:"itunes:type,omitempty"`

	// # Apple Podcasts:
	//
	// ItunesNewFeedURL (situational) is the new podcast RSS Feed URL.
	//
	// If you change the URL of your podcast feed, you should use this tag in
	// your new feed.
	//
	// Use the <itunes:new-feed-url> tag to manually change the URL where your
	// podcast is located.
	//  <itunes:new-feed-url>
	//    https://newlocation.com/example.rss
	//  </itunes:new-feed-url>
	//
	// You should maintain your old feed until you have migrated your existing
	// followers. [Learn how to update your podcast RSS feed URL].
	//
	// [Learn how to update your podcast RSS feed URL]: https://podcasters.apple.com/support/change-the-rss-feed-url
	//
	// Note: The <itunes:new-feed-url> tag reports new feed URLs to Apple
	// Podcasts and isn’t displayed in Apple Podcasts.
	ItunesNewFeedURL string `xml:"itunes:new-feed-url,omitempty"`

	// # Apple Podcasts:
	//
	// ItunesComplete (situational) is the podcast update status.
	//
	// If you will never publish another episode to your show, use this tag.
	//
	// Specifying the <itunes:complete> tag with a Yes value indicates that a
	// podcast is complete and you will not post any more episodes in the future.
	//
	// Specifying any value other than Yes has no effect.
	ItunesComplete ItunesYesType `xml:"itunes:complete,omitempty"`

	// # Apple Podcasts:
	//
	// ItunesBlock (situational) is the podcast show or hide status.
	//
	// If you want your show removed from the Apple directory, use this tag.
	//
	// Specifying the <itunes:block> tag with a Yes value, prevents the entire
	// podcast from appearing in Apple Podcasts.
	//
	// Specifying any value other than Yes has no effect.
	ItunesBlock ItunesYesType `xml:"itunes:block,omitempty"`

	// # Apple Podcasts:
	//
	// ItunesApplePodcastVerify (situational) is used to verify ownership of a show when a
	// podcast creator chooses to move it from one Apple Podcasts Connect
	// account to another. [Learn more about how to claim your show].
	//
	// [Learn more about how to claim your show]: https://podcasters.apple.com/support/5497-claim-your-show
	ItunesApplePodcastVerify string `xml:"itunes:applepodcastsverify,omitempty"`

	// # Apple Podcasts:
	//
	// PodcastTxtVerify (situational) is an alternate method to verify your show
	// . See ApplePodcast.ItunesApplePodcastVerify
	PodcastTxtVerify PodcastTxt `xml:"podcast:txt"`

	Items []PodcastEpisode `xml:"item"`
}

type ItunesImageTag struct {
	XMLName xml.Name `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd image"`
	Href    string   `xml:"href,attr"`
}

// ITunesShowType is an enum meant to be used with ITunesType <itunes:type> and
// specifies if a show is periodic or serial
type ItunesShowType string

const (
	ItunesShowTypeEpisodic ItunesShowType = "episodic"
	ItunesShowTypeSerial                  = "serial"
)

type PodcastEpisode struct {
	rss.Item

	// # Apple Podcasts:
	//
	// ItunesDuration (recommended) is the duration of an episode.
	//
	// Different duration formats are accepted however it is recommended to
	// convert the length of the episode into seconds.
	ItunesDuration string `xml:"itunes:duration,omitempty"`

	// # Apple Podcasts:
	//
	// ItunesImage (recommended) is the [episode artwork].
	//
	// You should use this tag when you have a high quality, episode-specific
	// image you would like listeners to see.
	//
	// Specify your episode artwork using the href attribute in the
	// <itunes:image> tag. [RSS Feed Sample].
	//
	// Depending on their device, listeners see your episode artwork in varying
	// sizes. Therefore, make sure your design is effective at both its original
	// size and at thumbnail size. You should include a title, brand, or source
	// name as part of your episode artwork. To avoid technical issues when you
	// update your episode artwork, be sure to:
	//
	// 	Change the artwork file name and URL at the same time.
	//
	// 	Confirm your art does not contain an Alpha Channel.
	//
	// 	Verify the web server hosting your artwork allows HTTP head requests
	// including Last Modified.
	//
	// Artwork must be a minimum size of 1400 x 1400 pixels and a maximum size
	// of 3000 x 3000 pixels, in JPEG or PNG format, 72 dpi, with appropriate
	// file extensions (.jpg, .png), and in the RGB colorspace.
	//
	// Make sure the file type in the URL matches the actual file type of the image file.
	//
	// [episode artwork]: https://podcasters.apple.com/support/896-artwork-requirements#episodes
	// [RSS Feed Sample]: https://help.apple.com/itc/podcasts_connect/#/itcbaf351599
	ItunesImage ItunesImageTag `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd image,omitempty"`

	// # Apple Podcasts:
	//
	// ItunesExplicit (recommended) is the podcast parental advisory information.
	//
	// The explicit value can be one of the following:
	// - True. If you specify true, indicating the presence of explicit content,
	//   Apple Podcasts displays an [Explicit] parental advisory graphic for your
	//   episode.
	//
	//   Episodes containing explicit material aren’t available in some Apple
	//   Podcasts territories.
	//
	// - False. If you specify false, indicating that the episode podcast doesn’t
	//   contain explicit language or adult content, Apple Podcasts displays a
	//   [Clean] parental advisory graphic for your episode.
	//
	// [Explicit]: https://help.apple.com/itc/podcasts_connect/#/itcfafb6d665
	// [Clean]: https://help.apple.com/itc/podcasts_connect/#/itcb343e8058
	ItunesExplicit *bool `xml:"itunes:explicit"`

	// # Apple Podcasts:
	//
	// ItunesTitle (situational) is an episode title specific for Apple Podcasts.
	//
	// <itunes:title> is a string containing a clear concise name of your
	// episode on Apple Podcasts.
	//
	// Don’t specify the episode number or season number in this tag. Instead,
	// specify those details in the appropriate tags ( <itunes:episode>,
	// <itunes:season>). Also, don’t repeat the title of your show within your
	// episode title.
	//
	// Separating episode and season number from the title makes it possible for
	//  Apple to easily index and order content from all shows.
	ItunesTitle string `xml:"itunes:title,omitempty"`

	// # Apple Podcasts:
	//
	// ItunesEpisode (situational) is an episode number.
	//
	// If all your episodes have numbers and you would like them to be ordered
	// based on them, use this tag for each one.
	//
	// Episode numbers are optional for <itunes:type>episodic shows, but are
	// mandatory for serial shows.
	//
	// Where episode is a non-zero integer (1, 2, 3, etc.) representing your
	// episode number.
	//
	// If you are using your RSS feed to distribute a free version of an episode
	// that is already available to Apple Podcasts paid subscribers, make sure
	// the episode numbers are the same so you don’t have duplicate episodes
	// appear on your show page. Learn more about how to
	// [set up your show for a subscription]
	//
	// [set up your show for a subscription]: https://podcasters.apple.com/support/set-up-your-show-for-a-subscription
	ItunesEpisode int `xml:"itunes:episode,omitempty"`

	// # Apple Podcasts:
	//
	// ItunesSeason (situational) is the episode season number.
	//
	// If an episode is within a season use this tag.
	//
	// Where season is a non-zero integer (1, 2, 3, etc.) representing your
	// season number.
	//
	// To allow the season feature for shows containing a single season, if
	// only one season exists in the RSS feed, Apple Podcasts doesn’t display a
	// season number. When you add a second season to the RSS feed, Apple
	// Podcasts displays the season numbers.
	ItunesSeason int `xml:"itunes:season,omitempty"`

	// # Apple Podcasts:
	//
	// ItunesEpisodeType (situational) is the episode type.
	//
	// If an episode is a trailer or bonus content, use this tag.
	//
	// Where the episodeType value can be one of the following:
	//
	//  -
	//  -
	//  -
	ItunesEpisodeType EpisodeType `xml:"itunes:episodeType,omitempty"`

	// # Apple Podcasts:
	//
	// A link to the episode transcript in the Closed Caption format. You should
	// use this tag when you have a valid transcript file available for users to
	// read.
	//
	// Specify the link to your transcript in the url attribute of the tag.
	//
	// Apple Podcasts will prefer VTT format over SRT format if multiple
	// instances are included. A valid type attribute is required. Accepted
	// types include text/vtt, application/srt, application/x-subrip.
	//
	// Learn more about <podcast> namespace RSS tags on the [Github repository].
	//
	// Options for displaying transcripts are available in Apple Podcasts
	// Connect for each show. [Learn more].
	//
	// [Github repository]: https://github.com/Podcastindex-org/podcast-namespace
	// [Learn more]: https://podcasters.apple.com/support/5316-transcripts-on-apple-podcasts
	PodcastTranscript string `xml:"podcast:transcript,omitempty"`

	// # Apple Podcasts (situational):
	//
	// ItunesBlock is the episode show or hide status.
	// If you want an episode removed from the Apple directory, use this tag.
	// Specifying the <itunes:block> tag with a Yes value prevents that episode from appearing in Apple Podcasts.
	// For example, you might want to block a specific episode if you know that its content would otherwise cause the entire podcast to be removed from Apple Podcasts.
	// Specifying any value other than Yes has no effect.
	ItunesBlock ItunesYesType `xml:"itunes:block,omitempty"`
}

// ItunesYes is meant to be used with various properties that either accept a
// a `Yes` value to opt in or anything else to opt out
type ItunesYesType string

// ItunesYesValue is the sentinel yes value
const ItunesYesValue ItunesYesType = "Yes"

// Alternate way to verify a podcast using <podcast:txt purpose=“applepodcastsverify”>
// vs <itunes:applepodcastsverify>
type PodcastTxt struct {
	XMLName xml.Name `xml:"podcast:txt"`
	Value   string   `xml:",chardata"`

	// You probably want
	// 	applepodcastsverify
	Purpose string `xml:"purpose,attr,omitempty"`
}

// The type of episode (full, trailer or bonus) to be used with itunes:episodeType
// / ItunesEpisodeType
type EpisodeType string

const (
	// FullEpisode (default). Specify full when you are submitting the complete
	// content of your show.
	FullEpisode EpisodeType = "full"

	// TrailerEpisode is when you are submitting a short, promotional piece of
	// content that represents a preview of your current show.
	//
	// The rules for using trailer and bonus tags depend on whether the
	// <itunes:season> and <itunes:episode> tags have values:
	//
	//  - No season or episode number: a show trailer
	//  - A season number and no episode number: a season trailer. (Note: an
	//    episode trailer should have a different <guid> than the actual
	//    episode)
	//  - Episode number and optionally a season number: an  episode
	//    trailer/teaser, later replaced with the actual episode
	TrailerEpisode = "trailer"

	// BonusEpisode is when you are submitting extra content for your show (for
	// example, behind the scenes information or interviews with the cast) or
	// cross-promotional content for another show.
	//
	// The rules for using trailer and bonus tags depend on whether the
	// <itunes:season> and <itunes:episode> tags have values:
	//
	//  - No season or episode number: a show bonus
	//  - A season number: a season bonus
	//  - Episode number and optionally a season number: a bonus episode related
	//    to a specific episode
	BonusEpisode = "bonus"
)
