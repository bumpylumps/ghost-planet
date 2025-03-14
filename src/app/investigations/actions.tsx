"use server"

import { createServerSupabaseClient } from "@/utils/supabase/server"
import { revalidatePath } from "next/cache"
import { redirect } from "next/navigation"



export async function handleSubmit(formData: FormData) {
    const supabase = createServerSupabaseClient()

    const investigation = {
        title: formData.get("title"),
        location: formData.get("location"),
        date: formData.get("date"),
        crew: formData.get("crew"),
    }

    const { data, error } = await supabase.from("investigations").insert(investigation)

    if(error) {
        console.error("Error inserting data: ", error.message)
        return { success:false, error: error.message }
    } else {
        console.log("Form Submission Successful: ", investigation)
        return { success:true }
    }

    revalidatePath("/investigations")

    redirect("/investigations")
}