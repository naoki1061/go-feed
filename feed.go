package go-feed

import (
)

type Item struct {
    Title         string
    Description   string
    Content       string
    Image         string
    Link          string
    Category      string
    Published     string
    Updated       string
    Subject       string
}

func Parse(feedXML string) ([]*Item, error) {

}

func ParseURL(feedURL string) ([]*Item, error) {

}
