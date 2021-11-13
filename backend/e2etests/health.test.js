import { expect } from "chai";
import request from "supertest";
import axios from "axios";
import dotenv from "dotenv";

dotenv.config()
const apiHost = process.env.API_HOST
const endpoint = "health"

describe(`${endpoint}`, function () {
  describe("GET", function () {
    it("returns 200 status", async function () {
      return request(apiHost)
        .get(endpoint)
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8");
    });
  });
});
