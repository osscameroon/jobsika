import { expect } from "chai";
import request from "supertest";
import dotenv from "dotenv";

dotenv.config();
const apiHost = process.env.API_HOST;
const endpoint = "cities";

describe(`${endpoint}`, function () {
  describe("GET", function () {
    it("return a list of cities", async function () {
      return request(apiHost)
        .get(endpoint)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
            expect(JSON.stringify(res.body[0])).equal('"Abong-Mbang"')
        });
    });
  });
});
