import React, { useMemo } from "react";
import SimpleMDE, { SimpleMdeReact } from "react-simplemde-editor";
import EasyMDE from "easymde";

const MDEditor = (props) => {
    //var md = customMarkdownParser(props.content);
    const spellCheckerOptions = useMemo(() => {
        return {
            spellChecker: false,
            //previewRender: (plainText, preview) => { return props.content },

        }
    })

    const getContent = (value) => {
        props.onGetContent(value);
    }

    return (
        <div className="editor">
            <SimpleMdeReact
                key={props.key}
                value={props.content}
                options={spellCheckerOptions}
                onChange={getContent}
            />
        </div>
    );

}


export default MDEditor;
