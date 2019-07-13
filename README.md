# godex
Experimental golang cli with DDD-ish design for pokedex item and pokemon information search


# godex

[![Go Report Card](https://goreportcard.com/badge/github.com/AdhityaRamadhanus/godex)](https://goreportcard.com/report/github.com/AdhityaRamadhanus/godex)  [![Build Status](https://travis-ci.org/AdhityaRamadhanus/godex.svg?branch=master)](https://travis-ci.org/AdhityaRamadhanus/godex)

Experimental golang cli with DDD-ish design for pokedex item and pokemon information search

Entities:
Item
Pokemon
Ability

<p>
  <a href="#installation">Installation |</a>
  <a href="#installation-with-docker">Installation (with docker) |</a>
  <a href="#Usage">Usage |</a>
  <a href="#licenses">License</a>
  <br><br>
  <blockquote>
	godex is pokedex information search using https://pokeapi.co/api/v2 as backend.
  </blockquote>
</p>

Installation
----------- 
* go get github.com/AdhityaRamadhanus/godex
* cd to project dir
* run build
```bash
make build
```
* run ./godex

Installation With Docker
----------- 
* go get github.com/AdhityaRamadhanus/godex
* cd to project dir
* run build image
```bash
docker build -t godex .
```
* run ./godex
```bash
docker run -it godex [pokemon/item name]
```

Usage
-----
```bash
./gode [pokemon/item name]
```

License
----

MIT © [Adhitya Ramadhanus]

