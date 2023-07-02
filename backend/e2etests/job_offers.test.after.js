import { expect } from "chai";
import request from "supertest";
import dotenv from "dotenv";
import fs from "fs";
import path from 'path';

dotenv.config();
const apiHost = process.env.API_HOST;
const endpoint = "jobs";


describe(`${endpoint}`, function () {
    describe("POST", function () {
        it("post a valid job offer without image", async function () {
            return request(apiHost)
                .post(`${endpoint}`)
                .set("Accept", "application/json")
                .send({
                    company_email: "ossdevs-cm@gmail.com",
                    company_name: "OssCameroon",
                    job_title: "Frontend Dev",
                    is_remote: true,
                    city: "Douala",
                    country: "Cameroon",
                    salary_range_min: 100000,
                    salary_range_max: 150000,
                    department: "Research and Development",
                    description: "OssCameroon is hiring a Remote Go Backend Engineer \n Remote - We are looking for a backend engineer who can work 30+ hr/weekOur ideal candidate has:- 1+ years experience writing in Go (golang.org)- 2-3 years experience writing REST APIs- Experience working at a small startup- A passion for building something meaningful â€¦ Salary and compensation No salary data published by company so we estimated salary based on similar jobs related to Golang, Engineer and Backend jobs that are similar: $70,000 — $120000/year",
                    benefits: "Health insurance, dental insurance, 401k",
                    how_to_apply: "Please submit your resume and cover letter.",
                    application_url: "",
                    application_email_address: "ossdevs-cm@gmail.com",
                    application_phone_number: "555-555-5555",
                    tags: "remote, golang, backend, engineer"
                })
                .expect(201)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    expect(res.body.company_name).equal("OssCameroon");
                    expect(res.body.has_image).equal(false);
                });
        });

        it("post a job offer without email", async function () {
            return request(apiHost)
                .post(`${endpoint}`)
                .set("Accept", "application/json")
                .send({
                    company_email: "",
                    company_name: "OssCameroon",
                    job_title: "Frontend Dev",
                    is_remote: true,
                    city: "",
                    country: "",
                    salary_range_min: 100000,
                    salary_range_max: 150000,
                    description: "OssCameroon is hiring a Remote Go Backend Engineer \n Remote - We are looking for a backend engineer who can work 30+ hr/weekOur ideal candidate has:- 1+ years experience writing in Go (golang.org)- 2-3 years experience writing REST APIs- Experience working at a small startup- A passion for building something meaningful â€¦ Salary and compensation No salary data published by company so we estimated salary based on similar jobs related to Golang, Engineer and Backend jobs that are similar: $70,000 — $120000/year",
                    how_to_apply: "",
                    application_url: "",
                    application_email_address: "ossdevs-cm@gmail.com"
                })
                .expect(400)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    expect(JSON.stringify(res.body)).contain("company email is not a valid email address");
                });
        });

        it("post a job offer without a valid email", async function () {
            return request(apiHost)
                .post(`${endpoint}`)
                .set("Accept", "application/json")
                .send({
                    company_email: "ossdevs-cm@gmail",
                    company_name: "OssCameroon",
                    job_title: "Frontend Dev",
                    is_remote: true,
                    city: "",
                    country: "",
                    salary_range_min: 100000,
                    salary_range_max: 150000,
                    description: "OssCameroon is hiring a Remote Go Backend Engineer \n Remote - We are looking for a backend engineer who can work 30+ hr/weekOur ideal candidate has:- 1+ years experience writing in Go (golang.org)- 2-3 years experience writing REST APIs- Experience working at a small startup- A passion for building something meaningful â€¦ Salary and compensation No salary data published by company so we estimated salary based on similar jobs related to Golang, Engineer and Backend jobs that are similar: $70,000 — $120000/year",
                    how_to_apply: "",
                    application_url: "",
                    application_email_address: "ossdevs-cm@gmail.com"
                })
                .expect(400)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    expect(JSON.stringify(res.body)).contain("company email is not a valid email address");
                });
        });

        it("post a job offer without company name", async function () {
            return request(apiHost)
                .post(`${endpoint}`)
                .set("Accept", "application/json")
                .send({
                    company_email: "ossdevs-cm@gmail.com",
                    company_name: "",
                    job_title: "Frontend Dev",
                    is_remote: true,
                    city: "",
                    country: "",
                    salary_range_min: 100000,
                    salary_range_max: 150000,
                    description: "OssCameroon is hiring a Remote Go Backend Engineer \n Remote - We are looking for a backend engineer who can work 30+ hr/weekOur ideal candidate has:- 1+ years experience writing in Go (golang.org)- 2-3 years experience writing REST APIs- Experience working at a small startup- A passion for building something meaningful â€¦ Salary and compensation No salary data published by company so we estimated salary based on similar jobs related to Golang, Engineer and Backend jobs that are similar: $70,000 — $120000/year",
                    how_to_apply: "",
                    application_url: "",
                    application_email_address: "ossdevs-cm@gmail.com"
                })
                .expect(400)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    expect(JSON.stringify(res.body)).contain("company name is mandatory");
                });
        });

        it("post a job offer without job title", async function () {
            return request(apiHost)
                .post(`${endpoint}`)
                .set("Accept", "application/json")
                .send({
                    company_email: "ossdevs-cm@gmail.com",
                    company_name: "OssCameroon",
                    job_title: "",
                    is_remote: true,
                    city: "",
                    country: "",
                    salary_range_min: 100000,
                    salary_range_max: 150000,
                    description: "OssCameroon is hiring a Remote Go Backend Engineer \n Remote - We are looking for a backend engineer who can work 30+ hr/weekOur ideal candidate has:- 1+ years experience writing in Go (golang.org)- 2-3 years experience writing REST APIs- Experience working at a small startup- A passion for building something meaningful â€¦ Salary and compensation No salary data published by company so we estimated salary based on similar jobs related to Golang, Engineer and Backend jobs that are similar: $70,000 — $120000/year",
                    how_to_apply: "",
                    application_url: "",
                    application_email_address: "ossdevs-cm@gmail.com"
                })
                .expect(400)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    expect(JSON.stringify(res.body)).contain("job title is mandatory");
                });
        });

        it("post a job offer without description", async function () {
            return request(apiHost)
                .post(`${endpoint}`)
                .set("Accept", "application/json")
                .send({
                    company_email: "ossdevs-cm@gmail.com",
                    company_name: "OssCameroon",
                    job_title: "Frontend Dev",
                    is_remote: true,
                    city: "",
                    country: "",
                    salary_range_min: 100000,
                    salary_range_max: 150000,
                    description: "",
                    how_to_apply: "",
                    application_url: "",
                    application_email_address: "ossdevs-cm@gmail.com"
                })
                .expect(400)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    expect(JSON.stringify(res.body)).contain("job description is mandatory");
                });
        });

        it("post a job offer without application url and email address and phone number", async function () {
            return request(apiHost)
                .post(`${endpoint}`)
                .set("Accept", "application/json")
                .send({
                    company_email: "ossdevs-cm@gmail.com",
                    company_name: "OssCameroon",
                    job_title: "Frontend Dev",
                    is_remote: true,
                    city: "",
                    country: "",
                    salary_range_min: 100000,
                    salary_range_max: 150000,
                    description: "OssCameroon is hiring a Remote Go Backend Engineer \n Remote - We are looking for a backend engineer who can work 30+ hr/weekOur ideal candidate has:- 1+ years experience writing in Go (golang.org)- 2-3 years experience writing REST APIs- Experience working at a small startup- A passion for building something meaningful â€¦ Salary and compensation No salary data published by company so we estimated salary based on similar jobs related to Golang, Engineer and Backend jobs that are similar: $70,000 — $120000/year",
                    how_to_apply: "",
                    application_url: "",
                    application_email_address: "",
                    application_phone_number: ""
                })
                .expect(400)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    expect(JSON.stringify(res.body)).contain("application url or email address or phone number is mandatory");
                });
        });

        it("post a job offer without a valid application email address", async function () {
            return request(apiHost)
                .post(`${endpoint}`)
                .set("Accept", "application/json")
                .send({
                    company_email: "ossdevs-cm@gmail.com",
                    company_name: "OssCameroon",
                    job_title: "Frontend Dev",
                    is_remote: true,
                    city: "",
                    country: "",
                    description: "OssCameroon is hiring a Remote Go Backend Engineer \n Remote - We are looking for a backend engineer who can work 30+ hr/weekOur ideal candidate has:- 1+ years experience writing in Go (golang.org)- 2-3 years experience writing REST APIs- Experience working at a small startup- A passion for building something meaningful â€¦ Salary and compensation No salary data published by company so we estimated salary based on similar jobs related to Golang, Engineer and Backend jobs that are similar: $70,000 — $120000/year",
                    how_to_apply: "",
                    application_url: "",
                    application_email_address: "ossdevs-cm@gmail"
                })
                .expect(400)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    expect(JSON.stringify(res.body)).contain("application email address is not a valid email address");
                });
        });

        it("post a job offer without a valid application phone number", async function () {
            return request(apiHost)
                .post(`${endpoint}`)
                .set("Accept", "application/json")
                .send({
                    company_email: "ossdevs-cm@gmail.com",
                    company_name: "OssCameroon",
                    job_title: "Frontend Dev",
                    is_remote: true,
                    city: "",
                    country: "",
                    salary_range_min: 100000,
                    salary_range_max: 150000,
                    description: "OssCameroon is hiring a Remote Go Backend Engineer \n Remote - We are looking for a backend engineer who can work 30+ hr/weekOur ideal candidate has:- 1+ years experience writing in Go (golang.org)- 2-3 years experience writing REST APIs- Experience working at a small startup- A passion for building something meaningful â€¦ Salary and compensation No salary data published by company so we estimated salary based on similar jobs related to Golang, Engineer and Backend jobs that are similar: $70,000 — $120000/year",
                    how_to_apply: "",
                    application_url: "",
                    application_email_address: "",
                    application_phone_number: "invalid phone number"
                })
                .expect(400)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    expect(JSON.stringify(res.body)).contain("application phone number is not a valid phone number");
                });
        });

        it("post a job offer without a location and no remote should fail", async function () {
            return request(apiHost)
                .post(`${endpoint}`)
                .set("Accept", "application/json")
                .send({
                    company_email: "ossdevs-cm@gmail.com",
                    company_name: "OssCameroon",
                    job_title: "Frontend Dev",
                    is_remote: false,
                    city: "",
                    country: "",
                    salary_range_min: 100000,
                    salary_range_max: 150000,
                    department: "Research and Development",
                    description: "OssCameroon is hiring a Remote Go Backend Engineer \n Remote - We are looking for a backend engineer who can work 30+ hr/weekOur ideal candidate has:- 1+ years experience writing in Go (golang.org)- 2-3 years experience writing REST APIs- Experience working at a small startup- A passion for building something meaningful â€¦ Salary and compensation No salary data published by company so we estimated salary based on similar jobs related to Golang, Engineer and Backend jobs that are similar: $70,000 — $120000/year",
                    benefits: "Health insurance, dental insurance, 401k",
                    how_to_apply: "Please submit your resume and cover letter.",
                    application_url: "",
                    application_email_address: "ossdevs-cm@gmail.com",
                    application_phone_number: "555-555-5555"
                })
                .expect(400)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    expect(JSON.stringify(res.body)).contain("city and country are mandatory when is not remote");
                });
        });

        it("post a valid job offer with image", async function () {
            const fileContent = fs.readFileSync("./oss.svg");
            const base64Content = fileContent.toString('base64');
            return request(apiHost)
                .post(`${endpoint}`)
                .set("Accept", "application/json")
                .send({
                    company_email: "ossdevs-cm@gmail.com",
                    company_name: "OssCameroon",
                    company_image: base64Content,
                    job_title: "Frontend Dev",
                    is_remote: true,
                    city: "Douala",
                    country: "Cameroon",
                    salary_range_min: 100000,
                    salary_range_max: 150000,
                    department: "Research and Development",
                    description: "OssCameroon is hiring a Remote Go Backend Engineer \n Remote - We are looking for a backend engineer who can work 30+ hr/weekOur ideal candidate has:- 1+ years experience writing in Go (golang.org)- 2-3 years experience writing REST APIs- Experience working at a small startup- A passion for building something meaningful â€¦ Salary and compensation No salary data published by company so we estimated salary based on similar jobs related to Golang, Engineer and Backend jobs that are similar: $70,000 — $120000/year",
                    benefits: "Health insurance, dental insurance, 401k",
                    how_to_apply: "Please submit your resume and cover letter.",
                    application_url: "",
                    application_email_address: "ossdevs-cm@gmail.com",
                    application_phone_number: "555-555-5555",
                    tags: "remote, golang, backend, engineer"
                })
                .expect(201)
                .expect("Content-Type", "application/json; charset=utf-8")
                .then((res) => {
                    expect(res.body.company_name).equal("OssCameroon");
                    expect(res.body.has_image).equal(true);
                });
        });


        it("post job offer with invalid image", async function () {
            return request(apiHost)
                .post(`${endpoint}`)
                .set("Accept", "application/json")
                .send({
                    company_email: "ossdevs-cm@gmail.com",
                    company_name: "OssCameroon",
                    company_image: "i'm not an image i'm just here to troll you",
                    job_title: "Frontend Dev",
                    is_remote: true,
                    city: "Douala",
                    country: "Cameroon",
                    salary_range_min: 100000,
                    salary_range_max: 150000,
                    department: "Research and Development",
                    description: "OssCameroon is hiring a Remote Go Backend Engineer \n Remote - We are looking for a backend engineer who can work 30+ hr/weekOur ideal candidate has:- 1+ years experience writing in Go (golang.org)- 2-3 years experience writing REST APIs- Experience working at a small startup- A passion for building something meaningful â€¦ Salary and compensation No salary data published by company so we estimated salary based on similar jobs related to Golang, Engineer and Backend jobs that are similar: $70,000 — $120000/year",
                    benefits: "Health insurance, dental insurance, 401k",
                    how_to_apply: "Please submit your resume and cover letter.",
                    application_url: "",
                    application_email_address: "ossdevs-cm@gmail.com",
                    application_phone_number: "555-555-5555",
                    tags: "remote, golang, backend, engineer"
                })
                .expect(400)
                .expect("Content-Type", "application/json; charset=utf-8")
        });
    });
});
