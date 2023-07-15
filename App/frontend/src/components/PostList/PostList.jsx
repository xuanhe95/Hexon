import React, { Component } from "react";
import "./PostList.scss";
import { sendMsg } from "../../api";

class PostList extends Component {
    constructor(props) {
        super(props);
        this.update = this.update.bind(this);
        this.curIndex = this.props.index;
    }

    update(index) {
        this.props.update(index);
    }


    render() {
        const messages = this.props.posts.map((msg, index) => (
            <li key={index}
                class={index === this.props.index ? "active" : ""}
                onClick={() => this.update(index)}>
                {this.props.posts[index]}</li >
        ));
        return (
            <div className="PostList">
                <ol >
                    {messages}
                </ol>
            </div >
        );
    }


    // () => { sendMsg("open_" + String(index))

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