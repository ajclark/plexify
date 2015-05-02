# plexify
Simple Go/Gin REST API to update a Plex media library

### Running
```plexify -p "/Applications/Plex Media Server.app/Contents/MacOS/Plex Media Scanner"```

### Sending an update request (pass in plex section id of the TV library, in this case)
```curl -H 'Content-type: application/json' homes-macbook-air.local:8080/scan -d '{"path":"/mnt/media/TV/Top Gear/Season 22", "section_id":2 }' -v```
