import {
  UsersApi,
  V1PostUserRequest,
} from 'services/api-ts-sdk/api';
import { consumers } from 'stream';

import { ApiAuthFixture } from './api.auth.fixture';

export class ApiUserFixture extends ApiAuthFixture {

  async createUser(req: V1PostUserRequest) {
    const userResp = await new UsersApi(
      { apiKey: await this.getBearerToken() },
      this.baseURL,
      fetch,
    )
      .postUser(req, {})
      .catch(async function (error) {
        const respBody = await consumers.json(error.body);
        throw new Error(`Create User Request failed. Request: ${JSON.stringify(req)} Response: ${JSON.stringify(respBody)}`);
      });
    console.log(`Successful create user response: ${userResp}`);
  }
}
