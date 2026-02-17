# You are here: 
- write insert func for:
	- audioNotes
	- photos 
- finish Evidence.FullSync()
- write createEvidenceHandler

## General Todos
- finish build as per lets-go-further
- finish flows for evidence inserts
- write package for audio uploads - audiocupp (backend audio processor)
- write package for photo uploads - photopup
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


 