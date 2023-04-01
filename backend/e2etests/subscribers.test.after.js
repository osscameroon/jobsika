import { expect } from "chai";
import request from "supertest";
import dotenv from "dotenv";

dotenv.config();
const apiHost = process.env.API_HOST;
const endpoint = "subscribers";

describe(`${endpoint}`, function () {
    describe("POST", function () {
        it("post a subscriber with valid email", async function () {
            return request(apiHost)
                .post(`${endpoint}`)
                .set("Accept", "application/json")
                .send({
                    email: "ossdevs-cm@gmail.com",
                })
                .expect(201)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    expect(JSON.stringify(res.body)).contains("ossdevs-cm@gmail.com");
                });
        });

        it("post a subscriber with existing email", async function () {
            return request(apiHost)
                .post(`${endpoint}`)
                .set("Accept", "application/json")
                .send({
                    email: "ossdevs-cm@gmail.com",
                })
                .expect(201)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    expect(JSON.stringify(res.body)).contains("ossdevs-cm@gmail.com");
                });
        });

        it("post a subscriber with empty email", async function () {
            return request(apiHost)
                .post(`${endpoint}`)
                .set("Accept", "application/json")
                .send({
                    email: "",
                })
                .expect(400)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    expect(JSON.stringify(res.body)).contain("email is not a valid email address");
                });
        });

        it("post a subscriber with invalid email", async function () {
            return request(apiHost)
                .post(`${endpoint}`)
                .set("Accept", "application/json")
                .send({
                    email: "ossdevs-cm@gmail",
                })
                .expect(400)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    expect(JSON.stringify(res.body)).contain("email is not a valid email address");
                });
        });
    });
});