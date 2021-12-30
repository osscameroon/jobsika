import { expect } from "chai";
import request from "supertest";
import dotenv from "dotenv";

dotenv.config();
const apiHost = process.env.API_HOST;
const endpoint = "companies";

describe(`${endpoint}`, function () {
  describe("GET", function () {
    it("return a list of companies", async function () {
      return request(apiHost)
        .get(endpoint)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body[0])).equal('{"id":1,"name":"Brainverse","rating":2,"createdat":"0001-01-01T00:00:00Z","updatedat":"0001-01-01T00:00:00Z"}');
        });
    });
  });
});
