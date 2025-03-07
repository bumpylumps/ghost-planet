async function handleSubmit(){
    console.log('submitted');
}


export default function Page() {
    
    // this could be a component instead 
    return (
        <form onSubmit={handleSubmit} >
            <label htmlFor="title">Add your Investigation!</label>
            <input name="title" />
            <input name="location" />
            <input name="date" />
            <input name="crew" />
            <button>Submit</button>
        </form>
    )
}