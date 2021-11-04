import { expect } from "chai";
import request from "supertest";
import axios from "axios";
import dotenv from "dotenv";

dotenv.config();
const apiHost = process.env.API_HOST;
const endpoint = "signin";

describe(`${endpoint}`, function () {
  describe("POST", function () {
    const email = "testsignin@email.com";
    const password = "testsigisigin-password";

    it("log user in with valid credentials", async function () {
    const email = "testsignin-1@email.com";
    const password = "testsigisigin-1-password";

      const data = {
        email: email,
        password: password,
      };

      //First sign user up
      await axios({
        url: `${apiHost}signup`,
        method: "post",
        data: data,
      })

      return request(apiHost)
        .post(endpoint)
        .send({ email: email, password: password })
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          expect(res.body.token).to.contain("Barear")
        });
    });

    it("fail to log user in with wrong credentials", async function () {
      const email = "testsignin-2@email.com";
      const password = "testsigisigin-2-password";

      const data = {
        email: email,
        password: "wrong-password",
      };

      //First sign user up
      await axios({
        url: `${apiHost}signup`,
        method: "post",
        data: data,
      })

      return request(apiHost)
        .post(endpoint)
        .send({ email: email, password: password })
        .expect(401)
        .expect("Content-Type", "application/json; charset=utf-8");
    });


    it("fail to create new user with invalid credentials, no email", async function () {
      return request(apiHost)
        .post(endpoint)
        .send({ password: password })
        .expect(400)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          expect(res.body.error).to.contain("wrong body");
        });
    });

    it("fail to create new user with invalid credentials, no password", async function () {
      return request(apiHost)
        .post(endpoint)
        .send({ email: email })
        .expect(400)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          expect(res.body.error).to.contain("wrong body");
        });
    });
  });
});
