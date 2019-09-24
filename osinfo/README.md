[<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">](https://github.com/multiverse-os)

## Multiverse OS: `libosinfo` Provides Access to Linux OS Infomation
**URL** [multiverse-os.org](https://multiverse-os.org)


A port of the C library `libosinfo` to Go, providing access to a variety of
information relating to major distros. Importantly releases, their URLs, minimum
requirements to operate, and other similar details. Used in Gnome Boxes to
provide automated downloading of VM images to simplify VM usage. 

This library provides a command-line tool to update the databases easily from
the Git repository. In addition, this tool imports the XML and rewrites it to
JSON because I found that much easier to work with and parse. The OS folder of
the database is now complete. The remainder could easily be imlemented using the
OS folder implementation as a reference. Other folders are much smaller. 

Additional tools to embed the database information in your release binary will
be included in the near future. 


