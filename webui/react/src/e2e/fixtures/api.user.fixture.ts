import { safeName } from 'e2e/utils/naming';
import _ from 'lodash';
import {
  UsersApi,
  V1PostUserRequest,
} from 'services/api-ts-sdk/api';
import streamConsumers from 'stream/consumers';

import { ApiAuthFixture } from './api.auth.fixture';

export class ApiUserFixture {
  readonly apiAuth: ApiAuthFixture
  constructor(apiAuth: ApiAuthFixture) {
    this.apiAuth=apiAuth;
  }

  newRandom(): V1PostUserRequest {
    return {
      isHashed: false,
      password: 'TestPassword1',
      user: {
        active: true,
        admin: true,
        username: safeName('test-user'),
      },
    };
  }

  /**
   * Creates a user with the given parameters via the API.
   * @param {V1PostUserRequest} req the user request with the config for the new user.
   * See apiUser.newRandom() for the default config.
   * @returns {Promise<V1PostUserRequest>} Representation of the created user. The request is returned since the
   * password is not stored on the V1User object and it is not returned in the response. However the Request is a
   * strict superset of the Response, so no info is lost.
   */
  async createUser(req: V1PostUserRequest): Promise<V1PostUserRequest> {
    const userResp = await new UsersApi(
      { apiKey: await this.apiAuth.getBearerToken() },
      'http://localhost:3001', //this.baseURL, - DNJ TODO extra slash handling is unfortunate
      fetch,
    )
      .postUser(req, {})
      .catch(async function (error) {
        const respBody = await streamConsumers.text(error.body);
        throw new Error(
          `Create User Request failed. Status: ${error.status} Request: ${JSON.stringify(
            req,
          )} Response: ${respBody}`,
        );
      });
    return _.merge(req, userResp);
  }
}
