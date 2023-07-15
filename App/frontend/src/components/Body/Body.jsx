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
            key: 0,
        };
        this.post_index = -1;
        this.setting = this.setting.bind(this);
        this.newPost = this.newPost.bind(this);
        this.updatePostIndex = this.updatePostIndex.bind(this);
        this.deletePost = this.deletePost.bind(this);
        this.openEditor = this.openEditor.bind(this);
        this.getTextContentFromEditor = this.getTextContentFromEditor.bind(this);
        this.save = this.save.bind(this);
        this.post_content = "";
    }
    render() {
        var content = null;

        if (this.state.setting) {
            content = this.showSetting()
        }
        else {
            content = this.openEditor();
        }
        return (
            <div class="body">
                <div className="BodyContainer">
                    <PostList posts={this.state.posts}
                        update={this.updatePostIndex}
                        index={this.post_index}
                    />
                    {content}
                </div>
                <div class="ButtonBar">
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
                        <SubmitButton text="Save" onClick={this.save} onSave={this.saveContent} />
                    </div>
                </div >
            </div>
        );
    }

    updatePostIndex = (index) => {
        sendMsg("open_" + String(index));
        this.post_index = index;
    }

    setting() {
        sendMsg("Setting");
        this.setState({ setting: !this.state.setting });
    }
    deletePost() {
        if (this.post_index < 0) return;
        if (this.state.posts.length === 0) return;
        sendMsg("delete_" + this.post_index.toString());    // 删除文章
        while (this.post_index >= this.state.posts.length - 1) {
            this.post_index--;
        }
        if (this.post_index >= 0) {
            this.updatePostIndex(this.post_index);  // 重新打开下一篇文章
        } else {
            this.setState({ posts: [] }); // 文章列表为空
        }

    }
    newPost() {
        sendMsg("new_post"); // 新建文章
        this.post_index = 0;    // 设置当前文章索引
        this.setState({ content: "" }); // 清空文章内容
        console.log("New Post");
    }
    publish() {
        sendMsg("publish_posts");  // 发布文章
    }


    save() {
        const msg = "save_content_to_" + this.post_index + ":" + this.post_content;
        sendMsg(msg);   // 保存文章
    }

    getTextContentFromEditor(content) {
        this.post_content = content;    // 获取文章内容
    }

    openPosts() {
        sendMsg("initilize_all_posts");   // 打开文章列表
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
            <MDEditor content={this.state.content}
                onGetContent={this.getTextContentFromEditor}
                key={this.state.key}
            />
        )
    }

    componentDidMount() {
        console.log("Connecting...");
        connect((msg) => {
            var contentCode = "post_content_";
            var titleCode = "post_title_";
            //console.log(msg.data)
            if (msg.data === "Hello From The Server!") {
                this.openPosts();
            }
            else if (msg.data.includes(contentCode)) {
                this.setState({ key: this.state.key + 1 });
                console.log(this.state.key);
                var content = msg.data.substr(contentCode.length);
                //if (content !== this.state.content) {
                console.log("Changed");
                this.setState({ content: content });
                //}
            }
            else if (msg.data.includes(titleCode)) {
                var titles = msg.data.split("\n");
                titles = titles.map((title) => {
                    return title.substr(titleCode.length);
                });
                console.log(titles.length);
                this.setState({ posts: titles });
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
