import "./app.scss"
import React from "react"

const App = () => {
    return (
        <div className={'content'}>
            <div className={'header'}></div>
            <div className={'body'}>
                <div className={'page-content'}></div>
                <div className={'side-menu'}></div>
            </div>
            <div className={'footer'}></div>
        </div>
    );
};

export default App;