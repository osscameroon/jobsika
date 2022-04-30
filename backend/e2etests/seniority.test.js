import { expect } from "chai";
import request from "supertest";
import dotenv from "dotenv";

dotenv.config();
const apiHost = process.env.API_HOST;
const endpoint = "seniority";

describe(`${endpoint}`, function () {
  describe("GET", function () {
    it("return a list of seniority", async function () {
      return request(apiHost)
        .get(endpoint)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
            expect(JSON.stringify(res.body)).equal('["Intern","Entry-level","Mid-level","Senior","Above senior-level","Executive"]')
        });
    });
  });
});
