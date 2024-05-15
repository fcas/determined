import { APIRequest, APIRequestContext, Browser, BrowserContext, Page } from '@playwright/test';

export class ApiAuthFixture {
  apiContext: APIRequestContext | undefined;
  readonly request: APIRequest;
  readonly browser: Browser;
  readonly baseURL: string;
  _page: Page | undefined;
  get page(): Page {
    if (this._page === undefined) {
      throw new Error('Accessing page object before initialization in authentication');
    }
    return this._page;
  }
  readonly #STATE_FILE = 'state.json';
  readonly #USERNAME: string;
  readonly #PASSWORD: string;
  context: BrowserContext | undefined;
  constructor(
    request: APIRequest,
    browser: Browser,
    baseURL: string | undefined,
    existingPage: Page | undefined = undefined,
  ) {
    if (process.env.PW_USER_NAME === undefined) {
      throw new Error('username must be defined');
    }
    if (process.env.PW_PASSWORD === undefined) {
      throw new Error('password must be defined');
    }
    if (baseURL === undefined) {
      throw new Error('baseURL must be defined in playwright config to use API requests.');
    }
    this.#USERNAME = process.env.PW_USER_NAME;
    this.#PASSWORD = process.env.PW_PASSWORD;
    this.request = request;
    this.browser = browser;
    this.baseURL = baseURL;
    this._page = existingPage;
  }

  protected async getBearerToken(): Promise<string> {
    const cookies = (await this.apiContext?.storageState())?.cookies ?? [];
    const authToken = cookies.find((cookie) => {
      return cookie.name === 'auth';
    })?.value;
    if (authToken === undefined) {
      throw new Error(
        'Attempted to retrieve the auth token from the PW apiContext, but it does not exist. Have you called apiAuth.login() yet?',
      );
    }
    return `Bearer ${authToken}`;
  }

  async login(): Promise<void> {
    this.apiContext = await this.request.newContext();
    await this.apiContext.post('/api/v1/auth/login', {
      data: {
        isHashed: false,
        password: this.#PASSWORD,
        username: this.#USERNAME,
      },
    });
    // Save cookie state into the file.
    const state = await this.apiContext.storageState({ path: this.#STATE_FILE });
    if (this._page !== undefined) {
      // add cookies to current page's existing context
      this.context = this._page.context();
      await this.context.addCookies(state.cookies);
    } else {
      // Create a new context for the browser with the saved token.
      this.context = await this.browser.newContext({ storageState: this.#STATE_FILE });
      this._page = await this.context.newPage();
    }
  }

  async logout(): Promise<void> {
    await this.apiContext?.dispose();
    await this.context?.close();
  }
}
