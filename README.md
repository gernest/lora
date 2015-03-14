# Lora [![Build Status](https://travis-ci.org/gernest/lora.svg?branch=master)](https://travis-ci.org/gernest/lora)
Create, manage and deploy  static website witth clicks of a buutton

# Features
* multiple website
* subdomain
* multi user
* Themes
* Templates

Lora uses an excelent static generator hugo. So, templates for lora are basically hugo website, and theme support is based on hugo themes.

# Dependencies
* Postgresql database
* hugo static site generator
* go programming language
* godep depenency manager

# Warning

This is an experimental, proof of concept. It works like a charm but use it at your own risks. I created this to get hands on experience
on go idioms and understand the language better.

# Installation
clone the repo in your `GOPATH`

    git clone https://github.com/gernest/lora
    
Then cd to the path and use godep to install all dependencies

    cd lora
    godep restore

Then build and run
    
    go build 
    ./lora

Dont forget to chance content of the file `conf/app.conf` to suit your settings.
Instructions for installing godep can be found [here](https://github.com/tools/godep)


Fork it, play nice with it!


# Author
Geofrey Ernest geofreyernest@live.com