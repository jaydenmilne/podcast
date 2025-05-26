// podcast package
//
// Types beginning with "Podcast" are from the podcasting 2.0 spec
package podcast

import (
	"encoding/xml"

	"github.com/jaydenmilne/podcast/rss"
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
	// Select the category that best reflects the content of your show. If
	// available, you can also define a subcategory.
	//
	// Although you can specify more than one category and subcategory in your
	// RSS feed, Apple Podcasts only recognizes the first category and
	// subcategory.
	//
	// When specifying categories and subcategories, be sure to properly escape
	// ampersands. For example:
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
	ItunesCategory []ItunesCategory `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd category"`

	// # Apple Podcasts:
	//
	// ItunesExplicit (required) is the podcast parental advisory information.
	//
	// The explicit value can be one of the following:
	// - True. If you specify true, indicating the presence of explicit content,
	//   Apple Podcasts displays an [Explicit] parental advisory graphic for your
	//   podcast.
	//
	//   Podcasts containing explicit material aren’t available in some Apple
	//   Podcasts territories.
	//
	// - False. If you specify false, indicating that your podcast doesn’t
	//   contain explicit language or adult content, Apple Podcasts displays a
	//   [Clean] parental advisory graphic for your podcast.
	//
	// [Explicit]: https://help.apple.com/itc/podcasts_connect/#/itcfafb6d665
	// [Clean]: https://help.apple.com/itc/podcasts_connect/#/itcb343e8058
	ItunesExplicit *bool `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd explicit,omitempty"`

	// # Apple Podcasts:
	//
	// ItunesAuthor (recommended) is the group responsible for creating the show
	//
	// Show author most often refers to the parent company or network of a
	// podcast, but it can also be used to identify the host(s) if none exists.
	//
	// Author information is especially useful if a company or organization
	// publishes multiple podcasts.
	ItunesAuthor string `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd author,omitempty"`

	// # Apple Podcasts:
	//
	// ItunesTitle (situational) is he show title specific for Apple Podcasts.
	//
	// <itunes:title> is a string containing a clear concise name of your show
	// on Apple Podcasts. Do not include episode or season number in the title.
	// There are dedicated tags for that information. See <itunes:episode> and
	// <itunes:season>.
	ItunesTitle string `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd title,omitempty"`

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
	ItunesType ItunesShowType `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd type,omitempty"`

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
	ItunesNewFeedURL string `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd new-feed-url,omitempty"`

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
	ItunesComplete ItunesYesType `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd complete,omitempty"`

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
	ItunesBlock ItunesYesType `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd block,omitempty"`

	// # Apple Podcasts:
	//
	// ItunesApplePodcastVerify (situational) is used to verify ownership of a show when a
	// podcast creator chooses to move it from one Apple Podcasts Connect
	// account to another. [Learn more about how to claim your show].
	//
	// [Learn more about how to claim your show]: https://podcasters.apple.com/support/5497-claim-your-show
	ItunesApplePodcastVerify string `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd applepodcastsverify,omitempty"`

	// # Apple Podcasts:
	//
	// PodcastTxt (situational) is an alternate method to verify your show
	// . See ApplePodcast.ItunesApplePodcastVerify
	PodcastTxt *PodcastTxt `xml:"https://podcastindex.org/namespace/1.0 txt,omitempty"`

	// PodcastPodroll allows you to recommend other podcasts
	PodcastPodroll *PodcastPodroll `xml:"https://podcastindex.org/namespace/1.0 podroll,omitempty"`

	// PodcastLocked requires permission to migrate a feed
	PodcastLocked *PodcastLocked `xml:"https://podcastindex.org/namespace/1.0 locked,omitempty"`

	// PodcastFunding are links to financially support a show
	PodcastFunding []PodcastFunding `xml:"https://podcastindex.org/namespace/1.0 funding,omitempty"`

	// PodcastPerson provides credits for hosts and guests. Publishers are
	// expected to use the podcast:person element in the <channel> parent to set
	// the regular people involved in the podcast: the detail that would be
	// expected to be seen in an overview of the show.
	//
	// It is suggested that <channel> is always populated, and <item> is populated
	// where needed for an individual episode.
	PodcastPeople []PodcastPerson `xml:"https://podcastindex.org/namespace/1.0 person,omitempty"`

	// PodcastLocation tells where is this podcast about
	PodcastLocation *PodcastLocation `xml:"https://podcastindex.org/namespace/1.0 location,omitempty"`

	// PodcastTrailer provides a show or season trailer
	PodcastTrailers []PodcastTrailer `xml:"https://podcastindex.org/namespace/1.0 trailer,omitempty"`

	// PodcastValue enables streaming micropayments to the podcast
	PodcastValue []PodcastValue `xml:"https://podcastindex.org/namespace/1.0 value,omitempty"`

	// PodcastMedium triggers a UI based on the type of content
	PodcastMedium PodcastMedium `xml:"https://podcastindex.org/namespace/1.0 medium,omitempty"`

	// PodcastLiveItem provides live-streaming within podcast apps
	PodcastLiveItem []PodcastLiveItem `xml:"https://podcastindex.org/namespace/1.0 liveItem,omitempty"`

	// PodcastGUID provides a globally unique identifier for the podcast (uuidv5)
	//
	// See [spec]
	//
	// [spec]: https://podcastindex.org/namespace/1.0#guid
	PodcastGUID string `xml:"https://podcastindex.org/namespace/1.0 guid,omitempty"`

	// PodcastBlock requests to not be consumable by an app or directory
	PodcastBlock []PodcastBlock `xml:"https://podcastindex.org/namespace/1.0 block,omitempty"`

	// PodcastLicense indicates the show's license
	PodcastLicense *PodcastLicense `xml:"https://podcastindex.org/namespace/1.0 license"`

	Items []Episode `xml:"https://www.rssboard.org/rss-specification item"`
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

type Episode struct {
	rss.Item

	// # Apple Podcasts:
	//
	// ItunesDuration (recommended) is the duration of an episode.
	//
	// Different duration formats are accepted however it is recommended to
	// convert the length of the episode into seconds.
	ItunesDuration string `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd duration,omitempty"`

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
	ItunesImage *ItunesImageTag `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd image,omitempty"`

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
	ItunesExplicit *bool `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd explicit"`

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
	ItunesTitle string `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd title,omitempty"`

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
	ItunesEpisode int `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd episode,omitempty"`

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
	ItunesSeason int `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd season,omitempty"`

	// # Apple Podcasts:
	//
	// ItunesEpisodeType (situational) is the episode type.
	//
	// If an episode is a trailer or bonus content, use this tag.
	ItunesEpisodeType EpisodeType `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd episodeType,omitempty"`

	// # Apple Podcasts (situational):
	//
	// ItunesBlock is the episode show or hide status.
	// If you want an episode removed from the Apple directory, use this tag.
	// Specifying the <itunes:block> tag with a Yes value prevents that episode from appearing in Apple Podcasts.
	// For example, you might want to block a specific episode if you know that its content would otherwise cause the entire podcast to be removed from Apple Podcasts.
	// Specifying any value other than Yes has no effect.
	ItunesBlock ItunesYesType `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd block,omitempty"`

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
	PodcastTranscript []PodcastTranscript `xml:"https://podcastindex.org/namespace/1.0 transcript,omitempty"`

	// PodcastChapters provides independently editable, enhanced chapters
	PodcastChapters *PodcastChapters `xml:"https://podcastindex.org/namespace/1.0 chapters,omitempty"`

	// PodcastSoundbite specifies suggested clips for sharing and promotion
	PodcastSoundbite []PodcastSoundbite `xml:"https://podcastindex.org/namespace/1.0 soundbite,omitempty"`

	// PodcastPerson provides credits for hosts and guests. Where present,
	// people information in <item> wholly replaces all information from the
	// <channel>.
	//
	// Publishers are expected to use the podcast:person in the <item> parent to
	// replace all existing information for an individual episode.
	PodcastPeople []PodcastPerson `xml:"https://podcastindex.org/namespace/1.0 person,omitempty"`

	// PodcastSeason provides named or numbered seasons/series
	PodcastSeason *PodcastSeason `xml:"https://podcastindex.org/namespace/1.0 season,omitempty"`

	// PodcastEpisode provides episode numbers and names
	PodcastEpisode *PodcastEpisode `xml:"https://podcastindex.org/namespace/1.0 episode,omitempty"`

	// PodcastLicense indicates the episode's license
	PodcastLicense *PodcastLicense `xml:"https://podcastindex.org/namespace/1.0 license,omitempty"`

	PodcastAlternateEnclosures []PodcastAlternateEnclosure `xml:"https://podcastindex.org/namespace/1.0 alternateEnclosure,omitempty"`

	// PodcastValue enables streaming micropayments to the podcast
	PodcastValue []PodcastValue `xml:"https://podcastindex.org/namespace/1.0 value,omitempty"`

	// PodcastImages allows multiple image resources for the podcast
	PodcastImages *PodcastImages `xml:"https://podcastindex.org/namespace/1.0 images"`

	// PodcastSocialInteracts provides comments on podcast episodes
	PodcastSocialInteracts []PodcastSocialInteract `xml:"https://podcastindex.org/namespace/1.0 socialInteract,omitempty"`

	// PodcastUpdateFrequency documents the intended release schedule
	PodcastUpdateFrequency *PodcastUpdateFrequency `xml:"https://podcastindex.org/namespace/1.0 updateFrequency"`

	// PodcastPodping indicates if the podcast uses Podping
	PodcastPodping *PodcastPodping `xml:"https://podcastindex.org/namespace/1.0 podping,omitempty"`
}

// ItunesYes is meant to be used with various properties that either accept a
// a `Yes` value to opt in or anything else to opt out
type ItunesYesType string

// ItunesYesValue is the sentinel yes value
const ItunesYesValue ItunesYesType = "Yes"

// ItunesCategory (required) is the show category information. For a
// complete list of categories and subcategories, see [Apple Podcast categories].
//
// [Apple Podcast Categories]: https://podcasters.apple.com/support/1691-apple-podcasts-categories
type ItunesCategory struct {
	XMLName xml.Name `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd category"`

	SubCategory *struct {
		XMLName xml.Name `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd category"`
		Text    string   `xml:"text,attr"`
	} `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd category,omitempty"`

	// Text is an enum of the actual category (or sub category). See [Podcast.ItunesCategory]
	// and the [list of categories] on Apple support for valid values
	//
	// [list of categories]: https://podcasters.apple.com/support/1691-apple-podcasts-categories
	Text string `xml:"text,attr"`
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

// PodcastTxt holds free-form text and is modeled after the DNS "TXT" record.
//
// See [spec]
//
// # Apple Podcasts:
//
// Alternate way to verify a podcast using <podcast:txt purpose=“applepodcastsverify”>
// vs <itunes:applepodcastsverify>
//
// [spec]: https://podcastindex.org/namespace/1.0#txt
type PodcastTxt struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 txt"`
	// Value is a free form string. Please do not exceed 4000 characters for the
	// node value or it may be truncated by aggregators.
	Value string `xml:",chardata"`

	// Purpose (optional) is a service specific string that will be used to
	// denote what purpose this tag serves. This could be something like
	// "example.com" if it's a third party hosting platform needing to insert
	// this data, or something like verify", "release" or any other free form
	// bit of info that is useful to the end consumer that needs it. The free
	// form nature of this tag requires that this attribute is also free formed.
	// This value should not exceed 128 characters.
	//
	// # Apple Podcasts
	//
	// You probably want
	// 	applepodcastsverify
	Purpose string `xml:"purpose,attr,omitempty"`
}

// PodcastTranscript (podcast 2.0) provides timestamped captions and transcripts
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#transcript
type PodcastTranscript struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 transcript"`

	// URL (required) is the url of the podcast transcript.
	URL string `xml:"url,attr"`

	// Type (required) is the Mime type of the file such as:
	// - text/plain
	// - text/html
	// - text/vtt
	// - application/json
	// - application/x-subrip
	Type string `xml:"type,attr"`

	// Language (optional) of the linked transcript. If there is no language
	// attribute given, the linked file is assumed to be the same language that
	// is specified by the RSS <language> element.
	Language string `xml:"language,attr,omitempty"`

	// Rel (optional) if rel="captions" attribute is present, the linked file is
	// considered to be a closed captions file, regardless of what the mime type
	// is. In  that scenario, time codes are assumed to be present in the file
	// in some capacity.
	Rel string `xml:"rel,attr,omitempty"`
}

// PodcastChapters provides independently editable, enhanced chapters
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#chapters
type PodcastChapters struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 chapters"`

	// URL (required) is where the chapters file is located.
	URL string `xml:"url,attr"`

	// Type (required) is the mime type of file - JSON preferred,
	//   'application/json+chapters'.
	Type string `xml:"type,attr"`
}

// Podroll allows for a podcaster to include references to one or more podcasts
// in it's <channel> as a way of "recommending" other podcasts to their listener.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#podroll
type PodcastPodroll struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 podroll"`

	RemoteItems []PodcastRemoteItem `xml:"https://podcastindex.org/namespace/1.0 remoteItem"`
}

// PodcastRemoteItem provides a way to "point" to another feed or an <item> in
// another feed in order to obtain some sort of data that the other feed or feed
// item has.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#remote-item
type PodcastRemoteItem struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 remoteItem"`

	// FeedGUID (required): The <podcast:guid> of the remote feed being pointed
	// to.
	FeedGUID string `xml:"feedGuid,attr"`

	// FeedURL (optional) is the url of the remote feed being pointed to.
	FeedURL string `xml:"feedUrl,attr,omitempty"`

	// ItemGUID (optional) If this remote item element is intended to point to
	// an <item> in the remote feed, this attribute should contain the value of
	// the <guid> of that <item>.
	ItemGUID string `xml:"itemGuid,attr,omitempty"`

	// Medium (optional) if the feed being pointed to is not of medium type
	// 'podcast', this attribute should contain it's <podcast:medium> type
	// from the [list] of types available in this document. The reason this is
	// helpful is to give the app a heads up on what type of data this is
	// expected to be since that may affect the way it approaches fetch and
	// display.
	//
	// [list]: https://podcastindex.org/namespace/1.0#medium
	Medium string `xml:"medium,attr,omitempty"`
}

// PodcastLocked requires permission to migrate a feed
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#locked
type PodcastLocked struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 locked"`

	// Values specifies if it is locked or not
	//
	// The node value must be yes or no.
	Value YesOrNo `xml:",chardata"`

	// Owner (optional): The owner attribute is an email address that can be
	// used to verify ownership of this feed during move and import operations.
	// This could be a public email or a virtual email address at the hosting
	// provider that redirects to the owner's true email address.
	Owner string `xml:"owner,attr,omitempty"`
}

type YesOrNo string

const (
	Yes YesOrNo = "yes"
	No  YesOrNo = "no"
)

// PodcastFunding lists possible donation/funding links for the podcast. The
// content of the tag is the recommended string to be used with the link.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#funding
type PodcastFunding struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 funding"`

	// This is a free form string supplied by the creator which they expect to
	// be displayed in the app next to the link. Please do not exceed 128
	// characters for the node value or it may be truncated by aggregators.
	Value YesOrNo `xml:",chardata"`

	// URL (required): The URL to be followed to fund the podcast.
	URL string `xml:"url,attr"`
}

// PodcastSoundbite Points to one or more soundbites within a podcast episode.
// The intended use includes episodes previews, discoverability, audiogram
// generation, episode highlights, etc. It should be assumed that the
// audio/video source of the soundbite is the audio/video given in the item's
// <enclosure> element.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#soundbite
type PodcastSoundbite struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 soundbite"`

	// StarTime (required): The time where the soundbite begins
	StartTime float32 `xml:"startTime,attr"`

	// Duration (required): How long is the soundbite (recommended between 15
	// and 120 seconds)
	Duration float32 `xml:"duration,attr"`

	// This is a free form string from the podcast creator to specify a title
	// for the soundbite. If the podcaster does not provide a value for the
	// soundbite title, then leave the value blank, and podcast apps can decide
	// to use the episode title or some other placeholder value in its place.
	// Please do not exceed 128 characters for the node value or it may be
	// truncated by aggregators.
	Value string `xml:",chardata"`
}

// PodcastPerson specifies a person of interest to the podcast. It is primarily
// intended to identify people like hosts, co-hosts and guests. Although, it is
// flexible enough to allow fuller credits to be given using the roles and
// groups that are listed in the Podcast Taxonomy Project
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#person
type PodcastPerson struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 person"`

	Text string `xml:",chardata"`

	// Group (optional): This should be a reference to an official group within
	// the Podcast Taxonomy Project list. If group is not present, then "cast"
	// is assumed.
	Group string `xml:"group,attr,omitempty"`

	// Role (optional) is used to identify what role the person serves on the
	// show or episode. This should be a reference to an official role within
	// the Podcast Taxonomy Project [list]. If role is missing then "host" is
	// assumed.
	//
	// [list]: https://github.com/Podcastindex-org/podcast-namespace/blob/main/taxonomy.json
	Role string `xml:"role,attr,omitempty"`

	// Href (optional) is the url to a relevant resource of information about
	// the person, such as a homepage of third-party profile platform. Please
	// see the [example feed] for possible choices of what to use here.
	//
	// [example feed]: https://github.com/Podcastindex-org/podcast-namespace/blob/main/example.xml
	Href string `xml:"href,attr"`

	// Img is the url of a picture or avatar of the person.
	Img string `xml:"img,attr"`
}

// PodcastLocation describe the location of editorial focus for a podcast's
// content (i.e. "what place is this podcast about?"). The tag has many use
// cases and is one of the more complex ones. You are highly encouraged to
// read the full [implementation document] before starting to code for it.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#location
// [implementation document]: https://github.com/Podcastindex-org/podcast-namespace/blob/main/location/location.md
type PodcastLocation struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 location"`

	// Text (required) is a free-form string meant to be a human readable
	// location. It may conform to conventional location verbiage
	// (i.e. "Austin, TX"), but it  shouldn't be depended on to be parseable in
	// any specific way. This value cannot be blank. Please do not exceed 128
	// characters for the node value or it may be truncated by aggregators.
	Text string `xml:",chardata"`

	// Geo (recommended): This is a latitude and longitude given in "geo"
	// notation (i.e. "geo:30.2672,97.7431").
	Geo string `xml:"geo,attr,omitempty"`

	// Osm(recommended): The Open Street Map identifier of this place, given using
	// the OSM notation (i.e. "R113314")
	Osm string `xml:"osm,attr,omitempty"`
}

// PodcastSeason allows for identifying which episodes in a podcast are part of
// a particular "season", with an optional season name attached.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#season
type PodcastSeason struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 season"`

	// Value (required) value is an integer, and represents the season "number"
	Value int `xml:",chardata"`

	// Name (optional) of the season. If this attribute is present,
	// applications are free to not show the season number to the end user,
	// and may use it simply for chronological sorting and grouping purposes.
	//
	// Please do not exceed 128 characters for the name attribute.
	Name string `xml:"name,attr,omitempty"`
}

// PodcastEpisode exists largely for compatibility with the season tag. But, it
// also allows for a similar idea to what "name" functions as in that element.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#episode
type PodcastEpisode struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 episode"`

	// Value (required): the node value is a decimal number
	//
	// The episode numbers are decimal, so numbering such as 100.5 is acceptable
	//  if there was a special mini-episode published between two other
	// episodes. In that scenario, the number would help with proper
	// chronological sorting, while the display attribute could specify an
	// alternate special "number" (a moniker) to display for the episode in a
	// podcast player app UI.
	Value float32 `xml:",chardata"`

	// Display (optional): If this attribute is present, podcast apps and
	// aggregators are encouraged to show its value instead of the purely
	// numerical node value.
	//
	// Please do not exceed 32 characters for the display attribute.
	Display string `xml:"display,attr,omitempty"`
}

// PodcastTrailer is used to define the location of an audio or video file to be
// used as a trailer for the entire podcast or a specific season. There can be
// more than one trailer present in the channel of the feed. This element is
// basically just like an <enclosure> with the extra pubdate and season
// attributes added.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#trailer
type PodcastTrailer struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 trailer"`

	// Text (required) is the title of the trailer.
	//
	// Please do not exceed 128 characters for the node value or it may be
	// truncated by aggregators.
	Text string `xml:",chardata"`

	// (required) Pubdate is the date the trailer was published. This attribute is an
	// RFC2822 formatted date string.
	Pubdate rss.RFC2822Date `xml:"pubdate,attr"`

	// URL (required) is a url that points to the audio or video file to be
	// played.
	URL string `xml:"url,attr"`

	// Length c(recommended) of the file in bytes
	Length int64 `xml:"length,attr,omitempty"`

	// Type (recommended): The mime type of the file.
	Type string `xml:"type,attr,omitempty"`

	// Season (optional): If this attribute is present it specifies that this
	// trailer is for a particular season number.
	//
	// If the the season attribute is present, it must be a number that matches
	// the format of the <podcast:season> tag. So, for a podcast that has 3
	// published seasons, a new <podcast:trailer season="4"> tag can be put in
	// the channel to later be matched up with a
	// <podcast:season>4<podcast:season> tag when it is published within a new
	// <item>.
	Season string `xml:"season,attr,omitempty"`
}

// PodcastLicense defines a license that is applied to the audio/video content
// of a single episode, or the audio/video of the podcast as a whole. Custom
// licenses must always include a url attribute. Implementors are encouraged
// to read the license tag [companion document] for a more complete picture of
// what this tag is intended to accomplish.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#license
// [companion document]: https://github.com/Podcastindex-org/podcast-namespace/blob/main/proposal-docs/license/license.md
type PodcastLicense struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 license"`

	// Value (required) must be a lower-cased reference to a license
	// "identifier" defined in the [SPDX License List] file if the license being
	// used is a well-known, public license. Or, if it is a custom license,
	// it must be a free form abbreviation of the name of the license as you
	// reference it publicly. Please do not exceed 128 characters for the node
	// value or it may be truncated by aggregators.
	//
	// [SPDX License List]: https://spdx.org/licenses/
	Value string `xml:",chardata"`

	// URL (optional): This is a url that points to the full, legal language of
	// the license being referenced. This attribute is optional for well-known
	// public licenses. For new, or custom licenses it is required.
	URL string `xml:"url,attr,omitempty"`
}

// This element is meant to provide different versions of, or companion media to
//  the main <enclosure> file. This could be an audio only version of a video
// podcast to allow apps to switch back and forth between audio/video, lower
// (or higher) bitrate versions for bandwidth constrained areas, alternative
// codecs for different device platforms, alternate URI schemes and download
// types such as IPFS or WebTorrent, commentary tracks or supporting source
// clips, etc.

// This is a complex tag, so implementors are highly encouraged to read the
// [companion document] for a fuller understanding of how this tag works and
// what it is capable of.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#alternate-enclosure
// [companion document]: https://github.com/Podcastindex-org/podcast-namespace/blob/main/proposal-docs/alternateEnclosure/alternateEnclosure.md
type PodcastAlternateEnclosure struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 alternateEnclosure"`

	// Type (required): Mime type of the media asset.
	Type string `xml:"type,attr"`

	// Length (recommended): Length of the file in bytes.
	Length string `xml:"length,attr,omitempty"`

	// Bitrate (optional): Average encoding bitrate of the media asset,
	// expressed in bits per second.
	Bitrate float32 `xml:"bitrate,attr,omitempty"`

	// Title (optional): A human-readable string identifying the name of the
	// media asset. Should be limited to 32 characters for UX.
	Title string `xml:"title,attr,omitempty"`

	// Height (optional): Height of the media asset for video formats.
	Height int `xml:"height,attr,omitempty"`

	// Lang (optional): An [IETF language tag (BCP 47) code] identifying the
	// language of this media.
	//
	// [IETF language tag (BCP 47) code]: https://en.wikipedia.org/wiki/BCP_47
	Lang int `xml:"lang,attr,omitempty"`

	// Rel (optional): Provides a method of offering and/or grouping together
	// different media elements. If not set, or set to "default", the media will
	// be grouped with the enclosure and assumed to be an alternative to the
	// enclosure's encoding/transport. This attribute can and should be the same
	// for items with the same content encoded by different means. Should be
	// limited to 32 characters for UX.
	Rel int `xml:"rel,attr,omitempty"`

	// Codecs (optional): An [RFC 6381] string specifying the codecs available in this media.
	//
	// [RFC 6381]: https://tools.ietf.org/html/rfc6381
	Codecs int `xml:"codecs,attr,omitempty"`

	// default: Boolean specifying whether or not the given media is the same as
	// the file from the enclosure element and should be the preferred media
	// element. The primary reason to set this is to offer alternative
	// transports for the enclosure. If not set, this should be assumed to be false.
	Default *bool `xml:"default,attr,omitempty"`

	// PodcastIntegrity provides a method of verifying integrity of the media
	PodcastIntegrity *PodcastIntegrity `xml:"https://podcastindex.org/namespace/1.0 integrity,omitempty"`

	Source []PodcastSource `xml:"https://podcastindex.org/namespace/1.0 source,omitempty"`
}

// PodcastSource defines a uri location for a <podcast:alternateEnclosure> media
// file. It is meant to be used as a child of the <podcast:alternateEnclosure>
// element. At least one <podcast:source> element must be present within every
// <podcast:alternateEnclosure> element.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#source
type PodcastSource struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 source"`

	// URI (required): This is the URI where the media file resides.
	URI string `xml:"uri,attr"`

	// ContentType: This is a string that declares the mime-type of the file.
	// It is useful if the transport mechanism is different than the file being
	// delivered, as is the case with a torrents.
	ContentType string `xml:"contentType,attr,omitempty"`
}

// PodcastIntegrity defines a method of verifying integrity of the media given
// either an [SRI-compliant integrity string] (preferred) or a base64 encoded
// PGP signature. This element is optional within a <podcast:alternateEnclosure>
// element. It allows to ensure that the file has not been tampered with.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#integrity
// [SRI-compliant integrity string]: (https://www.w3.org/TR/SRI/)
type PodcastIntegrity struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 integrity"`
	// Type (required): Type of integrity, either "sri" or "pgp-signature".
	Type string `xml:"type,attr"`
	// Value (required): Value of the sri string or base64 encoded pgp signature.
	Value string `xml:"value,attr"`
}

// PodcastValue enables crypto garbage to infect even podcasting
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#value
type PodcastValue struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 value"`

	// Type (required): This is the service slug of the cryptocurrency or
	// protocol layer.
	Type string `xml:"type,attr"`

	// Method (required): This is the transport mechanism that will be used.
	Method string `xml:"method,attr"`

	// Suggested (optional) suggestion on how much cryptocurrency to send with
	// each payment.
	Suggested string `xml:"suggested,attr,omitempty"`

	// Recipients get the crypto
	Recipients []PodcastValueRecipient `xml:"https://podcastindex.org/namespace/1.0 valueRecipient,omitempty"`

	// ValueTimeSplits enables time-based wallet-switching for streaming micropayments
	ValueTimeSplits []PodcastValueTimeSplit `xml:"https://podcastindex.org/namespace/1.0 valueTimeSplit,omitempty"`
}

// PodcastValueRecipient designates various destinations for payments to be sent
// to during consumption of the enclosed media. Each recipient is considered to
// receive a "split" of the total payment according to the number of shares
// given in the split attribute.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#value-recipient
type PodcastValueRecipient struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 valueRecipient"`
	// Name (recommended): A free-form string that designates who or what this
	// recipient is.
	Name string `xml:"name,attr"`

	// CustomKey (optional): The name of a custom record key to send along with
	// the payment.
	CustomKey string `xml:"customKey,attr,omitempty"`

	// customValue (optional):: A custom value to pass along with the payment.
	// This is considered the value that belongs to the customKey.
	CustomValue string `xml:"customValue,attr,omitempty"`

	// Type (required): A slug that represents the type of receiving address
	// that will receive the payment.
	Type string `xml:"type,attr"`

	// Address (required): This denotes the receiving address of the payee.
	Address string `xml:"address,attr"`

	// Split (required): The number of shares of the payment this recipient will
	// receive.
	Split string `xml:"split,attr"`

	// Fee: If this attribute is not specified, it is assumed to be false.
	Fee *bool `xml:"fee,attr,omitempty"`
}

// PodcastMedium tells an application what the content contained within the feed
// IS, as opposed to what the content is ABOUT in the case of a category.
//
// each medium also has a counterpart "list" variant, where the original medium
// name is suffixed by the letter "L" to indicate that it is a "List" of that
// type of content. For example, podcast would become podcastL, music would
// become musicL, audiobook would become audiobookL, etc.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#medium
type PodcastMedium string

const (
	// MediumPodcast (default) - Describes a feed for a podcast show. If no medium tag
	// is present in the channel, this medium is assumed.
	MediumPodcast PodcastMedium = "podcast"
	// MediumMusic - A feed of music organized into an "album" with each item a song
	// within the album.
	MediumMusic = "music"
	// MediumVideo - Like a "podcast" but used in a more visual experience. Something
	// akin to a dedicated video channel like would be found on YouTube.
	MediumVideo = "video"
	// MediumFilm - Specific types of videos with one item per feed. This is different
	// than a video medium because the content is considered to be cinematic;
	// like a movie or documentary.
	MediumFilm = "film"
	// MediumAudiobook - Specific types of audio with one item per feed, or where
	// items represent chapters within the book.
	MediumAudiobook = "audiobook"
	// MediumNewsletter - Describes a feed of curated written articles. Newsletter
	// articles now sometimes have an spoken version audio enclosure attached.
	MediumNewsletter = "newsletter"
	// MediumBlog - Describes a feed of informally written articles. Similar to
	// newsletter but more informal as in a traditional blog platform style.
	MediumBlog = "blog"

	MediumPodcastList    = "podcastL"
	MediumMusicList      = "musicL"
	MediumVideoList      = "videoL"
	MediumFilmList       = "filmL"
	MediumAudiobookList  = "audiobookL"
	MediumNewsletterList = "newsletterL"
	MediumBlogList       = "blogL"
)

// PodcastImages allows for specifying many different image sizes in a compact
// way at either the episode or channel level
//
// See [spec] and [srcset] syntax
//
// [spec]: https://podcastindex.org/namespace/1.0#images
// [srcset]: https://html.spec.whatwg.org/multipage/images.html#srcset-attributes
type PodcastImages struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 images"`
	//     srcset (required): A string that denotes each image url followed by a
	//  space and the pixel width, with each one separated by a comma. See the
	// example for a clear view of the syntax.
	Srcset string `xml:"srcset,attr"`
}

// PodcastLiveItem s used for a feed to deliver a live audio or video stream to
// podcast apps.
//
// A robust, well-written <podcast:liveItem> tag will include all three of:
// <podcast:alternateEnclosure>, <enclosure> and <podcast:contentLink> to ensure
// The broadest interopability with podcast apps.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#live-item
type PodcastLiveItem struct {
	Episode
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 liveItem"`

	// Status (required): A string that must be one of pending, live or ended.
	Status string `xml:"status,attr"`

	// Start (required): A string representing an ISO8601 timestamp that denotes
	//  the time when the stream is intended to start.
	Start ISO8601Timestamp `xml:"start,attr"`

	// End (recommended): A string representing an ISO8601 timestamp that
	// denotes the time when the stream is intended to end.
	End ISO8601Timestamp `xml:"end,attr"`

	// PodcastContentLinks provides links to content delivered outside the app
	PodcastContentLinks []PodcastContentLink `xml:"https://podcastindex.org/namespace/1.0 contentLink"`
}

// PodcastLiveStreamStatus tells you the status of a live stream
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#live-item
type PodcastLiveStreamStatus string

const (
	StatusPending PodcastLiveStreamStatus = "pending"
	StatusLive    PodcastLiveStreamStatus = "live"
	StatusEnded   PodcastLiveStreamStatus = "ended"
)

type ISO8601Timestamp string

// PodcastContentLink is used to indicate that the content begin delivered by
// the parent element can be found at an external location instead of, or in
// addition to, being delivered directly to the tag itself within an app
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#social-interact
type PodcastContentLink struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 contentLink"`

	// Value is a free form string meant to explain to the user where this
	// content link points and/or the nature of it's purpose.
	Value string `xml:",chardata"`

	// Href (required): A string that is the uri pointing to content outside of
	// the application.
	Href string `xml:"href,attr"`
}

// PodcastSocialInteract allows a podcaster to attach the url of a "root post"
// of a comment thread to an episode.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#social-interact
type PodcastSocialInteract struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 socialInteract"`

	// Priority (optional): When multiple socialInteract tags are present, this integer
	// gives order of priority. A lower number means higher priority.
	Priority int `xml:"priority,attr"`

	// URI (required): The uri/url of root post comment.
	URI string `xml:"uri,attr"`

	// Protocol (required): The [protocol] in use for interacting with the
	// comment root post.
	//
	// [protocol]: https://podcastindex.org/socialprotocols.txt
	Protocol string `xml:"protocol,attr"`

	// AccountId (recommended): The account id (on the commenting platform) of
	// the account that created this root post.
	AccountId string `xml:"accountId,attr,omitempty"`

	// AccountUrl (optional): The public url (on the commenting platform) of the
	// account that created this root post.
	AccountUrl string `xml:"accountUrl,attr,omitempty"`
}

// PodcastBlock allows a podcaster to express which platforms are allowed to
// publicly display this feed and its contents.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#block
type PodcastBlock struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 block"`

	Value YesOrNo `xml:",chardata"`

	// ID: A single entry from the [service slug] list.
	//
	// [service slug]: https://github.com/Podcastindex-org/podcast-namespace/blob/main/serviceslugs.txt
	ID string `xml:"id,attr,omitempty"`
}

// PodcastUpdateFrequency allows a podcaster to express their intended release
// schedule as structured data and text.
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#update-frequency
type PodcastUpdateFrequency struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 updateFrequency"`

	// Value is a free-form string, which might be displayed alongside other
	// information about the podcast. Please do not exceed 128 characters
	// for the node value or it may be truncated by aggregators.
	Value string `xml:,chardata`

	// Complete (optional): Boolean specifying if the podcast has no intention
	// to release further episodes. If not set, this should be assumed to be
	// false.
	Complete bool `xml:"complete,attr"`

	// Dtstart (optional): The date or datetime the recurrence rule begins as an
	// [ISO8601-fornmatted] string. If the rrule contains a value for COUNT,
	// then this attribute is required. If the rrule contains a value for UNTIL,
	//  then the value of this attribute must be formatted to the same
	// date/datetime standard.
	//
	// [ISO8601-fornmatted]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString
	Dtstart string `xml:"dtstart,attr"`

	// Rrule (recommended): A recurrence rule as defined in
	// [iCalendar RFC 5545 Section 3.3.10].
	//
	// [iCalendar RFC 5545 Section 3.3.10]: https://www.rfc-editor.org/rfc/rfc5545#section-3.3.10
	Rrule string `xml:"rrule,attr"`
}

// PodcastPodping allows feed owners to signal to aggregators that the feed
// sends out [Podping] notifications when changes are made to it, reducing the
// need for frequent speculative feed polling.
//
// See [spec] and [Podping]
//
// [spec]: https://podcastindex.org/namespace/1.0#podping
// [Podping]: https://github.com/Podcastindex-org/podping
type PodcastPodping struct {
	XMLName     xml.Name `xml:"https://podcastindex.org/namespace/1.0 podping"`
	UsesPodping bool     `xml:"usesPodping,attr"`
}

// PodcastUpdateFrequency allows different value splits for a certain period of
// time. For use with [PodcastValue]
//
// # I assume that sentence makes sense to crypto-pilled people
//
// See [spec]
//
// [spec]: https://podcastindex.org/namespace/1.0#value-time-split
type PodcastValueTimeSplit struct {
	XMLName xml.Name `xml:"https://podcastindex.org/namespace/1.0 valueTimeSplit"`

	PodcastRemoteItem      *PodcastRemoteItem      `xml:"https://podcastindex.org/namespace/1.0 remoteItem,omitempty"`
	PodcastValueRecipients []PodcastValueRecipient `xml:"https://podcastindex.org/namespace/1.0 valueRecipient"`

	// StartTime (required): The time, in seconds, to stop using the currently
	// active value recipient information and start using the value recipient
	// information contained in this element.
	StartTime int `xml:"startTime,attr"`

	// Duration (required): How many seconds the playback app should use this
	// element's value recipient information before switching back to the
	// value recipient information of the parent feed.
	Duration int `xml:"duration,attr"`

	// RemoteStartTime: The time in the remote item where the value split
	// begins. Allows the timestamp to be set correctly in value metadata. If
	// not defined, defaults to 0.
	RemoteStartTime string `xml:"remoteStartTime,attr"`

	// RemotePercentage: The percentage of the payment the remote recipients
	// will receive if a <podcast:remoteItem> is present. If not defined,
	// defaults to 100. If the value is less than 0, 0 is assumed. If the value
	// is greater than 100, 100 is assumed.
	RemotePercentage string `xml:"remotePercentage,attr"`
}
