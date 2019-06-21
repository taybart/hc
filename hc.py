import urllib.request
import sys

sys.tracebacklimit = 0
r = urllib.request.urlopen("http://example.com/")
sys.exit(r.getcode() != 200)
