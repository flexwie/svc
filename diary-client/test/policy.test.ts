import { authenticated } from "~/services/policy.server";
import { commitSession, getSession } from "~/services/session.server";
import { Request, Headers } from "@remix-run/node";

describe("authenticated policy", () => {
  it("should redirect unauthenticated", async () => {
    const session = await getSession("");

    const request = new Request("http://test.com/test", {
      headers: { Cookie: await commitSession(session) },
    });

    const cb = jest.fn();

    await authenticated(request, cb);

    expect(cb).not.toHaveBeenCalled();
  });

  it("should allow authenticated", async () => {
    const session = await getSession("");
    session.set("user", JSON.stringify({ id: 123 }));

    const request = new Request("http://test.com/test", {
      headers: { Cookie: await commitSession(session) },
    });

    const cb = jest.fn();

    await authenticated(request, cb);

    expect(cb).toHaveBeenCalled();
  });
});
