import { expect } from "chai";
import request from "supertest";
import dotenv from "dotenv";

dotenv.config();
const apiHost = process.env.API_HOST;
const endpoint = "jobs";


describe(`${endpoint}`, function () {
    describe("GET", function () {
        it("return a list of offers with page=1 & limit=2", async function () {
            return request(apiHost)
                .get(`${endpoint}?page=1&limit=2`)
                .send()
                .expect(200)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    const body = JSON.stringify(res.body)
                    expect(body).contains('"hits":');
                    expect(body).contains('"nbHits":5,"offset":0,"limit":2}');
                });
        });
    })

    describe("GET", function () {
        it("return a list of remote offers ", async function () {
            return request(apiHost)
                .get(`${endpoint}?isRemote=true`)
                .send()
                .expect(200)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    const body = JSON.stringify(res.body)
                    expect(body).contains('"hits":');
                    expect(body).contains('"nbHits":2,"offset":0,"limit":20}');
                });
        });
    })

    describe("GET", function () {
        it("return a list of non-remote offers ", async function () {
            return request(apiHost)
                .get(`${endpoint}?isRemote=false`)
                .send()
                .expect(200)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    const body = JSON.stringify(res.body)
                    expect(body).contains('"hits":');
                    expect(body).contains('"nbHits":3,"offset":0,"limit":20}');
                });
        });
    })
});