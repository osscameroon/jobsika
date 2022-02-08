import { expect } from "chai";
import request from "supertest";
import dotenv from "dotenv";

dotenv.config();
const apiHost = process.env.API_HOST;
const endpoint = "ratings";

describe(`${endpoint}`, function () {
  describe("GET", function () {
    it("return a list of ratings", async function () {
      return request(apiHost)
        .get(endpoint)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body[0])).equal(
            '{"salary_id":1,"company_id":994,"company_rating_id":569,"rating":2,"salary":1624669,"company_name":"Realbridge","seniority":"Seniority","comment":"Mauris enim leo, rhoncus sed, vestibulum sit amet, cursus id, turpis. Integer aliquet, massa id lobortis convallis, tortor risus dapibus augue, vel accumsan tellus nisi eu orci. Mauris lacinia sapien quis libero.","job_title":"Recruiting Manager","country":"Country","city":"Livefish"}'
          );
        });
    });

    it("return a rating of id == 1", async function () {
      return request(apiHost)
        .get(`${endpoint}/1`)
        .send()
        .expect(200)
        .expect("content-type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body)).to.equal('{"salary_id":1,"company_id":994,"company_rating_id":569,"rating":2,"salary":1624669,"company_name":"Realbridge","seniority":"Seniority","comment":"Mauris enim leo, rhoncus sed, vestibulum sit amet, cursus id, turpis. Integer aliquet, massa id lobortis convallis, tortor risus dapibus augue, vel accumsan tellus nisi eu orci. Mauris lacinia sapien quis libero.","job_title":"Recruiting Manager","country":"Country","city":"Livefish"}');
        });
    });

    it("return 404 for an non existing rating id", async function () {
      return request(apiHost)
        .get(`${endpoint}/1001`)
        .send()
        .expect(404)
        .expect("content-type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body)).contains('could not find rating');
        });
    });

  });
});

const endpoint2 = `average-rating`
describe(`${endpoint2}`, function () {
  describe("GET", function () {
    it("return a list of ratings", async function () {
      return request(apiHost)
        .get(endpoint2)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body)).equal(
            '{"rating":3,"salary":2543437}'
          );
        });
    });
  });
});
