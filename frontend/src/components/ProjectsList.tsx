import React from 'react'

type Props = {
    username: string
}

export const ProjectsList: React.FC<Props> = (props: Props) => {
    return (
        <div>
            {props.username}
        </div>
    )

}
