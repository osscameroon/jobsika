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
          expect(JSON.stringify(res.body)).equal(
            '["Douala","Yaoundé","Garoua","Bamenda","Maroua","Nkongsamba","Bafoussam","Ngaoundéré","Bertoua","Loum","Kumba","Edéa","Kumbo","Foumban","Mbouda","Dschang","Limbé","Ebolowa","Kousséri","Guider","Meiganga","Yagoua","Mbalmayo","Bafang","Tiko","Bafia","Wum","Kribi","Buea","Sangmélima","Foumbot","Bangangté","Batouri","Banyo","Nkambé","Bali","Mbanga","Mokolo","Melong","Manjo","Garoua-Boulaï","Mora","Kaélé","Tibati","Ndop","Akonolinga","Eséka","Mamfé","Obala","Muyuka","Nanga-Eboko","Abong-Mbang","Fundong","Nkoteng","Fontem","Mbandjock","Touboro","Ngaoundal","Yokadouma","Pitoa","Tombel","Kékem","Magba","Bélabo","Tonga","Maga","Koutaba","Blangoua","Guidiguis","Bogo","Batibo","Yabassi","Figuil","Makénéné","Gazawa","Tcholliré"]'
          );
        });
    });
  });
});
