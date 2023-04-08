# JLPT-Notify
A small Go program that checks the [JLPT Leiden](https://jlpt-leiden.nl) website periodically and notifies me by SMS whenever the source code changes.

## Why?
I missed out on registrations for the JLPT exam this year because I couldn't register in time. I want
to avoid a similar situation next year by getting a notification the moment registrations open up.

## Installation
To install and use this program yourself, you may refer to the following steps:
```bash
git clone https://github.com/bjvanbemmel/jlpt-notify
cd jlpt-notify
docker compose up -d
```

It's possible to run this program without Docker as well, but be aware that my implementation of Twilio uses environment variables to fetch the credentials.
You'd have to export the variables yourself, or include some way for the program to load an environment file by itself.

## Contribute
- For bug reports, please write an [issue](https://github.com/bjvanbemmel/jlpt-notify/issues/new).
- Want to add a new feature or improve an existing one? Fork the repository and create a [Pull Request](https://github.com/bjvanbemmel/jlpt-notify/compare).

## References and resources
- [JLPT Leiden website](https://www.jlpt-leiden.nl/)
- [JLPT international website](https://www.jlpt.jp/e/)
- [Twilio](https://www.twilio.com)
