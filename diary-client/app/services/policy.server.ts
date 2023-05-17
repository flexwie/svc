import { User } from "@prisma/client";
import { Request, redirect } from "@remix-run/node";
import { authenticator } from "./auth.server";
import { commitSession, getSession } from "./session.server";

export type Policy<PolicyResult, T> = (
  request: Request,
  callback: <T>(input: PolicyResult) => Promise<T>
) => Promise<T>;

export const authenticated: Policy<{ user: User }, any> = async (
  request,
  callback
) => {
  let user = await authenticator.isAuthenticated(request);
  if (!user) {
    const session = await getSession(request.headers.get("Cookie"));
    session.set("origin", request.url);
    const cookie = await commitSession(session);

    return redirect("/", { headers: { "Set-Cookie": cookie } });
  }

  return await callback({ user });
};
