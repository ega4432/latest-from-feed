# rss-feed

[![ci](https://github.com/ega4432/rss-feed/actions/workflows/ci.yaml/badge.svg)](https://github.com/ega4432/rss-feed/actions/workflows/ci.yaml)

- [rss-feed](#rss-feed)
  - [Overview](#overview)
  - [Installation](#installation)
  - [Usage](#usage)
  - [License](#license)

## Overview

A command line tool that allows you to retrieve feeds in json format.

## Installation

It will be added to `$GOPATH/bin`.

```sh
go install github.com/ega4432/rss-feed@latest
```

## Usage

```sh
$ rss-feed
rss-feed
Error: Required flag `url` has not been set.
Usage:
  rss-feed [flags]
  rss-feed [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Print version

Flags:
  -h, --help         help for rss-feed
  -l, --latest       Only print latest item.
  -u, --url string   (required)RSS feed URL to fetch data.

Use "rss-feed [command] --help" for more information about a command.
```

Below is an example of fetching new information provided by the Japanese Ministry of Health, Labor and Welfare website. You should specify the url with the `-u` or `--url` flag.

```sh
$ rss-feed --url https://www.mhlw.go.jp/stf/news.rdf
```

<details><summary>show details</summary>

```sh
$ rss-feed --url https://www.mhlw.go.jp/stf/news.rdf | jq .
[
  {
    "title": "第85回新型コロナウイルス感染症対策アドバイザリーボード資料（令和4年5月25日）を掲載しました",
    "link": "https://www.mhlw.go.jp/stf/seisakunitsuite/bunya/0000121431_00348.html",
    "links": [
      "https://www.mhlw.go.jp/stf/seisakunitsuite/bunya/0000121431_00348.html"
    ],
    "published": "2022-05-30T21:30:00+09:00",
    "publishedParsed": "2022-05-30T21:30:00+09:00",
    "author": {
      "name": "厚生労働省"
    },
    "authors": [
      {
        "name": "厚生労働省"
      }
    ],
    "dcExt": {
      "creator": [
        "厚生労働省"
      ],
      "date": [
        "2022-05-30T21:30:00+09:00"
      ]
    },
    "extensions": {
      "dc": {
        "creator": [
          {
            "name": "creator",
            "value": "厚生労働省",
            "attrs": {},
            "children": {}
          }
        ],
        "date": [
          {
            "name": "date",
            "value": "2022-05-30T21:30:00+09:00",
            "attrs": {},
            "children": {}
          }
        ]
      }
    }
  },
  {
    "title": "後藤厚生労働大臣　閣議後記者会見のお知らせ",
    "link": "https://www.mhlw.go.jp/stf/newpage_25966.html",
    "links": [
      "https://www.mhlw.go.jp/stf/newpage_25966.html"
    ],
    "published": "2022-05-30T18:47:00+09:00",
    "publishedParsed": "2022-05-30T18:47:00+09:00",
    "author": {
      "name": "厚生労働省"
    },
    "authors": [
      {
        "name": "厚生労働省"
      }
    ],
    "dcExt": {
      "creator": [
        "厚生労働省"
      ],
      "date": [
        "2022-05-30T18:47:00+09:00"
      ]
    },
    "extensions": {
      "dc": {
        "creator": [
          {
            "name": "creator",
            "value": "厚生労働省",
            "attrs": {},
            "children": {}
          }
        ],
        "date": [
          {
            "name": "date",
            "value": "2022-05-30T18:47:00+09:00",
            "attrs": {},
            "children": {}
          }
        ]
      }
    }
  },
  {
    "title": "労働政策審議会（職業安定分科会）軽量版",
    "link": "https://www.mhlw.go.jp/stf/shingi/shingi-rousei_126979.html",
    "links": [
      "https://www.mhlw.go.jp/stf/shingi/shingi-rousei_126979.html"
    ],
    "published": "2022-05-30T18:00:00+09:00",
    "publishedParsed": "2022-05-30T18:00:00+09:00",
    "author": {
      "name": "厚生労働省"
    },
    "authors": [
      {
        "name": "厚生労働省"
      }
    ],
    "dcExt": {
      "creator": [
        "厚生労働省"
      ],
      "date": [
        "2022-05-30T18:00:00+09:00"
      ]
    },
    "extensions": {
      "dc": {
        "creator": [
          {
            "name": "creator",
            "value": "厚生労働省",
            "attrs": {},
            "children": {}
          }
        ],
        "date": [
          {
            "name": "date",
            "value": "2022-05-30T18:00:00+09:00",
            "attrs": {},
            "children": {}
          }
        ]
      }
    }
  },
  {
    "title": "旧優生保護法一時金認定審査会",
    "link": "https://www.mhlw.go.jp/stf/newpage_05687.html",
    "links": [
      "https://www.mhlw.go.jp/stf/newpage_05687.html"
    ],
    "published": "2022-05-30T18:00:00+09:00",
    "publishedParsed": "2022-05-30T18:00:00+09:00",
    "author": {
      "name": "厚生労働省"
    },
    "authors": [
      {
        "name": "厚生労働省"
      }
    ],
    "dcExt": {
      "creator": [
        "厚生労働省"
      ],
      "date": [
        "2022-05-30T18:00:00+09:00"
      ]
    },
    "extensions": {
      "dc": {
        "creator": [
          {
            "name": "creator",
            "value": "厚生労働省",
            "attrs": {},
            "children": {}
          }
        ],
        "date": [
          {
            "name": "date",
            "value": "2022-05-30T18:00:00+09:00",
            "attrs": {},
            "children": {}
          }
        ]
      }
    }
  },
  {
    "title": "第２回健康増進に係る科学的な知見を踏まえた技術的事項に関するWG　資料",
    "link": "https://www.mhlw.go.jp/stf/newpage_25962.html",
    "links": [
      "https://www.mhlw.go.jp/stf/newpage_25962.html"
    ],
    "published": "2022-05-30T17:35:00+09:00",
    "publishedParsed": "2022-05-30T17:35:00+09:00",
    "author": {
      "name": "厚生労働省"
    },
    "authors": [
      {
        "name": "厚生労働省"
      }
    ],
    "dcExt": {
      "creator": [
        "厚生労働省"
      ],
      "date": [
        "2022-05-30T17:35:00+09:00"
      ]
    },
    "extensions": {
      "dc": {
        "creator": [
          {
            "name": "creator",
            "value": "厚生労働省",
            "attrs": {},
            "children": {}
          }
        ],
        "date": [
          {
            "name": "date",
            "value": "2022-05-30T17:35:00+09:00",
            "attrs": {},
            "children": {}
          }
        ]
      }
    }
  }
]
```

</details>

You can use the `-l` or `--latest` flag to get only the most recent one.

```sh
$ rss-feed --url https://www.mhlw.go.jp/stf/news.rdf --latest
```

<details><summary>show details</summary>

```sh
$ rss-feed --url https://www.mhlw.go.jp/stf/news.rdf --latest | jq .
{
  "title": "第85回新型コロナウイルス感染症対策アドバイザリーボード資料（令和4年5月25日）を掲載しました",
  "link": "https://www.mhlw.go.jp/stf/seisakunitsuite/bunya/0000121431_00348.html",
  "links": [
    "https://www.mhlw.go.jp/stf/seisakunitsuite/bunya/0000121431_00348.html"
  ],
  "published": "2022-05-30T21:30:00+09:00",
  "publishedParsed": "2022-05-30T21:30:00+09:00",
  "author": {
    "name": "厚生労働省"
  },
  "authors": [
    {
      "name": "厚生労働省"
    }
  ],
  "dcExt": {
    "creator": [
      "厚生労働省"
    ],
    "date": [
      "2022-05-30T21:30:00+09:00"
    ]
  },
  "extensions": {
    "dc": {
      "creator": [
        {
          "name": "creator",
          "value": "厚生労働省",
          "attrs": {},
          "children": {}
        }
      ],
      "date": [
        {
          "name": "date",
          "value": "2022-05-30T21:30:00+09:00",
          "attrs": {},
          "children": {}
        }
      ]
    }
  }
}
```

</details>

## License

`rss-feed` command is released under the MIT license. See [LICENSE](https://github.com/ega4432/rss-feed/blob/main/LICENSE).
