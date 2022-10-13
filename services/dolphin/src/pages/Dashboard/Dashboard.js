import React from "react";

import './Dashboard.scss';

const Dashboard = () => {
    return (

        <div className={'dashboard'}>
            <p className={'dashboardDesc'}>
                در حال حاضر هیچ صرافی فعالی در حال استفاده نیست. برای شروع استفاده از سیستم یک صرافی از لیست زیر انتخاب
                کنید یا یک صرافی جدید اضافه کنید.
            </p>
            <table>
                <ul>
                    <li>
                        <div>صرافی ۱</div>
                    </li>
                    <li>
                        <div>صرافی ۲</div>
                    </li>
                    <li>
                        <div>صرافی ۳</div>
                    </li>
                    <li>
                        <div>صرافی ۴</div>
                    </li>
                </ul>
            </table>

            <div className={'brokerageSelectionButtons'}>
                <button>افزودن</button>
                <button>انتخاب</button>
            </div>

        </div>
    )
}

export default Dashboard;