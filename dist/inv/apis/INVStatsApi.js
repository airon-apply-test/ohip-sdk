"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Inventory API
 * APIs to cater for Inventory functionality in OPERA Cloud. This includes sell limits for date ranges, viewing and updating the property&apos;s inventory, as well as item inventory (such as rollaways, microwaves etc.).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
exports.GetInventoryStatisticsParameterNameEnum = exports.GetInventoryStatisticsReportCodeEnum = exports.GetBlockInventoryStatisticsFetchInstructionsEnum = exports.GetBlockInventoryStatisticsDetailSortByCodeEnum = exports.GetBlockInventoryStatisticsSummarySortByCodeEnum = exports.INVStatsApi = void 0;
const runtime = __importStar(require("../runtime"));
const models_1 = require("../models");
/**
 *
 */
class INVStatsApi extends runtime.BaseAPI {
    /**
     * Clearing of cache in inventory statistics service <p><strong>OperationId:</strong>deleteinvStatsService</p>
     * Clear Cache
     */
    async deleteinvStatsServiceRaw(requestParameters, initOverrides) {
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
        if (requestParameters.xExternalsystem !== undefined && requestParameters.xExternalsystem !== null) {
            headerParameters['x-externalsystem'] = String(requestParameters.xExternalsystem);
        }
        if (requestParameters.acceptLanguage !== undefined && requestParameters.acceptLanguage !== null) {
            headerParameters['Accept-Language'] = String(requestParameters.acceptLanguage);
        }
        const response = await this.request({
            path: `/services/invStatsService/cache`,
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);
        return new runtime.JSONApiResponse(response, (jsonValue) => (0, models_1.StatusFromJSON)(jsonValue));
    }
    /**
     * Clearing of cache in inventory statistics service <p><strong>OperationId:</strong>deleteinvStatsService</p>
     * Clear Cache
     */
    async deleteinvStatsService(requestParameters, initOverrides) {
        const response = await this.deleteinvStatsServiceRaw(requestParameters, initOverrides);
        return await response.value();
    }
    /**
     * Use this API to get block statistics on a per-day basis across a specified date range. The result includes a summary section of rooms booked by status by date, details on deduct and non-deduct rooms booked per day, and Sales Allowance and House Availability figures per day. The details section of the results include block details such as name, block, code, start and end date, owner, and more. It also shows a breakdown of blocked picked-up and available rooms per day per block. The result set is used to feed the GRC (Group Rooms Control) page in the OPERA UI. <p><strong>OperationId:</strong>getBlockInventoryStatistics</p>
     * Get block inventory statistics
     */
    async getBlockInventoryStatisticsRaw(requestParameters, initOverrides) {
        const queryParameters = {};
        if (requestParameters.includeBlocksWithNoRoom !== undefined) {
            queryParameters['includeBlocksWithNoRoom'] = requestParameters.includeBlocksWithNoRoom;
        }
        if (requestParameters.includeOverbooking !== undefined) {
            queryParameters['includeOverbooking'] = requestParameters.includeOverbooking;
        }
        if (requestParameters.includeOpportunities !== undefined) {
            queryParameters['includeOpportunities'] = requestParameters.includeOpportunities;
        }
        if (requestParameters.includeTentativeInventory !== undefined) {
            queryParameters['includeTentativeInventory'] = requestParameters.includeTentativeInventory;
        }
        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }
        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }
        if (requestParameters.endDateRange !== undefined) {
            queryParameters['endDateRange'] = requestParameters.endDateRange.toISOString().substr(0, 10);
        }
        if (requestParameters.startDateRange !== undefined) {
            queryParameters['startDateRange'] = requestParameters.startDateRange.toISOString().substr(0, 10);
        }
        if (requestParameters.blockStatusCodes) {
            queryParameters['blockStatusCodes'] = requestParameters.blockStatusCodes;
        }
        if (requestParameters.originCodes) {
            queryParameters['originCodes'] = requestParameters.originCodes;
        }
        if (requestParameters.roomOwnerCodes) {
            queryParameters['roomOwnerCodes'] = requestParameters.roomOwnerCodes;
        }
        if (requestParameters.blockOwners) {
            queryParameters['blockOwners'] = requestParameters.blockOwners;
        }
        if (requestParameters.summarySortByCode !== undefined) {
            queryParameters['summarySortByCode'] = requestParameters.summarySortByCode;
        }
        if (requestParameters.detailSortByCode !== undefined) {
            queryParameters['detailSortByCode'] = requestParameters.detailSortByCode;
        }
        if (requestParameters.friday !== undefined) {
            queryParameters['friday'] = requestParameters.friday;
        }
        if (requestParameters.monday !== undefined) {
            queryParameters['monday'] = requestParameters.monday;
        }
        if (requestParameters.saturday !== undefined) {
            queryParameters['saturday'] = requestParameters.saturday;
        }
        if (requestParameters.sunday !== undefined) {
            queryParameters['sunday'] = requestParameters.sunday;
        }
        if (requestParameters.thursday !== undefined) {
            queryParameters['thursday'] = requestParameters.thursday;
        }
        if (requestParameters.tuesday !== undefined) {
            queryParameters['tuesday'] = requestParameters.tuesday;
        }
        if (requestParameters.wednesday !== undefined) {
            queryParameters['wednesday'] = requestParameters.wednesday;
        }
        if (requestParameters.fetchInstructions) {
            queryParameters['fetchInstructions'] = requestParameters.fetchInstructions;
        }
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
        if (requestParameters.xExternalsystem !== undefined && requestParameters.xExternalsystem !== null) {
            headerParameters['x-externalsystem'] = String(requestParameters.xExternalsystem);
        }
        if (requestParameters.acceptLanguage !== undefined && requestParameters.acceptLanguage !== null) {
            headerParameters['Accept-Language'] = String(requestParameters.acceptLanguage);
        }
        const response = await this.request({
            path: `/hotels/{hotelId}/blockInventoryStatistics`.replace(`{${"hotelId"}}`, encodeURIComponent(String(requestParameters.hotelId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);
        return new runtime.JSONApiResponse(response, (jsonValue) => (0, models_1.BlockInventoryStatisticFromJSON)(jsonValue));
    }
    /**
     * Use this API to get block statistics on a per-day basis across a specified date range. The result includes a summary section of rooms booked by status by date, details on deduct and non-deduct rooms booked per day, and Sales Allowance and House Availability figures per day. The details section of the results include block details such as name, block, code, start and end date, owner, and more. It also shows a breakdown of blocked picked-up and available rooms per day per block. The result set is used to feed the GRC (Group Rooms Control) page in the OPERA UI. <p><strong>OperationId:</strong>getBlockInventoryStatistics</p>
     * Get block inventory statistics
     */
    async getBlockInventoryStatistics(requestParameters, initOverrides) {
        const response = await this.getBlockInventoryStatisticsRaw(requestParameters, initOverrides);
        return await response.value();
    }
    /**
     * Use this API to retrieve the inventory data for a specified hotel. Narrow down your results using the query parameters such as a date range, room type, room class, and/or tentative inventory included. Maximum days limit with a single request is 30 days.<p><strong>OperationId:</strong>getInventoryStatistics</p>
     * Get hotel inventory
     */
    async getInventoryStatisticsRaw(requestParameters, initOverrides) {
        const queryParameters = {};
        if (requestParameters.dateRangeEnd !== undefined) {
            queryParameters['dateRangeEnd'] = requestParameters.dateRangeEnd.toISOString().substr(0, 10);
        }
        if (requestParameters.reportCode !== undefined) {
            queryParameters['reportCode'] = requestParameters.reportCode;
        }
        if (requestParameters.dateRangeStart !== undefined) {
            queryParameters['dateRangeStart'] = requestParameters.dateRangeStart.toISOString().substr(0, 10);
        }
        if (requestParameters.parameterName) {
            queryParameters['parameterName'] = requestParameters.parameterName;
        }
        if (requestParameters.parameterValue) {
            queryParameters['parameterValue'] = requestParameters.parameterValue;
        }
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
        if (requestParameters.xExternalsystem !== undefined && requestParameters.xExternalsystem !== null) {
            headerParameters['x-externalsystem'] = String(requestParameters.xExternalsystem);
        }
        if (requestParameters.acceptLanguage !== undefined && requestParameters.acceptLanguage !== null) {
            headerParameters['Accept-Language'] = String(requestParameters.acceptLanguage);
        }
        const response = await this.request({
            path: `/hotels/{hotelId}/inventoryStatistics`.replace(`{${"hotelId"}}`, encodeURIComponent(String(requestParameters.hotelId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);
        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(models_1.StatisticTypeFromJSON));
    }
    /**
     * Use this API to retrieve the inventory data for a specified hotel. Narrow down your results using the query parameters such as a date range, room type, room class, and/or tentative inventory included. Maximum days limit with a single request is 30 days.<p><strong>OperationId:</strong>getInventoryStatistics</p>
     * Get hotel inventory
     */
    async getInventoryStatistics(requestParameters, initOverrides) {
        const response = await this.getInventoryStatisticsRaw(requestParameters, initOverrides);
        return await response.value();
    }
    /**
     * Check inventory statistics service <p><strong>OperationId:</strong>pinginvStatsService</p>
     * Ping
     */
    async pinginvStatsServiceRaw(requestParameters, initOverrides) {
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
        if (requestParameters.xExternalsystem !== undefined && requestParameters.xExternalsystem !== null) {
            headerParameters['x-externalsystem'] = String(requestParameters.xExternalsystem);
        }
        if (requestParameters.acceptLanguage !== undefined && requestParameters.acceptLanguage !== null) {
            headerParameters['Accept-Language'] = String(requestParameters.acceptLanguage);
        }
        const response = await this.request({
            path: `/services/invStatsService/status`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);
        return new runtime.JSONApiResponse(response, (jsonValue) => (0, models_1.PingDetailFromJSON)(jsonValue));
    }
    /**
     * Check inventory statistics service <p><strong>OperationId:</strong>pinginvStatsService</p>
     * Ping
     */
    async pinginvStatsService(requestParameters, initOverrides) {
        const response = await this.pinginvStatsServiceRaw(requestParameters, initOverrides);
        return await response.value();
    }
}
exports.INVStatsApi = INVStatsApi;
/**
 * @export
 */
exports.GetBlockInventoryStatisticsSummarySortByCodeEnum = {
    Status: 'Status',
    StatusDescending: 'StatusDescending',
    ArrivalDate: 'ArrivalDate',
    BlockAccount: 'BlockAccount',
    Revenue: 'Revenue',
    RoomsOwner: 'RoomsOwner'
};
/**
 * @export
 */
exports.GetBlockInventoryStatisticsDetailSortByCodeEnum = {
    Status: 'Status',
    StatusDescending: 'StatusDescending',
    ArrivalDate: 'ArrivalDate',
    BlockAccount: 'BlockAccount',
    Revenue: 'Revenue',
    RoomsOwner: 'RoomsOwner'
};
/**
 * @export
 */
exports.GetBlockInventoryStatisticsFetchInstructionsEnum = {
    StatusSummary: 'StatusSummary',
    BookingSummary: 'BookingSummary',
    ForecastSummary: 'ForecastSummary',
    DetailSummary: 'DetailSummary',
    RoomSummary: 'RoomSummary'
};
/**
 * @export
 */
exports.GetInventoryStatisticsReportCodeEnum = {
    DetailedAvailabiltySummary: 'DetailedAvailabiltySummary',
    RoomCalendarStatistics: 'RoomCalendarStatistics',
    SellLimitSummary: 'SellLimitSummary',
    RoomsAvailabilitySummary: 'RoomsAvailabilitySummary'
};
/**
 * @export
 */
exports.GetInventoryStatisticsParameterNameEnum = {
    CancelledYn: 'CancelledYN',
    Channel: 'Channel',
    DeductRoomsYn: 'DeductRoomsYN',
    GroupBy: 'GroupBy',
    HouseArrPersonsYn: 'HouseArrPersonsYN',
    HouseArrRoomsYn: 'HouseArrRoomsYN',
    HouseAvailPercYn: 'HouseAvailPercYN',
    HouseAvailRoomsExcludingOverbookingYn: 'HouseAvailRoomsExcludingOverbookingYN',
    HouseAvailRoomsYn: 'HouseAvailRoomsYN',
    HouseAvailTenPercYn: 'HouseAvailTenPercYN',
    HouseBlkDeductNpuYn: 'HouseBlkDeductNpuYN',
    HouseBlkDeductPuYn: 'HouseBlkDeductPuYN',
    HouseBlkTentNpuYn: 'HouseBlkTentNpuYN',
    HouseBlkTentPuYn: 'HouseBlkTentPuYN',
    HouseDayUsePersonYn: 'HouseDayUsePersonYN',
    HouseDayUseRoomYn: 'HouseDayUseRoomYN',
    HouseDepPersonsYn: 'HouseDepPersonsYN',
    HouseDepRoomsYn: 'HouseDepRoomsYN',
    HouseInventoryRoomsYn: 'HouseInventoryRoomsYN',
    HouseMaxAvailabilityExcludingOverbookingYn: 'HouseMaxAvailabilityExcludingOverbookingYN',
    HouseMaxAvailabilityYn: 'HouseMaxAvailabilityYN',
    HouseMaxOccupancyYn: 'HouseMaxOccupancyYN',
    HouseMinAvailabilityExcludingOverbookingYn: 'HouseMinAvailabilityExcludingOverbookingYN',
    HouseMinAvailabilityYn: 'HouseMinAvailabilityYN',
    HouseOccPercYn: 'HouseOccPercYN',
    HouseOccTenPercYn: 'HouseOccTenPercYN',
    HouseOccupancyYn: 'HouseOccupancyYN',
    HouseOooyn: 'HouseOOOYN',
    HouseOosRoomsYn: 'HouseOOSRoomsYN',
    HouseOverBookingYn: 'HouseOverBookingYN',
    HousePeopleInHouseYn: 'HousePeopleInHouseYN',
    HousePhysicalRoomsYn: 'HousePhysicalRoomsYN',
    HouseRoomsSoldYn: 'HouseRoomsSoldYN',
    HouseSellLimitYn: 'HouseSellLimitYN',
    HouseTentRoomsExcludingOverbookingYn: 'HouseTentRoomsExcludingOverbookingYN',
    HouseTentRoomsSoldYn: 'HouseTentRoomsSoldYN',
    HouseTentRoomsYn: 'HouseTentRoomsYN',
    OutOfOrderRoomsYn: 'OutOfOrderRoomsYN',
    RestrictionsYn: 'RestrictionsYN',
    RoomArrPersonsYn: 'RoomArrPersonsYN',
    RoomArrRoomsYn: 'RoomArrRoomsYN',
    RoomAvailRoomsExcludingOverbookingYn: 'RoomAvailRoomsExcludingOverbookingYN',
    RoomAvailRoomsYn: 'RoomAvailRoomsYN',
    RoomBlkDeductNpuYn: 'RoomBlkDeductNpuYN',
    RoomBlkDeductPuYn: 'RoomBlkDeductPuYN',
    RoomBlkTentNpuYn: 'RoomBlkTentNpuYN',
    RoomBlkTentPuYn: 'RoomBlkTentPuYN',
    RoomCancelledYn: 'RoomCancelledYN',
    RoomClassList: 'RoomClassList',
    RoomDayUsePersonYn: 'RoomDayUsePersonYN',
    RoomDayUseRoomYn: 'RoomDayUseRoomYN',
    RoomDepPersonsYn: 'RoomDepPersonsYN',
    RoomDepRoomsYn: 'RoomDepRoomsYN',
    RoomInventoryRoomsYn: 'RoomInventoryRoomsYN',
    RoomMaxAvailabilityExcludingOverbookingYn: 'RoomMaxAvailabilityExcludingOverbookingYN',
    RoomMaxAvailabilityYn: 'RoomMaxAvailabilityYN',
    RoomMaxOccupancyYn: 'RoomMaxOccupancyYN',
    RoomMinAvailabilityExcludingOverbookingYn: 'RoomMinAvailabilityExcludingOverbookingYN',
    RoomMinAvailabilityYn: 'RoomMinAvailabilityYN',
    RoomOccupancyYn: 'RoomOccupancyYN',
    RoomOooyn: 'RoomOOOYN',
    RoomOosRoomsYn: 'RoomOOSRoomsYN',
    RoomOverBookingYn: 'RoomOverBookingYN',
    RoomPeopleInHouseYn: 'RoomPeopleInHouseYN',
    RoomPhysicalRoomsYn: 'RoomPhysicalRoomsYN',
    RoomRestrictionsYn: 'RoomRestrictionsYN',
    RoomRoomsSoldYn: 'RoomRoomsSoldYN',
    RoomSellLimitYn: 'RoomSellLimitYN',
    RoomTentRoomsExcludingOverbookingYn: 'RoomTentRoomsExcludingOverbookingYN',
    RoomTentRoomsYn: 'RoomTentRoomsYN',
    RoomTentYn: 'RoomTentYN',
    RoomTypeWildCardList: 'RoomTypeWildCardList'
};
