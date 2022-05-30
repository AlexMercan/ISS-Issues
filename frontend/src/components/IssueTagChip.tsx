import { Chip } from "@mui/material"
import { IssueTag } from "../model/Issue"
import { toRGB } from "../util/Util"


type Props = {
    tag: IssueTag
    onDelete: (name: string) => void
}

export const IssueTagChip: React.FC<Props> = (props) => {
    return (<Chip
        sx={{
            my: "2px",
            mx: "2px",
            bgcolor: toRGB(props.tag.name),
        }}
        variant="outlined"
        label={props.tag.name}
        onDelete={() => props.onDelete(props.tag.name)}
    />)
}
