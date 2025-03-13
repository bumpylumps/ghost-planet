'use client'

import * as React from 'react'





export default function Page() {
    const [title, setTitle] = React.useState('')
    const [location, setLocation] = React.useState('')
    const [date, setDate] = React.useState('')
    const [crew, setCrew] = React.useState('')



    async function handleSubmit(e: React.FormEvent){
        e.preventDefault()

        let investigation = {
            'title': title,
            'location': location,
            'date': date, 
            'crew': crew
        }
        console.log(investigation)
    }
    
    
    return (
        // this could be a component instead 
        <form onSubmit={handleSubmit} >
            <label htmlFor="title">Title:</label>
            <input 
                onChange={(e) => setTitle(e.target.value)}
                id="title"
                name="title"
                value={title} 
            />
            <label htmlFor="location">Location:</label>
            <input 
               onChange={(e) => setLocation(e.target.value)}
               id="location"
               name="location"
               value={location} 
            />
            <label htmlFor="date">Date:</label>
            <input 
               onChange={(e) => setDate(e.target.value)}
               id="date"
               name="date"
               value={date} 
            />
            <label htmlFor="crew">Crew:</label>
            <input 
               onChange={(e) => setCrew(e.target.value)}
               id="crew"
               name="crew"
               value={crew} 
            />
            <button>Submit</button>
        </form>
    )
}