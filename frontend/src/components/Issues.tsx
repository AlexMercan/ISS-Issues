import { Box, Container, createTheme, CssBaseline, ThemeProvider } from '@mui/material'
import React, { useEffect, useState } from 'react'
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import { IssueList } from "./IssueList";
import { useNavigate } from "react-router";
import { Issue } from '../model/Issue';
import { useUser } from '../UserContext';
import { UserRole } from '../model/User';

const theme = createTheme();

export const Issues = () => {
    const [issues, setIssues] = useState<Issue[]>([])
    const [searchText, setSearchText] = useState("")
    const currentUser = useUser().currentUser
    const navigate = useNavigate()
    useEffect(() => {
        const fetchData = async () => {
            const request = await fetch("http://localhost/api/issues");
            if (request.status === 200) {
                setIssues(await request.json())
            } else {
                navigate("/login")
            }
        }
        fetchData()
    }, [])
    let onIssueItemClick: ((id: number) => void) | undefined = (id: number) => {
        navigate("/issues/" + id)
    }
    if (currentUser && currentUser.role === UserRole.Tester) {
        onIssueItemClick = undefined
    }
    return (
        <ThemeProvider theme={theme}>
            <Container component="main" sx={{ maxWidth: 1300 }}>
                <CssBaseline />
                <Box component="div" sx={{
                    mt: 5,
                    bgcolor: "gray",
                    display: 'flex',
                    flexDirection: 'column',
                }}>
                    <Box component="div" sx={{
                        bgcolor: "cyan",
                        display: 'flex',
                        justifyContent: 'flex-start',
                    }}>
                        <h1>Issues</h1>
                    </Box>
                    <Box component="div" sx={{
                        bgcolor: "yellow",
                        display: 'flex',
                        justifyContent: 'flex-start',
                    }}>
                        <TextField id="outlined-basic" label="Search" variant="filled" sx={{ minWidth: "90%" }}
                            onChange={(e) => setSearchText(e.target.value)} />
                        <Button onClick={() => navigate("/issues/add")} sx={{ bgcolor: "#07B438", padding: 0 }} variant="contained">Open new
                            issue</Button>
                    </Box>
                    <IssueList onItemClick={onIssueItemClick} issueNameFilter={searchText} Issues={issues} />
                </Box>
            </Container>
        </ThemeProvider>
    )
}
