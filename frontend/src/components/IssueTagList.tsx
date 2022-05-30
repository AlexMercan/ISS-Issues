import { Box } from "@mui/material";
import React from "react";
import { IssueTag } from "../model/Issue";
import { IssueTagChip } from "./IssueTagChip";

type Props = {
    tags: IssueTag[]
    onIssueTagDelete: (name: string) => void
}

export const IssueTagList: React.FC<Props> = (props) => {
    const tagList = props.tags.map(tag => <IssueTagChip key={tag.id} tag={tag} onDelete={props.onIssueTagDelete} />)
    return (

        <Box sx={{
            display: "flex",
            flexDirection: "column",
            justifyContent: "flex-start",
            width: "170px"
        }}>
            {tagList}
        </Box>
    )
}
