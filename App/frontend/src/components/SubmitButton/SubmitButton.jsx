import React, { Component } from "react";
import "./SubmitButton.scss";


var SubmitButton = ({ text, onClick }) => {
    return (
        <div className="buttonContainer">
            <div className="button">
                <button onClick={onClick}>
                    {text}
                </button>
            </div >
        </div>
    );


}

export default SubmitButton;