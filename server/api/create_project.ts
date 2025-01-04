// import { useFetch } from "nuxt/app"

export default defineEventHandler(async (event) => {
    const { project_name, description } = await readBody(event)
    const { server_url } = useRuntimeConfig()

    if(!project_name) { 
        setResponseStatus(event, 400, "Project name is not defined") 
        return
    } 

    const response = await fetch(`${server_url}/create`, { 
        method: "POST",
        body: JSON.stringify({
            project_name,
            description
        })
    })



})