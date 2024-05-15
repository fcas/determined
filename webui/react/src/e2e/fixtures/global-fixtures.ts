import { test as base, Page } from '@playwright/test';

import { ApiAuthFixture } from './api.auth.fixture';
import { ApiUserFixture } from './api.user.fixture';
import { AuthFixture } from './auth.fixture';
import { DevFixture } from './dev.fixture';
import { UserFixture } from './user.fixture';

type CustomFixtures = {
  dev: DevFixture;
  auth: AuthFixture;
  apiAuth: ApiAuthFixture;
  backgroundApiAuth: ApiAuthFixture;
  user: UserFixture;
  apiUser: ApiUserFixture;
  backgroundApiUser: ApiUserFixture;
  authedPage: Page;
};

// https://playwright.dev/docs/test-fixtures
export const test = base.extend<CustomFixtures>({
  // get the auth but allow yourself to log in through the api manually.
  apiAuth: async ({ playwright, browser, dev, baseURL }, use) => {
    await dev.setServerAddress();
    const apiAuth = new ApiAuthFixture(playwright.request, browser, baseURL, dev.page);
    await apiAuth.login();
    await use(apiAuth);
  },

  apiUser: async ({ apiAuth }, use) => {
    const apiUser = new ApiUserFixture(apiAuth);
    await use(apiUser);
  },

  auth: async ({ page }, use) => {
    const auth = new AuthFixture(page);
    await use(auth);
  },

  // get the existing page but with auth cookie already logged in
  authedPage: async ({ apiAuth }, use) => {
    await use(apiAuth.page);
  },

  // does not require the pre-existing Playwright page and so can be called in beforeAll
  backgroundApiAuth: async ({ playwright, browser, baseURL }, use) => {
    const apiAuth = new ApiAuthFixture(playwright.request, browser, baseURL);
    const dev = new DevFixture(apiAuth.page);
    await dev.setServerAddress();
    await apiAuth.login();
    await use(apiAuth);
  },

  backgroundApiUser: async ({ backgroundApiAuth }, use) => {
    const backgroundApiUser = new ApiUserFixture(backgroundApiAuth);
    await use(backgroundApiUser);
  },

  dev: async ({ page }, use) => {
    const dev = new DevFixture(page);
    await use(dev);
  },

  user: async ({ page }, use) => {
    const user = new UserFixture(page);
    await use(user);
  },
});
