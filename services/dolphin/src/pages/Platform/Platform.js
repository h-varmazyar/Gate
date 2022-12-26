import React from "react";
import axios from "axios";

class Platform extends React.Component {
    componentDidMount() {
        const url = `${process.env.REACT_APP_BASE_URL}/app/platforms`
        axios.get(url).then((response) => {
            console.log(response.data)
        })
    }

    render() {
        return (
            <div>
                Hello
            </div>
        )
    }
}

export default Platform;