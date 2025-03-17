"use client"

import type * as React from "react"
import { useState, type FormEvent } from "react"
import { createInvestigation } from "./actions"
import { useRouter } from "next/navigation"

export default function InvestigationPage() {
  const router = useRouter()
  //keep track of form state for error reporting
  const [title, setTitle] = useState("")
  const [location, setLocation] = useState("")
  const [date, setDate] = useState("")
  const [crew, setCrew] = useState("")

  //keep track of form status to get validation in the way of submitting
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [serverError, setServerError] = useState<string | null>(null)

  //state for Errors
  const [errors, setErrors] = useState(({
    title: "",
    location: "",
    date: "",
    crew: "",
  }))


  // touched state for errors after user input
  //false by default (assume user hasn't interacted with form until page is loaded)
  const [touched, setTouched] = useState({
    title: false,
    location: false,
    date: false, 
    crew: false
  })

  // logic for checking individual form fields
  const validateField = (name: string, value: string) => {
    let error = ''

    switch (name) {
      case "title":
        //if field is empty, or too short
        if(!value.trim()) {
          error = "Title is required"
        } else if(value.length < 3) {
          error = "Title must be at least a word, no less than 3 chars"
        }
        break
      case "crew":
        // if field is empty, or too short
      if(!value.trim()) {
        error = "Crew is required"
      } else if(value.length < 3) {
        error = "There must be at least one crew member for the investigation"
      }
      break
      case "date":
        if(!value.trim()) {
          error = "Date is required"
        }
        break
      case "location":
        if(!value.trim) {
          error = "Location is required"
        } else if(value.length < 3){
          error = "Location must at least be a word"
        }
        break
      
    }

    return error
    }
  

  //watch input changes
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target

    //keep form state updated
    switch(name) {
      case "title":
        setTitle(value)
        break
      case "location": 
        setLocation(value)
        break
      case "date":
        setDate(value)
        break
      case "crew":
        setCrew(value)
        break
    }
  

  //validate on change after user input
  if (touched[name as keyof typeof touched]) {
    const error = validateField(name, value)
    setErrors((prev) => ({ ...prev, [name]: error }))
  }
}

// present error after user input + blur
const handleBlur = (e: React.FocusEvent<HTMLInputElement>) => {
  const { name, value } = e.target

  //mark field as touched 
  setTouched((prev) => ({ ...prev, [name]: true }))

  //validateField
  const error= validateField(name, value)
  setErrors((prev) => ({ ...prev, [name]: error }))
}

//all field validation
const validateForm = () => {
  const newErrors = {
    title: validateField("title", title),
    location: validateField("location", location),
    date: validateField("date", date),
    crew: validateField("crew", crew),
  }

  //keep track of errors
  setErrors(newErrors)

  //return true if there aren't any errors
  return !Object.values(newErrors).some((error) => error)
}

//form submission
async function handleSubmit(e: FormEvent){
    e.preventDefault()

    // reset server errors
    setServerError(null)

    // Mark all fields as touched
    setTouched({
      title: true,
      location: true,
      date: true, 
      crew: true,
    })

    if (!validateForm()) {
      console.log("Errors with form input, see form")
      return 
    }

    try {
      setIsSubmitting(true)

      const result = await createInvestigation({
        title: title,
        location: location || null,
        date: date || null,
        crew: crew || null,
      })

      if (result.error) {
        setServerError(result.error)
        return 
      }

      console.log("Investigation created:", result.data)

      //redirect to users investigation list
      //TODO: investigation list page
      router.push("/investigations")
      router.refresh()// this will refresh page to show new investigation added
    } catch(error){
      console.error("Error submitting form: ", error)
      setServerError("An unexpected error occured. Please try again.")
    } finally {
      setIsSubmitting(false)
    }
  }
    return (
    <div className="max-w-md mx-auto p-6 bg-white rounded-lg shadow-md">
        <h1 className="text-2xl font-bold mb-6">New Investigation</h1>
  
        <form onSubmit={handleSubmit} className="space-y-4">
          {/* conditional render for server error */}
          {serverError && (
            <div className="bg-red-100 border border-red-400 text-red-700
            px-4 py-3 rounded">{serverError}</div>
          )}


          <div className="space-y-2">
            <label htmlFor="title" className="block text-sm font-medium">
              Title:
            </label>
            <input 
              onChange={handleChange}
              onBlur={handleBlur}
              id="title" 
              name="title" 
              value={title}
              className={`w-full p-2 border rounded-md ${errors.title && touched.title ? "border-red-500" : "border-gray-300"}`} 
              />
              {/* Present errors to user as needed */}
              {errors.title && touched.title && <p className="text-red-500 text-sm">{errors.title}</p>}
          </div>
  
          <div className="space-y-2">
            <label htmlFor="location" className="block text-sm font-medium">
              Location:
            </label>
            <input 
              onChange={handleChange}
              onBlur={handleBlur}
              id="location" 
              name="location" 
              value={location}
              className="w-full p-2 border rounded-md border-gray-300"
              />
              {errors.location && touched.location && <p className="text-red-500 text-sm">{errors.location}</p>}
          </div>
  
          <div className="space-y-2">
            <label htmlFor="date" className="block text-sm font-medium">
              Date:
            </label>
            <input 
              onChange={handleChange}
              onBlur={handleBlur}
              id="date"
              name="date" 
              type="date" 
              value={date}
              className={`w-full p-2 border rounded-md ${errors.date && 
                touched.date ? "border-red-500" : "border-gray-300"}`} 
              />
              {errors.date && touched.date && <p className="text-red-500 text-sm">{errors.date}</p>}
          </div>
  
          <div className="space-y-2">
            <label htmlFor="crew" className="block text-sm font-medium">
              Crew:
            </label>
            <input 
              onChange={handleChange}
              onBlur={handleBlur}
              id="crew" 
              name="crew" 
              value={crew}
              className="w-full p-2 border rounded-md border-gray-300"  
            />
            {errors.crew && touched.crew && <p className="text-red-500 text-sm">{errors.crew}</p>}
          </div>
  
          <button
            type="submit"
            disabled={isSubmitting}
            className="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 transition-colors"
          >
            {isSubmitting ? "Submitting..." : "Submit"}
          </button>
        </form>
      </div>
    )
}