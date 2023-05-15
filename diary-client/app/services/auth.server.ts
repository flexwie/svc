import { Authenticator } from "remix-auth";
import { sessionStorage } from "./session.server";
import { User } from "@prisma/client";
import { prisma } from "~/services/prisma.server";
import { MicrosoftStrategy } from "remix-auth-microsoft";

export let authenticator = new Authenticator<User>(sessionStorage);

let msStrategy = new MicrosoftStrategy(
  {
    clientId: process.env["CLIENT_ID"]!,
    clientSecret: process.env["CLIENT_SECRET"]!,
    redirectUri: process.env["REDIRECT"]!,
    prompt: "login",
    tenantId: "common",
  },
  async ({ profile }) => {
    return prisma.user.upsert({
      where: { email: profile.emails[0].value },
      update: {
        name: profile._json.givenname + " " + profile._json.familyname,
      },
      create: {
        name: profile._json.givenname + " " + profile._json.familyname,
        email: profile.emails[0].value,
      },
    });
  }
);

authenticator.use(msStrategy);
