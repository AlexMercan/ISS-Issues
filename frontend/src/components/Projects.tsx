import React, { useEffect, useState } from 'react'
import { ProjectsList } from './ProjectsList'

const getUser = async () => {
    var request = await fetch("http://localhost/api/projects", {
        credentials: 'include',
        mode: "cors"
    })
    var text = await request.text()
    return text
}

export const Projects = () => {
    const [user, setUser] = useState("")
    useEffect(() => {
        const fetchData = async () => {
            const body = await getUser()
            setUser(body)
        }
        fetchData()
    }, [])
    return (
        <div>
            <ProjectsList username={user} />
        </div>
    )
}
