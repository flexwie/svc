import { Editor as DEditor } from "draft-js";
import {
  DraftHandleValue,
  EditorState,
  RichUtils,
  convertFromRaw,
  convertToRaw,
} from "draft-js";
import { FunctionComponent, useEffect, useRef, useState } from "react";

interface EditorProps {
  maxLength?: number;
  content: any;
  onChange: (raw: string) => void;
  readonly: boolean;
}

export const Editor: FunctionComponent<EditorProps> = ({
  readonly,
  onChange,
  content,
  maxLength = -1,
}) => {
  const [editorState, setEditorState] = useState(() =>
    content
      ? EditorState.createWithContent(convertFromRaw(JSON.parse(content)))
      : EditorState.createWithContent(
        convertFromRaw({
          entityMap: {},
          blocks: [
            {
              text: "",
              key: "foo",
              type: "unstyled",
              entityRanges: [],
            },
          ],
        })
      )
  );

  const ref = useRef<DEditor>();

  useEffect(() => {
    ref.current?.focus();
  }, []);

  useEffect(() => {
    // update onchange handler
    const content = editorState.getCurrentContent();
    const md = JSON.stringify(convertToRaw(content));

    onChange(md);
  }, [editorState]);

  const handleBeforeInput = () => {
    if (maxLength == -1) return "not-handled";

    const content = editorState.getCurrentContent();
    const length = content.getPlainText("").length;

    if (length + 1 > maxLength) {
      return "handled";
    }

    return "not-handled";
  };

  // handle draft editor commands
  const handleKeyCommand = (
    command: string,
    editorState: EditorState
  ): DraftHandleValue => {
    const newState = RichUtils.handleKeyCommand(editorState, command);

    if (newState) {
      setEditorState(newState);
      return "handled";
    }

    return "not-handled";
  };

  return (
    <div className="px-6 py-4 h-fit border rounded-lg -mb-7">
      <span className="relative -top-7 -left-2 bg-white px-2 text-slate-400 font-light">
        Entry
      </span>
      <div className="-mt-6">
        <DEditor
          readOnly={readonly}
          ref={ref}
          editorState={editorState}
          editorKey="editor"
          handleKeyCommand={handleKeyCommand}
          onChange={setEditorState}
          handleBeforeInput={handleBeforeInput}
        />
      </div>
    </div>
  );
};

export const getContentLength = (raw: string): number => {
  const data = JSON.parse(raw);
  const state = EditorState.createWithContent(convertFromRaw(data));
  return state.getCurrentContent().getPlainText("").length;
};
