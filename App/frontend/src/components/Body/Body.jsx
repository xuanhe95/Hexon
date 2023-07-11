import React, { Component } from "react";
import SubmitButton from "../SubmitButton/SubmitButton";
import "./Body.scss";
import { connect, sendMsg, isConnected } from "../../api";
import Editor from "../Editor/Editor";
import Setting from "../Setting/Setting";
import PostList from "../PostList/PostList";
import MDEditor from "../MDEditor/MDEditor";

class Body extends Component {

    constructor(props) {
        super(props);
        this.state = {
            posts: [],
            setting: false,
            content: "",
        };
        this.setting = this.setting.bind(this);

    }
    render() {
        var content = null;

        if (this.state.setting) {
            content = this.showSetting()
        }
        else {
            console.log("content")
            content = this.openEditor();

        }
        return (

            <div class="body">
                <div className="BodyContainer">
                    <PostList posts={this.state.posts}
                        func={(index) => { sendMsg("open" + index.toString()) }} />
                    {content}
                </div>
                <div class="settingButton">
                    <SubmitButton text="Setting" onClick={this.setting} />
                </div>

                <div class="deleteButton">
                    <SubmitButton text="Delete" onClick={this.deletePost} />
                </div>
                <div class="newButton">
                    <SubmitButton text="New" onClick={this.newPost} />
                </div>
                <div class="publishButton">
                    <SubmitButton text="Publish" onClick={this.publish} />
                </div>
                <div class="saveButton">
                    <SubmitButton text="Save" onClick={this.open} />
                </div>
            </div>
        );
    }

    setting() {
        sendMsg("Setting");
        this.setState({ setting: !this.state.setting });
    }
    deletePost() {
        sendMsg("Delete");
    }
    newPost() {
        sendMsg("New");
    }
    publish() {
        sendMsg("Publish");
    }
    save() {
        sendMsg("Save");
    }
    open(index) {
        sendMsg("open" + index.toString());
    }
    openPosts() {
        sendMsg("openPosts");
    }
    showSetting() {
        return (
            <Setting />
        )
    }

    showEditor() {
        return (
            <MDEditor />
        )
    }
    openEditor() {
        return (
            <MDEditor content={this.state.content} />
        )
    }

    componentDidMount() {
        console.log("Connecting...");
        connect((msg) => {
            var contentCode = "post_content: ";
            //console.log(msg.data)
            if (msg.data === "Hello From The Server!") {
                this.openPosts();
            }
            else if (msg.data.includes(contentCode)) {
                var content = msg.data.substr(contentCode.length);
                if (content !== this.state.content) {
                    console.log("Changed");
                    this.setState({ content: content });
                }

            }
            else if (msg.data.includes("title")) {
                var titles = msg.data.split("\n");
                console.log(titles.length);
                this.setState(prevState => ({
                    posts: [...this.state.posts, ...titles]
                }))
                console.log(this.state.posts);
                console.log(titles);
            }
            else if (msg.data === "Delete") {
                this.setState(prevState => ({
                    posts: [...this.state.posts.slice(0, -1)]
                }));
            }
            //console.log(this.state);
        });
    }
};

export default Body;
