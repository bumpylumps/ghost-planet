## Custom JSON templates
- ``` investigations: {
		location: {
			*see below*
		}
		phenomena: string
		date/timestamp: date
		evidence: {
			*see below*
		}
		notes: []string
		buddy/emergency contact: {}string
	}
- ``` evidence: {
		text-notes: []string
		audio-notes: mp3/flac/etc. *TODO figure out best*
		photos: .img/.png/ etc. *TODO figure out best*
		evps: .mp3/flac/etc. *TODO figure out best*
		visibility: public/private bool
	}
-  ``` locations: {
		 story/lore: {}string *TODO figure out structure*
		 lat/long/address: Mappy.coordinates
		 past-investigations: [{},{}]string
		 popularity: int/float? *TODO figure out DB metric for this.. hidden field that updates on investigation form submit?*
	} 
	
-  ```forum-post: {
		author: string (username)
		timestamp: date/time
		body: string
		header/title: string/ OP title or reply title
		footer/signature box: string
	}
- ```user: {
		name: string(hidden)
		username: string
		profile-page: stringurl
		status: online/offline/idle
		forum-posts: []string/url
		evidence: [{public evidences}]
	}