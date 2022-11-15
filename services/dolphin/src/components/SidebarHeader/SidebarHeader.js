import React from "react";
import avatar from '../../assets/images/avatar.png';

import './SidebarHeader.scss';

const SidebarHeader = (props) => {
    return (
        <div className={'user'}>
            <img className={'avatar'} src={avatar} alt="Avatar"/>
            <div>{props.username}</div>
        </div>
    )
}

export default SidebarHeader;