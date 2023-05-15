import { useSWEffect } from "~/services/client/sw-hook";
import { cssBundleHref } from "@remix-run/css-bundle";
import type { LinksFunction } from "@remix-run/node";
import {
  Links,
  LiveReload,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
} from "@remix-run/react";
import stylesheet from "~/tailwind.css";
import draftStyle from "draft-js/dist/Draft.css";
export const links: LinksFunction = () => [
  ...(cssBundleHref ? [{ rel: "stylesheet", href: cssBundleHref }] : []),
  { rel: "stylesheet", href: stylesheet },
  { rel: "stylesheet", href: draftStyle },
];
export default function App() {
  useSWEffect();
  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta
          name="viewport"
          content="width=device-width,initial-scale=1,maximum-scale=1,user-scalable=0"
        />
        <Meta />
        <link rel="manifest" href="/resources/manifest.webmanifest" />
        <Links />
      </head>
      <body>
        <div className="mb-5 mt-8 flex justify-around">
          <h1 className="font-hand text-5xl">Diary</h1>
        </div>
        <hr className="mx-6" />
        <Outlet /> <ScrollRestoration /> <Scripts /> <LiveReload />
      </body>
    </html>
  );
}
