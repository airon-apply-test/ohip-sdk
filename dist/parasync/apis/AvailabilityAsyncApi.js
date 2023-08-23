"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate Async API
 * APIs to cater for Price and Rate Availability Asynchronous functionality in OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.AvailabilityAsyncApi = void 0;
const runtime = __importStar(require("../runtime"));
const models_1 = require("../models");
/**
 *
 */
class AvailabilityAsyncApi extends runtime.BaseAPI {
    /**
     * Use this to check whether the restrictions sent have been processed. <p><strong>OperationId:</strong>getRestrictions</p>
     * Get status for created restrictions.
     */
    async getRestrictionsRaw(requestParameters, initOverrides) {
        const queryParameters = {};
        const headerParameters = {};
        if (requestParameters.authorization !== undefined && requestParameters.authorization !== null) {
            headerParameters['authorization'] = String(requestParameters.authorization);
        }
        if (requestParameters.xAppKey !== undefined && requestParameters.xAppKey !== null) {
            headerParameters['x-app-key'] = String(requestParameters.xAppKey);
        }
        if (requestParameters.xHotelid !== undefined && requestParameters.xHotelid !== null) {
            headerParameters['x-hotelid'] = String(requestParameters.xHotelid);
        }
        if (requestParameters.acceptLanguage !== undefined && requestParameters.acceptLanguage !== null) {
            headerParameters['Accept-Language'] = String(requestParameters.acceptLanguage);
        }
        const response = await this.request({
            path: `/externalSystems/{extSystemCode}/hotels/{hotelId}/restrictions/{requestId}`.replace(`{${"hotelId"}}`, encodeURIComponent(String(requestParameters.hotelId))).replace(`{${"extSystemCode"}}`, encodeURIComponent(String(requestParameters.extSystemCode))).replace(`{${"requestId"}}`, encodeURIComponent(String(requestParameters.requestId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);
        return new runtime.JSONApiResponse(response, (jsonValue) => (0, models_1.RestrictionsStatusFromJSON)(jsonValue));
    }
    /**
     * Use this to check whether the restrictions sent have been processed. <p><strong>OperationId:</strong>getRestrictions</p>
     * Get status for created restrictions.
     */
    async getRestrictions(requestParameters, initOverrides) {
        const response = await this.getRestrictionsRaw(requestParameters, initOverrides);
        return await response.value();
    }
    /**
     * Use this to check whether the restrictions sent have been processed. <p><strong>OperationId:</strong>getRestrictionsProcessStatus</p>
     * Check status of Restrictions
     */
    async getRestrictionsProcessStatusRaw(requestParameters, initOverrides) {
        const queryParameters = {};
        const headerParameters = {};
        if (requestParameters.authorization !== undefined && requestParameters.authorization !== null) {
            headerParameters['authorization'] = String(requestParameters.authorization);
        }
        if (requestParameters.xAppKey !== undefined && requestParameters.xAppKey !== null) {
            headerParameters['x-app-key'] = String(requestParameters.xAppKey);
        }
        if (requestParameters.xHotelid !== undefined && requestParameters.xHotelid !== null) {
            headerParameters['x-hotelid'] = String(requestParameters.xHotelid);
        }
        if (requestParameters.acceptLanguage !== undefined && requestParameters.acceptLanguage !== null) {
            headerParameters['Accept-Language'] = String(requestParameters.acceptLanguage);
        }
        const response = await this.request({
            path: `/externalSystems/{extSystemCode}/hotels/{hotelId}/restrictions/{requestId}`.replace(`{${"requestId"}}`, encodeURIComponent(String(requestParameters.requestId))).replace(`{${"extSystemCode"}}`, encodeURIComponent(String(requestParameters.extSystemCode))).replace(`{${"hotelId"}}`, encodeURIComponent(String(requestParameters.hotelId))),
            method: 'HEAD',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);
        return new runtime.VoidApiResponse(response);
    }
    /**
     * Use this to check whether the restrictions sent have been processed. <p><strong>OperationId:</strong>getRestrictionsProcessStatus</p>
     * Check status of Restrictions
     */
    async getRestrictionsProcessStatus(requestParameters, initOverrides) {
        await this.getRestrictionsProcessStatusRaw(requestParameters, initOverrides);
    }
    /**
     * A user can send various restrictions to OPERA by specifying restriction details in the request. <p><strong>OperationId:</strong>postRestrictionsProcess</p>
     * Create restrictions in OPERA Cloud.
     */
    async postRestrictionsProcessRaw(requestParameters, initOverrides) {
        const queryParameters = {};
        const headerParameters = {};
        headerParameters['Content-Type'] = 'application/json;charset=UTF-8';
        if (requestParameters.authorization !== undefined && requestParameters.authorization !== null) {
            headerParameters['authorization'] = String(requestParameters.authorization);
        }
        if (requestParameters.xAppKey !== undefined && requestParameters.xAppKey !== null) {
            headerParameters['x-app-key'] = String(requestParameters.xAppKey);
        }
        if (requestParameters.xHotelid !== undefined && requestParameters.xHotelid !== null) {
            headerParameters['x-hotelid'] = String(requestParameters.xHotelid);
        }
        if (requestParameters.acceptLanguage !== undefined && requestParameters.acceptLanguage !== null) {
            headerParameters['Accept-Language'] = String(requestParameters.acceptLanguage);
        }
        const response = await this.request({
            path: `/externalSystems/{extSystemCode}/hotels/{hotelId}/restrictions`.replace(`{${"hotelId"}}`, encodeURIComponent(String(requestParameters.hotelId))).replace(`{${"extSystemCode"}}`, encodeURIComponent(String(requestParameters.extSystemCode))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: (0, models_1.RestrictionsToJSON)(requestParameters.restrictions),
        }, initOverrides);
        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(models_1.InstanceLinkFromJSON));
    }
    /**
     * A user can send various restrictions to OPERA by specifying restriction details in the request. <p><strong>OperationId:</strong>postRestrictionsProcess</p>
     * Create restrictions in OPERA Cloud.
     */
    async postRestrictionsProcess(requestParameters, initOverrides) {
        const response = await this.postRestrictionsProcessRaw(requestParameters, initOverrides);
        return await response.value();
    }
}
exports.AvailabilityAsyncApi = AvailabilityAsyncApi;
