- finish build as per lets-go-further
- Set up SQL migrations for investigations, locations, users, tokens
- Once Users and locations can be created and stored, finish createInvestigationsHandler func
- Figure out better data structure for Evidence: 
	- don't like slices of whatever strings
	- would rather have a structure for:
		- Text notes
		- audio notes
		- EVPS - can share audio notes structure 
		- photos 
- Figure out structure for Lore 
	- can share text notes structure
	
- figure out proper validation for requests on Locations, Evidence, Investigations, Users, etc. 

- Change user/location/evidence refs in Investigation schema to Id's instead of full entities
	- create route to get full investigation with entitie refs if really want it

- Maybe break up Investigations.go into seperate files for each entity (location, user, etc)
- Add user id field to locations and evidence? 
- For public locations, use Ghostbuddy user as owner

## Testing todos
- add stuff to mock functions for unit testing 
- add funcs for mock locations, investigations, etc. 