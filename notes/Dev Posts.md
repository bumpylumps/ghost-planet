
9-6-25
I love it when the Saturday morning motivation hits! I'm currently back in the lab, transforming GhostBuds into GhostPlanet with a Go backend. 

The Express stack wasn't scaling well for Ghostbud's use case. The potential for overhead on concurrent API calls and the lack of type safety was threatening to get unwieldy as location, evidence, and user data was introduced. Go's concurrency model fits much better for GhostPlanet's workload patterns and will be much easier to maintain as GhostPlanet scales up. Not to mention its type system will provide some serious data integrity. Overall it will be a lot easier to squash bugs in the future (hopefully before they hit production!). 

I can't wait for this project to get unleashed on the world to help people expand their paranormal horizons and better connect with the weirder side of our shared reality.