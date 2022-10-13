import './Card.scss'
import '../../styles/css/Columns.scss'
import React, {Component} from "react"

class Card extends Component {
    constructor(props) {
        super(props);
    }

    render() {
        console.log("st:", this.props.Styles)
        return (
            <div className={`card ${this.props.Styles}`}>
                {this.props.children}
            </div>
        )
    }
}

Card.defaultProps = {
    Styles: null
}

export default Card
