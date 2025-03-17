
## TODO: 

- [ ] Understand why result in /investigations/page.tsx is throwing an error
- [ ] implement server side form validation
- [ ] create schema for: 
	- [ ] evidence
- [ ] figure out collecting images and audio for evidence storage
- [ ] figure out how to handle clerk user info and ghostplanet user info
	- [ ] start with clerk users
		- [ ] test: add user_id field to investigation schema, populate w/ clerk userid to keep track of user on form submissions
- [ ] add typescript types for form validation:
	- [ ] investigations
	- [ ] evidence
	- [ ] user
	- [ ] crews
- [x] Throw Errors for form validation as well
- [x] switching to onSubmit instead of action needed to happen for client side form validation anyway,
	but type error was being thrown for action on form
- [ ] design UI
- [ ] build UI
- [ ] integrate UI
- [ ] create user dashboard
- [ ] configure app to protect routes with clerk, handle sign ins w/ clerk (custom flow)
	- [ ] set up middleware properly for clerk auth
- [x] Set up form with 4 inputs
- [x] connect app to a db (supabase)
- [x] switch input types in investi-form to match the content they are collecting
- [x] insert investigation into db row using form submit


## GOALS down the road
- [ ] integrate an interactive map to capture location instead of string
- [ ] incorporate clerk orgs to keep track of crews
- [ ] mobile version
- [ ] get newkirk or Tenney to user test and offer feedback