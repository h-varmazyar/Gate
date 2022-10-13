import './TextInput.scss'
import '../../styles/css/Columns.scss'
import React, {Component} from "react"


export const InputDisplayStyles = {
    Inline: 'Inline',
    Outline: 'Outline',
    Basic: 'Basic',
}

export class TextInput extends Component {
    constructor(props) {
        super(props);
    }

    render() {
        let placeHolder;
        let containerClass;
        let inputClass;
        let labelClass;
        switch (this.props.displayStyle) {
            case InputDisplayStyles.Basic:
                placeHolder = this.props.placeholder;
                containerClass = 'basic-container';
                inputClass = 'basic-input';
                labelClass = 'basic-label';
                break
            case InputDisplayStyles.Inline:
                placeHolder = this.props.label;
                containerClass = 'inline-container';
                inputClass = 'inline-input';
                labelClass = 'inline-label';
                break
            case InputDisplayStyles.Outline:
                placeHolder = " ";
                containerClass = 'outline-container';
                inputClass = 'outline-input';
                labelClass = 'outline-label';
                break
        }
        return (
            <div
                className={`${containerClass} ${this.props.Styles}`}>
                {
                    this.props.displayStyle === InputDisplayStyles.Basic ?
                        <label className={labelClass}>
                            {this.props.label}
                        </label> : null
                }
                <input className={inputClass}
                       type={this.props.type}
                       id={this.props.id}
                       name={this.props.name}
                       disabled={this.disabled}
                       value={this.props.value}
                       {...this.props}
                       placeholder={placeHolder}
                />
                {
                    this.props.displayStyle === InputDisplayStyles.Outline ?
                        <label className={labelClass}>
                            {this.props.label}
                        </label> : null
                }

            </div>
        )
    }
}

TextInput.defaultProps = {
    Styles: null,
    isOutlined: false,
    disabled: false,
    type: 'text',
    displayStyle: InputDisplayStyles.Outline,
}