import React from "react";

import './Dashboard.scss';
import Button from "../../components/Button/Button";

const Dashboard = () => {
    return (

        <div className={'dashboard'}>
            <p className={'dashboardDesc'}>
                در حال حاضر هیچ صرافی فعالی در حال استفاده نیست. برای شروع استفاده از سیستم یک صرافی از لیست زیر انتخاب
                کنید یا یک صرافی جدید اضافه کنید.
            </p>

            <form>
                <div>
                    <input type="radio" value="Male" name="gender" /> Male
                    <input type="radio" value="Female" name="gender" /> Female
                    <input type="radio" value="Other" name="gender" /> Other
                </div>

                {/*<table>*/}
                {/*    <ul>*/}
                {/*        <li>*/}
                {/*            <div>صرافی ۱</div>*/}
                {/*        </li>*/}
                {/*        <li>*/}
                {/*            <div>صرافی ۲</div>*/}
                {/*        </li>*/}
                {/*        <li>*/}
                {/*            <div>صرافی ۳</div>*/}
                {/*        </li>*/}
                {/*        <li>*/}
                {/*            <div>صرافی ۴</div>*/}
                {/*        </li>*/}
                {/*    </ul>*/}
                {/*</table>*/}

                <div className={'brokerageSelectionButtons'}>
                    <Button Type="info">افزودن</Button>
                    <Button Type="success">انتخاب</Button>
                </div>
            </form>

        </div>
    )
}

export default Dashboard;