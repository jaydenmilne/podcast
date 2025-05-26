# Go Podcast Feed Generator

This package provides type definitions to create podcasts RSS feeds. 

It fully supports:

- [RSS 2.0](https://www.rssboard.org/rss-specification)
- [Apple Podcasts](https://help.apple.com/itc/podcasts_connect/#/itcb54353390)
- [Podcasting 2.0 / `podcast:` namespace](https://podcastindex.org/namespace/1.0)

You should also be able to use this package to parse well-formed podcasts, but
you may want to use a more robust parser such as [gofeed](https://github.com/mmcdole/gofeed).

This package aims to:

* Be unopinionated, structs with little to no associated code
* Not be stringly typed
* Include snippets from the standards docs in the godocs

## Examples

### Generate a podcast

### Parse a well-formed podcast
