# Tracking Tracker

Program to track package status on:

- Fedex
- USPS
- Cainiao (Alibaba)

## Features
- Not much
- Gotify notifications

## Installation 
Compile with:
```shell
go build .
```

Create a config file:
```shell
cp config.json.sample config.json
```

## Usage
Usage help:
```shell
./tracking-tracker
```

Register a new tracking code: 
```shell
./tracking-tracker register [fedex|cainiao|usps] [tracking code]
```

Remove a tracking code: 
```shell
./tracking-tracker remove [tracking code]
```

Crawl all registered tracking codes: 
```shell
./tracking-tracker crawl
```
 
## Notes
USPS requires headless chrome|chromium:
```
google-chrome --headless --remote-debugging-port=9222
```

```
chromium-browser-chrome --headless --remote-debugging-port=9222
```






