"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * Opera Cloud Inventory Asynchronous API
 * APIs to cater for Inventory related asynchronous functionality in OPERA Cloud. This includes viewing inventory data along with its revenue and updating inventory&apos;s sell limits for date ranges. <p>This API follows an async pattern where</p><ul><li>You make an initial request, which returns a Location header</li><li>You poll HEAD on the Location header returned to obtain the status of the process</li><li>Once the process completes HEAD will return in the Location header the URL that must be called to obtain the results of the process</li><li>You call the URL to obtain the results of the process</li></ul><br /><br /> Compatible with OPERA Cloud release 22.4.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.4.0.0
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
exports.InventoryAsyncApi = void 0;
const runtime = __importStar(require("../runtime"));
const models_1 = require("../models");
/**
 *
 */
class InventoryAsyncApi extends runtime.BaseAPI {
    /**
     * This API returns inventory revenue&apos;s statistics for a hotel within dates given in the startRevenueInventoryStatisticsProcess API request. You can get the value of the summaryId from the Location header returned by the getRevenueInventoryStatisticsProcessStatus operation after the process is completed.<p><strong>OperationId:</strong>getRevenueInventoryStatistics</p>
     * Get results of a revenue inventory statistics process
     */
    async getRevenueInventoryStatisticsRaw(requestParameters, initOverrides) {
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
            path: `/externalSystems/{extSystemCode}/hotels/{hotelId}/revenueInventoryStatistics/{requestId}`.replace(`{${"requestId"}}`, encodeURIComponent(String(requestParameters.requestId))).replace(`{${"extSystemCode"}}`, encodeURIComponent(String(requestParameters.extSystemCode))).replace(`{${"hotelId"}}`, encodeURIComponent(String(requestParameters.hotelId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);
        return new runtime.JSONApiResponse(response, (jsonValue) => (0, models_1.RevenueInventoryStatisticsDetailsFromJSON)(jsonValue));
    }
    /**
     * This API returns inventory revenue&apos;s statistics for a hotel within dates given in the startRevenueInventoryStatisticsProcess API request. You can get the value of the summaryId from the Location header returned by the getRevenueInventoryStatisticsProcessStatus operation after the process is completed.<p><strong>OperationId:</strong>getRevenueInventoryStatistics</p>
     * Get results of a revenue inventory statistics process
     */
    async getRevenueInventoryStatistics(requestParameters, initOverrides) {
        const response = await this.getRevenueInventoryStatisticsRaw(requestParameters, initOverrides);
        return await response.value();
    }
    /**
     * Use this API to check whether the process to collate inventory revenue statistics is complete. You can get value of summaryId from the Location header returned by the startRevenueInventoryStatisticsProcess operation.<p><strong>OperationId:</strong>getRevenueInventoryStatisticsProcessStatus</p>
     * Check status of Inventory Revenue Statistic process
     */
    async getRevenueInventoryStatisticsProcessStatusRaw(requestParameters, initOverrides) {
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
            path: `/externalSystems/{extSystemCode}/hotels/{hotelId}/revenueInventoryStatistics/{requestId}`.replace(`{${"requestId"}}`, encodeURIComponent(String(requestParameters.requestId))).replace(`{${"extSystemCode"}}`, encodeURIComponent(String(requestParameters.extSystemCode))).replace(`{${"hotelId"}}`, encodeURIComponent(String(requestParameters.hotelId))),
            method: 'HEAD',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);
        return new runtime.VoidApiResponse(response);
    }
    /**
     * Use this API to check whether the process to collate inventory revenue statistics is complete. You can get value of summaryId from the Location header returned by the startRevenueInventoryStatisticsProcess operation.<p><strong>OperationId:</strong>getRevenueInventoryStatisticsProcessStatus</p>
     * Check status of Inventory Revenue Statistic process
     */
    async getRevenueInventoryStatisticsProcessStatus(requestParameters, initOverrides) {
        await this.getRevenueInventoryStatisticsProcessStatusRaw(requestParameters, initOverrides);
    }
    /**
     * This API returns the status log if any of sell limit data creation is failed. You can get the value of summaryId with the getSellLimitsProcessStatus API response (under header location). <p><strong>OperationId:</strong>getSellLimits</p>
     * Get status for sell limits for a property.
     */
    async getSellLimitsRaw(requestParameters, initOverrides) {
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
            path: `/externalSystems/{extSystemCode}/hotels/{hotelId}/sellLimits/{requestId}`.replace(`{${"requestId"}}`, encodeURIComponent(String(requestParameters.requestId))).replace(`{${"hotelId"}}`, encodeURIComponent(String(requestParameters.hotelId))).replace(`{${"extSystemCode"}}`, encodeURIComponent(String(requestParameters.extSystemCode))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);
        return new runtime.JSONApiResponse(response, (jsonValue) => (0, models_1.SellLimitByDateStatusFromJSON)(jsonValue));
    }
    /**
     * This API returns the status log if any of sell limit data creation is failed. You can get the value of summaryId with the getSellLimitsProcessStatus API response (under header location). <p><strong>OperationId:</strong>getSellLimits</p>
     * Get status for sell limits for a property.
     */
    async getSellLimits(requestParameters, initOverrides) {
        const response = await this.getSellLimitsRaw(requestParameters, initOverrides);
        return await response.value();
    }
    /**
     * Use this API to check whether the sell limit data creation is completed yet. You can get the value of the summaryId with the postSellLimitsProcess API response (under header location). <p><strong>OperationId:</strong>getSellLimitsProcessStatus</p>
     * Get the status for sell limits asynchronous process
     */
    async getSellLimitsProcessStatusRaw(requestParameters, initOverrides) {
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
            path: `/externalSystems/{extSystemCode}/hotels/{hotelId}/sellLimits/{requestId}`.replace(`{${"requestId"}}`, encodeURIComponent(String(requestParameters.requestId))).replace(`{${"extSystemCode"}}`, encodeURIComponent(String(requestParameters.extSystemCode))).replace(`{${"hotelId"}}`, encodeURIComponent(String(requestParameters.hotelId))),
            method: 'HEAD',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);
        return new runtime.VoidApiResponse(response);
    }
    /**
     * Use this API to check whether the sell limit data creation is completed yet. You can get the value of the summaryId with the postSellLimitsProcess API response (under header location). <p><strong>OperationId:</strong>getSellLimitsProcessStatus</p>
     * Get the status for sell limits asynchronous process
     */
    async getSellLimitsProcessStatus(requestParameters, initOverrides) {
        await this.getSellLimitsProcessStatusRaw(requestParameters, initOverrides);
    }
    /**
     * You can use this API to create sell limits in OPERA by date. <p><strong>OperationId:</strong>postSellLimitsProcess</p>
     * Create sell limit by date.
     */
    async postSellLimitsProcessRaw(requestParameters, initOverrides) {
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
            path: `/externalSystems/{extSystemCode}/hotels/{hotelId}/sellLimits`.replace(`{${"hotelId"}}`, encodeURIComponent(String(requestParameters.hotelId))).replace(`{${"extSystemCode"}}`, encodeURIComponent(String(requestParameters.extSystemCode))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters.sellLimitsByDate.map(models_1.SellLimitByDateTypeToJSON),
        }, initOverrides);
        return new runtime.VoidApiResponse(response);
    }
    /**
     * You can use this API to create sell limits in OPERA by date. <p><strong>OperationId:</strong>postSellLimitsProcess</p>
     * Create sell limit by date.
     */
    async postSellLimitsProcess(requestParameters, initOverrides) {
        await this.postSellLimitsProcessRaw(requestParameters, initOverrides);
    }
    /**
     * Use this API to start the process to get a hotel inventory&apos;s revenue statistics for a given date range. Maximum limit of the date range is 94 days with a single request, but the recommendation is to use a date range proportionate to the size of the hotel. Returns a header parameter Location that can be used in the getRevenueInventoryStatisticsProcessStatus operation.  Use event driven APIs (see https://docs.oracle.com/cd/F29336_01/doc.201/f27480/c_business_events.htm#OHIPU-BusinessEvents-F0AC1F1C) to get real time inventory updates.<p><strong>OperationId:</strong>startRevenueInventoryStatisticsProcess</p>
     * Start process to get hotel inventory revenue statistics
     */
    async startRevenueInventoryStatisticsProcessRaw(requestParameters, initOverrides) {
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
            path: `/externalSystems/{extSystemCode}/hotels/{hotelId}/revenueInventoryStatistics`.replace(`{${"extSystemCode"}}`, encodeURIComponent(String(requestParameters.extSystemCode))).replace(`{${"hotelId"}}`, encodeURIComponent(String(requestParameters.hotelId))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: (0, models_1.RevenueInventoryStatisticsSearchTypeToJSON)(requestParameters.criteria),
        }, initOverrides);
        return new runtime.VoidApiResponse(response);
    }
    /**
     * Use this API to start the process to get a hotel inventory&apos;s revenue statistics for a given date range. Maximum limit of the date range is 94 days with a single request, but the recommendation is to use a date range proportionate to the size of the hotel. Returns a header parameter Location that can be used in the getRevenueInventoryStatisticsProcessStatus operation.  Use event driven APIs (see https://docs.oracle.com/cd/F29336_01/doc.201/f27480/c_business_events.htm#OHIPU-BusinessEvents-F0AC1F1C) to get real time inventory updates.<p><strong>OperationId:</strong>startRevenueInventoryStatisticsProcess</p>
     * Start process to get hotel inventory revenue statistics
     */
    async startRevenueInventoryStatisticsProcess(requestParameters, initOverrides) {
        await this.startRevenueInventoryStatisticsProcessRaw(requestParameters, initOverrides);
    }
}
exports.InventoryAsyncApi = InventoryAsyncApi;
