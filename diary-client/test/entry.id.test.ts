import { loader } from "../app/routes/entry/$id";

describe("loader", () => {
  it("returns on unauthorised", async () => {
    const response = await loader({
      request: new Request("http://test.com/entry/01-01-2020"),
      params: { id: "01-01-2020" },
      context: {},
    });

    expect(response).toHaveProperty("status");
  });
});
