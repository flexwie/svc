import { PencilIcon } from "@heroicons/react/24/solid";
import {
  ActionArgs,
  LoaderArgs,
  V2_MetaFunction,
  json,
  redirect,
} from "@remix-run/node";
import {
  Form,
  useLoaderData,
  useNavigation,
  useParams,
  useTransition,
} from "@remix-run/react";
import { FunctionComponent, useEffect, useMemo, useRef, useState } from "react";
import { prisma } from "~/services/prisma.server";
import { Editor, getContentLength } from "~/components/editor";
import { authenticator } from "~/services/auth.server";

export async function loader({ request, params }: LoaderArgs) {
  const user = await authenticator.isAuthenticated(request, {
    failureRedirect: "/",
  });

  if (!params.id) {
    // no date found; this should not happen but better safe than sorry
    throw new Error("could not read date from query");
  }

  const date = new Date(params.id);

  const entries = await prisma.entry.findUnique({
    where: { entryIdentifier: { date, ownerId: user.id } },
  });

  return entries;
}

export async function action({ request, params }: ActionArgs) {
  const user = await authenticator.isAuthenticated(request);
  if (!user) return redirect("/");

  const body = await request.formData();

  if (!body.has("data"))
    return json({ error: "missing data property" }, { status: 400 });
  if (!params.id) throw Error("could not get id from url");

  const content = body.get("data");
  const date = new Date(params.id);

  await prisma.entry.upsert({
    where: { entryIdentifier: { date, ownerId: user.id } },
    update: { content: content?.valueOf() },
    create: {
      date,
      content: content!.valueOf().toString(),
      ownerId: user.id,
    },
  });

  return redirect(`/entry/${params.id}`);
}

export const meta: V2_MetaFunction = () => {
  return [{ title: `Diary` }];
};

export default function Entry() {
  const data = useLoaderData<typeof loader>();
  const { id } = useParams();
  const transition = useNavigation();

  const [editorContent, setEditorContent] = useState<string | undefined>(
    data?.content
  );
  const [editorOpen, setEditorOpen] = useState(false);
  const [contentLength, setContentLength] = useState(0);

  const handleButtonClick = () => {
    setEditorOpen(true);
  };

  const handleEditorChange = (c: string) => {
    setEditorContent(c);
    setContentLength(getContentLength(c));
  };

  const readonly = useMemo(
    () => !editorOpen || transition.state === "submitting",
    [editorOpen, transition]
  );

  return (
    <div>
      {data?.content || editorOpen ? (
        <Form method="post" action={`/entry/${id}`}>
          <div className="mb-4">
            <Editor
              readonly={readonly}
              content={data?.content}
              onChange={handleEditorChange}
              maxLength={160}
            />
          </div>
          <input hidden readOnly name="data" value={editorContent || ""} />
          {editorOpen && (
            <div className="flex justify-between items-center mt-10">
              <button
                className="border-black font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center hover:shadow-md border-2"
                type="submit"
                disabled={
                  transition.state == "submitting" || contentLength == 0
                }
              >
                {transition.state == "submitting" ? "Saving..." : "Save"}
              </button>
              <span className="font-thin text-sm">{contentLength}/160</span>
            </div>
          )}
        </Form>
      ) : (
        <CreateNewButton onClick={handleButtonClick} />
      )}
    </div>
  );
}

interface CreateNewButtonProps {
  onClick: () => void;
}

const CreateNewButton: FunctionComponent<CreateNewButtonProps> = ({
  onClick,
}) => {
  return (
    <div className="flex items-center justify-around">
      <button
        onClick={() => onClick()}
        type="button"
        className="text-white bg-black font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center hover:shadow-md"
      >
        <PencilIcon className="w-4 h-4 mr-2 -ml-1" />
        Create new
      </button>
    </div>
  );
};
