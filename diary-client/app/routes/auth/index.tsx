import { ActionFunction, redirect } from "@remix-run/node";
import { authenticator } from "~/services/auth.server";

export const loader = () => {
  return redirect("/");
};

export const action: ActionFunction = ({ request }) => {
  return authenticator.authenticate("microsoft", request);
};
