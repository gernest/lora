# Lora [![Build Status](https://travis-ci.org/gernest/lora.svg?branch=master)](https://travis-ci.org/gernest/lora)


# What is it

Generates, manages, builds and deploy static sites for multiple users in a single application.
 Currently we support hugo based websites.

Or cut it short to a web hosting platform for static sites 

### Really?

Oh yeah, you can run this and provide a hosting platform for your clients

# Features

* Identicons for profiles
* Multi User Architecture
* In house resouce access clearance system. This allows giving access to certain resouces for only
users who have that priviledge.   For instance you want only users with level 6 clearance to access the
route `/admin` , this will make sure everyone who is below the level gets `permission denied` when he/she tries to access
the page.
* Templates Support
* Themes support
* Versioned Builds
* Download your site as zip
* deploy to the cloud with dokku

# Requirements

* A Postgresql database
* [hugo](http://gohugo.io/) static site generator (Will be removed later )
* git 

# How to hack it?

Lora uses postgresql database as its backend, and hugo for builing. It is based on  [beego](http://beego.me/) framework
and is written in go programming language a.k.a golang

I assume you have a working golang environment you can skip this part if its true, if not I have included a script `setup.sh` which helps
install golang and sets up the environment for you. To setup do the following.

    $ git clone https://github.com/gernest/lora.git
    $ cd lora
    $ chmod +x setup.sh
    $ ,/setup.sh
    
Then relax the heavy golang setup will  be done for you.

Now you can hack on lora whatever you want and if you are done build your work with the `build.sh` script

    $ cd /path/to/lora
    $ chmod +x build.sh
    $ ./build.sh

Which will create a binary `lora` at the current working directory
you can then execute by command `./lora`. Visit `http://localhost:4567` to view your site.

Remember to read `conf/app.config` and change it to suit your needs

Lora usess compass and sass for generating `css` If you want to hack on the styles, make sure sass and compass
is installed and you can run the command `compass watch` and enjoy hacking.

# Releases

Not yet


# Contributing

Contributions are welcome

# Author

Geofrey Ernest