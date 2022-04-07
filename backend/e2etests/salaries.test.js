import { expect } from "chai";
import request from "supertest";
import dotenv from "dotenv";

dotenv.config();
const apiHost = process.env.API_HOST;
const endpoint = "salaries";

describe(`${endpoint}`, function () {
  describe("GET", function () {
    it("return a list of salairies", async function () {
      return request(apiHost)
        .get(endpoint)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body[0])).equal(
            '{"id":1,"title":2,"salary":1624669,"seniority":"Seniority","city":"Livefish","country":"Country","company_id":99,"company_rating_id":56,"createdat":"0001-01-01T00:00:00Z","updatedat":"0001-01-01T00:00:00Z"}'
          );
        });
    });

    it("return a salary of id == 1", async function () {
      return request(apiHost)
        .get(`${endpoint}/1`)
        .send()
        .expect(200)
        .expect("content-type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body)).to.equal('{"id":1,"title":2,"salary":1624669,"seniority":"Seniority","city":"Livefish","country":"Country","company_id":99,"company_rating_id":56,"createdat":"0001-01-01T00:00:00Z","updatedat":"0001-01-01T00:00:00Z"}');
        });
    });

    it("return 404 for an non existing rating id", async function () {
      return request(apiHost)
        .get(`${endpoint}/1001`)
        .send()
        .expect(404)
        .expect("content-type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body)).contains('could not find salary');
        });
    });

  });
});
