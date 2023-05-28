import { expect } from "chai";
import request from "supertest";
import dotenv from "dotenv";

dotenv.config();
const apiHost = process.env.API_HOST;
const endpoint = "pay";

describe(`${endpoint}`, function () {
  describe("POST", function () {
    it("fails to make a payment with invalid email address", async function () {
      return request(apiHost)
        .post(`${endpoint}`)
        .set("Accept", "application/json")
        .send({
          email: "wrong email",
          tier: "new jobsika tier",
          job_offer_id: "1",
        })
        .expect(400)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body)).contain("email field is invalid");
        });
    });

    it("proceed with a payment", async function () {
      return request(apiHost)
        .post(`${endpoint}`)
        .set("Accept", "application/json")
        .send({
          email: "test@email.com",
            tier: "new jobsika tier",
          job_offer_id: "1",
        })
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body)).contain(
            "opencollective.com"
          );
        });
    });

  });
});
