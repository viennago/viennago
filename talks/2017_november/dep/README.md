<html>
	<head>
		<link rel="stylesheet" href="css/reveal.css">
		<link rel="stylesheet" href="css/theme/white.css">
	</head>
	<body>
		<div class="reveal">
			<div class="slides">
				<section>Slide 1</section>
				<section>Slide 2</section>
<section>

# Dep

##

## The Problem

* The Go standard installation has a single source tree (`~/go/src/...`)
* not all organisations are using mono-repos
* not all organisations are using build-systems that immediately detect dependency breakages


## What is Dep?

"Our goal is that phase 1 will be wrapped by the end of the 1.9 cycle, and dep
will be merged into the toolchain when 1.10 development begins."


## Dep Setup

## Why use Dep?

## Gopkg.toml / Gopkg.lock

* Documentation: https://github.com/golang/dep/blob/master/docs/Gopkg.toml.md
* Hands off the `Gopkg.lock` file

## How does dep decide what version of a dependency to use?

<small>

* All semver versions come first, and sort mostly according to the semver 2.0 spec, with one exception:
    * Semver versions with a prerelease are sorted after all non-prerelease semver. Within this subset they are sorted first by their numerical component, then lexicographically by their prerelease version.
* The default branch(es) are next; the semantics of what "default branch" means are specific to the underlying source type, but this is generally what you'd get from a go get.
* All other branches come next, sorted lexicographically.
* All non-semver versions (tags) are next, sorted lexicographically.
* Revisions, if any, are last, sorted lexicographically. Revisions do not typically appear in version lists, so the only invariant we maintain is determinism - deeper semantics, like chronology or topology, do not matter.

</small>

# commands

## dep commands

* tip: add the `-v` flag; `dep` is currently slow, it's good to see it's doing something.

* `dep init -v`
* `dep init -gopath -v`
* `dep ensure -v`
* `dep ensure -update github.com/`
* `dep version`


## dep init

## dep ensure

* dep ensure
* dep ensure -update
* dep ensure -add

## dep status

## dep prune

## commands: dep ensure

```sh
$ dep ensure -v
Gopkg.lock was already in sync with imports and Gopkg.toml
(1/63) Wrote github.com/go-logfmt/logfmt@v0.3.0
(2/63) Wrote golang.org/x/time@master
(3/63) Wrote github.com/spf13/cast@v1.1.0
(4/63) Wrote github.com/spf13/jwalterweatherman@master
(5/63) Wrote github.com/sony/gobreaker@0.3.0
(6/63) Wrote github.com/openzipkin/zipkin-go-opentracing@v0.3.2
(7/63) Wrote github.com/opentracing-contrib/go-observer@master
(8/63) Wrote github.com/spf13/afero@v1.0.0
(9/63) Wrote github.com/spf13/cobra@v0.0.1
(10/63) Wrote github.com/PuerkitoBio/purell@v1.1.0
(11/63) Wrote github.com/pelletier/go-toml@v1.0.1
(12/63) Wrote github.com/pierrec/lz4@v1.0.1
(13/63) Wrote github.com/satori/go.uuid@v1.1.0
(14/63) Wrote github.com/spf13/pflag@v1.0.0
(15/63) Wrote github.com/opentracing/opentracing-go@v1.0.2
(16/63) Wrote github.com/spf13/viper@v1.0.0
...
(62/63) Wrote github.com/prometheus/procfs@master
(63/63) Wrote github.com/rcrowley/go-metrics@master
```

## migrating to dep

It is possible to import the dependencies from other managers during the `dep init` phase.

* glide (https://glide.sh/)
* godep (https://godoc.org/github.com/tools/godep)
* vndr (https://github.com/LK4D4/vndr)
* govend (https://github.com/govend/govend)
* gb (https://getgb.io/)
* gvt (https://github.com/FiloSottile/gvt)

Otherwise, to skip this, use `dep init -skip-tools`

## caveats

* as always in go, read the docs, then you won't be surprised.
* version-tag over revision


## Closing notes

### Help Sources

* `#vendor` channel in [gophers-slack](https://invite.slack.golangbridge.org/)
* issues on github https://github.com/golang/dep/issues

### roadmap

https://github.com/golang/dep/wiki/Roadmap


</section>
			</div>
		</div>
        <script src="lib/js/head.min.js"></script>
		<script src="js/reveal.js"></script>
		<script>
			Reveal.initialize({
                dependencies: [
                    { src: 'plugin/markdown/marked.js', condition: function() { return !!document.querySelector( '[data-markdown]' ); } },
                    { src: 'plugin/markdown/markdown.js', condition: function() { return !!document.querySelector( '[data-markdown]' ); } },
                    { src: 'plugin/highlight/highlight.js', async: true, callback: function() { hljs.initHighlightingOnLoad(); } },
                    { src: 'plugin/zoom-js/zoom.js', async: true },
                ],
                markdown: {
                    smartypants: true
                }
            });
		</script>
	</body>
</html>

