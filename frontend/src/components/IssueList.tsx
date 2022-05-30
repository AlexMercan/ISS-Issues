import React from 'react'
import { createTheme, CssBaseline, ThemeProvider } from "@mui/material";
import { IssueItem } from "./IssueItem";
import { Issue } from '../model/Issue';
import { useUser } from '../UserContext';
import { UserRole } from '../model/User';

const theme = createTheme();

type Props = {
    issueNameFilter: string
    Issues: Issue[]
    onItemClick?: (id: number) => void
}

export const IssueList: React.FC<Props> = (props: Props) => {
    const user = useUser().currentUser
    let onClick = (_: number) => { }
    if (props.onItemClick) {
        onClick = props.onItemClick
    }
    const issueItems = props.Issues
        .filter(item => item.name.match(props.issueNameFilter))
        .map((item) => <IssueItem onItemClick={onClick} key={item.id} Issue={item} />)
    return (
        <ThemeProvider theme={theme}>
            <CssBaseline />
            {issueItems}
        </ThemeProvider>
    )

}
