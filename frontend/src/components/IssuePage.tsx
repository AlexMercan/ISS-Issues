import React, { useEffect, useState } from 'react'
import { Autocomplete, Box, Button, Chip, Container, createTheme, CssBaseline, ThemeProvider } from "@mui/material";
import { useNavigate, useParams } from "react-router";
import TextField from "@mui/material/TextField";
import { Issue, IssueStatus, IssueTag } from '../model/Issue';
import { IssueTagList } from './IssueTagList';
import { toRGB } from '../util/Util';
import { Unique } from '../services/Service'
import { UpdateIssue } from '../services/IssueService';
import { useUser } from '../UserContext';
import { UserRole } from '../model/User';

const theme = createTheme();

export const IssuePage = () => {
    const navigate = useNavigate()
    const [issue, setIssue] = useState<Issue | null>(null)
    const [issueTagList, setIssueTagList] = useState<IssueTag[]>([])
    const [newIssueTags, setNewIssueTags] = useState<IssueTag[]>([])
    const [originalTags, setOriginalTags] = useState<IssueTag[]>([])
    const [issueStatus, setIssueStatus] = useState<IssueStatus>(IssueStatus.Open)
    const { issueId } = useParams()
    const user = useUser().currentUser
    let role = UserRole.Programmer
    if (user) {
        role = user.role
    }
    const [buttonStatus, _] = useState(role === UserRole.Tester)
    const onChipDelete = (name: string) => {
        if (!buttonStatus) {
            setOriginalTags(old => old.filter(t => t.name !== name))
        }
    }

    const handleSelectionChange = (_: React.SyntheticEvent, values: string[]) => {
        setNewIssueTags(
            values.map(val => issueTagList.filter(t => t.name === val)[0])
        )
    }

    const onChangeIssueStatus = () => {
        const currentStatus = issueStatus;
        if (currentStatus === IssueStatus.Open) {
            setIssueStatus(IssueStatus.Closed)
        } else {
            setIssueStatus(IssueStatus.Open)
        }
    }

    const onUpdateButtonClick = async () => {
        const uniqueTags = Unique(newIssueTags.concat(originalTags));
        let newIssue = issue;
        if (newIssue) {
            newIssue.edges.assignedTags = uniqueTags;
            newIssue.status = issueStatus;
            await UpdateIssue(newIssue)
            navigate("/issues")
        }
    }


    useEffect(() => {
        if (issue) {
            setOriginalTags(issue?.edges.assignedTags)
            setIssueStatus(issue?.status)
        }
    }, [issue])
    useEffect(() => {
        const fetchData = async () => {
            const request = await fetch("http://localhost/api/issues/" + issueId);
            if (request.status === 200) {
                setIssue(await request.json())
            } else {
                navigate("/login")
            }
        }
        fetchData()
    }, [])
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
                        <h1>{issue != null && issue.name}</h1>
                    </Box>
                    <Box component="div" sx={{
                        bgcolor: "blue",
                        display: 'flex',
                        justifyContent: 'flex-start',
                    }}>

                        <Box sx={{
                            borderRadius: "75px",
                            backgroundColor: "red",
                            padding: "10px",
                            width: "auto",
                        }}>
                            {issue != null && issue.name}
                        </Box>
                    </Box>
                    <Box sx={{ display: "flex" }}>
                        <Box sx={{
                            width: "1200px",
                            height: "288px",
                            backgroundColor: "#976161"
                        }}>
                            <h4>Description</h4>
                            <p>{issue != null && issue.description}</p>
                        </Box>
                        <IssueTagList onIssueTagDelete={onChipDelete} tags={originalTags} />
                    </Box>
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
                                bgcolor: toRGB(issueStatus),
                                color: "red",
                            }}
                            disabled={buttonStatus}
                            onClick={onChangeIssueStatus}>
                            {issueStatus} Issue
                        </Button>
                        <Button
                            sx={{
                                mx: "5px",
                                bgcolor: "cyan",
                            }}
                            disabled={buttonStatus}
                            onClick={onUpdateButtonClick}>
                            Update
                        </Button>
                    </Box>
                </Box>
            </Container >
        </ThemeProvider >
    )
}
