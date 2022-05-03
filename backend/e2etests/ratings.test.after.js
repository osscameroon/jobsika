import { expect } from "chai";
import request from "supertest";
import dotenv from "dotenv";

dotenv.config();
const apiHost = process.env.API_HOST;
const endpoint = "ratings";

describe("POST", function () {
  it("POST add a new rating entry", async function () {
    return request(apiHost)
      .post(`${endpoint}?page=1&limit=2`)
      .set("Accept", "application/json")
      .send({
        company_name: "La Mater",
        //we set the salary to zero to avoid breaking the average-ratings tests
        //as the salary set -1 are not counted in the calculation of the average
        salary: 0,
        //we set the rating to zero to avoid breaking the average-ratings tests
        //as the rating set -1 are not counted in the calculation of the average
        rating: 0,
        job_title: "a job_title",
        comment: "my comment",
        seniority: "Seniority",
        city: "Maroua",
        //the country field is omitted here as we always set it to Cameroon for now
      })
      .expect(201)
      .expect("Content-Type", "application/json; charset=utf-8");
  });

  it("POST, fail to add a new rating entry without job_title", async function () {
    return request(apiHost)
      .post(`${endpoint}?page=1&limit=2`)
      .set("Accept", "application/json")
      .send({
        company_name: "La Mater",
        //we set the salary to zero to avoid breaking the average-ratings tests
        //as the salary set -1 are not counted in the calculation of the average
        salary: 0,
        //we set the rating to zero to avoid breaking the average-ratings tests
        //as the rating set -1 are not counted in the calculation of the average
        rating: 0,
        comment: "my comment",
        seniority: "Seniority",
        city: "Maroua",
      })
      .expect(400)
      .expect("Content-Type", "application/json; charset=utf-8");
  });

  it("POST, fail to add a new rating entry without city", async function () {
    return request(apiHost)
      .post(`${endpoint}?page=1&limit=2`)
      .set("Accept", "application/json")
      .send({
        company_name: "La Mater",
        //we set the salary to zero to avoid breaking the average-ratings tests
        //as the salary set -1 are not counted in the calculation of the average
        salary: 0,
        //we set the rating to zero to avoid breaking the average-ratings tests
        //as the rating set -1 are not counted in the calculation of the average
        rating: 0,
        job_title: "a job_title",
        comment: "my comment",
        seniority: "Seniority",
      })
      .expect(400)
      .expect("Content-Type", "application/json; charset=utf-8");
  });

  it("POST add a new rating entry with lowercase cities, company_name, and jobtitle", async function () {
    const sendRequest = () =>
      request(apiHost)
        .post(`${endpoint}?page=1&limit=2`)
        .set("Accept", "application/json")
        .send({
          company_name: "la mater",
          job_title: "technicien de surface",
          //we set the salary to zero to avoid breaking the average-ratings tests
          //as the salary set -1 are not counted in the calculation of the average
          salary: 0,
          //we set the rating to zero to avoid breaking the average-ratings tests
          //as the rating set -1 are not counted in the calculation of the average
          rating: 0,
          comment: "my comment",
          seniority: "Seniority",
          city: "maroua",
          //the country field is omitted here as we always set it to Cameroon for now
        })
        .expect(201)
        .expect("Content-Type", "application/json; charset=utf-8");

    for (let i = 0; i < 3; i++) {
      await sendRequest();
    }
    return request(apiHost)
      .get(`${endpoint}?limit=200`)
      .set("Accept", "application/json")
      .send()
      .expect(200)
      .expect("Content-Type", "application/json; charset=utf-8")
      .then((res) => {
        const lastElem = res.body.hits.length - 1;
        const rating = JSON.stringify(res.body.hits[lastElem]);
        expect(rating).contains("La Mater");
        expect(rating).contains("Technicien De Surface");
        expect(rating).contains("Maroua");
      });
  });

  it("Only show ratings who has 3 entry with the same job title and company when company filter is apply", async function () {
    const sendRequest = (jobtitle) =>
      request(apiHost)
        .post(`${endpoint}`)
        .set("Accept", "application/json")
        .send({
          company_name: "Osstest",
          job_title: jobtitle,
          //we set the salary to zero to avoid breaking the average-ratings tests
          //as the salary set -1 are not counted in the calculation of the average
          salary: 0,
          //we set the rating to zero to avoid breaking the average-ratings tests
          //as the rating set -1 are not counted in the calculation of the average
          rating: 0,
          comment: "my comment",
          seniority: "Seniority",
          city: "maroua",
          //the country field is omitted here as we always set it to Cameroon for now
        })
        .expect(201)
        .expect("Content-Type", "application/json; charset=utf-8");

    for (let i = 0; i < 1; i++) {
      await sendRequest("Designer");
    }
    for (let i = 0; i < 3; i++) {
      await sendRequest("Dev");
    }
    return request(apiHost)
      .get(`${endpoint}?company=Osstest`)
      .set("Accept", "application/json")
      .send()
      .expect(200)
      .expect("Content-Type", "application/json; charset=utf-8")
      .then((res) => {
        expect(res.body.hits.every((r) => r.job_title === 'Dev')).equal(true)
      });
  });

  it("POST add a new rating entry for one company, fields will be hide", async function () {
    return request(apiHost)
      .post(endpoint)
      .set("Accept", "application/json")
      .send({
        company_name: "Company1rating",
        //we set the salary to zero to avoid breaking the average-ratings tests
        //as the salary set -1 are not counted in the calculation of the average
        salary: 0,
        //we set the rating to zero to avoid breaking the average-ratings tests
        //as the rating set -1 are not counted in the calculation of the average
        rating: 0,
        comment: "my comment",
        seniority: "Seniority",
        city: "Maroua",
        job_title: "Cleaner",
        //the country field is omitted here as we always set it to Cameroon for now
      })
      .expect(201)
      .expect("Content-Type", "application/json; charset=utf-8")
      .then(async () => {
        return request(apiHost)
          .get(`${endpoint}?page=1&limit=5000`)
          .set("Accept", "application/json")
          .send()
          .expect(200)
          .expect("Content-Type", "application/json; charset=utf-8")
          .then((res) => {
            const lastElem = res.body.hits.length - 1;
            const rating = res.body.hits[lastElem];
            expect(rating.company_name).equal("");
            expect(rating.comment).equal("");
            expect(rating.company_id).equal(0);
            expect(rating.rating).equal(0);
          });
      });
  });

  it("POST add 3 new rating entry for one company, fields will not be hide", async function () {
    // TODO: add this entry 3 times
    const sendRequest = () => request(apiHost)
      .post(`${endpoint}?page=3`)
      .set("Accept", "application/json")
      .send({
        company_name: "Company3rating",
        //we set the salary to zero to avoid breaking the average-ratings tests
        //as the salary set -1 are not counted in the calculation of the average
        salary: 0,
        //we set the rating to zero to avoid breaking the average-ratings tests
        //as the rating set -1 are not counted in the calculation of the average
        rating: 3,
        comment: "my comment",
        seniority: "Seniority",
        city: "Maroua",
        job_title: "Cleaner",
        //the country field is omitted here as we always set it to Cameroon for now
      })
      .expect(201)
      .expect("Content-Type", "application/json; charset=utf-8");

    for (let i = 0; i < 3; i++) {
      await sendRequest();
    }

    return request(apiHost)
      .get(`${endpoint}?page=1&limit=10&company=Company3rating`)
      .set("Accept", "application/json")
      .send()
      .expect(200)
      .expect("Content-Type", "application/json; charset=utf-8")
      .then((res) => {
        const lastElem = res.body.hits.length - 1;
        const rating = res.body.hits[lastElem];
        expect(rating.company_name).equal("Company3rating");
        expect(rating.comment).equal("my comment");
        expect(rating.company_id !== 0).equal(true);
        expect(rating.rating).equal(3);
      });
  });
});
