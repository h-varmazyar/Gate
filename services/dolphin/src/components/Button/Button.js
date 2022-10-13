import './Button.scss'
import '../../styles/css/Columns.scss'
import React, {Component} from "react"
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'

class Button extends Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <button
                type="button"
                disabled={this.disabled}
                onClick={this.props.onClick}
                className={`button ${this.props.Styles}`}
                {...this.props}
            >
                <div className={'vertical-center'}>
                    {this.props.iconTop !== null &&
                        <FontAwesomeIcon className={'icon-top'} icon={this.props.iconTop}></FontAwesomeIcon>
                    }
                    <div className={'horizontal-center'}>
                        {this.props.iconLeft !== null &&
                            <FontAwesomeIcon className={'icon-left'} icon={this.props.iconLeft}></FontAwesomeIcon>
                        }
                        <span className={'label'}>{this.props.label}</span>
                        {this.props.iconRight !== null &&
                            <FontAwesomeIcon className={'icon-right'} icon={this.props.iconRight}></FontAwesomeIcon>
                        }
                    </div>
                    {this.props.iconBottom !== null &&
                        <FontAwesomeIcon className={'icon-bottom'} icon={this.props.iconBottom}></FontAwesomeIcon>
                    }
                </div>
            </button>
        )
    }
}

Button.defaultProps = {
    Styles: null,
    disabled: false,
    label: "",
    iconTop: null,
    iconLeft: null,
    iconRight: null,
    iconBottom: null,
}

export default Button
