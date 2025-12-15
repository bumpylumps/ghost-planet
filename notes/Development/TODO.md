- finish build as per lets-go-further
- write package for audio uploads
- combine servemux + multipart form processing + image/ imaging core package
- Set up image processing on the server, which parses JSON sent by client, generates a thumbnail, and provides a source url

Data:
- Set up SQL migrations for investigations, locations, users, tokens
- Once Users and locations can be created and stored, finish createInvestigationsHandler func
- Figure out structure for Lore 
	- can share text notes structure
	
Doing this now - 12/15/25
- test all new validators in evidence structs

- For public locations, use Ghostbuddy user as owner

## Testing todos
- add stuff to mock functions for unit testing 
- add funcs for mock locations, investigations, etc. 