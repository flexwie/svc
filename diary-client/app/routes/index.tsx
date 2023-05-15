import { V2_MetaFunction } from "@remix-run/node";
import { Outlet } from "@remix-run/react";

export const meta: V2_MetaFunction = () => {
  return [{ title: "New Remix App" }];
};

export default function Index() {
  return (
    <div className="flex flex-col justify-center">
      <form action="/auth/" method="post">
        <button className="text-white bg-black font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center hover:shadow-md">
          Login
        </button>
      </form>
      <div>
        <img src="img/diary.svg" className="w-52" />
      </div>
      <Outlet />
    </div>
  );
}
