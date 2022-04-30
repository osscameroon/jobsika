import { expect } from "chai";
import request from "supertest";
import dotenv from "dotenv";

dotenv.config();
const apiHost = process.env.API_HOST;
const endpoint = "cities";

describe(`${endpoint}`, function () {
  describe("GET", function () {
    it("return a list of cities", async function () {
      return request(apiHost)
        .get(endpoint)
        .send()
        .expect(200)
        .expect("Content-Type", "application/json; charset=utf-8")
        .then((res) => {
            expect(JSON.stringify(res.body)).equal('["Abong-Mbang","Akonolinga","Bafang","Bafia","Bafoussam","Bali","Bamenda","Bangangté","Banyo","Batibo","Batouri","Bertoua","Blangoua","Bogo","Buea","Bélabo","Douala","Dschang","Ebolowa","Edéa","Eséka","Figuil","Fontem","Foumban","Foumbot","Fundong","Garoua","Garoua-Boulaï","Gazawa","Guider","Guidiguis","Kaélé","Kousséri","Koutaba","Kribi","Kumba","Kumbo","Kékem","Limbé","Loum","Maga","Magba","Makénéné","Mamfé","Manjo","Maroua","Mbalmayo","Mbandjock","Mbanga","Mbouda","Meiganga","Melong","Mokolo","Mora","Muyuka","Nanga-Eboko","Ndop","Ngaoundal","Ngaoundéré","Nkambé","Nkongsamba","Nkoteng","Obala","Pitoa","Sangmélima","Tcholliré","Tibati","Tiko","Tombel","Tonga","Touboro","Wum","Yabassi","Yagoua","Yaoundé","Yokadouma"]')
        });
    });
  });
});
