package main

import (
	"errors"
	"fmt"
	"html"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync/atomic"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/caarlos0/env/v6"
	"github.com/fatih/color"
	_ "github.com/joho/godotenv/autoload" // load .env
	"github.com/kjk/notionapi"
	"github.com/kjk/notionapi/tomarkdown"
	"golang.org/x/sync/errgroup"
)

type Config struct {
	Token          string `env:"NOTION_TOKEN,required"`
	BlogColID      string `env:"BLOG_COLLECTION_ID,required"`
	BlogColViewID  string `env:"BLOG_COLLECTION_VIEW_ID,required"`
	OtherColID     string `env:"OTHER_COLLECTION_ID,required"`
	OtherColViewID string `env:"OTHER_COLLECTION_VIEW_ID,required"`
}

func init() {
	log.SetHandler(cli.Default)
	color.NoColor = false
}

func main() {
	var config Config
	if err := env.Parse(&config); err != nil {
		log.WithError(err).Fatal("invalid config")
	}

	client := &notionapi.Client{}
	client.AuthToken = config.Token
	// client.Logger = os.Stderr

	index, err := queryCollection(client, config.BlogColID, config.BlogColViewID)
	if err != nil {
		log.WithError(err).Fatal("failed to query blog index")
	}

	g := New(10)
	total := len(index)
	var done int64

	var progressFn = func() string {
		return fmt.Sprintf("%d/%d", atomic.AddInt64(&done, 1), total)
	}

	for _, k := range index {
		k := k
		g.Go(func() error {
			return renderPage(
				client,
				k,
				progressFn,
				func(page *notionapi.Page) string {
					return toString(page.Root().Prop("properties.S6_\""))
				},
				func(page *notionapi.Page) string {
					slug := toString(page.Root().Prop("properties.S6_\""))
					return fmt.Sprintf("content/posts/%s.md", strings.ReplaceAll(slug, "/", ""))
				},
				func(page *notionapi.Page) string {
					slug := toString(page.Root().Prop("properties.S6_\""))
					date := toDateString(page.Root().Prop("properties.a`af"))
					draft := !toBool(page.Root().Prop("properties.la`A"))
					city := toString(page.Root().Prop("properties.%]Hm"))
					tags := toList(page.Root().Prop("properties.h|dn"))
					title := page.Root().Title
					return blogHeader(title, date, draft, slug, city, tags)
				},
				func(page *notionapi.Page) bool {
					return !toBool(page.Root().Prop("properties.la`A"))
				},
				func(page *notionapi.Page) error {
					if toString(page.Root().Prop("properties.S6_\"")) == "" {
						return errors.New("missing slug")
					}
					if toDateString(page.Root().Prop("properties.a`af")) == "" {
						return errors.New("missing date")
					}
					if page.Root().Title == "" {
						return errors.New("title")
					}

					return nil
				},
			)
		})
	}

	if err := g.Wait(); err != nil {
		log.WithError(err).Fatal("failed to build blog files")
	}

	index, err = queryCollection(client, config.OtherColID, config.OtherColViewID)
	if err != nil {
		log.WithError(err).Fatal("failed to query other pages index")
	}

	total = len(index)
	done = 0
	for _, k := range index {
		k := k
		g.Go(func() error {
			return renderPage(
				client,
				k,
				progressFn,
				func(page *notionapi.Page) string {
					return toString(page.Root().Prop("properties.7F2|"))
				},
				func(page *notionapi.Page) string {
					slug := toString(page.Root().Prop("properties.7F2|"))
					return fmt.Sprintf("content/%s.md", strings.ReplaceAll(slug, "/", ""))
				},
				func(page *notionapi.Page) string {
					return pageHeader(page.Root().Title)
				},
				func(page *notionapi.Page) bool {
					return false
				},
				func(page *notionapi.Page) error {
					if toString(page.Root().Prop("properties.7F2|")) == "" {
						return errors.New("missing slug")
					}
					if page.Root().Title == "" {
						return errors.New("title")
					}

					return nil
				},
			)
		})
	}

	if err := g.Wait(); err != nil {
		log.WithError(err).Fatal("failed to build other files")
	}
}

func queryCollection(client *notionapi.Client, colID, colViewID string) ([]string, error) {
	log.WithField("collection", colID).Info("querying")
	req := notionapi.QueryCollectionRequest{}
	req.Collection.ID = colID
	req.CollectionView.ID = colViewID
	query := &notionapi.Query{
		Aggregate: []notionapi.QueryAggregate{
			{
				AggregationType: "count",
				ID:              "count",
				Type:            "title",
				Property:        "title",
				ViewType:        "table",
			},
		},
		Filter: map[string]interface{}{
			"operator": "and",
		},
		Sort: []notionapi.QuerySort{
			{
				Direction: "descending",
				Property:  "a`af",
			},
		},
	}
	loader := notionapi.MakeLoaderReducer(query)
	loader.Reducers[notionapi.ReducerCollectionGroupResultsName].(*notionapi.ReducerCollectionGroupResults).Limit = 5000
	req.Loader = loader
	resp, err := client.QueryCollection(req, query)
	if err != nil {
		if strings.Contains(err.Error(), "429") {
			log.Info("got a 429, sleeping 1m")
			time.Sleep(time.Minute)
			return queryCollection(client, colID, colViewID)
		}
		return []string{}, err
	}
	if len(resp.RecordMap.Blocks) == 0 {
		return []string{}, fmt.Errorf("no results querying collection")
	}

	var blocks []string
	for k, v := range resp.RecordMap.Blocks {
		if v == nil {
			continue
		}
		if k == colID {
			continue
		}
		if v.Block.ParentID != colID {
			continue
		}
		if v.Block.Type != "page" {
			log.WithField("id", k).WithField("type", v.Block.Type).Warn("not a page")
			continue
		}
		blocks = append(blocks, k)
	}

	return blocks, nil
}

var tweetExp = regexp.MustCompile(`^https://twitter.com/.*/status/(\d+).*$`)
var youtubeExp = regexp.MustCompile(`https://www.youtube.com/watch?v=(.+)&.*`)
var youtubeShortExp = regexp.MustCompile(`https://youtu.be/(.+)`)

func renderPage(
	client *notionapi.Client,
	k string,
	progressProvider func() string,
	slugProvider func(p *notionapi.Page) string,
	filenameProvider func(p *notionapi.Page) string,
	headerProvider func(p *notionapi.Page) string,
	pageSkipper func(p *notionapi.Page) bool,
	pageValidator func(p *notionapi.Page) error,
) error {
	var ctx = log.WithField("page", progressProvider())

	page, err := client.DownloadPage(k)
	if err != nil {
		if strings.Contains(err.Error(), "429") {
			log.Info("got a 429, sleeping 1m")
			time.Sleep(time.Minute)
			return renderPage(client, k, progressProvider, slugProvider, filenameProvider, headerProvider, pageSkipper, pageValidator)
		}
		return fmt.Errorf("failed to download page %s: %w", k, err)
	}

	ctx = ctx.WithField("title", page.Root().Title)

	if pageSkipper(page) {
		ctx.Warn("skipping")
		return nil
	}

	if err := pageValidator(page); err != nil {
		return fmt.Errorf("invalid page %s ('%s'): %w", k, page.Root().Title, err)
	}

	slug := slugProvider(page)

	ctx.Info("rendering")

	converter := tomarkdown.NewConverter(page)
	converter.RenderBlockOverride = func(block *notionapi.Block) bool {
		switch block.Type {
		case notionapi.BlockHeader:
			converter.Newline()
			converter.RenderHeaderLevel(block, 2)
			return true
		case notionapi.BlockSubHeader:
			converter.Newline()
			converter.RenderHeaderLevel(block, 3)
			return true
		case notionapi.BlockSubSubHeader:
			converter.Newline()
			converter.RenderHeaderLevel(block, 4)
			return true
		case notionapi.BlockCode:
			// hack: create an html block that starts with !!!EMBED!!! and it gets actually really embedded for realz in real life
			if strings.HasPrefix(block.Code, "!!!EMBED!!!") {
				converter.Printf("{{< rawhtml >}}\n<p>\n\t")
				converter.Printf(strings.Replace(block.Code, "!!!EMBED!!!", "", 1) + "\n")
				converter.Printf("</p>\n{{< /rawhtml >}}\n")
				return true
			}
			converter.Printf("```" + toLang(block.CodeLanguage) + "\n")
			converter.Printf(block.Code + "\n")
			converter.Printf("```\n")
			return true
		case notionapi.BlockEmbed:
			if strings.HasPrefix(block.Source, "https://speakerdeck.com/") || strings.HasPrefix(block.Source, "https://slides.com") {
				converter.Newline()
				converter.Printf("[See slides](%s).", block.Source)
				converter.Newline()
				return true
			}
			ctx.WithField("src", block.Source).Warn("unhandled embed")
		case notionapi.BlockTweet:
			converter.Newline()
			converter.Printf("{{< tweet %s >}}", tweetExp.FindStringSubmatch(block.Source)[1])
			converter.Newline()
			ctx.Warn("Tweets might be deleted anytime, consider using something else instead")
			return true
		case notionapi.BlockVideo:
			if strings.HasPrefix(block.Source, "https://youtube.com") {
				converter.Newline()
				converter.Printf("{{< youtube %s >}}", youtubeExp.FindStringSubmatch(block.Source)[1])
				converter.Newline()
				return true
			} else if strings.HasPrefix(block.Source, "https://youtu.be") {
				converter.Newline()
				converter.Printf("{{< youtube %s >}}", youtubeShortExp.FindStringSubmatch(block.Source)[1])
				converter.Newline()
				return true
			}
			ctx.WithField("src", block.Source).Warn("unhandled video")
		case notionapi.BlockImage:
			file, err := client.DownloadFile(block.Source, block)
			if err != nil {
				if strings.Contains(err.Error(), "429") {
					log.Info("got a 429, sleeping 1m")
					time.Sleep(time.Minute)
					return converter.RenderBlockOverride(block)
				}
				ctx.WithError(err).WithField("src", block.Source).Fatal("couldn't download file")
			}
			imgPath := fmt.Sprintf("static/public/images/%s/%s%s", slug, block.ID, path.Ext(block.Source))
			imgCtx := ctx.WithField("path", imgPath).WithField("src", block.Source)
			imgCtx.Debug("downloading image")
			if err := os.MkdirAll(filepath.Dir(imgPath), 0750); err != nil {
				imgCtx.WithError(err).Fatal("couldn't create dirs for file")
			}
			if err := ioutil.WriteFile(imgPath, file.Data, 0644); err != nil {
				imgCtx.WithError(err).Fatal("couldn't write file")
			}
			converter.Printf(
				`{{< figure caption="%s" src="%s" >}}`,
				html.EscapeString(toCaption(block)),
				strings.Replace(imgPath, "static/", "/", 1),
			)
			return true
		}
		return false
	}

	return ioutil.WriteFile(
		filenameProvider(page),
		buildMarkdown(headerProvider(page), converter.ToMarkdown()),
		0644,
	)
}

func toCaption(block *notionapi.Block) string {
	if block.GetCaption() == nil {
		return ""
	}

	var caption = ""
	for _, t := range block.GetCaption() {
		caption += t.Text
	}
	return caption
}

func toLang(s string) string {
	if s == "Plain Text" {
		return ""
	}
	return strings.NewReplacer(
		"shell", "sh", // less diffs
		"docker", "dockerfile", // less diffs
	).Replace(strings.ToLower(s))
}

var postURLRegex = regexp.MustCompile(`\(https://carlosbecker.com/posts/(.+?)/\)`)

func buildMarkdown(header string, content []byte) []byte {
	var ss = strings.NewReplacer(
		"“", "\"",
		"”", "\"",
		"’", "'",
		"‘", "'",
		"…", "...",
	).Replace(string(content))

	ss = postURLRegex.ReplaceAllString(ss, `({{< ref "$1.md" >}})`)

	return []byte(strings.Join(append([]string{header}, strings.Split(ss, "\n")[1:]...), "\n") + "\n")
}

func blogHeader(title, date string, draft bool, slug, city string, tags []string) string {
	return fmt.Sprintf(`---
title: "%s"
date: %s
draft: %v
slug: %s
city: %s
toc: true
tags: [%s]
---`, title, date, draft, slug, city, strings.Join(tags, ", "))
}

func pageHeader(title string) string {
	return fmt.Sprintf(`---
title: "%s"
type: page
hideMetadata: true
hideAuthorBio: true
hideSuggestions: true
---`, title)
}

func toBool(v interface{}, ok bool) bool {
	return toString(v, ok) == "Yes"
}

func toString(v interface{}, ok bool) string {
	if !ok {
		return ""
	}

	return v.([]interface{})[0].([]interface{})[0].(string)
}

func toList(v interface{}, ok bool) []string {
	if !ok {
		return []string{}
	}
	return strings.Split(toString(v, ok), ",")
}

func toDateString(v interface{}, ok bool) string {
	if !ok {
		return ""
	}

	// may god have mercy on my soul
	return v.([]interface{})[0].([]interface{})[1].([]interface{})[0].([]interface{})[1].(map[string]interface{})["start_date"].(string)
}

//
// copied from goreleaser codebase
//

// Group is the Semphore ErrorGroup itself.
type Group interface {
	Go(func() error)
	Wait() error
}

// New returns a new Group of a given size.
func New(size int) Group {
	return &parallelGroup{
		ch: make(chan bool, size),
		g:  errgroup.Group{},
	}
}

var _ Group = &parallelGroup{}

type parallelGroup struct {
	ch chan bool
	g  errgroup.Group
}

// Go execs one function respecting the group and semaphore.
func (s *parallelGroup) Go(fn func() error) {
	s.g.Go(func() error {
		s.ch <- true
		defer func() {
			<-s.ch
		}()
		return fn()
	})
}

// Wait waits for the group to complete and return an error if any.
func (s *parallelGroup) Wait() error {
	return s.g.Wait()
}
