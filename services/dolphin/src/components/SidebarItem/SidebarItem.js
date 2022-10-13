import React from "react";
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'

import "./SidebarItem.scss"

const SidebarItem = (props) => {
    return (
        <a className={'item'} href={props.url}>
            <FontAwesomeIcon className={'itemIcon'} icon={props.icon}/>
            <span className={'itemText'}>{props.text}</span>
        </a>
    );
}

export default SidebarItem;