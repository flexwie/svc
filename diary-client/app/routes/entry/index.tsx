import { LoaderFunction, redirect } from "@remix-run/node";
import { DateTime } from "luxon";

export const loader: LoaderFunction = () => {
  const now = DateTime.now().toISODate();
  return redirect(`/entry/${now}`);
};
