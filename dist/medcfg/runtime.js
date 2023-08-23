"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Content Service
 * Opera Content Service offers capability to manage large content such as images and files.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.TextApiResponse = exports.BlobApiResponse = exports.VoidApiResponse = exports.JSONApiResponse = exports.canConsumeForm = exports.mapValues = exports.querystring = exports.exists = exports.COLLECTION_FORMATS = exports.RequiredError = exports.FetchError = exports.ResponseError = exports.BaseAPI = exports.DefaultConfig = exports.Configuration = exports.BASE_PATH = void 0;
const index_1 = require("../oauth/index");
exports.BASE_PATH = "/med/config/v1".replace(/\/+$/, "");
class Configuration {
    constructor(configuration = {}) {
        this.configuration = configuration;
    }
    set config(configuration) {
        this.configuration = configuration;
    }
    get basePath() {
        return this.configuration.basePath != null ? this.configuration.basePath : exports.BASE_PATH;
    }
    get hotelId() {
        return this.configuration.hotelId;
    }
    get ohipCredentials() {
        return this.configuration.ohipCredentials;
    }
    get fetchApi() {
        return this.configuration.fetchApi;
    }
    get middleware() {
        return this.configuration.middleware || [];
    }
    get queryParamsStringify() {
        return this.configuration.queryParamsStringify || querystring;
    }
    get username() {
        return this.configuration.username;
    }
    get password() {
        return this.configuration.password;
    }
    get apiKey() {
        const apiKey = this.configuration.apiKey;
        if (apiKey) {
            return typeof apiKey === 'function' ? apiKey : () => apiKey;
        }
        return undefined;
    }
    get accessToken() {
        const accessToken = this.configuration.accessToken;
        if (accessToken) {
            return typeof accessToken === 'function' ? accessToken : async () => accessToken;
        }
        return undefined;
    }
    get headers() {
        return this.configuration.headers;
    }
    get credentials() {
        return this.configuration.credentials;
    }
}
exports.Configuration = Configuration;
exports.DefaultConfig = new Configuration();
/**
 * This is the base class for all generated API classes.
 */
class BaseAPI {
    constructor(configuration = exports.DefaultConfig) {
        this.configuration = configuration;
        this.backoffTimeout = 0;
        this.authTries = 0;
        this.maxBackOffTimeoutToWait = 1000;
        this.baseBackOffTimeout = 1;
        this.maxBackOffTimeout = 60 * 60 * 1000;
        this.fetchApi = async (url, init) => {
            let fetchParams = { url, init };
            for (const middleware of this.middleware) {
                if (middleware.pre) {
                    fetchParams = await middleware.pre(Object.assign({ fetch: this.fetchApi }, fetchParams)) || fetchParams;
                }
            }
            let response = undefined;
            try {
                response = await (this.configuration.fetchApi || fetch)(fetchParams.url, fetchParams.init);
            }
            catch (e) {
                for (const middleware of this.middleware) {
                    if (middleware.onError) {
                        response = await middleware.onError({
                            fetch: this.fetchApi,
                            url: fetchParams.url,
                            init: fetchParams.init,
                            error: e,
                            response: response ? response.clone() : undefined,
                        }) || response;
                    }
                }
                if (response === undefined) {
                    if (e instanceof Error) {
                        throw new FetchError(e, 'The request failed and the interceptors did not return an alternative response');
                    }
                    else {
                        throw e;
                    }
                }
            }
            for (const middleware of this.middleware) {
                if (middleware.post) {
                    response = await middleware.post({
                        fetch: this.fetchApi,
                        url: fetchParams.url,
                        init: fetchParams.init,
                        response: response.clone(),
                    }) || response;
                }
            }
            return response;
        };
        this.getAuthApi = () => {
            const appKey = this.configuration.apiKey('');
            const authConfig = new index_1.Configuration({
                apiKey: appKey,
                basePath: this.configuration.basePath,
            });
            return new index_1.AuthenticationApi(authConfig);
        };
        // ADDED
        this.handleClientRequest = async (context) => {
            if (!this.token) {
                await this.requestNewAuthToken();
            }
            else if (this.isAuthTokenExpired()) {
                await this.renewAuthToken();
            }
            if (!this.token)
                throw new Error('Failed getting OHIP authentication token.');
            context.init.headers['authorization'] = `Bearer ${this.token}`;
            context.init.headers['x-app-key'] = this.configuration.apiKey('');
            context.init.headers['x-hotel-id'] = this.configuration.hotelId;
            return context;
        };
        this.handleClientRequestError = async (error) => {
            if (error && ![401, 403].includes(error === null || error === void 0 ? void 0 : error.response.status))
                return Promise.reject(error);
            await this.requestNewAuthToken();
            if (!this.token)
                return Promise.reject(error);
            error.init.headers['authorization'] = `Bearer ${this.token}`;
            return error.fetch(error.url, error.init);
        };
        let authenticationMiddleware = [];
        if (!exports.BASE_PATH.includes('oauth')) {
            authenticationMiddleware = [
                {
                    pre: this.handleClientRequest,
                    onError: this.handleClientRequestError,
                    post: async (context) => {
                        return context.response;
                    },
                },
            ];
        }
        this.middleware = [...authenticationMiddleware, ...configuration.middleware];
    }
    withMiddleware(...middlewares) {
        const next = this.clone();
        next.middleware = next.middleware.concat(...middlewares);
        return next;
    }
    withPreMiddleware(...preMiddlewares) {
        const middlewares = preMiddlewares.map((pre) => ({ pre }));
        return this.withMiddleware(...middlewares);
    }
    withPostMiddleware(...postMiddlewares) {
        const middlewares = postMiddlewares.map((post) => ({ post }));
        return this.withMiddleware(...middlewares);
    }
    /**
     * Check if the given MIME is a JSON MIME.
     * JSON MIME examples:
     *   application/json
     *   application/json; charset=UTF8
     *   APPLICATION/JSON
     *   application/vnd.company+json
     * @param mime - MIME (Multipurpose Internet Mail Extensions)
     * @return True if the given MIME is JSON, false otherwise.
     */
    isJsonMime(mime) {
        if (!mime) {
            return false;
        }
        return BaseAPI.jsonRegex.test(mime);
    }
    async request(context, initOverrides) {
        const { url, init } = await this.createFetchParams(context, initOverrides);
        const response = await this.fetchApi(url, init);
        if (response && (response.status >= 200 && response.status < 300)) {
            return response;
        }
        throw new ResponseError(response, 'Response returned an error code');
    }
    async createFetchParams(context, initOverrides) {
        let url = this.configuration.basePath + exports.BASE_PATH + context.path;
        if (context.query !== undefined && Object.keys(context.query).length !== 0) {
            // only add the querystring to the URL if there are query parameters.
            // this is done to avoid urls ending with a "?" character which buggy webservers
            // do not handle correctly sometimes.
            url += '?' + this.configuration.queryParamsStringify(context.query);
        }
        const headers = Object.assign({}, this.configuration.headers, context.headers);
        Object.keys(headers).forEach(key => headers[key] === undefined ? delete headers[key] : {});
        const initOverrideFn = typeof initOverrides === "function"
            ? initOverrides
            : async () => initOverrides;
        const initParams = {
            method: context.method,
            headers,
            body: context.body,
            credentials: this.configuration.credentials,
        };
        const overriddenInit = Object.assign(Object.assign({}, initParams), (await initOverrideFn({
            init: initParams,
            context,
        })));
        const init = Object.assign(Object.assign({}, overriddenInit), { body: isFormData(overriddenInit.body) ||
                overriddenInit.body instanceof URLSearchParams ||
                isBlob(overriddenInit.body)
                ? overriddenInit.body
                : JSON.stringify(overriddenInit.body) });
        return { url, init };
    }
    async requestNewAuthToken() {
        // a request for a new token is already in progress
        // just wait for it to finish instead of trying to request a new one
        if (this.requestingNewAuthToken) {
            await this.requestingNewAuthToken;
            return;
        }
        if (this.shouldBackOff()) {
            const backOffMsRemaining = this.backoffTimeout - new Date().getTime();
            if (backOffMsRemaining > this.maxBackOffTimeoutToWait) {
                console.log(`Backing off for ${backOffMsRemaining}ms: Max backoff timeout to wait reached...`);
            }
            else {
                console.log(`Backing off for ${backOffMsRemaining}ms: Waiting for new token...`);
                await delay(backOffMsRemaining);
                await this.requestNewAuthToken();
            }
            return;
        }
        this.requestingNewAuthToken = new Promise(async (resolve) => {
            this.authTries++;
            this.incrementActiveCrendentialIndex();
            this.clearTokens();
            try {
                console.log(`Requesting new OHIP session token using credentials[${this.activeCredentialIndex}]`);
                const authApi = this.getAuthApi();
                const data = await authApi.getToken({
                    xAppKey: this.configuration.apiKey(''),
                    grantType: 'password',
                    username: this.configuration.ohipCredentials[this.activeCredentialIndex]
                        .username,
                    password: this.configuration.ohipCredentials[this.activeCredentialIndex]
                        .password,
                });
                this.setTokenHeaders(data);
                this.authTries = 0;
            }
            catch (error) {
                console.error('Requesting new OHIP session token failed', error);
                this.setBackoffTimeout();
            }
            resolve();
            this.requestingNewAuthToken = undefined;
        });
        return this.requestingNewAuthToken;
    }
    async renewAuthToken() {
        try {
            const authApi = this.getAuthApi();
            const data = await authApi.getToken({
                // @ts-ignore
                grantType: 'refresh_token',
                refresh_token: this.refreshToken,
                xAppKey: this.configuration.apiKey(''),
            });
            this.setTokenHeaders(data);
        }
        catch (error) {
            console.error('Refreshing OHIP session token failed', error);
            await this.requestNewAuthToken();
        }
    }
    setTokenHeaders(response) {
        var _a;
        // @ts-ignore
        this.refreshToken = response.refresh_token;
        this.token = response.accessToken;
        // response.expires_in is in number of seconds. Multiply it by 90% of 1000ms
        this.tokenExpiration = Date.now() + ((_a = response.expiresIn) !== null && _a !== void 0 ? _a : 0) * 900;
    }
    clearTokens() {
        this.refreshToken = undefined;
        this.token = undefined;
        this.tokenExpiration = undefined;
    }
    isAuthTokenExpired() {
        return Date.now() > this.tokenExpiration;
    }
    shouldBackOff() {
        if (this.authTries === 0)
            return false;
        return new Date().getTime() < this.backoffTimeout;
    }
    setBackoffTimeout() {
        if (this.activeCredentialIndex !==
            this.configuration.ohipCredentials.length - 1)
            return;
        const delay = Math.min(this.baseBackOffTimeout *
            Math.pow(10, Math.floor(this.authTries / this.configuration.ohipCredentials.length)), this.maxBackOffTimeout);
        if (delay === this.maxBackOffTimeout) {
            // if this happens we probably have the wrong password set in the config
            console.error('Max backoff timeout reached. Check OHIP credentials.');
        }
        this.backoffTimeout = new Date().getTime() + delay;
    }
    async incrementActiveCrendentialIndex() {
        if (this.activeCredentialIndex === undefined ||
            this.activeCredentialIndex + 1 >=
                this.configuration.ohipCredentials.length) {
            this.activeCredentialIndex = 0;
        }
        else {
            this.activeCredentialIndex += 1;
        }
    }
    /**
     * Create a shallow clone of `this` by constructing a new instance
     * and then shallow cloning data members.
     */
    clone() {
        const constructor = this.constructor;
        const next = new constructor(this.configuration);
        next.middleware = this.middleware.slice();
        return next;
    }
}
exports.BaseAPI = BaseAPI;
BaseAPI.jsonRegex = new RegExp('^(:?application\/json|[^;/ \t]+\/[^;/ \t]+[+]json)[ \t]*(:?;.*)?$', 'i');
;
function isBlob(value) {
    return typeof Blob !== 'undefined' && value instanceof Blob;
}
function isFormData(value) {
    return typeof FormData !== "undefined" && value instanceof FormData;
}
class ResponseError extends Error {
    constructor(response, msg) {
        super(msg);
        this.response = response;
        this.name = "ResponseError";
    }
}
exports.ResponseError = ResponseError;
class FetchError extends Error {
    constructor(cause, msg) {
        super(msg);
        this.cause = cause;
        this.name = "FetchError";
    }
}
exports.FetchError = FetchError;
class RequiredError extends Error {
    constructor(field, msg) {
        super(msg);
        this.field = field;
        this.name = "RequiredError";
    }
}
exports.RequiredError = RequiredError;
exports.COLLECTION_FORMATS = {
    csv: ",",
    ssv: " ",
    tsv: "\t",
    pipes: "|",
};
function exists(json, key) {
    const value = json[key];
    return value !== null && value !== undefined;
}
exports.exists = exists;
function querystring(params, prefix = '') {
    return Object.keys(params)
        .map(key => querystringSingleKey(key, params[key], prefix))
        .filter(part => part.length > 0)
        .join('&');
}
exports.querystring = querystring;
function querystringSingleKey(key, value, keyPrefix = '') {
    const fullKey = keyPrefix + (keyPrefix.length ? `[${key}]` : key);
    if (value instanceof Array) {
        const multiValue = value.map(singleValue => encodeURIComponent(String(singleValue)))
            .join(`&${encodeURIComponent(fullKey)}=`);
        return `${encodeURIComponent(fullKey)}=${multiValue}`;
    }
    if (value instanceof Set) {
        const valueAsArray = Array.from(value);
        return querystringSingleKey(key, valueAsArray, keyPrefix);
    }
    if (value instanceof Date) {
        return `${encodeURIComponent(fullKey)}=${encodeURIComponent(value.toISOString())}`;
    }
    if (value instanceof Object) {
        return querystring(value, fullKey);
    }
    return `${encodeURIComponent(fullKey)}=${encodeURIComponent(String(value))}`;
}
function mapValues(data, fn) {
    return Object.keys(data).reduce((acc, key) => (Object.assign(Object.assign({}, acc), { [key]: fn(data[key]) })), {});
}
exports.mapValues = mapValues;
function canConsumeForm(consumes) {
    for (const consume of consumes) {
        if ('multipart/form-data' === consume.contentType) {
            return true;
        }
    }
    return false;
}
exports.canConsumeForm = canConsumeForm;
class JSONApiResponse {
    constructor(raw, transformer = (jsonValue) => jsonValue) {
        this.raw = raw;
        this.transformer = transformer;
    }
    async value() {
        return this.transformer(await this.raw.json());
    }
}
exports.JSONApiResponse = JSONApiResponse;
class VoidApiResponse {
    constructor(raw) {
        this.raw = raw;
    }
    async value() {
        return undefined;
    }
}
exports.VoidApiResponse = VoidApiResponse;
class BlobApiResponse {
    constructor(raw) {
        this.raw = raw;
    }
    async value() {
        return await this.raw.blob();
    }
    ;
}
exports.BlobApiResponse = BlobApiResponse;
class TextApiResponse {
    constructor(raw) {
        this.raw = raw;
    }
    async value() {
        return await this.raw.text();
    }
    ;
}
exports.TextApiResponse = TextApiResponse;
async function delay(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
}
