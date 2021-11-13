import { expect } from "chai";
import request from "supertest";
import axios from "axios";
import dotenv from "dotenv";

dotenv.config();
const apiHost = process.env.API_HOST;
const endpoint = "signup";

describe(`${endpoint}`, function () {
  describe("POST", function () {
    const email = "test@email.com";
    const password = "test-password";

    it("creates new user with valid credentials", async function () {
      return request(apiHost)
        .post(endpoint)
        .send({ email: email, password: password })
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8");
    });

    it("fail to create a user with same email twice", async function () {
      const sameEmail = "sameEmail";
      const password = "password";

      const data = {
        email: sameEmail,
        password: password,
      };

      await axios({
        url: `${apiHost}${endpoint}`,
        method: "post",
        data: data,
      })
        .then(() => {})
        .catch(() => {});

      return request(apiHost)
        .post(endpoint)
        .send({ email: sameEmail, password: password })
        .expect(400)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          expect(res.body.error).to.contain("email already exists");
        });
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
