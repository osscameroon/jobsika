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
        .set('Accept', 'application/json')
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
          city: "Maroua"
          //the country field is omitted here as we always set it to Cameroon for now
        })
        .expect(201)
        .expect("Content-Type", "application/json; charset=utf-8");
    });
});
