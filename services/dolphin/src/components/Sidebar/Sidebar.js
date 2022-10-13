import React from "react";
import SidebarItem from "../SidebarItem/SidebarItem";
import {faDashboard, faExchange, faStore} from "@fortawesome/free-solid-svg-icons";

import './Sidebar.scss';
import SidebarHeader from "../SidebarHeader/SidebarHeader";

const Sidebar = (props) => {
    return (
        <nav className={'sidebar'}>
            <SidebarHeader
                username={props.username}/>
            <ul className={'list'}>
                <li>
                    <SidebarItem
                        text={'داشبورد'}
                        url={'./dashboard'}
                        icon={faDashboard}
                    ></SidebarItem>
                </li>
                <li>
                    <SidebarItem
                        text={'صرافی ها'}
                        url={'./brokerages'}
                        icon={faExchange}
                    ></SidebarItem>
                </li>
                <li>
                    <SidebarItem
                        text={'بازارها'}
                        url={'./markets'}
                        icon={faStore}
                    ></SidebarItem>
                </li>
            </ul>
        </nav>
    )
}

export default Sidebar;