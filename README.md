geocaching-saturation-checker
=============================

A way to check if your cache follows the 0.1mile saturation guidelines

WIP, but I hope to work on it more now. The idea is a map that places geocaches on it where each has an overlayed circle with a radius of 512' / 0.1miles. If the circles overlapped they would be red, green if not. This way a geocacher can easily test coordinates for a potential cache before placing it and having to move it when it doesn't follow guidelines. 

Currently, since the Geocaching.com api is not public, I'm going to try to populate the map with parse .loc files grabbed from the geocaching.com website and then overlay circle elements centered on geocaches for their "reserved area".

Any contributions or steps in the right direction would be greatly welcome. If you want to contact me you can check out my website, http://benreichert.com/contact.html for means to do so. 

I hope this works out and it is a usable project!


See it in progress
==================

In the directory where you cloned the repo, do:

``python -m SimpleHTTPServer``

and connect to localhost:8000/map.html
