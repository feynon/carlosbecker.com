---
title: "Archive"
date: 1990-04-16T00:00:00Z
permalink: /archive
menu: nav
disableComments: true
---


<ul>
  <!-- orders content according to the "date" field in front matter -->
  {{ range .Data.Pages.ByDate }}
    <li>
      <h1><a href="{{ .Permalink }}">{{ .Title }}</a></h1>
      <time>{{ .Date.Format "Mon, Jan 2, 2006" }}</time>
    </li>
  {{ end }}
</ul>
