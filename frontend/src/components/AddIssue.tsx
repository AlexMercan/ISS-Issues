import { Autocomplete, Box, Button, Chip, Container, CssBaseline, TextField, ThemeProvider, createTheme, TextareaAutosize } from "@mui/material"
import { useEffect, useState } from "react";
import { useNavigate } from "react-router";
import { Issue, IssueStatus, IssueTag } from "../model/Issue";
import { SaveIssue } from "../services/IssueService";

const theme = createTheme();

export const AddIssue = () => {
    const navigate = useNavigate()
    const [issueTagList, setIssueTagList] = useState<IssueTag[]>([])
    const [newIssueTags, setNewIssueTags] = useState<IssueTag[]>([])
    const [issueName, setIssueName] = useState("")
    const [issueDescription, setIssueDescription] = useState("")

    const handleSelectionChange = (e: React.SyntheticEvent, values: string[]) => {
        setNewIssueTags(
            values.map(val => issueTagList.filter(t => t.name === val)[0])
        )
    }

    const onSubmitClick = async () => {
        const issue: Issue = {
            name: issueName,
            description: issueDescription,
            edges: {
                assignedTags: newIssueTags
            },
            id: -1,
            owner_id: -1,
            status: IssueStatus.Open
        }
        await SaveIssue(issue)
        navigate("/issues")
    }
    useEffect(() => {
        const fetchData = async () => {
            const request = await fetch("http://localhost/api/issuetags");
            if (request.status === 200) {
                setIssueTagList(await request.json())
            } else {
                navigate("/login")
            }
        }
        fetchData()
    }, [])
    return (
        <ThemeProvider theme={theme}>
            <Container component="main" sx={{ maxWidth: 1300 }}>
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
                        <h1>
                            Open new issue
                        </h1>
                    </Box>
                    <TextField variant="filled" label="Issue name" placeholder="New issue"
                        onChange={(e) => setIssueName(e.target.value)}
                    ></TextField>
                    <TextareaAutosize
                        aria-label="description"
                        placeholder="Description"
                        style={{
                            width: "1146px",
                            height: "288px",
                            background: "cyan"
                        }}
                        onChange={(e) => setIssueDescription(e.target.value)}
                    />


                    <Autocomplete
                        multiple
                        id="tags-filled"
                        options={issueTagList.map((tag) => tag.name)}
                        renderTags={(value: readonly string[], getTagProps) =>
                            value.map((option: string, index: number) => (
                                <Chip variant="outlined" label={option} {...getTagProps({ index })} />
                            ))
                        }
                        renderInput={(params) => (
                            <TextField
                                {...params}
                                variant="filled"
                                label="Tags"
                                placeholder="Favorites"
                            />
                        )}
                        onChange={handleSelectionChange}
                    />
                    <Box sx={{
                        display: "flex",
                        justifyContent: "flex-start"
                    }}>
                        <Button
                            sx={{
                                bgcolor: "lime"
                            }}
                            onClick={onSubmitClick}
                        >
                            Submit
                        </Button>
                    </Box>
                </Box>
            </Container >
        </ThemeProvider >
    )
}
