import React, {Component} from "react"
import logo from '../../assets/images/logo.jpeg';
import Sidebar from '../../components/Sidebar/Sidebar'
import Dashboard from '../../pages/Dashboard/Dashboard'
import Platform from '../../pages/Platform/Platform'

import './MainLayout.scss'

class MainLayout extends Component {
    render() {
        return (
            <div className={'layoutContainer'}>

                <Sidebar
                    username={"hossein varmazyar"}
                />

                <div className={'body'}>
                    <header className={'header'}>
                        <img alt='header Gate logo' src={logo}/>
                    </header>
                    <div className={'content'}>
                        {
                            this.props.pageTitle ?
                                <h2 className='contentHeader'>
                                    {this.props.pageTitle}
                                </h2> : null
                        }
                        <div className='contentWrapper'>
                            {this.props.children}
                            {/*<Dashboard/>*/}
                            <Platform/>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default MainLayout;

