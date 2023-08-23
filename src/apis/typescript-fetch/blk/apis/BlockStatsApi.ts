/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  BlockStats,
  ExceptionDetailType,
  OperaVersion,
  Status,
} from '../models';
import {
    BlockStatsFromJSON,
    BlockStatsToJSON,
    ExceptionDetailTypeFromJSON,
    ExceptionDetailTypeToJSON,
    OperaVersionFromJSON,
    OperaVersionToJSON,
    StatusFromJSON,
    StatusToJSON,
} from '../models';

export interface DeleteBlockStatsServiceCacheRequest {
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    xExternalsystem?: string;
    acceptLanguage?: string;
}

export interface GetBlockStatsRequest {
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    hotelId?: string;
    reportCode?: Set<GetBlockStatsReportCodeEnum>;
    reportEndDate?: Array<Date>;
    reportStartDate?: Array<Date>;
    statisticalCode?: Set<GetBlockStatsStatisticalCodeEnum>;
    reportParametersParameterName?: Array<string>;
    reportParametersParameterValue?: Array<string>;
    blockOwnersCode?: Set<GetBlockStatsBlockOwnersCodeEnum>;
    blockStatusCode?: Set<GetBlockStatsBlockStatusCodeEnum>;
    xExternalsystem?: string;
    acceptLanguage?: string;
}

export interface PingBlockStatsServiceRequest {
    authorization?: string;
    xAppKey?: string;
    xHotelid?: string;
    xExternalsystem?: string;
    acceptLanguage?: string;
}

/**
 * 
 */
export class BlockStatsApi extends runtime.BaseAPI {

    /**
     * Use this API to delete Block Stat Service cache. <p><strong>OperationId:</strong>deleteBlockStatsServiceCache</p>
     * Delete Block Stat Service cache
     */
    async deleteBlockStatsServiceCacheRaw(requestParameters: DeleteBlockStatsServiceCacheRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Status>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

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
            path: `/services/blockStats/cache`,
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => StatusFromJSON(jsonValue));
    }

    /**
     * Use this API to delete Block Stat Service cache. <p><strong>OperationId:</strong>deleteBlockStatsServiceCache</p>
     * Delete Block Stat Service cache
     */
    async deleteBlockStatsServiceCache(requestParameters: DeleteBlockStatsServiceCacheRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Status> {
        const response = await this.deleteBlockStatsServiceCacheRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * This API retrieves block statistics. <p><strong>OperationId:</strong>getBlockStats</p>
     * Get Block statistics
     */
    async getBlockStatsRaw(requestParameters: GetBlockStatsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<BlockStats>> {
        const queryParameters: any = {};

        if (requestParameters.hotelId !== undefined) {
            queryParameters['hotelId'] = requestParameters.hotelId;
        }

        if (requestParameters.reportCode) {
            queryParameters['reportCode'] = requestParameters.reportCode;
        }

        if (requestParameters.reportEndDate) {
            queryParameters['reportEndDate'] = requestParameters.reportEndDate;
        }

        if (requestParameters.reportStartDate) {
            queryParameters['reportStartDate'] = requestParameters.reportStartDate;
        }

        if (requestParameters.statisticalCode) {
            queryParameters['statisticalCode'] = requestParameters.statisticalCode;
        }

        if (requestParameters.reportParametersParameterName) {
            queryParameters['reportParametersParameterName'] = requestParameters.reportParametersParameterName;
        }

        if (requestParameters.reportParametersParameterValue) {
            queryParameters['reportParametersParameterValue'] = requestParameters.reportParametersParameterValue;
        }

        if (requestParameters.blockOwnersCode) {
            queryParameters['blockOwnersCode'] = requestParameters.blockOwnersCode;
        }

        if (requestParameters.blockStatusCode) {
            queryParameters['blockStatusCode'] = requestParameters.blockStatusCode;
        }

        const headerParameters: runtime.HTTPHeaders = {};

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
            path: `/blocks/statistics`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => BlockStatsFromJSON(jsonValue));
    }

    /**
     * This API retrieves block statistics. <p><strong>OperationId:</strong>getBlockStats</p>
     * Get Block statistics
     */
    async getBlockStats(requestParameters: GetBlockStatsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<BlockStats> {
        const response = await this.getBlockStatsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Use this API to retrieve block stats service status. <p><strong>OperationId:</strong>pingBlockStatsService</p>
     * Retrieve/Ping  Block stats service
     */
    async pingBlockStatsServiceRaw(requestParameters: PingBlockStatsServiceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<OperaVersion>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

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
            path: `/services/blockStats/status`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => OperaVersionFromJSON(jsonValue));
    }

    /**
     * Use this API to retrieve block stats service status. <p><strong>OperationId:</strong>pingBlockStatsService</p>
     * Retrieve/Ping  Block stats service
     */
    async pingBlockStatsService(requestParameters: PingBlockStatsServiceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<OperaVersion> {
        const response = await this.pingBlockStatsServiceRaw(requestParameters, initOverrides);
        return await response.value();
    }

}

/**
 * @export
 */
export const GetBlockStatsReportCodeEnum = {
    HouseSummary: 'HouseSummary',
    HouseSummaryMonthToDate: 'HouseSummaryMonthToDate',
    HouseSummaryYearToDate: 'HouseSummaryYearToDate',
    ReservationActivity: 'ReservationActivity',
    ReservationActivityDaily: 'ReservationActivityDaily',
    ReservationActivityMonthToDate: 'ReservationActivityMonthToDate',
    ReservationActivityYearToDate: 'ReservationActivityYearToDate',
    ComplimentaryOrHouseUse: 'ComplimentaryOrHouseUse',
    DailyProjection: 'DailyProjection',
    RoomStatus: 'RoomStatus',
    LastHourStatus: 'LastHourStatus',
    Turndown: 'Turndown',
    InHouse: 'InHouse',
    CheckIns: 'CheckIns',
    CheckOuts: 'CheckOuts',
    AvailableRooms: 'AvailableRooms',
    RoomMaintenance: 'RoomMaintenance',
    ReservationStatistics: 'ReservationStatistics',
    VipGuests: 'VIPGuests',
    AdvanceCheckIn: 'AdvanceCheckIn',
    BlockArrivals: 'BlockArrivals',
    ReservationsCancellationsToday: 'ReservationsCancellationsToday',
    AiRoomAssignment: 'AIRoomAssignment',
    CompRouting: 'CompRouting'
} as const;
export type GetBlockStatsReportCodeEnum = typeof GetBlockStatsReportCodeEnum[keyof typeof GetBlockStatsReportCodeEnum];
/**
 * @export
 */
export const GetBlockStatsStatisticalCodeEnum = {
    ApprovedCompPostings: 'ApprovedCompPostings',
    StagedCompPostings: 'StagedCompPostings',
    DeclinedCompPostings: 'DeclinedCompPostings',
    CompRoutingInstructionsRequests: 'CompRoutingInstructionsRequests',
    TotalPhysicalRooms: 'TotalPhysicalRooms',
    TotalRoomsToSell: 'TotalRoomsToSell',
    TotalOutOfOrder: 'TotalOutOfOrder',
    TotalOutOfService: 'TotalOutOfService',
    TotalRevenue: 'TotalRevenue',
    PercentRoomsOccupied: 'PercentRoomsOccupied',
    StayoverRooms: 'StayoverRooms',
    StayoverPersons: 'StayoverPersons',
    StayoverVip: 'StayoverVIP',
    DeparturesExpectedRooms: 'DeparturesExpectedRooms',
    DeparturesExpectedPersons: 'DeparturesExpectedPersons',
    DeparturesExpectedVip: 'DeparturesExpectedVIP',
    DeparturesActualRooms: 'DeparturesActualRooms',
    DeparturesActualPersons: 'DeparturesActualPersons',
    DeparturesActualVip: 'DeparturesActualVIP',
    DeparturesInLastHour: 'DeparturesInLastHour',
    ArrivalsInLastHour: 'ArrivalsInLastHour',
    ArrivalsTotal: 'ArrivalsTotal',
    ArrivalsExpectedRooms: 'ArrivalsExpectedRooms',
    ArrivalsExpectedPersons: 'ArrivalsExpectedPersons',
    ArrivalsExpectedVip: 'ArrivalsExpectedVIP',
    ArrivalsExpectedRoomsMadeToday: 'ArrivalsExpectedRoomsMadeToday',
    ArrivalsExpectedPersonsMadeToday: 'ArrivalsExpectedPersonsMadeToday',
    ArrivalsExpectedVipMadeToday: 'ArrivalsExpectedVIPMadeToday',
    ArrivalsActualRooms: 'ArrivalsActualRooms',
    ArrivalsActualPersons: 'ArrivalsActualPersons',
    ArrivalsActualVip: 'ArrivalsActualVIP',
    ExtendedStaysRooms: 'ExtendedStaysRooms',
    ExtendedStaysPersons: 'ExtendedStaysPersons',
    ExtendedStaysVip: 'ExtendedStaysVIP',
    DeparturesTotal: 'DeparturesTotal',
    EarlyDeparturesRooms: 'EarlyDeparturesRooms',
    EarlyDeparturesPersons: 'EarlyDeparturesPersons',
    EarlyDeparturesVip: 'EarlyDeparturesVIP',
    DayUseRooms: 'DayUseRooms',
    DayUsePersons: 'DayUsePersons',
    DayUseVip: 'DayUseVIP',
    WalkInRooms: 'WalkInRooms',
    WalkInPersons: 'WalkInPersons',
    WalkInVip: 'WalkInVIP',
    CanceledOnArrivalRooms: 'CanceledOnArrivalRooms',
    CanceledOnArrivalPersons: 'CanceledOnArrivalPersons',
    CanceledOnArrivalVip: 'CanceledOnArrivalVIP',
    ComplimentaryArrivalRooms: 'ComplimentaryArrivalRooms',
    ComplimentaryArrivalPersons: 'ComplimentaryArrivalPersons',
    ComplimentaryArrivalVip: 'ComplimentaryArrivalVIP',
    ComplimentaryStayoverRooms: 'ComplimentaryStayoverRooms',
    ComplimentaryStayoverPersons: 'ComplimentaryStayoverPersons',
    ComplimentaryStayoverVip: 'ComplimentaryStayoverVIP',
    ComplimentaryDepartureRooms: 'ComplimentaryDepartureRooms',
    ComplimentaryDeparturePersons: 'ComplimentaryDeparturePersons',
    ComplimentaryDepartureVip: 'ComplimentaryDepartureVIP',
    HouseUseArrivalVip: 'HouseUseArrivalVIP',
    HouseUseStayoverRooms: 'HouseUseStayoverRooms',
    HouseUseStayoverPersons: 'HouseUseStayoverPersons',
    HouseUseStayoverVip: 'HouseUseStayoverVIP',
    HouseUseDepartureRooms: 'HouseUseDepartureRooms',
    HouseUseDeparturePersons: 'HouseUseDeparturePersons',
    HouseUseDepartureVip: 'HouseUseDepartureVIP',
    MinAvailableTonightRooms: 'MinAvailableTonightRooms',
    MaxOccupiedTonightRooms: 'MaxOccupiedTonightRooms',
    MaxOccupiedTonightPersons: 'MaxOccupiedTonightPersons',
    MaxOccupiedTonightVip: 'MaxOccupiedTonightVIP',
    MaxPercentageOccupiedTonightRooms: 'MaxPercentageOccupiedTonightRooms',
    BlocksNotPickedUp: 'BlocksNotPickedUp',
    IndividualRooms: 'IndividualRooms',
    IndividualPersons: 'IndividualPersons',
    IndividualVip: 'IndividualVIP',
    GroupAndBlockRooms: 'GroupAndBlockRooms',
    GroupAndBlockPersons: 'GroupAndBlockPersons',
    GroupAndBlockVip: 'GroupAndBlockVIP',
    RoomRevenue: 'RoomRevenue',
    AverageRoomRevenue: 'AverageRoomRevenue',
    InspectedRooms: 'InspectedRooms',
    InspectedVacant: 'InspectedVacant',
    InspectedAssigned: 'InspectedAssigned',
    InspectedOccupied: 'InspectedOccupied',
    CleanedRooms: 'CleanedRooms',
    CleanVacant: 'CleanVacant',
    CleanAssigned: 'CleanAssigned',
    CleanOccupied: 'CleanOccupied',
    DirtyVacant: 'DirtyVacant',
    DirtyAssigned: 'DirtyAssigned',
    DirtyOccupied: 'DirtyOccupied',
    PickupVacant: 'PickupVacant',
    PickupAssigned: 'PickupAssigned',
    PickupOccupied: 'PickupOccupied',
    OutOfOrderVacant: 'OutOfOrderVacant',
    OutOfOrderAssigned: 'OutOfOrderAssigned',
    OutOfOrderOccupied: 'OutOfOrderOccupied',
    OutOfServiceVacant: 'OutOfServiceVacant',
    OutOfServiceAssigned: 'OutOfServiceAssigned',
    OutOfServiceOccupied: 'OutOfServiceOccupied',
    QueueRooms: 'QueueRooms',
    TurndownRequired: 'TurndownRequired',
    TurndownNotRequired: 'TurndownNotRequired',
    TurndownCompletedRequired: 'TurndownCompletedRequired',
    AdultsInHouse: 'AdultsInHouse',
    ChildrenInHouse: 'ChildrenInHouse',
    AdultsExpectedCheckedOut: 'AdultsExpectedCheckedOut',
    ChildrenExpectedCheckedOut: 'ChildrenExpectedCheckedOut',
    AdultsDeparted: 'AdultsDeparted',
    ChildrenDeparted: 'ChildrenDeparted',
    InHouseRooms: 'InHouseRooms',
    InHouse: 'InHouse',
    MaxOccupancyPercentage: 'MaxOccupancyPercentage',
    Stayover: 'Stayover',
    TotalRoomsReserved: 'TotalRoomsReserved',
    RevPar: 'RevPar',
    SkipRooms: 'SkipRooms',
    SleepRooms: 'SleepRooms',
    HouseUseArrivalRooms: 'HouseUseArrivalRooms',
    HouseUseArrivalPersons: 'HouseUseArrivalPersons',
    AverageCheckInTime: 'AverageCheckInTime',
    CurrentAveWaitTime: 'CurrentAveWaitTime',
    CheckedInsTotal: 'CheckedInsTotal',
    ExpectedCheckInsTotal: 'ExpectedCheckInsTotal',
    CheckedOutsTotal: 'CheckedOutsTotal',
    ExpectedCheckOutsTotal: 'ExpectedCheckOutsTotal',
    ScheduledCheckOutsTotal: 'ScheduledCheckOutsTotal',
    RoomMaintenanceResolvedTotal: 'RoomMaintenanceResolvedTotal',
    RoomMaintenanceUnResolvedTotal: 'RoomMaintenanceUnResolvedTotal',
    PreRegisteredTotal: 'PreRegisteredTotal',
    VipPreRegisteredTotal: 'VIPPreRegisteredTotal',
    OpenFolioTotal: 'OpenFolioTotal',
    TurndownTotal: 'TurndownTotal',
    VipTurndownTotal: 'VIPTurndownTotal',
    VipGuestsArriving: 'VIPGuestsArriving',
    VipGuestsDeparting: 'VIPGuestsDeparting',
    VipGuestsTotal: 'VIPGuestsTotal',
    IndividualAdvanceCheckedInCurrent: 'IndividualAdvanceCheckedInCurrent',
    BlockAdvanceCheckedInCurrent: 'BlockAdvanceCheckedInCurrent',
    IndividualAdvanceCheckedInInhouse: 'IndividualAdvanceCheckedInInhouse',
    BlockAdvanceCheckedInInhouse: 'BlockAdvanceCheckedInInhouse',
    IndividualAdvanceCheckedInTotal: 'IndividualAdvanceCheckedInTotal',
    BlockAdvanceCheckedInTotal: 'BlockAdvanceCheckedInTotal',
    InHouseBlocksTotal: 'InHouseBlocksTotal',
    CancellationsTotal: 'CancellationsTotal',
    NewReservationsTotal: 'NewReservationsTotal',
    ArrivalResvs: 'ArrivalResvs',
    ArrivalVipResvs: 'ArrivalVIPResvs',
    ArrivalMemberResvs: 'ArrivalMemberResvs',
    ArrivalUnallocResvs: 'ArrivalUnallocResvs',
    ArrivalManualAssgnResvs: 'ArrivalManualAssgnResvs',
    ArrivalAiAssgnResvs: 'ArrivalAIAssgnResvs',
    ArrivalAiUpgResvs: 'ArrivalAIUpgResvs',
    ArrivalAiAssgnVipResvs: 'ArrivalAIAssgnVIPResvs',
    ArrivalAiAssgnMemberResvs: 'ArrivalAIAssgnMemberResvs',
    ArrivalAiAssgnOverridden: 'ArrivalAIAssgnOverridden'
} as const;
export type GetBlockStatsStatisticalCodeEnum = typeof GetBlockStatsStatisticalCodeEnum[keyof typeof GetBlockStatsStatisticalCodeEnum];
/**
 * @export
 */
export const GetBlockStatsBlockOwnersCodeEnum = {
    HouseSummary: 'HouseSummary',
    HouseSummaryMonthToDate: 'HouseSummaryMonthToDate',
    HouseSummaryYearToDate: 'HouseSummaryYearToDate',
    ReservationActivity: 'ReservationActivity',
    ReservationActivityDaily: 'ReservationActivityDaily',
    ReservationActivityMonthToDate: 'ReservationActivityMonthToDate',
    ReservationActivityYearToDate: 'ReservationActivityYearToDate',
    ComplimentaryOrHouseUse: 'ComplimentaryOrHouseUse',
    DailyProjection: 'DailyProjection',
    RoomStatus: 'RoomStatus',
    LastHourStatus: 'LastHourStatus',
    Turndown: 'Turndown',
    InHouse: 'InHouse',
    CheckIns: 'CheckIns',
    CheckOuts: 'CheckOuts',
    AvailableRooms: 'AvailableRooms',
    RoomMaintenance: 'RoomMaintenance',
    ReservationStatistics: 'ReservationStatistics',
    VipGuests: 'VIPGuests',
    AdvanceCheckIn: 'AdvanceCheckIn',
    BlockArrivals: 'BlockArrivals',
    ReservationsCancellationsToday: 'ReservationsCancellationsToday',
    AiRoomAssignment: 'AIRoomAssignment',
    CompRouting: 'CompRouting'
} as const;
export type GetBlockStatsBlockOwnersCodeEnum = typeof GetBlockStatsBlockOwnersCodeEnum[keyof typeof GetBlockStatsBlockOwnersCodeEnum];
/**
 * @export
 */
export const GetBlockStatsBlockStatusCodeEnum = {
    HouseSummary: 'HouseSummary',
    HouseSummaryMonthToDate: 'HouseSummaryMonthToDate',
    HouseSummaryYearToDate: 'HouseSummaryYearToDate',
    ReservationActivity: 'ReservationActivity',
    ReservationActivityDaily: 'ReservationActivityDaily',
    ReservationActivityMonthToDate: 'ReservationActivityMonthToDate',
    ReservationActivityYearToDate: 'ReservationActivityYearToDate',
    ComplimentaryOrHouseUse: 'ComplimentaryOrHouseUse',
    DailyProjection: 'DailyProjection',
    RoomStatus: 'RoomStatus',
    LastHourStatus: 'LastHourStatus',
    Turndown: 'Turndown',
    InHouse: 'InHouse',
    CheckIns: 'CheckIns',
    CheckOuts: 'CheckOuts',
    AvailableRooms: 'AvailableRooms',
    RoomMaintenance: 'RoomMaintenance',
    ReservationStatistics: 'ReservationStatistics',
    VipGuests: 'VIPGuests',
    AdvanceCheckIn: 'AdvanceCheckIn',
    BlockArrivals: 'BlockArrivals',
    ReservationsCancellationsToday: 'ReservationsCancellationsToday',
    AiRoomAssignment: 'AIRoomAssignment',
    CompRouting: 'CompRouting'
} as const;
export type GetBlockStatsBlockStatusCodeEnum = typeof GetBlockStatsBlockStatusCodeEnum[keyof typeof GetBlockStatsBlockStatusCodeEnum];
