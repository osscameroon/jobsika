import { expect } from "chai";
import request from "supertest";
import dotenv from "dotenv";

dotenv.config();
const apiHost = process.env.API_HOST;
const endpoint = "ratings";

describe(`${endpoint}`, function () {
  describe("GET", function () {
    it("return a list of ratings with page=1 & limit=2", async function () {
      return request(apiHost)
        .get(`${endpoint}?page=1&limit=2`)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          const body = JSON.stringify(res.body)
          expect(body).contains('"hits":');
          expect(body).contains('"limit":2,"nbHits":100,"offset":0}');
        });
    });

    it("return a list of ratings with page=3", async function () {
      return request(apiHost)
        .get(`${endpoint}?page=3`)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          const body = JSON.stringify(res.body)
          expect(body).contains('"hits":');
          expect(body).contains('"limit":20,"nbHits":100,"offset":40}');
        });
    });

    it("return a list of ratings with jobtitle=Statistician II", async function () {
      return request(apiHost)
        .get(`${endpoint}?jobtitle=Statistician II`)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          const body = JSON.stringify(res.body)
          expect(body).contains('"hits":');
          expect(body).contains('"limit":20,"nbHits":16,"offset":0}');
        });
    });

    it("return a list of ratings with company=Fliptune and jobtitle=Statistician II", async function () {
      return request(apiHost)
        .get(`${endpoint}?company=Fliptune&jobtitle=Statistician II`)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          const body = JSON.stringify(res.body)
          expect(body).contains('"hits":');
          expect(body).contains('"limit":20,"nbHits":1,"offset":0}');
        });
    });

    it("return a list of ratings with company=Jaxbean", async function () {
      return request(apiHost)
        .get(`${endpoint}?company=Jaxbean`)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          const body = JSON.stringify(res.body)
          expect(body).contains('"hits":');
          expect(body).contains('"limit":20,"nbHits":1,"offset":0}');
        });
    });

    it("return a rating of id == 1", async function () {
      return request(apiHost)
        .get(`${endpoint}/1`)
        .send()
        .expect(200)
        .expect("content-type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body)).to.equal('{"salary_id":1,"company_id":99,"company_rating_id":0,"rating":0,"salary":1624669,"company_name":"Jaxbean","seniority":"Seniority","comment":"","job_title":"Assistant Manager","country":"Country","city":"Livefish","createdat":"0001-01-01T00:00:00Z"}');
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
    it("return the average of ratings", async function () {
      return request(apiHost)
        .get(endpoint2)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body)).equal(
            '{"rating":4,"salary":2494171}'
          );
        });
    });

    it("return the average of ratings with jobtitle=Statistician II", async function () {
      return request(apiHost)
        .get(`${endpoint2}?jobtitle=Statistician II`)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body)).equal(
            '{"rating":4,"salary":2824939}'
          );
        });
    });

    it("return the average of ratings with company=Jaxbean", async function () {
      return request(apiHost)
        .get(`${endpoint2}?company=Jaxbean`)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body)).equal(
            '{"rating":0,"salary":1624669}'
          );
        });
    });

    it("return a list of ratings with company=Fliptune and jobtitle=Statistician II", async function () {
      return request(apiHost)
        .get(`${endpoint2}?company=Fliptune&jobtitle=Statistician II`)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
          expect(JSON.stringify(res.body)).equal(
            '{"rating":0,"salary":841396}'
          );
        });
    });

  });
});
