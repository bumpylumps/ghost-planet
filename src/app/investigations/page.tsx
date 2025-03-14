import { handleSubmit } from "./actions"

export default function InvestigationPage() {
    return (
    <div className="max-w-md mx-auto p-6 bg-white rounded-lg shadow-md">
        <h1 className="text-2xl font-bold mb-6">New Investigation</h1>
        
        {/* eslint-disable-next-line @typescript-eslint/no-explicit-any */}
        <form action={handleSubmit} className="space-y-4">
          <div className="space-y-2">
            <label htmlFor="title" className="block text-sm font-medium">
              Title:
            </label>
            <input id="title" name="title" className="w-full p-2 border rounded-md" required />
          </div>
  
          <div className="space-y-2">
            <label htmlFor="location" className="block text-sm font-medium">
              Location:
            </label>
            <input id="location" name="location" className="w-full p-2 border rounded-md" required />
          </div>
  
          <div className="space-y-2">
            <label htmlFor="date" className="block text-sm font-medium">
              Date:
            </label>
            <input id="date" name="date" type="date" className="w-full p-2 border rounded-md" required />
          </div>
  
          <div className="space-y-2">
            <label htmlFor="crew" className="block text-sm font-medium">
              Crew:
            </label>
            <input id="crew" name="crew" className="w-full p-2 border rounded-md" required />
          </div>
  
          <button
            type="submit"
            className="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700 transition-colors"
          >
            Submit
          </button>
        </form>
      </div>
    )
}