- finish build as per lets-go-further
- write package for audio uploads - audiocupp (backend audio processor)
- write package for photo uploads - photopup
- combine servemux + multipart form processing + image/ imaging core package
- Set up image processing on the server, which parses JSON sent by client, generates a thumbnail, and provides a source url
- set up GET routes for:
	- all evidence
	- photos
	- audionotes
	- textnotes
	- evps

Data:
- Set up seperate tables for:
	- text notes
	- audio notes/ evps
	- photos
- Set up SQL migrations for investigations, locations, users, tokens
- Once Users and locations can be created and stored, finish createInvestigationsHandler func
- Figure out structure for Lore 
	- can share text notes structure
	
Doing this now - 1/15/25
- finished evidence insert
	- make sure that the request body without every evidence type works 
- test all new validators in evidence structs

- For public locations, use Ghostbuddy user as owner

## Testing todos
- add stuff to mock functions for unit testing 
- add funcs for mock locations, investigations, etc. 


 