---
date: 2014-10-07T00:00:00Z
slug: elections
title: Elections, in Ruby
---

> Updated with second round script in Oct 26, 2014.

Well, last sunday (Oct 5) was the brazilian elections. I was doing nothing, so
I decided to write a simple ruby script to parse the results and show the
top 3 candidates.

Besides the reverse clean-code done by the TSE (Superior Electoral Court),
it was pretty easy:

<script src="https://gist.github.com/caarlos0/88ec017ad687c883581a.js"></script>

Maybe the only interesting thing here is the metaprogramming in the String
class.
