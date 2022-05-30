import { Box } from '@mui/material';
import React from 'react'
import { useNavigate } from "react-router";
import { Issue, IssueTag } from '../model/Issue';
import { toRGB } from '../util/Util';

type Props = {
    Issue: Issue
    onItemClick: (id: number) => void
}

export const IssueItem: React.FC<Props> = (props: Props): JSX.Element => {
    let tags: JSX.Element[] = []
    if (props.Issue.edges?.assignedTags) {
        tags = props.Issue.edges.assignedTags.map((tag: IssueTag) => <div key={tag.id}
            style={{
                marginLeft: "10px",
                backgroundColor: toRGB(tag.name),
                borderRadius: "100"
            }}>{tag.name}</div>)
    }
    return (
        <Box component={"div"} sx={{
            bgcolor: "gray",
            display: 'flex',
            flexWrap: 'wrap',
            minHeight: '50px',
            justifyContent: 'flex-start',
            alignItems: "center",
            "&:hover": { bgcolor: "red", cursor: "pointer" }
        }} onClick={() => props.onItemClick(props.Issue.id)}>
            <span style={{ minHeight: "5" }}>{props.Issue.name}</span>
            {tags}
        </Box>
    )
}
