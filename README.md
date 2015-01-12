# Lora [![Build Status](https://travis-ci.org/gernest/lora.svg)](https://travis-ci.org/gernest/lora)

A website as a service platform powered by Go docs : [here](http://gernest.github.io/lora/)

# Requirements

* A working postgres database
* A working Go Environment
* Hugo (This is a static site generator, lora themes and templates are compiled by this)
* A Dokku server(Optional)

# Installation

Clone this repo and go get dependencies

    $ git clone https://github.com/gernest/lora
    $ cd lora && go get -v ./..
 
# How to run

Edit `conf/app.conf` to reflect your database settings 
    
    $ go build && ./lora

or 

    $ go run lora.go

Open your browser at `http://localhost:4567` to view the running instance.

Contributing
------------

Contributions are welcome. 

Before writing code, send mail to geofreyernest@live.com to discuss what you
plan to do. This gives me a chance to validate the design, avoid duplication of
effort and ensure that the changes fit the goals of the project. Do not start
the discussion with a pull request.
