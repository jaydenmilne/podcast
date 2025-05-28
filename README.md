# Go Podcast RSS Feed Types

This package provides type definitions to create podcasts RSS feeds using go's
`encoding/xml`.

It fully supports:

- [RSS 2.0](https://www.rssboard.org/rss-specification)
- [Apple Podcasts](https://help.apple.com/itc/podcasts_connect/#/itcb54353390)
- [Podcasting 2.0 / `podcast:` namespace](https://podcastindex.org/namespace/1.0)

You should also be able to use this package to parse well-formed podcasts, but
you may want to use a more robust parser such as [gofeed](https://github.com/mmcdole/gofeed).

This package aims to be:

* Just types. Be unopinionated, annotated structs with little to no associated 
  code
* Not stringly typed
* Include snippets from the standards docs in the godocs so you know what you
  need to include
* No dependencies (there is one test dependency)

## Examples

### Generate a podcast

```go
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
```

### Parse a well-formed podcast
```go
decoder := xml.NewDecoder(strings.NewReader(feedXML))
decoder.DefaultSpace = "https://www.rssboard.org/rss-specification"

var decoded RSSPodcast
decoder.Decode(&decoded)
```

## RSS Package

It also provides an RSS package that you should also be able to use to parse 
and generate simple RSS feeds
