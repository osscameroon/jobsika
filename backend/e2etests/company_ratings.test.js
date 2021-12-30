import { expect } from "chai";
import request from "supertest";
import dotenv from "dotenv";

dotenv.config();
const apiHost = process.env.API_HOST;
const endpoint = "company-ratings";

describe(`${endpoint}`, function () {
  describe("GET", function () {
    it("return a list of company-ratings", async function () {
      return request(apiHost)
        .get(endpoint)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body[0])).equal('{"id":1,"company_id":761,"rating":4,"comment":"Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Proin risus. Praesent lectus.","createdat":"0001-01-01T00:00:00Z","updatedat":"0001-01-01T00:00:00Z"}');
        });
    });

    it("return a company-ratings of rating id == 1", async function () {
      return request(apiHost)
        .get(`${endpoint}/1`)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body)).equal('{"id":1,"company_id":761,"rating":4,"comment":"Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Proin risus. Praesent lectus.","createdat":"0001-01-01T00:00:00Z","updatedat":"0001-01-01T00:00:00Z"}');
        });
    });

    it("return company-ratings 404 for an non existing rating id", async function () {
      return request(apiHost)
        .get(`${endpoint}/1001`)
        .send()
        .expect(404)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body)).contains('could not find company ratings');
        });
    });

    it("return a company-ratings of company id == 763", async function () {
      return request(apiHost)
        .get(`${endpoint}?company_id=763`)
        .send()
        .expect(200)
        .expect("content-type", "application/json; charset=utf-8")
        .then((res) => {
          expect(res.body.length).to.equal(2);
        });
    });
  });
});
