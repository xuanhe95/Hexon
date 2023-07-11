import React, { useMemo } from "react";
import SimpleMDE, { SimpleMdeReact } from "react-simplemde-editor";
import EasyMDE from "easymde";

const MDEditor = (props) => {
    console.log(props.content)
    console.log("TEST")
    //var md = customMarkdownParser(props.content);
    const spellCheckerOptions = useMemo(() => {
        return {
            spellChecker: false,
            //previewRender: (plainText, preview) => { return props.content },

        }
    })
    return (
        <div className="editor">
            <SimpleMdeReact
                value={props.content}
                options={spellCheckerOptions}
            />
        </div>
    );

}


export default MDEditor;
