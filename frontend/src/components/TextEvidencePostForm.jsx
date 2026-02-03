import React from 'react'

const TextEvidencePostForm = () => {
  return (
    <div>
      <form>
        <h1>Text Evidence</h1>
        <input type="text" id="subject"  placeholder="subject" />
        <textarea id="body" name="body" rows="5" cols="33"></textarea> 
        <label for="location-select">Choose a location: </label>
        <select name="locations" id="location-select">
          <option value="">Please Shoose an option</option>
          <option value="123">Shorehand</option>
          <option value="124">Thorehand</option>
        </select>
        <input type="button" value="Submit"/>
      </form>
    </div>
  )
}

export default TextEvidencePostForm
