# Crawler

This is a web crawler project.

## Running the Crawler

To run the crawler, use the following command with three attributes: `url`, `max_concurrency`, and `max_posts`.

```sh
./crawler <url> <max_concurrency> <max_posts>
```

- `url`: The starting URL for the crawler.
- `max_concurrency`: The maximum number of concurrent requests.
- `max_posts`: The maximum number of posts to crawl.

Example:

```sh
./crawler http://example.com 5 100
```

This command will start crawling `http://example.com` with a maximum of 5 concurrent requests and will stop after crawling 100 posts.