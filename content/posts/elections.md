---
title: "Elections, in Ruby"
date: 2014-10-10
draft: false
slug: elections
city: Joinville
toc: true
tags: []
---

> Updated with second round script in Oct 26, 2014.

Well, last sunday (Oct 5) was the brazilian elections. I was doing nothing, so I decided to write a simple ruby script to parse the results and show the top 3 candidates.

Besides the reverse clean-code done by the TSE (Superior Electoral Court), it was pretty easy:

```ruby
require 'net/http'
require 'uri'
require 'json'
uri = URI('http://divulga.tse.jus.br/2014/divulgacao/oficial/143/dadosdivweb/br/br-0001-e001431-w.js')
class String; def percent_of(n) "#{(self.to_f / n.to_f * 100.0).round(2)}%"; end; end
begin
  data = JSON(Net::HTTP.get_response(uri).body)
  system('clear')
  puts "\n\n<!--more-->-\n#{data['ht']} - #{data['ea'].percent_of(data['e'])} dos votos apurados\n----"
  data['cand'].take(3).each do |candidate|
    puts "[#{candidate['n']}] #{candidate['nm']} - #{candidate['v'].percent_of(data['vv'])}"
  end
end while sleep(60)
```
```ruby
require 'net/http'
require 'uri'
require 'json'
uri = URI('http://divulga.tse.jus.br/2014/divulgacao/oficial/144/dadosdivweb/br/br-0001-e001441-w.js')
class String; def percent_of(n) "#{(self.to_f / n.to_f * 100.0).round(2)}%"; end; end
begin
  data = JSON(Net::HTTP.get_response(uri).body)
  system('clear')
  puts "\n\n----\n#{data['ht']} - #{data['ea'].percent_of(data['e'])} dos votos apurados\n----"
  data['cand'].each do |candidate|
    puts "[#{candidate['n']}] #{candidate['nm']} - #{candidate['v'].percent_of(data['vv'])}"
  end
end while sleep(60)
```

Maybe the only interesting thing here is the metaprogramming in the String class.
