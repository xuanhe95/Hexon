import React, { Component } from "react";
import "./App.css";
import { connect, sendMsg } from "./api";
import Header from "./components/Header/Header";
import Footer from "./components/Footer/Footer";
import PostList from "./components/PostList/PostList"
import SubmitButton from "./components/SubmitButton";
import Tiptap from "./components/Tiptap/Tiptap";

import Body from "./components/Body/Body";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      posts: [],

    };
  }

  send() {
    console.log("hello");
    sendMsg("Posts");
  }

  render() {

    return (
      <div className="App">
        <Header />
        <Body />
        <Footer />

      </div>
    );
  }


}
export default App;


// import logo from './logo.svg';
// import './App.css';

// function App() {
//   return (
//     <div className="App">
//       <header className="App-header">
//         <img src={logo} className="App-logo" alt="logo" />
//         <p>
//           Edit <code>src/App.js</code> and save to reload.
//         </p>
//         <a
//           className="App-link"
//           href="https://reactjs.org"
//           target="_blank"
//           rel="noopener noreferrer"
//         >
//           Learn React
//         </a>
//       </header>
//     </div>
//   );
// }