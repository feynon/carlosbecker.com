module github.com/caarlos0/notion-2-md

go 1.15

require (
	github.com/caarlos0/env/v6 v6.3.0
	github.com/joho/godotenv v1.3.0
	github.com/kjk/notionapi v0.0.0-20200903081654-eafa3ed70a1b
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
)

replace github.com/kjk/notionapi => github.com/caarlos0/notionapi v0.0.0-20200916220720-d56e013818ad
