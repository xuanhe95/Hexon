import React from "react";
import "./Setting.scss";

const Setting = () => (
    <div className="setting">
        <form className="form" action="" >
            <label for="address">Local Hexo Adress: </label>
            <br></br>
            <input type="text" id="address" name="address" placeholder="请输入本地Hexo文件地址..." required>
            </input>
            <br></br>
            <input type="text" id="blog" name="blog" placeholder="请输入博客地址..." required>
            </input>
            <br></br>
            <input type="submit" value="提交"></input>

        </form>
    </div >
);

export default Setting;