import React, { Component } from "react";
import "./PostList.scss";
import { sendMsg } from "../../api";

class PostList extends Component {
    render() {
        const messages = this.props.posts.map((msg, index) => (
            < li key={index} onClick={() => { sendMsg("open_" + String(index)) }} >
                {this.props.posts[index]}</li >
        ));
        return (

            // console.log(id),
            // console.log(id[id.length - 1]),
            <div className="PostList">
                <ol >
                    {messages}
                </ol>
            </div >
        );
    }

    // toList(messages) {
    //     var content = "";
    //     for (var i = 0; i < messages.length; i++) {
    //         content += "<li>"
    //         content += messages[i].data;
    //         content += "</li>"
    //     }
    //     return content;
    // }
}

export default PostList;